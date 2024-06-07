<?php

$installPath = "/usr/local/bin/gophp-parser";
$phpBin      = PHP_BINARY;
$script      = __DIR__ . '/parser.php';
$content     = <<<CODE
#!/usr/bin/env bash
{$phpBin} {$script} "$@"
CODE;

file_put_contents($installPath, $content);
chmod($installPath, 0744);
echo "Written to {$installPath}.\n";
echo "Done!";
