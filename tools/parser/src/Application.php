<?php declare(strict_types=1);

namespace GoPhp\Tools;

use PhpParser\Error;
use PhpParser\NodeTraverser;
use PhpParser\NodeVisitor\NameResolver;
use PhpParser\Parser;
use PhpParser\ParserFactory;

class Application
{
    private Parser      $parser;
    private NodeEncoder $encoder;

    public function __construct()
    {
        $this->parser  = (new ParserFactory())->create(ParserFactory::ONLY_PHP7);
        $this->encoder = new NodeEncoder();
    }

    public function parseCode(string $code): int
    {
        try {
            $ast = $this->parser->parse($code);

            // resolve namespaced name
            $traverser = new NodeTraverser();
            $traverser->addVisitor(new NameResolver());
            $traverser->traverse($ast);

            $json = $this->encoder->encode($ast);

            echo $this->jsonOutput($json);
            return 0;
        } catch (Error $e) {
            $error = "Parse Fail: " . $e->getMessage();
            echo $this->jsonOutput("", $error);
            return 1;
        } catch (\Throwable $e) {
            $error = "Unexpected error:" . $e->getMessage();
            echo $this->jsonOutput("", $error);
            return 2;
        }
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