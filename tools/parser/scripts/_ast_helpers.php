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
    public string    $rawName;
    public string    $newName;
    public string    $type;
    public string    $docComment;
    public ?TypeHint $typeHint = null;
}

class TypeHint
{
    const SIMPLE = 1;
    const ARRAY  = 2;
    const UNION  = 3;

    public int  $mode; // 1: simple, 2: array, 3: union
    public bool $nullable;
    // simple type
    public string $name;
    // array type
    public TypeHint $item;
    // union type
    /** @var TypeHint[] */
    public array $types;

    public function toGoType(string $pkg = ""): string
    {
        switch ($this->mode) {
            case self::SIMPLE:
                if ($pkg && str_starts_with($this->name, '*')) {
                    return '*' . $pkg . substr($this->name, 1);
                } elseif ($this->isBuiltin($this->name)) {
                    return $this->name;
                } else {
                    return $pkg . $this->name;
                }
            case self::ARRAY:
                return '[]' . $this->item->toGoType($pkg);
            default:
                return 'any';
        }
    }

    private function isBuiltin(string $type): bool
    {
        return ($type[0] >= 'a' && $type[0] <= 'z');
    }

    public static function simple(string $name, bool $nullable = false): TypeHint
    {
        $type           = new TypeHint();
        $type->mode     = self::SIMPLE;
        $type->name     = $name;
        $type->nullable = $nullable;
        return $type;
    }

    public static function array(TypeHint $item, bool $nullable = false): TypeHint
    {
        $type           = new TypeHint();
        $type->mode     = self::ARRAY;
        $type->item     = $item;
        $type->nullable = $nullable;
        return $type;
    }

    /**
     * @param TypeHint[] $types
     * @param bool       $nullable
     * @return TypeHint
     */
    public static function union(array $types, bool $nullable = false): TypeHint
    {
        $type           = new TypeHint();
        $type->mode     = self::UNION;
        $type->types    = $types;
        $type->nullable = $nullable;
        return $type;
    }
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

        $field->rawName    = $property->getName();
        $field->newName    = ucfirst(trim($property->getName(), '_'));
        $field->type       = $property->getType() ? $property->getType()->getName() : 'any';
        $field->docComment = $property->getDocComment() ?: "";
        if (preg_match('/@var ([^ ]+)/', $field->docComment, $matches)) {
            $field->typeHint = parseTypeHint($matches[1]);
        }

        $type->fields[] = $field;
    }

    return $type;
}

function parseTypeHint(string $str): ?TypeHint
{
    if (empty($str)) {
        return null;
    }

    if (preg_match("/^([\\w\\\\]+(\[\])?)(\|[\\w\\\\]+(\[\])?)*$/", $str)) {
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
                    $types[] = TypeHint::array(parseSimpleTypeHint(substr($part, 0, -2)));
                } else {
                    $types[] = parseSimpleTypeHint($part);
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
        $itemType = parseTypeHint(substr($str, 1, strlen($str) - 4));
        return TypeHint::array($itemType);
    }

    if (str_starts_with($str, 'array<') && str_ends_with($str, '>')) {
        $itemType = parseTypeHint(substr($str, 6, strlen($str) - 7));
        return TypeHint::array($itemType);
    }

    throw new Exception("解析 TypeHint 失败:" . $str);
}

function parseSimpleTypeHint(string $str): TypeHint
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
