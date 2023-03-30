<?php

namespace GoPhp\Tools;

use PhpParser\Node;

class NodeEncoder
{
    function encode(array $ast): string
    {
        return json_encode($this->transform($ast), JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE);
    }

    /** 定制 json 格式 */
    private function transform(mixed $data): mixed
    {
        if (is_array($data)) {
            foreach ($data as $i => $item) {
                $data[$i] = $this->transform($item);
            }
        } elseif ($data instanceof Node) {
            return ['nodeType' => $this->getNodeType($data)] + $this->transform(get_object_vars($data));
        }
        return $data;
    }

    private function getNodeType(Node $node): string
    {
        $name = get_class($node);
        if (str_starts_with($name, Node::class . '\\')) {
            $name = substr($name, strlen(Node::class) + 1);
        }
        return str_replace(['\\', '_'], '', $name);
    }
}