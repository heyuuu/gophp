--TEST--
version string
--FILE--
<?php
$a = 1;
$b = &$a;
var_dump($a, $b, [$a, $b]);

$c = [1, 2];
$d = &$c[0];
var_dump($c, $d, [$c, $d]);

--EXPECT--
int(1)
int(1)
array(2) {
  [0]=>
  int(1)
  [1]=>
  int(1)
}
array(2) {
  [0]=>
  &int(1)
  [1]=>
  int(2)
}
int(1)
array(2) {
  [0]=>
  array(2) {
    [0]=>
    &int(1)
    [1]=>
    int(2)
  }
  [1]=>
  int(1)
}