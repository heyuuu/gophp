// <<generate>>

package standard

import (
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/core"
	"sik/core/streams"
	"sik/sapi/cli"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func _phpArrayToEnvp(environment *types.Zval, is_persistent int) PhpProcessEnvT {
	var element *types.Zval
	var env PhpProcessEnvT
	var key *types.String
	var str *types.String
	var ep **byte
	var p *byte
	var cnt int
	var sizeenv int = 0
	var env_hash *types.Array
	memset(&env, 0, b.SizeOf("env"))
	if environment == nil {
		return env
	}
	cnt = types.Z_ARRVAL_P(environment).GetNNumOfElements()
	if cnt < 1 {
		env.SetEnvarray((**byte)(zend.Pecalloc(1, b.SizeOf("char *"), is_persistent)))
		env.SetEnvp((*byte)(zend.Pecalloc(4, 1, is_persistent)))
		return env
	}
	zend.ALLOC_HASHTABLE(env_hash)
	env_hash = types.MakeArrayEx(cnt, nil, 0)

	/* first, we have to get the size of all the elements in the hash */

	var __ht *types.Array = environment.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		element = _z
		str = zend.ZvalGetString(element)
		if str.GetLen() == 0 {
			types.ZendStringReleaseEx(str, 0)
			continue
		}
		sizeenv += str.GetLen() + 1
		if key != nil && key.GetLen() != 0 {
			sizeenv += key.GetLen() + 1
			types.ZendHashAddPtr(env_hash, key, str)
		} else {
			types.ZendHashNextIndexInsertPtr(env_hash, str)
		}
	}
	env.SetEnvarray((**byte)(zend.Pecalloc(cnt+1, b.SizeOf("char *"), is_persistent)))
	ep = env.GetEnvarray()
	env.SetEnvp((*byte)(zend.Pecalloc(sizeenv+4, 1, is_persistent)))
	p = env.GetEnvp()
	var __ht__1 *types.Array = env_hash
	for _, _p := range __ht__1.foreachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		str = _z.GetPtr()
		*ep = p
		ep++
		if key != nil {
			memcpy(p, key.GetVal(), key.GetLen())
			p += key.GetLen()
			b.PostInc(&(*p)) = '='
		}
		memcpy(p, str.GetVal(), str.GetLen())
		p += str.GetLen()
		b.PostInc(&(*p)) = '0'
		types.ZendStringReleaseEx(str, 0)
	}
	b.Assert(uint32(p-env.GetEnvp()) <= sizeenv)
	env_hash.Destroy()
	zend.FREE_HASHTABLE(env_hash)
	return env
}
func _phpFreeEnvp(env PhpProcessEnvT, is_persistent int) {
	if env.GetEnvarray() != nil {
		zend.Pefree(env.GetEnvarray(), is_persistent)
	}
	if env.GetEnvp() != nil {
		zend.Pefree(env.GetEnvp(), is_persistent)
	}
}
func ProcOpenRsrcDtor(rsrc *types.ZendResource) {
	var proc *PhpProcessHandle = (*PhpProcessHandle)(rsrc.GetPtr())
	var i int
	var wstatus int
	var waitpid_options int = 0
	var wait_pid pid_t

	/* Close all handles to avoid a deadlock */

	for i = 0; i < proc.GetNpipes(); i++ {
		if proc.GetPipes()[i] != 0 {
			proc.GetPipes()[i].DelRefcount()
			zend.ZendListClose(proc.GetPipes()[i])
			proc.GetPipes()[i] = 0
		}
	}
	if !(FG(pclose_wait)) {
		waitpid_options = WNOHANG
	}
	for {
		wait_pid = waitpid(proc.GetChild(), &wstatus, waitpid_options)
		if !(wait_pid == -1 && errno == EINTR) {
			break
		}
	}
	if wait_pid <= 0 {
		FG(pclose_ret) = -1
	} else {
		if WIFEXITED(wstatus) {
			wstatus = WEXITSTATUS(wstatus)
		}
		FG(pclose_ret) = wstatus
	}
	_phpFreeEnvp(proc.GetEnv(), proc.GetIsPersistent())
	zend.Pefree(proc.GetPipes(), proc.GetIsPersistent())
	zend.Pefree(proc.GetCommand(), proc.GetIsPersistent())
	zend.Pefree(proc, proc.GetIsPersistent())
}
func ZmStartupProcOpen(type_ int, module_number int) int {
	LeProcOpen = zend.ZendRegisterListDestructorsEx(ProcOpenRsrcDtor, nil, "process", module_number)
	return types.SUCCESS
}
func ZifProcTerminate(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zproc *types.Zval
	var proc *PhpProcessHandle
	var sig_no zend.ZendLong = SIGTERM
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			zproc = fp.ParseResource()
			fp.StartOptional()
			sig_no = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if b.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.GetRes(), "process", LeProcOpen))) == nil {
		return_value.SetFalse()
		return
	}
	if kill(proc.GetChild(), sig_no) == 0 {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifProcClose(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zproc *types.Zval
	var proc *PhpProcessHandle
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zproc = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if b.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.GetRes(), "process", LeProcOpen))) == nil {
		return_value.SetFalse()
		return
	}
	FG(pclose_wait) = 1
	zend.ZendListClose(zproc.GetRes())
	FG(pclose_wait) = 0
	return_value.SetLong(FG(pclose_ret))
	return
}
func ZifProcGetStatus(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zproc *types.Zval
	var proc *PhpProcessHandle
	var wstatus int
	var wait_pid pid_t
	var running int = 1
	var signaled int = 0
	var stopped int = 0
	var exitcode int = -1
	var termsig int = 0
	var stopsig int = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zproc = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if b.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.GetRes(), "process", LeProcOpen))) == nil {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	zend.AddAssocString(return_value, "command", proc.GetCommand())
	zend.AddAssocLong(return_value, "pid", zend.ZendLong(proc.GetChild()))
	errno = 0
	wait_pid = waitpid(proc.GetChild(), &wstatus, WNOHANG|WUNTRACED)
	if wait_pid == proc.GetChild() {
		if WIFEXITED(wstatus) {
			running = 0
			exitcode = WEXITSTATUS(wstatus)
		}
		if WIFSIGNALED(wstatus) {
			running = 0
			signaled = 1
			termsig = WTERMSIG(wstatus)
		}
		if WIFSTOPPED(wstatus) {
			stopped = 1
			stopsig = WSTOPSIG(wstatus)
		}
	} else if wait_pid == -1 {
		running = 0
	}
	zend.AddAssocBool(return_value, "running", running)
	zend.AddAssocBool(return_value, "signaled", signaled)
	zend.AddAssocBool(return_value, "stopped", stopped)
	zend.AddAssocLong(return_value, "exitcode", exitcode)
	zend.AddAssocLong(return_value, "termsig", termsig)
	zend.AddAssocLong(return_value, "stopsig", stopsig)
}
func CloseDescriptor(fd PhpFileDescriptorT) __auto__ { return close(fd) }
func GetValidArgString(zv *types.Zval, elem_num int) *types.String {
	var str *types.String = zend.ZvalGetString(zv)
	if str == nil {
		return nil
	}
	if strlen(str.GetVal()) != str.GetLen() {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Command array element %d contains a null byte", elem_num)
		types.ZendStringRelease(str)
		return nil
	}
	return str
}
func ZifProcOpen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var command_zv *types.Zval
	var command *byte = nil
	var cwd *byte = nil
	var cwd_len int = 0
	var descriptorspec *types.Zval
	var pipes *types.Zval
	var environment *types.Zval = nil
	var other_options *types.Zval = nil
	var env PhpProcessEnvT
	var ndesc int = 0
	var i int
	var descitem *types.Zval = nil
	var str_index *types.String
	var nindex zend.ZendUlong
	var descriptors *PhpProcOpenDescriptorItem = nil
	var ndescriptors_array int
	var argv **byte = nil
	var child PhpProcessIdT
	var proc *PhpProcessHandle
	var is_persistent int = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 6, 0)
			command_zv = fp.ParseZval()
			descriptorspec = fp.ParseArray()
			pipes = fp.ParseZval()
			fp.StartOptional()
			cwd, cwd_len = fp.ParseStringEx(true, false)
			environment = fp.ParseArrayEx(true, false)
			other_options = fp.ParseArrayEx(true, false)
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	memset(&env, 0, b.SizeOf("env"))
	if command_zv.IsType(types.IS_ARRAY) {
		var arg_zv *types.Zval
		var num_elems uint32 = types.Z_ARRVAL_P(command_zv).GetNNumOfElements()
		if num_elems == 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Command array must have at least one element")
			return_value.SetFalse()
			return
		}
		argv = zend.SafeEmalloc(b.SizeOf("char *"), num_elems+1, 0)
		i = 0
		var __ht *types.Array = command_zv.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			arg_zv = _z
			var arg_str *types.String = GetValidArgString(arg_zv, i+1)
			if arg_str == nil {
				argv[i] = nil
				goto exit_fail
			}
			if i == 0 {
				command = zend.Pestrdup(arg_str.GetVal(), is_persistent)
			}
			argv[b.PostInc(&i)] = zend.Estrdup(arg_str.GetVal())
			types.ZendStringRelease(arg_str)
		}
		argv[i] = nil

		/* As the array is non-empty, we should have found a command. */

		b.Assert(command != nil)

		/* As the array is non-empty, we should have found a command. */

	} else {
		zend.ConvertToString(command_zv)
		command = zend.Pestrdup(command_zv.GetStr().GetVal(), is_persistent)
	}
	if environment != nil {
		env = _phpArrayToEnvp(environment, is_persistent)
	}
	ndescriptors_array = types.Z_ARRVAL_P(descriptorspec).GetNNumOfElements()
	descriptors = zend.SafeEmalloc(b.SizeOf("struct php_proc_open_descriptor_item"), ndescriptors_array, 0)
	memset(descriptors, 0, b.SizeOf("struct php_proc_open_descriptor_item")*ndescriptors_array)

	/* walk the descriptor spec and set up files/pipes */

	var __ht *types.Array = descriptorspec.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		nindex = _p.GetH()
		str_index = _p.GetKey()
		descitem = _z
		var ztype *types.Zval
		if str_index != nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "descriptor spec must be an integer indexed array")
			goto exit_fail
		}
		descriptors[ndesc].SetIndex(int(nindex))
		if descitem.IsType(types.IS_RESOURCE) {

			/* should be a stream - try and dup the descriptor */

			var stream *core.PhpStream
			var fd core.PhpSocketT
			core.PhpStreamFromZval(stream, descitem)
			if types.FAILURE == core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD, (*any)(&fd), core.REPORT_ERRORS) {
				goto exit_fail
			}
			descriptors[ndesc].SetChildend(dup(fd))
			if descriptors[ndesc].GetChildend() < 0 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "unable to dup File-Handle for descriptor "+zend.ZEND_ULONG_FMT+" - %s", nindex, strerror(errno))
				goto exit_fail
			}
			descriptors[ndesc].SetMode(DESC_FILE)
		} else if descitem.GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Descriptor item must be either an array or a File-Handle")
			goto exit_fail
		} else {
			if b.Assign(&ztype, descitem.GetArr().IndexFindH(0)) != nil {
				if zend.TryConvertToString(ztype) == 0 {
					goto exit_fail
				}
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Missing handle qualifier in array")
				goto exit_fail
			}
			if strcmp(ztype.GetStr().GetVal(), "pipe") == 0 {
				var newpipe []PhpFileDescriptorT
				var zmode *types.Zval
				if b.Assign(&zmode, descitem.GetArr().IndexFindH(1)) != nil {
					if zend.TryConvertToString(zmode) == 0 {
						goto exit_fail
					}
				} else {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Missing mode parameter for 'pipe'")
					goto exit_fail
				}
				descriptors[ndesc].SetMode(DESC_PIPE)
				if 0 != pipe(newpipe) {
					core.PhpErrorDocref(nil, faults.E_WARNING, "unable to create pipe %s", strerror(errno))
					goto exit_fail
				}
				if strncmp(zmode.GetStr().GetVal(), "w", 1) != 0 {
					descriptors[ndesc].SetParentend(newpipe[1])
					descriptors[ndesc].SetChildend(newpipe[0])
					descriptors[ndesc].SetMode(descriptors[ndesc].GetMode() | DESC_PARENT_MODE_WRITE)
				} else {
					descriptors[ndesc].SetParentend(newpipe[0])
					descriptors[ndesc].SetChildend(newpipe[1])
				}
				if (descriptors[ndesc].GetMode() & DESC_PARENT_MODE_WRITE) != 0 {
					descriptors[ndesc].SetModeFlags(O_WRONLY)
				} else {
					descriptors[ndesc].SetModeFlags(O_RDONLY)
				}
			} else if strcmp(ztype.GetStr().GetVal(), "file") == 0 {
				var zfile *types.Zval
				var zmode *types.Zval
				var fd core.PhpSocketT
				var stream *core.PhpStream
				descriptors[ndesc].SetMode(DESC_FILE)
				if b.Assign(&zfile, descitem.GetArr().IndexFindH(1)) != nil {
					if zend.TryConvertToString(zfile) == 0 {
						goto exit_fail
					}
				} else {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Missing file name parameter for 'file'")
					goto exit_fail
				}
				if b.Assign(&zmode, descitem.GetArr().IndexFindH(2)) != nil {
					if zend.TryConvertToString(zmode) == 0 {
						goto exit_fail
					}
				} else {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Missing mode parameter for 'file'")
					goto exit_fail
				}

				/* try a wrapper */

				stream = core.PhpStreamOpenWrapper(zfile.GetStr().GetVal(), zmode.GetStr().GetVal(), core.REPORT_ERRORS|core.STREAM_WILL_CAST, nil)

				/* force into an fd */

				if stream == nil || types.FAILURE == core.PhpStreamCast(stream, core.PHP_STREAM_CAST_RELEASE|core.PHP_STREAM_AS_FD, (*any)(&fd), core.REPORT_ERRORS) {
					goto exit_fail
				}
				descriptors[ndesc].SetChildend(fd)
			} else if strcmp(ztype.GetStr().GetVal(), "redirect") == 0 {
				var ztarget *types.Zval = types.ZendHashIndexFindDeref(descitem.GetArr(), 1)
				var target *PhpProcOpenDescriptorItem = nil
				var childend PhpFileDescriptorT
				if ztarget == nil {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Missing redirection target")
					goto exit_fail
				}
				if ztarget.GetType() != types.IS_LONG {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Redirection target must be an integer")
					goto exit_fail
				}
				for i = 0; i < ndesc; i++ {
					if descriptors[i].GetIndex() == ztarget.GetLval() {
						target = &descriptors[i]
						break
					}
				}
				if target != nil {
					childend = target.GetChildend()
				} else {
					if ztarget.GetLval() < 0 || ztarget.GetLval() > 2 {
						core.PhpErrorDocref(nil, faults.E_WARNING, "Redirection target "+zend.ZEND_LONG_FMT+" not found", ztarget.GetLval())
						goto exit_fail
					}

					/* Support referring to a stdin/stdout/stderr pipe adopted from the parent,
					 * which happens whenever an explicit override is not provided. */

					childend = ztarget.GetLval()

					/* Support referring to a stdin/stdout/stderr pipe adopted from the parent,
					 * which happens whenever an explicit override is not provided. */

				}
				descriptors[ndesc].SetChildend(dup(childend))
				if descriptors[ndesc].GetChildend() < 0 {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Failed to dup() for descriptor "+zend.ZEND_LONG_FMT+" - %s", nindex, strerror(errno))
					goto exit_fail
				}
				descriptors[ndesc].SetMode(DESC_REDIRECT)
			} else if strcmp(ztype.GetStr().GetVal(), "null") == 0 {
				descriptors[ndesc].SetChildend(open("/dev/null", O_RDWR))
				if descriptors[ndesc].GetChildend() < 0 {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Failed to open /dev/null - %s", strerror(errno))
					goto exit_fail
				}
				descriptors[ndesc].SetMode(DESC_FILE)
			} else if strcmp(ztype.GetStr().GetVal(), "pty") == 0 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "pty pseudo terminal not supported on this system")
				goto exit_fail
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "%s is not a valid descriptor spec/mode", ztype.GetStr().GetVal())
				goto exit_fail
			}
		}
		ndesc++
	}

	/* the unix way */

	child = fork()
	if child == 0 {

		/* this is the child process */

		/* close those descriptors that we just opened for the parent stuff,
		 * dup new descriptors into required descriptors and close the original
		 * cruft */

		for i = 0; i < ndesc; i++ {
			switch descriptors[i].GetMode() & ^DESC_PARENT_MODE_WRITE {
			case DESC_PIPE:
				close(descriptors[i].GetParentend())
			}
			if dup2(descriptors[i].GetChildend(), descriptors[i].GetIndex()) < 0 {
				r.Perror("dup2")
			}
			if descriptors[i].GetChildend() != descriptors[i].GetIndex() {
				close(descriptors[i].GetChildend())
			}
		}
		if cwd != nil {
			core.PhpIgnoreValue(chdir(cwd))
		}
		if argv != nil {

			/* execvpe() is non-portable, use environ instead. */

			if env.GetEnvarray() != nil {
				cli.Environ = env.GetEnvarray()
			}
			execvp(command, argv)
		} else {
			if env.GetEnvarray() != nil {
				execle("/bin/sh", "sh", "-c", command, nil, env.GetEnvarray())
			} else {
				execl("/bin/sh", "sh", "-c", command, nil)
			}
		}
		_exit(127)
	} else if child < 0 {

		/* failed to fork() */

		for i = 0; i < ndesc; i++ {
			close(descriptors[i].GetChildend())
			if descriptors[i].GetParentend() != 0 {
				close(descriptors[i].GetParentend())
			}
		}
		core.PhpErrorDocref(nil, faults.E_WARNING, "fork failed - %s", strerror(errno))
		goto exit_fail
	}

	/* we forked/spawned and this is the parent */

	pipes = zend.ZendTryArrayInit(pipes)
	if pipes == nil {
		goto exit_fail
	}
	proc = (*PhpProcessHandle)(zend.Pemalloc(b.SizeOf("struct php_process_handle"), is_persistent))
	proc.SetIsPersistent(is_persistent)
	proc.SetCommand(command)
	proc.SetPipes(zend.Pemalloc(b.SizeOf("zend_resource *")*ndesc, is_persistent))
	proc.SetNpipes(ndesc)
	proc.SetChild(child)
	proc.SetEnv(env)

	/* clean up all the child ends and then open streams on the parent
	 * ends, where appropriate */

	for i = 0; i < ndesc; i++ {
		var mode_string *byte = nil
		var stream *core.PhpStream = nil
		CloseDescriptor(descriptors[i].GetChildend())
		switch descriptors[i].GetMode() & ^DESC_PARENT_MODE_WRITE {
		case DESC_PIPE:
			switch descriptors[i].GetModeFlags() {
			case O_WRONLY:
				mode_string = "w"
			case O_RDONLY:
				mode_string = "r"
			case O_RDWR:
				mode_string = "r+"
			}
			stream = streams.PhpStreamFopenFromFd(descriptors[i].GetParentend(), mode_string, nil)
			if stream != nil {
				var retfp types.Zval

				/* nasty hack; don't copy it */

				stream.AddFlags(core.PHP_STREAM_FLAG_NO_SEEK)
				core.PhpStreamToZval(stream, &retfp)
				zend.AddIndexZval(pipes, descriptors[i].GetIndex(), &retfp)
				proc.GetPipes()[i] = retfp.GetRes()
				retfp.AddRefcount()
			}
		default:
			proc.GetPipes()[i] = nil
		}
	}
	if argv != nil {
		var arg **byte = argv
		for (*arg) != nil {
			zend.Efree(*arg)
			arg++
		}
		zend.Efree(argv)
	}
	zend.Efree(descriptors)
	return_value.SetResource(zend.ZendRegisterResource(proc, LeProcOpen))
	return
exit_fail:
	if descriptors != nil {
		zend.Efree(descriptors)
	}
	_phpFreeEnvp(env, is_persistent)
	if command != nil {
		zend.Pefree(command, is_persistent)
	}
	if argv != nil {
		var arg **byte = argv
		for (*arg) != nil {
			zend.Efree(*arg)
			arg++
		}
		zend.Efree(argv)
	}
	return_value.SetFalse()
	return
}
