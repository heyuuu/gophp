// <<generate>>

package standard

import (
	"sik/core"
)

// Source: <ext/standard/credits.h>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define CREDITS_H

// Source: <ext/standard/credits.c>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "info.h"

// # include "SAPI.h"

// #define CREDIT_LINE(module,authors) php_info_print_table_row ( 2 , module , authors )

func PhpPrintCredits(flag int) {
	if core.sapi_module.phpinfo_as_text == 0 && (flag&1<<5) != 0 {
		PhpPrintInfoHtmlhead()
	}
	if core.sapi_module.phpinfo_as_text == 0 {
		var __str *byte = "<h1>PHP Credits</h1>\n"
		core.PhpOutputWrite(__str, strlen(__str))
	} else {
		var __str *byte = "PHP Credits\n"
		core.PhpOutputWrite(__str, strlen(__str))
	}
	if (flag & 1 << 0) != 0 {

		/* Group */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(1, "PHP Group")
		PhpInfoPrintTableRow(1, "Thies C. Arntzen, Stig Bakken, Shane Caraveo, Andi Gutmans, Rasmus Lerdorf, Sam Ruby, Sascha Schumann, Zeev Suraski, Jim Winstead, Andrei Zmievski")
		PhpInfoPrintTableEnd()
	}
	if (flag & 1 << 1) != 0 {

		/* Design & Concept */

		PhpInfoPrintTableStart()
		if core.sapi_module.phpinfo_as_text == 0 {
			PhpInfoPrintTableHeader(1, "Language Design &amp; Concept")
		} else {
			PhpInfoPrintTableHeader(1, "Language Design & Concept")
		}
		PhpInfoPrintTableRow(1, "Andi Gutmans, Rasmus Lerdorf, Zeev Suraski, Marcus Boerger")
		PhpInfoPrintTableEnd()

		/* PHP Language */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "PHP Authors")
		PhpInfoPrintTableHeader(2, "Contribution", "Authors")
		PhpInfoPrintTableRow(2, "Zend Scripting Language Engine", "Andi Gutmans, Zeev Suraski, Stanislav Malyshev, Marcus Boerger, Dmitry Stogov, Xinchen Hui, Nikita Popov")
		PhpInfoPrintTableRow(2, "Extension Module API", "Andi Gutmans, Zeev Suraski, Andrei Zmievski")
		PhpInfoPrintTableRow(2, "UNIX Build and Modularization", "Stig Bakken, Sascha Schumann, Jani Taskinen, Peter Kokot")
		PhpInfoPrintTableRow(2, "Windows Support", "Shane Caraveo, Zeev Suraski, Wez Furlong, Pierre-Alain Joye, Anatol Belski, Kalle Sommer Nielsen")
		PhpInfoPrintTableRow(2, "Server API (SAPI) Abstraction Layer", "Andi Gutmans, Shane Caraveo, Zeev Suraski")
		PhpInfoPrintTableRow(2, "Streams Abstraction Layer", "Wez Furlong, Sara Golemon")
		PhpInfoPrintTableRow(2, "PHP Data Objects Layer", "Wez Furlong, Marcus Boerger, Sterling Hughes, George Schlossnagle, Ilia Alshanetsky")
		PhpInfoPrintTableRow(2, "Output Handler", "Zeev Suraski, Thies C. Arntzen, Marcus Boerger, Michael Wallner")
		PhpInfoPrintTableRow(2, "Consistent 64 bit support", "Anthony Ferrara, Anatol Belski")
		PhpInfoPrintTableEnd()
	}
	if (flag & 1 << 2) != 0 {

		/* SAPI Modules */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "SAPI Modules")
		PhpInfoPrintTableHeader(2, "Contribution", "Authors")

		// # include "credits_sapi.h"

		PhpInfoPrintTableEnd()

		// # include "credits_sapi.h"

	}
	if (flag & 1 << 3) != 0 {

		/* Modules */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "Module Authors")
		PhpInfoPrintTableHeader(2, "Module", "Authors")

		// # include "credits_ext.h"

		PhpInfoPrintTableEnd()

		// # include "credits_ext.h"

	}
	if (flag & 1 << 4) != 0 {
		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "PHP Documentation")
		PhpInfoPrintTableRow(2, "Authors", "Mehdi Achour, Friedhelm Betz, Antony Dovgal, Nuno Lopes, Hannes Magnusson, Philip Olson, Georg Richter, Damien Seguy, Jakub Vrana, Adam Harvey")
		PhpInfoPrintTableRow(2, "Editor", "Peter Cowburn")
		PhpInfoPrintTableRow(2, "User Note Maintainers", "Daniel P. Brown, Thiago Henrique Pojda")
		PhpInfoPrintTableRow(2, "Other Contributors", "Previously active authors, editors and other contributors are listed in the manual.")
		PhpInfoPrintTableEnd()
	}
	if (flag & 1 << 6) != 0 {
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(1, "PHP Quality Assurance Team")
		PhpInfoPrintTableRow(1, "Ilia Alshanetsky, Joerg Behrens, Antony Dovgal, Stefan Esser, Moriyoshi Koizumi, Magnus Maatta, Sebastian Nohn, Derick Rethans, Melvyn Sopacua, Pierre-Alain Joye, Dmitry Stogov, Felipe Pena, David Soria Parra, Stanislav Malyshev, Julien Pauli, Stephen Zarkos, Anatol Belski, Remi Collet, Ferenc Kovacs")
		PhpInfoPrintTableEnd()
	}
	if (flag & 1 << 7) != 0 {

		/* Websites and infrastructure */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "Websites and Infrastructure team")

		/* www., wiki., windows., master., and others, I guess pecl. too? */

		PhpInfoPrintTableRow(2, "PHP Websites Team", "Rasmus Lerdorf, Hannes Magnusson, Philip Olson, Lukas Kahwe Smith, Pierre-Alain Joye, Kalle Sommer Nielsen, Peter Cowburn, Adam Harvey, Ferenc Kovacs, Levi Morrison")
		PhpInfoPrintTableRow(2, "Event Maintainers", "Damien Seguy, Daniel P. Brown")

		/* Mirroring */

		PhpInfoPrintTableRow(2, "Network Infrastructure", "Daniel P. Brown")

		/* Windows build boxes and such things */

		PhpInfoPrintTableRow(2, "Windows Infrastructure", "Alex Schoenmaker")
		PhpInfoPrintTableEnd()
	}
	if core.sapi_module.phpinfo_as_text == 0 && (flag&1<<5) != 0 {
		var __str *byte = "</div></body></html>\n"
		core.PhpOutputWrite(__str, strlen(__str))
	}
}

/* }}} */
