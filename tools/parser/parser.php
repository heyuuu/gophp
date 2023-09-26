<?php

// parse args
use GoPhp\Tools\Application;

$opts = getopt("c:f:s:o:", ["code", "file", "src:", "output:"]);
$code = $opts["code"] ?? $opts["c"] ?? "";
$file = $opts["file"] ?? $opts["f"] ?? "";

// main
if ($code || $file) {
    require_once __DIR__ . '/vendor/autoload.php';
    $application = new Application();
    if ($code) {
        $status = $application->parseCode($code);
    } elseif ($file) {
        $code   = file_get_contents($file);
        $status = $application->parseCode($code);
    }
    // exit($status);
    exit();
}

// Show help
echo <<<'HELP'
Usage:
    php parser.php [arguments]
The arguments are:
    -c|--code       source code string
    -f|--file       source file path
For example:
    php parser.php -c "var_dump(1);"
HELP;
