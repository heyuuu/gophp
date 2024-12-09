<?php

// lit
null;
false;
true;
1;
1.0;
'one string';
"another string";
[
    1, 2, 1.0, 'abc'
];

// variable
$var;
$$var;

// index
$var['a']['b']['c'];

// cast
(bool)$var;
(int)$var;
(double)$var;
(string)$var;
(array)$var;
(object)$var;
(unset)$var;

// new
new stdObject;
new $class;

// unary op
+$var;
-$var;
!$var;
~$var;
++$var;
--$var;
$var++;
$var--;

// binary op
$v1 & $v2;
$v1 | $v2;
$v1 ^ $v2;
$v1 && $v2;
$v1 || $v2;
$v1 xor $v2;
$v1 ?? $v2;
$v1 . $v2;
$v1 / $v2;
$v1 == $v2;
$v1 > $v2;
$v1 >= $v2;
$v1 === $v2;
$v1 - $v2;
$v1 % $v2;
$v1 * $v2;
$v1 != $v2;
$v1 !== $v2;
$v1 + $v2;
$v1 ** $v2;
$v1 << $v2;
$v1 >> $v2;
$v1 < $v2;
$v1 <= $v2;
$v1 <=> $v2;




