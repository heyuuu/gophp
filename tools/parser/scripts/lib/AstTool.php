<?php

namespace GoPhp\Tools\Scripts;

use PhpParser\Node;
use ReflectionClass;
use Symfony\Component\Finder\Finder;
use Symfony\Component\Finder\SplFileInfo;

class AstTool
{
    /**
     * @return NodeType[]
     */
    public static function allTypes(): array {
        $types  = [];
        $finder = new Finder();
        $finder->in(PHP_ROOT . '/vendor/nikic/php-parser/lib/PhpParser/Node')->files();
        /** @var SplFileInfo $file */
        foreach ($finder as $file) {
            $className = Node::class . "\\" . str_replace(["/", ".php"], ["\\", ""], $file->getRelativePathname());
            if (class_exists($className) || interface_exists($className)) {
                $refClass = new ReflectionClass($className);
                $types[]  = new NodeType($refClass);
            }
        }
        usort($types, function (NodeType $a, NodeType $b) {
            return $a->typeName < $b->typeName ? -1 : 1;
        });
        return $types;
    }
}