<?php

// parse args
use GoPhp\Tools\Application;

$opts = getopt("c:", ["code"]);
$code = $opts["code"] ?? $opts["c"] ?? "";

// main
if ($code) {
    require_once __DIR__ . '/vendor/autoload.php';
    $app = new Application();
    echo $app->parseCode($code);
    exit();
}

// Show help
echo <<<'HELP'
Usage:
    php parser.php [arguments]
The arguments are:
    -c|--code       source code string
For example:
    php parser.php -c "<?php var_dump(1);"
HELP;
