<?php

// parse args
use GoPhp\Tools\Application;

$opts = getopt("c:f:s:o:", ["code"]);
$code = $opts["code"] ?? $opts["c"] ?? "";

// main
if ($code) {
    require_once __DIR__ . '/vendor/autoload.php';
    $application = new Application();
    $status      = $application->parseCode($code);
    exit($status);
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
