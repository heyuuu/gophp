<?php

namespace GoPhp\Tools\Scripts;

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
