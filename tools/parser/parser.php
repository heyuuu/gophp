<?php

use GoPhp\Tools\Commands\ParseCommand;
use Symfony\Component\Console\Application;

require_once __DIR__ . '/vendor/autoload.php';

$commands = [
    new ParseCommand(),
];

$app = new Application('gophp-parser', '0.2.0');
$app->addCommands($commands);
$app->run();
