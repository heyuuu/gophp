package php

func PhpInitConfig(ctx *Context) {
	ig := ctx.INI()

	ig.ConfigInit()

	//openBasedir := PG__(ctx).OpenBasedir()
	//PG__(ctx).openBasedir = ""
	//filename := findIniFile(ctx, ig)
	//PG__(ctx).openBasedir = openBasedir

	//if filename != "" {
	//	var fh = NewFileHandleByFilename(filename)
	//	ig.LoadIniFile(fh)
	//	ig.ConfigSet("cfg_file_path", fh.Filename())
	//	ig.SetIniOpenedPath(fh.Filename())
	//}

	///* Check for PHP_INI_SCAN_DIR environment variable to override/set config file scan directory */
	//ig.SetIniScannedPath("")

	if ig.IniEntries() != "" {
		/* Reset active ini section */
		ig.LoadIniStr(ig.IniEntries())
	}
}
