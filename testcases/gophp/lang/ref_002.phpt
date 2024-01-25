--TEST--
version string
--FILE--
<?php

$aa = [1];
$bb = &$aa[0];
$cc  = [100, &$bb];
$dd  = &$cc[1];
echo "=======\n";
var_dump($aa, $bb, $cc, $dd);

$cc[1] = 222;
echo "=======\n";
var_dump($aa, $bb, $cc, $dd);

$dd = 111;
echo "=======\n";
var_dump($aa, $bb, $cc, $dd);

function test(array $b)
{
    $b[1] = 444;
}

test($cc);

echo "=======\n";
var_dump($aa, $bb, $cc, $dd);

function test2($v)
{
    $v = 555;
}

test2($dd);

echo "=======\n";
var_dump($aa, $bb, $cc, $dd);

--EXPECT--
=======
array(1) {
  [0]=>
  &int(1)
}
int(1)
array(2) {
  [0]=>
  int(100)
  [1]=>
  &int(1)
}
int(1)
=======
array(1) {
  [0]=>
  &int(222)
}
int(222)
array(2) {
  [0]=>
  int(100)
  [1]=>
  &int(222)
}
int(222)
=======
array(1) {
  [0]=>
  &int(111)
}
int(111)
array(2) {
  [0]=>
  int(100)
  [1]=>
  &int(111)
}
int(111)
=======
array(1) {
  [0]=>
  &int(444)
}
int(444)
array(2) {
  [0]=>
  int(100)
  [1]=>
  &int(444)
}
int(444)
=======
array(1) {
  [0]=>
  &int(444)
}
int(444)
array(2) {
  [0]=>
  int(100)
  [1]=>
  &int(444)
}
int(444)