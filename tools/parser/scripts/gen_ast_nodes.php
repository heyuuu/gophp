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
        foreach ($types as $type) {
            if ($type->isInterface) {
                $interfaces[] = $this->printInterface($type);
            } else {
                $classes[] = $this->printClass($type);
            }
        }

        $code = "package ast\n\n";
        $code .= "type Node interface {}\n\n";
        $code .= join("\n", $interfaces) . "\n";
        $code .= join("\n", $classes);

        file_put_contents($this->outputFile, $code);
    }

    private function printInterface(Type $type): string
    {
        $code = $this->buildClassComment($type) . "\n";
        $code .= "type {$type->newName} interface {\n";
        $code .= "}\n";
        return $code;
    }

    private function printClass(Type $type): string
    {
        $code = $this->buildClassComment($type) . "\n";
        $code .= "type {$type->newName} struct {\n";
        foreach ($type->fields as $field) {
            $code .= "    {$field->newName} {$field->type}\n";
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
}