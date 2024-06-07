<?php

use GoPhp\Tools\Commands\GenerateCommand;
use GoPhp\Tools\Commands\ParseCommand;
use Symfony\Component\Console\Application;

require_once __DIR__ . '/vendor/autoload.php';

const MAIN_ROOT   = __DIR__ . '/../../';
const PARSER_ROOT = __DIR__;

$commands = [
    new ParseCommand(),
    new GenerateCommand(),
];

$app = new Application('gophp-parser', '0.2.0');
$app->addCommands($commands);
$app->run();
