<?php

namespace GoPhp\Tools\Generate;

abstract class BaseGenerator
{
    abstract public function generate();

    protected function writeFile(string $file, string $content)
    {
        if (!file_exists($dir = dirname($file))) {
            mkdir($dir, 0755, true);
        }
        file_put_contents($file, $content);
    }
}