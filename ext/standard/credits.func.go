// <<generate>>

package standard

import (
	"sik/core"
)

func CREDIT_LINE(module string, authors string) { PhpInfoPrintTableRow(2, module, authors) }
func PhpPrintCredits(flag int) {
	if core.SM__().GetPhpinfoAsText() == 0 && (flag&PHP_CREDITS_FULLPAGE) != 0 {
		PhpPrintInfoHtmlhead()
	}
	if core.SM__().GetPhpinfoAsText() == 0 {
		core.PUTS("<h1>PHP Credits</h1>\n")
	} else {
		core.PUTS("PHP Credits\n")
	}
	if (flag & PHP_CREDITS_GROUP) != 0 {

		/* Group */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(1, "PHP Group")
		PhpInfoPrintTableRow(1, "Thies C. Arntzen, Stig Bakken, Shane Caraveo, Andi Gutmans, Rasmus Lerdorf, Sam Ruby, Sascha Schumann, Zeev Suraski, Jim Winstead, Andrei Zmievski")
		PhpInfoPrintTableEnd()
	}
	if (flag & PHP_CREDITS_GENERAL) != 0 {

		/* Design & Concept */

		PhpInfoPrintTableStart()
		if core.SM__().GetPhpinfoAsText() == 0 {
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
		CREDIT_LINE("Zend Scripting Language Engine", "Andi Gutmans, Zeev Suraski, Stanislav Malyshev, Marcus Boerger, Dmitry Stogov, Xinchen Hui, Nikita Popov")
		CREDIT_LINE("Extension Module API", "Andi Gutmans, Zeev Suraski, Andrei Zmievski")
		CREDIT_LINE("UNIX Build and Modularization", "Stig Bakken, Sascha Schumann, Jani Taskinen, Peter Kokot")
		CREDIT_LINE("Windows Support", "Shane Caraveo, Zeev Suraski, Wez Furlong, Pierre-Alain Joye, Anatol Belski, Kalle Sommer Nielsen")
		CREDIT_LINE("Server API (SAPI) Abstraction Layer", "Andi Gutmans, Shane Caraveo, Zeev Suraski")
		CREDIT_LINE("Streams Abstraction Layer", "Wez Furlong, Sara Golemon")
		CREDIT_LINE("PHP Data Objects Layer", "Wez Furlong, Marcus Boerger, Sterling Hughes, George Schlossnagle, Ilia Alshanetsky")
		CREDIT_LINE("Output Handler", "Zeev Suraski, Thies C. Arntzen, Marcus Boerger, Michael Wallner")
		CREDIT_LINE("Consistent 64 bit support", "Anthony Ferrara, Anatol Belski")
		PhpInfoPrintTableEnd()
	}
	if (flag & PHP_CREDITS_SAPI) != 0 {

		/* SAPI Modules */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "SAPI Modules")
		PhpInfoPrintTableHeader(2, "Contribution", "Authors")

		// # include "credits_sapi.h"

		PhpInfoPrintTableEnd()

		// # include "credits_sapi.h"

	}
	if (flag & PHP_CREDITS_MODULES) != 0 {

		/* Modules */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "Module Authors")
		PhpInfoPrintTableHeader(2, "Module", "Authors")

		// # include "credits_ext.h"

		PhpInfoPrintTableEnd()

		// # include "credits_ext.h"

	}
	if (flag & PHP_CREDITS_DOCS) != 0 {
		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "PHP Documentation")
		CREDIT_LINE("Authors", "Mehdi Achour, Friedhelm Betz, Antony Dovgal, Nuno Lopes, Hannes Magnusson, Philip Olson, Georg Richter, Damien Seguy, Jakub Vrana, Adam Harvey")
		CREDIT_LINE("Editor", "Peter Cowburn")
		CREDIT_LINE("User Note Maintainers", "Daniel P. Brown, Thiago Henrique Pojda")
		CREDIT_LINE("Other Contributors", "Previously active authors, editors and other contributors are listed in the manual.")
		PhpInfoPrintTableEnd()
	}
	if (flag & PHP_CREDITS_QA) != 0 {
		PhpInfoPrintTableStart()
		PhpInfoPrintTableHeader(1, "PHP Quality Assurance Team")
		PhpInfoPrintTableRow(1, "Ilia Alshanetsky, Joerg Behrens, Antony Dovgal, Stefan Esser, Moriyoshi Koizumi, Magnus Maatta, Sebastian Nohn, Derick Rethans, Melvyn Sopacua, Pierre-Alain Joye, Dmitry Stogov, Felipe Pena, David Soria Parra, Stanislav Malyshev, Julien Pauli, Stephen Zarkos, Anatol Belski, Remi Collet, Ferenc Kovacs")
		PhpInfoPrintTableEnd()
	}
	if (flag & PHP_CREDITS_WEB) != 0 {

		/* Websites and infrastructure */

		PhpInfoPrintTableStart()
		PhpInfoPrintTableColspanHeader(2, "Websites and Infrastructure team")

		/* www., wiki., windows., master., and others, I guess pecl. too? */

		CREDIT_LINE("PHP Websites Team", "Rasmus Lerdorf, Hannes Magnusson, Philip Olson, Lukas Kahwe Smith, Pierre-Alain Joye, Kalle Sommer Nielsen, Peter Cowburn, Adam Harvey, Ferenc Kovacs, Levi Morrison")
		CREDIT_LINE("Event Maintainers", "Damien Seguy, Daniel P. Brown")

		/* Mirroring */

		CREDIT_LINE("Network Infrastructure", "Daniel P. Brown")

		/* Windows build boxes and such things */

		CREDIT_LINE("Windows Infrastructure", "Alex Schoenmaker")
		PhpInfoPrintTableEnd()
	}
	if core.SM__().GetPhpinfoAsText() == 0 && (flag&PHP_CREDITS_FULLPAGE) != 0 {
		core.PUTS("</div></body></html>\n")
	}
}
