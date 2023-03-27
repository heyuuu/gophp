<?php

namespace GoPhp\Tools;

use PhpParser\Parser;
use PhpParser\ParserFactory;
use Symfony\Component\Finder\Finder;
use Symfony\Component\Finder\SplFileInfo;

class Application
{
    private Parser      $parser;
    private NodeEncoder $encoder;

    public function __construct()
    {
        $this->parser  = (new ParserFactory())->create(ParserFactory::ONLY_PHP7);
        $this->encoder = new NodeEncoder();
    }

    function run(string $src, string $output): int
    {
        foreach ($this->eachFile($src) as [$file, $relativeFile]) {
            $ast = $this->parse($file);
            if (!$ast) {
                throw new \Exception("解析文件语法失败: " . $relativeFile);
            }
            $json = $this->encoder->encode($ast);

            $outputFile = $output . DIRECTORY_SEPARATOR . $relativeFile . '.json';
            $this->safeWriteFile($outputFile, $json);
        }

        return 0;
    }

    private function eachFile(string $src)
    {
        $src = realpath($src);
        if (!$src) {
            throw new \Exception("src 文件地址不存在");
        }

        if (is_file($src)) {
            yield [$src, basename($src)];
        } elseif (is_dir($src)) {
            $finder = new Finder();
            $finder->in($src)->files()->name("*.php");
            /** @var SplFileInfo $fileInfo */
            foreach ($finder as $fileInfo) {
                yield [$fileInfo->getPathname(), $fileInfo->getRelativePathname()];
            }
        }
    }

    private function parse(string $file): ?array
    {
        $code = file_get_contents($file);
        return $this->parser->parse($code);
    }

    private function safeWriteFile(string $outputFile, string $content)
    {
        if (!file_exists($dir = dirname($outputFile))) {
            mkdir($dir, 0755, true);
        }
        file_put_contents($outputFile, $content);
    }
}