// <<generate>>

package standard

// Source: <ext/standard/net.c>

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
   | Authors: Sara Golemon <pollita@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* {{{ proto array|false net_get_interfaces()
Returns an array in the form:
array(
  'ifacename' => array(
    'description' => 'Awesome interface', // Win32 only
    'mac' => '00:11:22:33:44:55',         // Win32 only
    'mtu' => 1234,                        // Win32 only
    'unicast' => array(
      0 => array(
        'family' => 2,                    // e.g. AF_INET, AF_INET6, AF_PACKET
        'address' => '127.0.0.1',
        'netmnask' => '255.0.0.0',
        'broadcast' => '127.255.255.255', // POSIX only
        'ptp' => '127.0.0.2',             // POSIX only
      ), // etc...
    ),
  ), // etc...
)
*/
