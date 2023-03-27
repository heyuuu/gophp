<?php

namespace GoPhp\Tools;

class NodeEncoder
{
    function encode(array $ast): string
    {
        return json_encode($ast, JSON_PRETTY_PRINT);
    }
}