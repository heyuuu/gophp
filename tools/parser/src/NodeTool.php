<?php

namespace GoPhp\Tools;

use PhpParser\Node;

class NodeTool
{
    public static function getTypeName(string $className): string
    {
        if ($className == Node::class) {
            return "Node";
        }
        if (str_starts_with($className, Node::class . '\\')) {
            $className = substr($className, strlen(Node::class) + 1);
        }
        return str_replace(['\\', '_'], '', $className);
    }

    public static function getNewTypeName(string $className): string
    {
        $typeName = self::getTypeName($className);

        foreach (["Stmt", "Expr", "Scalar"] as $type) {
            if (str_starts_with($className, $type)) {
                $typeName = substr($className, strlen($type)) . $type;
                break;
            }
        }

        return $typeName;
    }
}