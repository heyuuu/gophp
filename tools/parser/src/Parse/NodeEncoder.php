<?php

namespace GoPhp\Tools\Parse;

use PhpParser\Node;

class NodeEncoder
{
    public function encode(array $ast): string
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
            $objectVars = $this->transform(get_object_vars($data));
            if ($data instanceof Node\Scalar\String_ || $data instanceof Node\Scalar\EncapsedStringPart) {
                $objectVars["value"] = base64_encode($objectVars["value"]);
            }

            $objectVars['attributes'] = $data->getAttributes();

            return ['nodeType' => $this->getNodeType($data)] + $objectVars;
        }
        return $data;
    }

    private function getNodeType(Node $node): string
    {
        $name = get_class($node);
        return NodeTool::getTypeName($name);
    }
}