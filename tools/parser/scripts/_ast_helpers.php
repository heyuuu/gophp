<?php

/* types */

use PhpParser\Node;
use Symfony\Component\Finder\Finder;

class Type
{
    public ReflectionClass $reflection;
    public bool            $isInterface = false;
    public string          $rawName;
    /** @var string[] */
    public array  $supers = [];
    public string $newName;

    /** @var Field[] */
    public array $fields = [];
}

class Field
{
    public string $rawName;
    public string $newName;
    public string $type;
}

/* functions */

/**
 * @return Type[]
 */
function getAllTypes(): array
{
    $types  = [];
    $finder = new Finder();
    $finder->in(PHP_ROOT . '/vendor/nikic/php-parser/lib/PhpParser/Node')->files();
    /** @var \Symfony\Component\Finder\SplFileInfo $file */
    foreach ($finder as $file) {
        $className = Node::class . "\\" . str_replace(["/", ".php"], ["\\", ""], $file->getRelativePathname());
        if (class_exists($className) || interface_exists($className)) {
            $refClass = new ReflectionClass($className);
            $types[]  = buildType($refClass);
        }
    }
    usort($types, function (Type $a, Type $b) {
        return $a->rawName < $b->rawName ? -1 : 1;
    });
    return $types;
}

function buildType(ReflectionClass $class): Type
{
    $type = new Type();

    $type->reflection  = $class;
    $type->isInterface = $class->isInterface() || $class->isAbstract();
    $type->rawName     = toTypeName($class->name);
    $type->newName     = toNewTypeName($class->name);

    // supers
    if ($class->getParentClass()) {
        $type->supers[] = toTypeName($class->getParentClass()->getName());
    }
    foreach ($class->getInterfaceNames() as $interfaceName) {
        if ($interfaceName == JsonSerializable::class || $interfaceName == Node::class) {
            continue;
        }
        $type->supers[] = toTypeName($interfaceName);
    }

    // fields
    foreach ($class->getProperties() as $property) {
        $field = new Field();

        $field->rawName = $property->getName();
        $field->newName = ucfirst(trim($property->getName(), '_'));
        $field->type    = $property->getType() ? $property->getType()->getName() : 'any';

        $type->fields[] = $field;
    }

    return $type;
}

function toNewTypeName(string $className): string
{
    $className = toTypeName($className);

    foreach (["Stmt", "Expr", "Scalar"] as $type) {
        if (str_starts_with($className, $type)) {
            $className = substr($className, strlen($type)) . $type;
            break;
        }
    }

    return $className;
}

function toTypeName(string $className): string
{
    if (str_starts_with($className, Node::class . '\\')) {
        $className = substr($className, strlen(Node::class) + 1);
    }
    return str_replace(['\\', '_'], '', $className);
}
