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

    public function parseCode(string $code): string
    {
        try {
            $ast = $this->parser->parse($code);

            // resolve namespaced name
            $traverser = new NodeTraverser();
            $traverser->addVisitor(new NameResolver());
            $traverser->traverse($ast);

            $data = $this->encoder->encode($ast);

            return $this->jsonOutput($data);
        } catch (Error $e) {
            return $this->jsonOutput("", "Parse Fail: " . $e->getMessage());
        } catch (\Throwable $e) {
            return $this->jsonOutput("", "Unexpected error: " . $e->getMessage());
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