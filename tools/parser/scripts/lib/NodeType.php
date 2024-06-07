<?php

namespace GoPhp\Tools\Scripts;

use GoPhp\Tools\Parse\NodeTool;
use JsonSerializable;
use PhpParser\Node;
use ReflectionClass;
use Stringable;

class NodeType
{
    public readonly string $typeName;
    public readonly bool   $isInterface;
    /** @var string[] */
    public readonly array $supers;
    /** @var NodeTypeField[] */
    public readonly array $fields;

    public function __construct(ReflectionClass $refClass)
    {
        $this->typeName    = NodeTool::getTypeName($refClass->name);
        $this->isInterface = $refClass->isInterface() || $refClass->isAbstract();
        $this->supers      = $this->initSupers($refClass);
        $this->fields      = $this->initFields($refClass);
    }

    private function initSupers(ReflectionClass $refClass): array
    {
        $supers = [];
        if ($refClass->getParentClass()) {
            $supers[] = NodeTool::getTypeName($refClass->getParentClass()->name);
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
            $fields[] = new NodeTypeField($property);
        }
        return $fields;
    }
}