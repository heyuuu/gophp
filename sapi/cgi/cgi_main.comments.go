// <<generate>>

package cgi

// Source: <sapi/cgi/cgi_main.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                       |
   |          Stig Bakken <ssb@php.net>                                   |
   |          Zeev Suraski <zeev@php.net>                                 |
   | FastCGI: Ben Mansell <php@slimyhorror.com>                           |
   |          Shane Caraveo <shane@caraveo.com>                           |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* XXX this will need to change later when threaded fastcgi is implemented.  shane */

/* these globals used for forking children on unix systems */

/**
 * Set to non-zero if we are the parent process
 */

/* Did parent received exit signals SIG_TERM/SIG_INT/SIG_QUIT */

/* Is Parent waiting for children to exit */

/**
 * Process group
 */

/* {{{ user_config_cache
 *
 * Key for each cache entry is dirname(PATH_TRANSLATED).
 *
 * NOTE: Each cache entry config_hash contains the combination from all user ini files found in
 *       the path starting from doc_root through to dirname(PATH_TRANSLATED).  There is no point
 *       storing per-file entries as it would not be possible to detect added / deleted entries
 *       between separate files.
 */

/* }}} */

/* {{{ php_cgi_ini_activate_user_config
 */

/* }}} */

/* {{{ sapi_module_struct cgi_sapi_module
 */

/* }}} */

/* }}} */

/* {{{ php_cgi_usage
 */

/* }}} */

/* }}} */

/* {{{ init_request_info

initializes request_info structure

specificly in this section we handle proper translations
for:

PATH_INFO
  derived from the portion of the URI path following
  the script name but preceding any query data
  may be empty

PATH_TRANSLATED
  derived by taking any path-info component of the
  request URI and performing any virtual-to-physical
  translation appropriate to map it onto the server's
  document repository structure

  empty if PATH_INFO is empty

  The env var PATH_TRANSLATED **IS DIFFERENT** than the
  request_info.path_translated variable, the latter should
  match SCRIPT_FILENAME instead.

SCRIPT_NAME
  set to a URL path that could identify the CGI script
  rather than the interpreter.  PHP_SELF is set to this

REQUEST_URI
  uri section following the domain:port part of a URI

SCRIPT_FILENAME
  The virtual-to-physical translation of SCRIPT_NAME (as per
  PATH_TRANSLATED)

These settings are documented at
http://cgi-spec.golux.com/


Based on the following URL request:

http://localhost/info.php/test?a=b

should produce, which btw is the same as if
we were running under mod_cgi on apache (ie. not
using ScriptAlias directives):

PATH_INFO=/test
PATH_TRANSLATED=/docroot/test
SCRIPT_NAME=/info.php
REQUEST_URI=/info.php/test?a=b
SCRIPT_FILENAME=/docroot/info.php
QUERY_STRING=a=b

but what we get is (cgi/mod_fastcgi under apache):

PATH_INFO=/info.php/test
PATH_TRANSLATED=/docroot/info.php/test
SCRIPT_NAME=/php/php-cgi  (from the Action setting I suppose)
REQUEST_URI=/info.php/test?a=b
SCRIPT_FILENAME=/path/to/php/bin/php-cgi  (Action setting translated)
QUERY_STRING=a=b

Comments in the code below refer to using the above URL in a request

*/

/* }}} */

/**
 * Clean up child processes upon exit
 */

/* {{{ php_cgi_globals_ctor
 */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ main
 */

/* }}} */
