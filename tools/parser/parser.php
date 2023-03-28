<?php

// parse args
use GoPhp\Tools\Application;

$opts   = getopt("c:f:s:o:", ["code", "file", "src:", "output:"]);
$code   = $opts["code"] ?? $opts["c"] ?? "";
$file   = $opts["file"] ?? $opts["f"] ?? "";
$dir    = $opts["dir"] ?? $opts["d"] ?? "";
$output = $opts["output"] ?? $opts["o"] ?? null;

// main
if ($code || $file || $dir) {
    require_once __DIR__ . '/vendor/autoload.php';
    $application = new Application();
    if ($code) {
        $status = $application->parseCode($code, $output);
    } elseif ($file) {
        $status = $application->parseFile($file, $output);
    } else { // $src
        if (!$output) {
            die("Require --output when --dir is used");
        }
        $status = $application->parseDir($dir, $output);
    }
    exit($status);
}

// Show help
echo <<<'HELP'
Usage:
    php parser.php [arguments]
The arguments are:
    -r|--code       source code string
    -f|--file       source file path
    -d|--dir        source directory path
    -o|--output     output directory path
For example:
    php parser.php -r "var_dump(1);"
    php parser.php -s code/ -o dist/
HELP;
