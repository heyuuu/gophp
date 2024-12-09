<?php

$fn1 = function () {
    return 1;
};

$fn2 = static function&() {
    $a = 2;
    return $a;
};

$fnArrow = fn() => 3;