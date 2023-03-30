<?php

require_once __DIR__ . '/_helpers.php';
require_once __DIR__ . '/_ast_helpers.php';

(new GenAstDecoder)->run();

class GenAstDecoder
{
    private string $outputFile = PROJ_ROOT . '/php/parser/internal/phpparse/decode_node.go';
    private string $template   = <<<'CODE'
package phpparse

import "gophp/php/ast"

func decodeNode(data map[string]any) (node ast.Node, err error) {
	nodeType := data["nodeType"].(string)
	switch nodeType {
%s
	}

	return node, nil
}
CODE;

    public function run()
    {
        $cases = [];
        foreach (getAllTypes() as $type) {
            if (!$type->isInterface) {
                $cases[] = $this->caseType($type);
            }
        }

        $code = sprintf($this->template, join("\n", $cases));
        file_put_contents($this->outputFile, $code);
    }

    private function caseType(Type $type): string
    {
        $indent = str_repeat('    ', 3);
        $fields = [];
        foreach ($type->fields as $field) {
            $value     = $this->castValue($field->typeHint, "data[\"{$field->rawName}\"]");
            $fields [] = "{$field->newName}: $value,\n";
        }
        $fieldsStr = join($indent, $fields);

        return <<<CASE
    case "{$type->rawName}":
        node = &ast.{$type->newName}{
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
                    return "asTypeOrNil[{$name}](${value})";
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