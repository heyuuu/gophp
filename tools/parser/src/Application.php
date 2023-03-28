<?php declare(strict_types=1);

namespace GoPhp\Tools;

use PhpParser\Error;
use PhpParser\ErrorHandler\Throwing;
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

    function parseCode(string $code, string $output = null): int
    {
        try {
            $json = $this->parseCodeToJson($code);
            if ($output) {
                $outputFile = $output . DIRECTORY_SEPARATOR . "ast.php.json";
                $this->safeWriteFile($outputFile, $json);
            } else {
                echo $this->jsonOutput($json);
            }
            return 0;
        } catch (\Exception $exception) {
            if ($output) {
                echo $exception->getMessage();
            } else {
                echo $this->jsonOutput("", $exception->getMessage());
            }
            return 1;
        }
    }

    function parseFile(string $outputFile, string $output = null): int
    {
        try {
            $json = $this->parseFileToJson($outputFile);
            if ($output) {
                $outputFile = $output . DIRECTORY_SEPARATOR . basename($outputFile) . ".json";
                $this->safeWriteFile($outputFile, $json);
            } else {
                echo $this->jsonOutput($json);
            }
        } catch (\Exception $exception) {
            if ($output) {
                echo $exception->getMessage();
            } else {
                echo $this->jsonOutput("", $exception->getMessage());
            }
            return 1;
        }
        return 0;
    }

    function parseDir(string $dir, string $output): int
    {
        try {
            foreach ($this->eachFile($dir) as [$file, $relativeFile]) {
                $json       = $this->parseFileToJson($file);
                $outputFile = $output . DIRECTORY_SEPARATOR . $relativeFile . '.json';
                $this->safeWriteFile($outputFile, $json);
            }
        } catch (\Exception $exception) {
            echo $exception->getMessage();
            return 1;
        }
        return 0;
    }

    private function eachFile(string $dir)
    {
        $dir = realpath($dir);
        if (!$dir) {
            throw new \Exception("目标路径不存在: " . $dir);
        } elseif (!is_dir($dir)) {
            throw new \Exception("目标路径不是个文件夹: " . $dir);
        }

        $finder = new Finder();
        $finder->in($dir)->files()->name("*.php");
        /** @var SplFileInfo $fileInfo */
        foreach ($finder as $fileInfo) {
            yield [$fileInfo->getPathname(), $fileInfo->getRelativePathname()];
        }
    }

    private function parseFileToJson(string $file): string
    {
        $code = file_get_contents($file);
        $ast  = $this->parser->parse($code);
        if (!$ast) {
            throw new \Exception("解析文件语法失败: " . $file);
        }
        return $this->encoder->encode($ast);
    }

    private function parseCodeToJson(string $code): string
    {
        try {
            $ast = $this->parser->parse($code, new Throwing());
            return $this->encoder->encode($ast);
        } catch (Error $exception) {
            throw new \Exception("解析文件语法失败: " . $exception->getMessage());
        }
    }

    private function safeWriteFile(string $outputFile, string $content)
    {
        if (!file_exists($dir = dirname($outputFile))) {
            mkdir($dir, 0755, true);
        }
        file_put_contents($outputFile, $content);
    }

    private function jsonOutput(string $data, string $error = ""): string
    {
        return json_encode([
            'ok'    => empty($error),
            'data'  => $data,
            'error' => $error,
        ], JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE);
    }
}