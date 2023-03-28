<?php

// parse args
use GoPhp\Tools\Application;

$opts   = getopt("s:o:", ["src:", "output:"]);
$src    = $opts["src"] ?? $opts["s"] ?? "";
$output = $opts["output"] ?? $opts["o"] ?? "";

// main
if ($src && $output) {
    require_once __DIR__ . '/vendor/autoload.php';
    $code = (new Application())->run($src, $output);
    exit($code);
}

// Show help
echo <<<'HELP'
Usage:
    php parser.php [arguments]
The arguments are:
    -s|--src        sources file path (file or directory)
    -o|--output     output path (file or directory)
HELP;
