<?php

namespace GoPhp\Tools;

use PhpParser\Node;
use PhpParser\NodeAbstract;

class NodeTool
{
    public static function getTypeName(string $className): string
    {
        if ($className == Node::class || $className == NodeAbstract::class) {
            return "Node";
        }

        // typeName 为去除前缀和分割符的类名
        if (str_starts_with($className, Node::class . '\\')) {
            $className = substr($className, strlen(Node::class) + 1);
        }
        $typeName = str_replace(['\\', '_'], '', $className);

        // 特殊类型，父类型名后移
        foreach (["Stmt", "Expr", "Scalar"] as $type) {
            if (str_starts_with($typeName, $type)) {
                $typeName = substr($typeName, strlen($type)) . $type;
                break;
            }
        }

        return $typeName;
    }
}
