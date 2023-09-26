<?php

namespace GoPhp\Tools\Scripts;

use Exception;
use GoPhp\Tools\NodeTool;
use JsonSerializable;
use PhpParser\Node;
use ReflectionClass;
use Stringable;

class NodeType
{
    public readonly string $typeName;
    public readonly string $newTypeName;
    public readonly bool   $isInterface;
    /** @var string[] */
    public readonly array $supers;
    /** @var Field[] */
    public readonly array $fields;

    public function __construct(ReflectionClass $refClass)
    {
        $this->typeName    = NodeTool::getTypeName($refClass->name);
        $this->newTypeName = NodeTool::getNewTypeName($refClass->name);
        $this->isInterface = $refClass->isInterface() || $refClass->isAbstract();
        $this->supers      = $this->initSupers($refClass);
        $this->fields      = $this->initFields($refClass);
    }

    private function initSupers(ReflectionClass $refClass): array
    {
        $supers = [];
        if ($refClass->getParentClass()) {
            $supers[] = NodeTool::getTypeName($refClass->getParentClass()->getName());
        }
        foreach ($refClass->getInterfaceNames() as $interfaceName) {
            if ($interfaceName == JsonSerializable::class || $interfaceName == Node::class || $interfaceName == Stringable::class) {
                continue;
            }
            $supers[] = NodeTool::getTypeName($interfaceName);
        }
        return $supers;
    }

    private function initFields(ReflectionClass $refClass): array
    {
        $fields = [];
        foreach ($refClass->getProperties() as $property) {
            if ($property->isStatic() || $property->name === 'attributes') {
                continue;
            }

            $field = new Field();

            $field->rawName    = $property->getName();
            $field->newName    = ucfirst(trim($property->getName(), '_'));
            $field->type       = $property->getType() ? $property->getType()->getName() : 'any';
            $field->docComment = $property->getDocComment() ?: "";
            if (preg_match('/@var ([^ ]+)/', $field->docComment, $matches)) {
                $field->typeHint = $this->parseTypeHint($matches[1]);
            }

            $fields[] = $field;
        }
        return $fields;
    }

    private function parseTypeHint(string $str): ?TypeHint
    {
        if (empty($str)) {
            return null;
        }

        if (preg_match("/^([\\w\\\\]+(\[])?)(\|[\\w\\\\]+(\[])?)*$/", $str)) {
            $parts = explode('|', $str);

            $types    = [];
            $nullable = false;
            foreach ($parts as $part) {
                if ($part == 'null') {
                    $nullable = true;
                } else {
                    if ($part == 'array') {
                        $types[] = TypeHint::array(TypeHint::simple('any'));
                    } elseif (str_ends_with($part, '[]')) {
                        $types[] = TypeHint::array($this->parseSimpleTypeHint(substr($part, 0, -2)));
                    } else {
                        $types[] = $this->parseSimpleTypeHint($part);
                    }
                }
            }
            if (count($types) == 0) {
                throw new Exception("解析 TypeHint 失败:" . $str);
            } elseif (count($types) == 1) {
                $type = $types[0];
            } else {
                $type = TypeHint::union($types);
            }
            $type->nullable = $nullable;
            return $type;
        }

        if (str_starts_with($str, '(') && str_ends_with($str, ')[]')) {
            $itemType = $this->parseTypeHint(substr($str, 1, strlen($str) - 4));
            return TypeHint::array($itemType);
        }

        if (str_starts_with($str, 'array<') && str_ends_with($str, '>')) {
            $itemType = $this->parseTypeHint(substr($str, 6, strlen($str) - 7));
            return TypeHint::array($itemType);
        }

        throw new Exception("解析 TypeHint 失败:" . $str);
    }

    private function parseSimpleTypeHint(string $str): TypeHint
    {
        if (str_starts_with($str, 'Node\\')) {
            $str = substr($str, 5);
        }
        if (str_starts_with($str, 'Expr\\')) {
            $str = substr($str, 5) . 'Expr';
        }
        $str = str_replace(['\\', '_'], '', $str);

        $name = match ($str) {
            // base type
            'bool' => 'bool',
            'int' => 'int',
            'float' => 'float64',
            'string' => 'string',
            // interface
            'Node' => 'Node',
            'Stmt' => 'Stmt',
            'Expr' => 'Expr',
            // others
            'ArrayItem' => '*ArrayItemExpr',
            'ClosureUse' => '*ClosureUseExpr',
            'Case' => '*CaseStmt',
            'DeclareDeclare' => '*DeclareDeclareStmt',
            'UseUse' => '*UseUseStmt',
            'ElseIf' => '*ElseIfStmt',
            'Else' => '*ElseStmt',
            'PropertyProperty' => '*PropertyPropertyStmt',
            'StaticVar' => '*StaticVarStmt',
            'TraitUseAdaptation' => '*TraitUseAdaptationStmt',
            'Catch' => '*CatchStmt',
            'Finally' => '*FinallyStmt',
            default => '*' . $str
        };

        return TypeHint::simple($name);
    }
}