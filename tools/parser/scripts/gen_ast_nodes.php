<?php

require_once __DIR__ . '/_helpers.php';
require_once __DIR__ . '/_ast_helpers.php';

(new GenAstNodes)->run();

class GenAstNodes
{
    private string $outputFile = PROJ_ROOT . '/php/ast/ast.go';

    public function run()
    {
        $types      = getAllTypes();
        $interfaces = [];
        $classes    = [];
        $extends    = [];
        foreach ($types as $type) {
            if ($type->isInterface) {
                $interfaces[] = $this->printInterface($type);
            } else {
                $classes[] = $this->printClass($type);
                foreach ($type->supers as $super) {
                    $superMethod       = lcfirst($super) . "Node";
                    $extends[$super][] = "func (*{$type->newName}) {$superMethod}() {}\n";
                }
            }
        }

        $code = "package ast\n\n";
        $code .= "type Node interface {}\n\n";
        $code .= "type (\n" . join("\n", $interfaces) . "\n)\n";
        $code .= "type (\n" . join("\n", $classes) . "\n)\n";
        foreach ($extends as $super => $types) {
            $code .= "\n// $super\n";
            $code .= join("", $types);
        }

        file_put_contents($this->outputFile, $code);
    }

    private function printInterface(Type $type): string
    {
        $code = $this->buildClassComment($type) . "\n";
        $code .= "{$type->newName} interface {\n";
        foreach ($type->supers as $super) {
            if ($super === 'PhpParserNodeAbstract') {
                $super = 'Node';
            }
            $code .= "    $super\n";
        }
        $code .= "    " . lcfirst($type->newName) . "Node()\n";
        $code .= "}\n";
        return $code;
    }

    private function printClass(Type $type): string
    {
        $code = $this->buildClassComment($type) . "\n";
        $code .= "{$type->newName} struct {\n";
        foreach ($type->fields as $field) {
            $docComment = $this->clearPropertyDocComment($field->docComment);
            $goType     = $field->typeHint?->toGoType() ?: 'any';
            if ($docComment) {
                $code .= "    {$field->newName} {$goType} // {$docComment}\n";
            } else {
                $code .= "    {$field->newName} {$goType}\n";
            }
        }
        $code .= "}\n";

        return $code;
    }

    private function buildClassComment(Type $type): string
    {
        $comment = '// ' . $type->rawName;
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