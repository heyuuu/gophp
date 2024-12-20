<?php

namespace GoPhp\Tools\Generate;

use GoPhp\Tools\Common\NodeTool;
use GoPhp\Tools\Common\NodeType;
use GoPhp\Tools\Common\TypeHint;

class AstDecodeGenerator extends BaseGenerator
{
    private string $outputFile = MAIN_ROOT . '/compile/parser/internal/phpparse/decode_node.go';
    private string $template   = <<<'CODE'
package phpparse

import "github.com/heyuuu/gophp/compile/ast"

func decodeNode(data map[string]any) (node ast.Node, err error) {
	nodeType := data["nodeType"].(string)
	switch nodeType {
%s
	}

	return node, nil
}
CODE;

    public function generate()
    {
        $cases = [];
        foreach (NodeTool::allTypes() as $type) {
            if (!$type->isInterface) {
                $cases[] = $this->caseType($type);
            }
        }

        $code = sprintf($this->template, join("\n", $cases));
        $this->writeFile($this->outputFile, $code);
    }

    private function caseType(NodeType $type): string
    {
        $indent = str_repeat('    ', 3);
        $fields = [];
        foreach ($type->fields as $field) {
            $value    = $this->castValue($field->typeHint, "data[\"{$field->rawName}\"]");
            $fields[] = "{$field->newName}: $value,";
        }
        $fieldsStr = join("\n" . $indent, $fields);

        return <<<CASE
    case "{$type->typeName}":
        node = &ast.{$type->typeName}{
            {$fieldsStr}
        }
CASE;
    }

    private function castValue(?TypeHint $typeHint, string $value): string
    {
        if ($typeHint === null) {
            return $value;
        }

        switch ($typeHint->mode) {
            case TypeHint::SIMPLE:
                if ($typeHint->name === 'int') {
                    return "asInt({$value})";
                } elseif ($typeHint->name == 'float64') {
                    return "asFloat({$value})";
                }

                $name = $typeHint->toGoType('ast.');
                if ($typeHint->nullable) {
                    return "asTypeOrNil[{$name}]({$value})";
                } else {
                    return $value . ".({$name})";
                }
            case TypeHint::ARRAY:
                $item = $typeHint->item->toGoType('ast.');
                return "asSlice[{$item}]({$value})";
        }

        return $value;
    }
}