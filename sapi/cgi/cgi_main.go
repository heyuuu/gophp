package cgi

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"log"
)

func main(argc int, argv []*byte) int {
	var free_query_string int = 0
	var exit_status int = types.SUCCESS
	var cgi int = 0
	var c int
	var i int
	var len_ int
	var file_handle zend.ZendFileHandle
	var s *byte

	/* temporary locals */

	var behavior int = PHP_MODE_STANDARD
	var no_headers int = 0
	var orig_optind int = PhpOptind
	var orig_optarg *byte = PhpOptarg
	var script_file *byte = nil
	var ini_entries_len int = 0

	/* end of temporary locals */

	var max_requests int = 500
	var requests int = 0
	var fastcgi int
	var bindpath *byte = nil
	var fcgi_fd int = 0
	var request *core.FcgiRequest = nil
	var warmup_repeats int = 0
	var repeats int = 1
	var benchmark int = 0
	var start __struct__timeval
	var end __struct__timeval
	var status int = 0
	var query_string *byte
	var decoded_query_string *byte
	var skip_getopt int = 0

	app := core.NewApp()

	zend.ZendSignalStartup()
	PhpCgiGlobalsCtor(&php_cgi_globals)
	app.SapiStartup(CgiModule)
	fastcgi = core.FcgiIsFastcgi()
	CgiModule.SetPhpIniPathOverride(nil)
	if fastcgi == 0 {

		/* Make sure we detect we are a cgi - a bit redundancy here,
		 * but the default case is that we have to check only the first one. */
		if getenv("SERVER_SOFTWARE") || getenv("SERVER_NAME") || getenv("GATEWAY_INTERFACE") || getenv("REQUEST_METHOD") {
			cgi = 1
		}
	}
	if b.Assign(&query_string, getenv("QUERY_STRING")) != nil && strchr(query_string, '=') == nil {

		/* we've got query string that has no = - apache CGI will pass it to command line */

		var p *uint8
		decoded_query_string = strdup(query_string)
		streams.PhpUrlDecode(decoded_query_string, strlen(decoded_query_string))
		for p = (*uint8)(decoded_query_string); (*p) != 0 && (*p) <= ' '; p++ {

		}
		if (*p) == '-' {
			skip_getopt = 1
		}
		zend.Free(decoded_query_string)
	}
	for skip_getopt == 0 && b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 0, 2)) != -1 {
		switch c {
		case 'c':
			if CgiModule.GetPhpIniPathOverride() != nil {
				zend.Free(CgiModule.GetPhpIniPathOverride())
			}
			CgiModule.SetPhpIniPathOverride(strdup(PhpOptarg))
			break
		case 'n':
			CgiModule.SetPhpIniIgnore(1)
			break
		case 'd':

			/* define ini __special__  entries on command line */

			var len_ int = strlen(PhpOptarg)
			var val *byte
			if b.Assign(&val, strchr(PhpOptarg, '=')) {
				val++
				if !(isalnum(*val)) && (*val) != '"' && (*val) != '\'' && (*val) != '0' {
					CgiModule.SetIniEntries(realloc(CgiModule.GetIniEntries(), ini_entries_len+len_+b.SizeOf("\"\\\"\\\"\\n\\0\"")))
					memcpy(CgiModule.GetIniEntries()+ini_entries_len, PhpOptarg, val-PhpOptarg)
					ini_entries_len += val - PhpOptarg
					memcpy(CgiModule.GetIniEntries()+ini_entries_len, "\"", 1)
					ini_entries_len++
					memcpy(CgiModule.GetIniEntries()+ini_entries_len, val, len_-(val-PhpOptarg))
					ini_entries_len += len_ - (val - PhpOptarg)
					memcpy(CgiModule.GetIniEntries()+ini_entries_len, "\"\n0", b.SizeOf("\"\\\"\\n\\0\""))
					ini_entries_len += b.SizeOf("\"\\n\\0\\\"\"") - 2
				} else {
					CgiModule.SetIniEntries(realloc(CgiModule.GetIniEntries(), ini_entries_len+len_+b.SizeOf("\"\\n\\0\"")))
					memcpy(CgiModule.GetIniEntries()+ini_entries_len, PhpOptarg, len_)
					memcpy(CgiModule.GetIniEntries()+ini_entries_len+len_, "\n0", b.SizeOf("\"\\n\\0\""))
					ini_entries_len += len_ + b.SizeOf("\"\\n\\0\"") - 2
				}
			} else {
				CgiModule.SetIniEntries(realloc(CgiModule.GetIniEntries(), ini_entries_len+len_+b.SizeOf("\"=1\\n\\0\"")))
				memcpy(CgiModule.GetIniEntries()+ini_entries_len, PhpOptarg, len_)
				memcpy(CgiModule.GetIniEntries()+ini_entries_len+len_, "=1\n0", b.SizeOf("\"=1\\n\\0\""))
				ini_entries_len += len_ + b.SizeOf("\"=1\\n\\0\"") - 2
			}
			break
		case 'b':
			if fastcgi == 0 {
				bindpath = strdup(PhpOptarg)
			}
			break
		case 's':
			behavior = PHP_MODE_HIGHLIGHT
			break
		}
	}
	PhpOptind = orig_optind
	PhpOptarg = orig_optarg
	if fastcgi != 0 || bindpath != nil {
		/* Override SAPI callbacks */
		CgiModule.IsFastCgi = true
	}
	CgiModule.SetExecutableLocation(argv[0])

	/* startup after we get the above ini override se we get things right */

	if !CgiModule.Startup() {
		zend.Free(bindpath)
		return types.FAILURE
	}

	/* check force_cgi after startup, so we have proper output */

	if cgi != 0 && CGIG(force_redirect) {

		/* Apache will generate REDIRECT_STATUS,
		 * Netscape and redirect.so will generate HTTP_REDIRECT_STATUS.
		 * redirect.so and installation instructions available from
		 * http://www.koehntopp.de/php.
		 *   -- kk@netuse.de
		 */

		if !(getenv("REDIRECT_STATUS")) && !(getenv("HTTP_REDIRECT_STATUS")) && (!(CGIG(redirect_status_env)) || !(getenv(CGIG(redirect_status_env)))) {
			faults.Try(func() {
				core.SG__().sapi_headers.http_response_code = 400
				core.PUTS("<b>Security Alert!</b> The PHP CGI cannot be accessed directly.\n\n\n<p>This PHP CGI binary was compiled with force-cgi-redirect enabled.  This\n\nmeans that a page will only be served up if the REDIRECT_STATUS CGI variable is\n\nset, e.g. via an Apache Action directive.</p>\n\n<p>For more information as to <i>why</i> this behaviour exists, see the <a href=\"http://php.net/security.cgi-bin\">\nmanual page for CGI security</a>.</p>\n\n<p>For more information about changing this behaviour or re-enabling this webserver,\n\nconsult the installation file that came with this distribution, or visit \n\n<a href=\"http://php.net/install.windows\">the manual page</a>.</p>\n")
			})
			return types.FAILURE
		}

		/* Apache will generate REDIRECT_STATUS,
		 * Netscape and redirect.so will generate HTTP_REDIRECT_STATUS.
		 * redirect.so and installation instructions available from
		 * http://www.koehntopp.de/php.
		 *   -- kk@netuse.de
		 */

	}
	core.FcgiSetLogger(FcgiLog)
	if bindpath != nil {
		var backlog int = 128
		if getenv("PHP_FCGI_BACKLOG") {
			backlog = atoi(getenv("PHP_FCGI_BACKLOG"))
		}
		fcgi_fd = core.FcgiListen(bindpath, backlog)
		if fcgi_fd < 0 {
			log.Printf("Couldn't create FastCGI listen socket on port %s\n", bindpath)
			return types.FAILURE
		}
		fastcgi = core.FcgiIsFastcgi()
	}

	/* make php call us to get _ENV vars */

	PhpPhpImportEnvironmentVariables = core.PhpImportEnvironmentVariables
	core.PhpImportEnvironmentVariables = CgiPhpImportEnvironmentVariables
	if fastcgi != 0 {

		/* How many times to run PHP scripts before dying */

		if getenv("PHP_FCGI_MAX_REQUESTS") {
			max_requests = atoi(getenv("PHP_FCGI_MAX_REQUESTS"))
			if max_requests < 0 {
				log.Printf("PHP_FCGI_MAX_REQUESTS is not valid\n")
				return types.FAILURE
			}
		}

		/* library is already initialized, now init our request */

		request = core.FcgiInitRequest(fcgi_fd, nil, nil, nil)

		/* Pre-fork or spawn, if required */

		if getenv("PHP_FCGI_CHILDREN") {
			var children_str *byte = getenv("PHP_FCGI_CHILDREN")
			Children = atoi(children_str)
			if Children < 0 {
				log.Printf("PHP_FCGI_CHILDREN is not valid\n")
				return types.FAILURE
			}
			core.FcgiSetMgmtVar("FCGI_MAX_CONNS", b.SizeOf("\"FCGI_MAX_CONNS\"")-1, children_str, strlen(children_str))

			/* This is the number of concurrent requests, equals FCGI_MAX_CONNS */

			core.FcgiSetMgmtVar("FCGI_MAX_REQS", b.SizeOf("\"FCGI_MAX_REQS\"")-1, children_str, strlen(children_str))

			/* This is the number of concurrent requests, equals FCGI_MAX_CONNS */

		} else {
			core.FcgiSetMgmtVar("FCGI_MAX_CONNS", b.SizeOf("\"FCGI_MAX_CONNS\"")-1, "1", b.SizeOf("\"1\"")-1)
			core.FcgiSetMgmtVar("FCGI_MAX_REQS", b.SizeOf("\"FCGI_MAX_REQS\"")-1, "1", b.SizeOf("\"1\"")-1)
		}
		if Children != 0 {
			var running int = 0
			var pid pid_t

			/* Create a process group for ourself & children */

			setsid()
			Pgroup = getpgrp()

			/* Set up handler to kill children upon exit */

			Act.sa_flags = 0
			Act.sa_handler = FastcgiCleanup
			if sigaction(SIGTERM, &Act, &OldTerm) || sigaction(SIGINT, &Act, &OldInt) || sigaction(SIGQUIT, &Act, &OldQuit) {
				r.Perror("Can't set signals")
				exit(1)
			}
			if core.FcgiInShutdown() != 0 {
				goto parent_out
			}
			for Parent != 0 {
				for {
					pid = fork()
					switch pid {
					case 0:

						/* One of the children.
						 * Make sure we don't go round the
						 * fork loop any more
						 */

						Parent = 0

						/* don't catch our signals */

						sigaction(SIGTERM, &OldTerm, 0)
						sigaction(SIGQUIT, &OldQuit, 0)
						sigaction(SIGINT, &OldInt, 0)
						zend.ZendSignalInit()
						break
					case -1:
						r.Perror("php (pre-forking)")
						exit(1)
						break
					default:

						/* Fine */

						running++
						break
					}
					if !(Parent != 0 && running < Children) {
						break
					}
				}
				if Parent != 0 {
					ParentWaiting = 1
					for true {
						if wait(&status) >= 0 {
							running--
							break
						} else if ExitSignal != 0 {
							break
						}
					}
					if ExitSignal != 0 {
						goto parent_out
					}
				}
			}
		} else {
			Parent = 0
			zend.ZendSignalInit()
		}
	}

	faults.TryCatch(func() {
		for skip_getopt == 0 && b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 1, 2)) != -1 {
			switch c {
			case 'T':
				benchmark = 1
				var comma *byte = strchr(PhpOptarg, ',')
				if comma != nil {
					warmup_repeats = atoi(PhpOptarg)
					repeats = atoi(comma + 1)
				} else {
					repeats = atoi(PhpOptarg)
				}
				gettimeofday(&start, nil)
				break
			case 'h':

			case '?':

			case core.PHP_GETOPT_INVALID_ARG:
				if request != nil {
					core.FcgiDestroyRequest(request)
				}
				core.FcgiShutdown()
				no_headers = 1
				core.SG__().headers_sent = 1
				PhpCgiUsage(argv[0])
				core.PhpOutputEndAll()
				exit_status = 0
				if c == core.PHP_GETOPT_INVALID_ARG {
					exit_status = 1
				}
				goto out
			}
		}
		PhpOptind = orig_optind
		PhpOptarg = orig_optarg

		/* start of FAST CGI loop */

		for fastcgi == 0 || core.FcgiAcceptRequest(request) >= 0 {
			if fastcgi != 0 {
				core.SG__().server_context = any(request)
			} else {
				core.SG__().server_context = any(1)
			}
			InitRequestInfo(request)
			if cgi == 0 && fastcgi == 0 {
				for b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 0, 2)) != -1 {
					switch c {
					case 'a':
						r.Printf("Interactive mode enabled\n\n")
						break
					case 'C':
						core.SG__().options |= core.SAPI_OPTION_NO_CHDIR
						break
					case 'e':
						zend.CG__().SetCompilerOptions(zend.CG__().GetCompilerOptions() | zend.ZEND_COMPILE_EXTENDED_INFO)
						break
					case 'f':
						if script_file != nil {
							zend.Efree(script_file)
						}
						script_file = zend.Estrdup(PhpOptarg)
						no_headers = 1
						break
					case 'i':
						if script_file != nil {
							zend.Efree(script_file)
						}
						if core.PhpRequestStartup() == types.FAILURE {
							core.SG__().server_context = nil
							core.PhpModuleShutdown()
							zend.Free(bindpath)
							return types.FAILURE
						}
						if no_headers != 0 {
							core.SG__().headers_sent = 1
							core.SG__().RequestInfo.no_headers = 1
						}
						standard.PhpPrintInfo(0xffffffff)
						core.PhpRequestShutdown(any(0))
						core.FcgiShutdown()
						exit_status = 0
						goto out
					case 'l':
						no_headers = 1
						behavior = PHP_MODE_LINT
						break
					case 'm':
						if script_file != nil {
							zend.Efree(script_file)
						}
						core.SG__().headers_sent = 1
						core.PhpPrintf("[PHP Modules]\n")
						PrintModules()
						core.PhpPrintf("\n[Zend Modules]\n")
						PrintExtensions()
						core.PhpPrintf("\n")
						core.PhpOutputEndAll()
						core.FcgiShutdown()
						exit_status = 0
						goto out
					case 'q':
						no_headers = 1
						break
					case 'v':
						if script_file != nil {
							zend.Efree(script_file)
						}
						no_headers = 1
						if core.PhpRequestStartup() == types.FAILURE {
							core.SG__().server_context = nil
							core.PhpModuleShutdown()
							zend.Free(bindpath)
							return types.FAILURE
						}
						core.SG__().headers_sent = 1
						core.SG__().RequestInfo.no_headers = 1
						core.PhpPrintf("PHP %s (%s) (built: %s %s)\nCopyright (c) The PHP Group\n%s", core.PHP_VERSION, core.SM__().Name(), __DATE__, __TIME__, zend.GetZendVersion())
						core.PhpRequestShutdown(any(0))
						core.FcgiShutdown()
						exit_status = 0
						goto out
					case 'w':
						behavior = PHP_MODE_STRIP
						break
					case 'z':
						zend.ZendLoadExtension(PhpOptarg)
						break
					default:
						break
					}
				}
				if script_file != nil {

					/* override path_translated if -f on command line */

					if core.SG__().RequestInfo.path_translated {
						zend.Efree(core.SG__().RequestInfo.path_translated)
					}
					core.SG__().RequestInfo.path_translated = script_file

					/* before registering argv to module exchange the *new* argv[0] */

					core.SG__().RequestInfo.argc = argc - (PhpOptind - 1)
					core.SG__().RequestInfo.argv = &argv[PhpOptind-1]
					core.SG__().RequestInfo.argv[0] = script_file
				} else if argc > PhpOptind {

					/* file is on command line, but not in -f opt */

					if core.SG__().RequestInfo.path_translated {
						zend.Efree(core.SG__().RequestInfo.path_translated)
					}
					core.SG__().RequestInfo.path_translated = zend.Estrdup(argv[PhpOptind])

					/* arguments after the file are considered script args */

					core.SG__().RequestInfo.argc = argc - PhpOptind
					core.SG__().RequestInfo.argv = &argv[PhpOptind]
				}
				if no_headers != 0 {
					core.SG__().headers_sent = 1
					core.SG__().RequestInfo.no_headers = 1
				}

				/* all remaining arguments are part of the query string
				 * this section of code concatenates all remaining arguments
				 * into a single string, separating args with a &
				 * this allows command lines like:
				 *
				 *  test.php v1=test v2=hello+world!
				 *  test.php "v1=test&v2=hello world!"
				 *  test.php v1=test "v2=hello world!"
				 */

				if !(core.SG__().RequestInfo.query_string) && argc > PhpOptind {
					var slen int = strlen(core.PG__().arg_separator.input)
					len_ = 0
					for i = PhpOptind; i < argc; i++ {
						if i < argc-1 {
							len_ += strlen(argv[i]) + slen
						} else {
							len_ += strlen(argv[i])
						}
					}
					len_ += 2
					s = zend.Malloc(len_)
					*s = '0'
					for i = PhpOptind; i < argc; i++ {
						strlcat(s, argv[i], len_)
						if i < argc-1 {
							strlcat(s, core.PG__().arg_separator.input, len_)
						}
					}
					core.SG__().RequestInfo.query_string = s
					free_query_string = 1
				}

				/* all remaining arguments are part of the query string
				 * this section of code concatenates all remaining arguments
				 * into a single string, separating args with a &
				 * this allows command lines like:
				 *
				 *  test.php v1=test v2=hello+world!
				 *  test.php "v1=test&v2=hello world!"
				 *  test.php v1=test "v2=hello world!"
				 */

			}

			/*
			   we never take stdin if we're (f)cgi, always
			   rely on the web server giving us the info
			   we need in the environment.
			*/

			if core.SG__().RequestInfo.path_translated || cgi != 0 || fastcgi != 0 {
				zend.ZendStreamInitFilename(&file_handle, core.SG__().RequestInfo.path_translated)
			} else {
				zend.ZendStreamInitFp(&file_handle, stdin, "Standard input code")
			}

			/* request startup only after we've done all we can to
			 * get path_translated */

			if core.PhpRequestStartup() == types.FAILURE {
				if fastcgi != 0 {
					core.FcgiFinishRequest(request, 1)
				}
				core.SG__().server_context = nil
				core.PhpModuleShutdown()
				return types.FAILURE
			}
			if no_headers != 0 {
				core.SG__().headers_sent = 1
				core.SG__().RequestInfo.no_headers = 1
			}

			/*
			   at this point path_translated will be set if:
			   1. we are running from shell and got filename was there
			   2. we are running as cgi or fastcgi
			*/

			if cgi != 0 || fastcgi != 0 || core.SG__().RequestInfo.path_translated {
				if core.PhpFopenPrimaryScript(&file_handle) == types.FAILURE {
					faults.Try(func() {
						if errno == EACCES {
							core.SG__().sapi_headers.http_response_code = 403
							core.PUTS("Access denied.\n")
						} else {
							core.SG__().sapi_headers.http_response_code = 404
							core.PUTS("No input file specified.\n")
						}
					})

					/* we want to serve more requests if this is fastcgi
					 * so cleanup and continue, request shutdown is
					 * handled later */

					if fastcgi != 0 {
						goto fastcgi_request_done
					}
					if core.SG__().RequestInfo.path_translated {
						zend.Efree(core.SG__().RequestInfo.path_translated)
						core.SG__().RequestInfo.path_translated = nil
					}
					if free_query_string != 0 && core.SG__().RequestInfo.query_string {
						zend.Free(core.SG__().RequestInfo.query_string)
						core.SG__().RequestInfo.query_string = nil
					}
					core.PhpRequestShutdown(any(0))
					core.SG__().server_context = nil
					core.PhpModuleShutdown()
					app.SapiShutdown()
					zend.Free(bindpath)
					return types.FAILURE
				}
			}
			if CGIG(check_shebang_line) {
				zend.CG__().SetSkipShebang(1)
			}
			switch behavior {
			case PHP_MODE_STANDARD:
				core.PhpExecuteScript(&file_handle)
				break
			case PHP_MODE_LINT:
				core.PG__().during_request_startup = 0
				exit_status = core.PhpLintScript(&file_handle)
				if exit_status == types.SUCCESS {
					zend.ZendPrintf("No syntax errors detected in %s\n", file_handle.GetFilename())
				} else {
					zend.ZendPrintf("Errors parsing %s\n", file_handle.GetFilename())
				}
				break
			case PHP_MODE_STRIP:
				if zend.OpenFileForScanning(&file_handle) == types.SUCCESS {
					zend.ZendStrip()
					zend.ZendFileHandleDtor(&file_handle)
					core.PhpOutputTeardown()
				}
				return types.SUCCESS
				break
			case PHP_MODE_HIGHLIGHT:
				var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
				if zend.OpenFileForScanning(&file_handle) == types.SUCCESS {
					standard.PhpGetHighlight(&syntax_highlighter_ini)
					zend.ZendHighlight(&syntax_highlighter_ini)
					if fastcgi != 0 {
						goto fastcgi_request_done
					}
					zend.ZendFileHandleDtor(&file_handle)
					core.PhpOutputTeardown()
				}
				return types.SUCCESS
				break
			}
		fastcgi_request_done:
			if core.SG__().RequestInfo.path_translated {
				zend.Efree(core.SG__().RequestInfo.path_translated)
				core.SG__().RequestInfo.path_translated = nil
			}
			core.PhpRequestShutdown(any(0))
			if exit_status == 0 {
				exit_status = zend.EG__().GetExitStatus()
			}
			if free_query_string != 0 && core.SG__().RequestInfo.query_string {
				zend.Free(core.SG__().RequestInfo.query_string)
				core.SG__().RequestInfo.query_string = nil
			}
			if fastcgi == 0 {
				if benchmark != 0 {
					if warmup_repeats != 0 {
						warmup_repeats--
						if warmup_repeats == 0 {
							gettimeofday(&start, nil)
						}
						continue
					} else {
						repeats--
						if repeats > 0 {
							script_file = nil
							PhpOptind = orig_optind
							PhpOptarg = orig_optarg
							continue
						}
					}
				}
				break
			}

			/* only fastcgi will get here */

			requests++
			if max_requests != 0 && requests == max_requests {
				core.FcgiFinishRequest(request, 1)
				zend.Free(bindpath)
				if max_requests != 1 {

					/* no need to return exit_status of the last request */

					exit_status = 0

					/* no need to return exit_status of the last request */

				}
				break
			}
		}
		if request != nil {
			core.FcgiDestroyRequest(request)
		}
		core.FcgiShutdown()
		if CgiModule.GetPhpIniPathOverride() != nil {
			zend.Free(CgiModule.GetPhpIniPathOverride())
		}
		if CgiModule.GetIniEntries() != nil {
			zend.Free(CgiModule.GetIniEntries())
		}
	}, func() {
		exit_status = 255
	})
out:
	if benchmark != 0 {
		var sec int
		var usec int
		gettimeofday(&end, nil)
		sec = int(end.tv_sec - start.tv_sec)
		if end.tv_usec >= start.tv_usec {
			usec = int(end.tv_usec - start.tv_usec)
		} else {
			sec -= 1
			usec = int(end.tv_usec + 1000000 - start.tv_usec)
		}
		log.Printf("\nElapsed time: %d.%06d sec\n", sec, usec)
	}
parent_out:
	core.SG__().server_context = nil
	core.PhpModuleShutdown()
	app.SapiShutdown()
	return exit_status
}
