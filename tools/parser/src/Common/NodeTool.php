<?php

namespace GoPhp\Tools\Common;

use PhpParser\Node;
use ReflectionClass;
use Symfony\Component\Finder\Finder;
use Symfony\Component\Finder\SplFileInfo;

class NodeTool
{
    /**
     * @return NodeType[]
     */
    public static function allTypes(): array
    {
        $types  = [];
        $finder = new Finder();
        $finder->in(PARSER_ROOT . '/vendor/nikic/php-parser/lib/PhpParser/Node')->files();
        /** @var SplFileInfo $file */
        foreach ($finder as $file) {
            $className = Node::class . '\\' . str_replace(['/', '.php'], ['\\', ''], $file->getRelativePathname());
            if (class_exists($className) || interface_exists($className)) {
                $types[] = new NodeType(new ReflectionClass($className));
            }
        }
        usort($types, function (NodeType $a, NodeType $b) {
            return $a->typeName < $b->typeName ? -1 : 1;
        });
        return $types;
    }
}