<?php

namespace GoPhp\Tools\Generate;

use GoPhp\Tools\Common\AstTool;
use GoPhp\Tools\Common\NodeType;

class AstNodeGenerator extends BaseGenerator
{
    private string $outputFile = MAIN_ROOT . '/compile/ast/ast.go';

    public function generate()
    {
        $types      = AstTool::allTypes();
        $interfaces = [];
        $classes    = [];
        $extends    = [];

        $interfaces[] = "    Node interface {\n        node()\n    }\n";
        foreach ($types as $type) {
            if ($type->isInterface) {
                $interfaces[] = $this->printInterface($type);
            } else {
                $classes[]         = $this->printClass($type);
                $extends["node"][] = "func (*{$type->typeName}) node() {}\n";
                foreach ($type->supers as $super) {
                    if ($super === "Node") {
                        continue;
                    }
                    $superMethod       = lcfirst($super) . "Node";
                    $extends[$super][] = "func (*{$type->typeName}) {$superMethod}() {}\n";
                }
            }
        }

        $code = "package ast\n\n";
        $code .= "type (\n" . join("\n", $interfaces) . "\n)\n";
        $code .= "type (\n" . join("\n", $classes) . "\n)\n";
        foreach ($extends as $super => $types) {
            $code .= "\n// $super\n";
            $code .= join("", $types);
        }

        $this->writeFile($this->outputFile, $code);
    }

    private function printInterface(NodeType $type): string
    {
        $code = "    " . $this->buildClassComment($type) . "\n";
        $code .= "    {$type->typeName} interface {\n";
        foreach ($type->supers as $super) {
            if ($super == "PhpParserNodeAbstract") {
                $super = "Node";
            }
            $code .= "        " . $super . "\n";
        }
        $code .= "        " . lcfirst($type->typeName) . "Node()\n";
        $code .= "    }\n";
        return $code;
    }

    private function printClass(NodeType $type): string
    {
        $code = "    " . $this->buildClassComment($type) . "\n";
        $code .= "    {$type->typeName} struct {\n";
        foreach ($type->fields as $field) {
            $docComment = $this->clearPropertyDocComment($field->docComment);
            $goType     = $field->typeHint?->toGoType() ?: 'any';
            if ($docComment) {
                $code .= "        {$field->newName} {$goType} // {$docComment}\n";
            } else {
                $code .= "        {$field->newName} {$goType}\n";
            }
        }
        $code .= "    }\n";
        return $code;
    }

    private function buildClassComment(NodeType $type): string
    {
        $comment = '// ' . $type->typeName;
        if ($type->supers) {
            $comment .= ' : ' . join(', ', $type->supers);
        }
        return $comment;
    }

    private function clearPropertyDocComment(string $comment): string
    {
        return trim(substr($comment, 3, strlen($comment) - 5));
    }
}