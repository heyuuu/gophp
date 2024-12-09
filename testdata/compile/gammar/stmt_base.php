<?php

// empty stmt
;
// expr stmt
1;
// const stmt
const A = 1;

// echo stmt
echo "abc";
echo $a;

// global stmt
global $a;
global $$a;

// static var stmt
static $a, $b = 1;

// unset stmt
unset($a);

// if stmt
if ($a === 1) {
    // branch if
    echo 'if';
} elseif ($b == 2) {
    // branch elseif
    echo 'elseif';
} else {
    // branch else
    echo 'else';
}

// switch stmt
switch ($a === 2) {
    case 1:
        // branch case 1
        break;
    case 2:
        // branch case 2
        // fallthrough
    case 3:
        break;
    default:
        // branch default
}

// for stmt
for ($i = 0; $i < 10; $i++) {
    // branch for
    echo "for";
}

// foreach stmt
foreach ($arr as $key => $value) {
    // branch foreach
    echo "foreach";
}

// while stmt
while (true) {
    while (true) {
        break 2;
    }
}

// do-while stmt
do {
    continue;
    break;
} while (true);

// try-catch-finally stmt
try {
    // branch try
    echo 1;
} catch (\Exception $e) {
    // branch catch
    echo "catch";
} finally {
    // branch finally
    echo "finally";
}
