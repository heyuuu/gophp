<?php

namespace GoPhp\Tools\Scripts;

class Field
{
    public string    $rawName;
    public string    $newName;
    public string    $type;
    public string    $docComment;
    public ?TypeHint $typeHint = null;
}