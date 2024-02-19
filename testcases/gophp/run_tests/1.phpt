--TEST--
Bug #54238 (use-after-free in substr_replace())
--INI--
error_reporting=E_ALL&~E_NOTICE
--FILE--
<?php
var_dump(ini_get("error_reporting"));
var_dump(error_reporting());

$f = array(array('A', 'A'));

$z = substr_replace($f, $f, $f, 1);
var_dump($z, $f);
?>
--EXPECT--
string(5) "32759"
int(32759)
array(1) {
  [0]=>
  string(9) "AArrayray"
}
array(1) {
  [0]=>
  array(2) {
    [0]=>
    string(1) "A"
    [1]=>
    string(1) "A"
  }
}
