package standard

const BROWSCAP_NUM_CONTAINS = 5

/* browser data defined in startup phase, eagerly loaded in MINIT */

var GlobalBdata BrowserData = MakeBrowserData(0)

/* browser data defined in activation phase, lazily loaded in get_browser.
 * Per request and per thread, if applicable */

var BrowscapGlobals ZendBrowscapGlobals

const DEFAULT_SECTION_NAME = "Default Browser Capability Settings"

/* OBJECTS_FIXME: This whole extension needs going through. The use of objects looks pretty broken here */

/* Length of prefix not containing any wildcards */

/* Length of regex, including escapes, anchors, etc. */
