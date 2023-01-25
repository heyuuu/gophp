// <<generate>>

package streams

// Source: <main/streams/userspace.c>

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
   | Authors: Wez Furlong <wez@thebrainroom.com>                          |
   |          Sara Golemon <pollita@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* names of methods */

/* {{{ class should have methods like these:

   function stream_open($path, $mode, $options, &$opened_path)
   {
         return true/false;
   }

   function stream_read($count)
   {
          return false on error;
       else return string;
   }

   function stream_write($data)
   {
          return false on error;
       else return count written;
   }

   function stream_close()
   {
   }

   function stream_flush()
   {
       return true/false;
   }

   function stream_seek($offset, $whence)
   {
       return true/false;
   }

   function stream_tell()
   {
       return (int)$position;
   }

   function stream_eof()
   {
       return true/false;
   }

   function stream_stat()
   {
       return array( just like that returned by fstat() );
   }

   function stream_cast($castas)
   {
       if ($castas == STREAM_CAST_FOR_SELECT) {
           return $this->underlying_stream;
       }
       return false;
   }

   function stream_set_option($option, $arg1, $arg2)
   {
       switch($option) {
       case STREAM_OPTION_BLOCKING:
           $blocking = $arg1;
           ...
       case STREAM_OPTION_READ_TIMEOUT:
           $sec = $arg1;
           $usec = $arg2;
           ...
       case STREAM_OPTION_WRITE_BUFFER:
           $mode = $arg1;
           $size = $arg2;
           ...
       default:
           return false;
       }
   }

   function url_stat(string $url, int $flags)
   {
       return array( just like that returned by stat() );
   }

   function unlink(string $url)
   {
       return true / false;
   }

   function rename(string $from, string $to)
   {
       return true / false;
   }

   function mkdir($dir, $mode, $options)
   {
       return true / false;
   }

   function rmdir($dir, $options)
   {
       return true / false;
   }

   function dir_opendir(string $url, int $options)
   {
       return true / false;
   }

   function dir_readdir()
   {
       return string next filename in dir ;
   }

   function dir_closedir()
   {
       release dir related resources;
   }

   function dir_rewinddir()
   {
       reset to start of dir list;
   }

   function stream_lock($operation)
   {
       return true / false;
   }

    function stream_truncate($new_size)
   {
       return true / false;
   }

   }}} **/

/* {{{ proto bool stream_wrapper_register(string protocol, string classname[, int flags])
   Registers a custom URL protocol handler class */

/* }}} */

/* }}} */

/* }}} */

/* parse the return value from one of the stat functions and store the
 * relevant fields into the statbuf provided */
