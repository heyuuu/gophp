<?php

namespace GoPhp\Tools\Generate;

use GoPhp\Tools\Common\NodeTool;
use GoPhp\Tools\Common\NodeType;
use PhpParser\Node\Expr;
use PhpParser\Node\Stmt;

class AstNodeGenerator extends BaseGenerator
{
    private string $outputFile = MAIN_ROOT . '/compile/ast/ast.go';
    private string $template   = <<<'CODE'
package ast

type File struct {
	Declares   []*DeclareStmt
	Namespaces []*NamespaceStmt
}

// Node
type Node interface {
	node()
}

// baseNode
type baseNode struct {
	Meta map[string]any `json:"@"`
}

func (*baseNode) node() {}

// baseExpr
type baseExpr struct {
	baseNode
}

func (*baseExpr) exprNode() {}

// baseStmt
type baseStmt struct {
	baseNode
}

func (*baseStmt) stmtNode() {}


CODE;

    public function generate()
    {
        $types      = NodeTool::allTypes();
        $groupTypes = $this->groupTypes($types);

        $code = $this->template;
        $code .= $this->printTypes('node interfaces', $groupTypes['interface'] ?? [], [$this, 'printInterface']);
        $code .= $this->printTypes('misc', $groupTypes['misc'] ?? [], fn(NodeType $type) => $this->printClass($type));
        $code .= $this->printTypes('Expr', $groupTypes['expr'] ?? [], fn(NodeType $type) => $this->printClass($type, 'baseExpr'));
        $code .= $this->printTypes('Stmt', $groupTypes['stmt'] ?? [], fn(NodeType $type) => $this->printClass($type, 'baseStmt'));

        $extends = [];
        foreach ($types as $type) {
            if (!$type->isInterface) {
                foreach ($type->supers as $super) {
                    if ($super === "Node" || $super === 'Expr' || $super === 'Stmt') {
                        continue;
                    }
                    $superMethod       = lcfirst($super) . "Node";
                    $extends[$super][] = "func (*{$type->typeName}) {$superMethod}() {}\n";
                }
            }
        }
        foreach ($extends as $super => $types) {
            $code .= "\n// $super\n";
            $code .= join("", $types);
        }

        $this->writeFile($this->outputFile, $code);
    }

    /**
     * @param  NodeType[]  $types
     * @return array<string, NodeType[]>
     */
    private function groupTypes(array $types): array
    {
        $groupTypes = [];
        foreach ($types as $type) {
            if ($type->isInterface) {
                $groupTypes['interface'][] = $type;
            } elseif (is_subclass_of($type->className, Expr::class)) {
                $groupTypes['expr'][] = $type;
            } elseif (is_subclass_of($type->className, Stmt::class)) {
                $groupTypes['stmt'][] = $type;
            } else {
                $groupTypes['misc'][] = $type;
            }
        }

        return $groupTypes;
    }

    /**
     * @param  string  $lineComment
     * @param  NodeType[]  $types
     * @param  callable  $typePrinter
     * @return string
     */
    private function printTypes(string $lineComment, array $types, callable $typePrinter): string
    {
        $typeCodes = array_map($typePrinter, $types);
        return "// {$lineComment}\ntype (\n" . join("\n", $typeCodes) . "\n)\n";
    }

    private function printInterface(NodeType $type): string
    {
        $properties = [];
        foreach ($type->supers as $super) {
            if ($super == "PhpParserNodeAbstract") {
                $super = "Node";
            }
            $properties[] = $super;
        }
        $properties[]    = lcfirst($type->typeName) . 'Node()';
        $propertiesLines = join("\n        ", $properties);

        $typeComment = $this->buildClassComment($type);
        $typeName    = $type->typeName;

        return <<<CODE
    {$typeComment}
    {$typeName} interface {
        {$propertiesLines}
    }
CODE;
    }

    private function printClass(NodeType $type, string $parent = 'baseNode'): string
    {
        $properties = [];
        foreach ($type->fields as $field) {
            $docComment = $this->clearPropertyDocComment($field->docComment);
            $goType     = $field->typeHint?->toGoType() ?: 'any';
            if ($docComment) {
                $properties[] = "{$field->newName} {$goType} // {$docComment}";
            } else {
                $properties[] = "{$field->newName} {$goType}";
            }
        }
        $propertiesLines = join("\n        ", $properties);

        $typeComment = $this->buildClassComment($type);
        $typeName    = $type->typeName;

        return <<<CODE
    {$typeComment}
    {$typeName} struct {
        {$parent}
        {$propertiesLines}
    }
CODE;
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
        $comment = trim(substr($comment, 3, strlen($comment) - 5));
        return str_replace("\n", '\\n', $comment);
    }
}