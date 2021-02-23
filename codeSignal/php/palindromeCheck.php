<?php
function palindromeCheck(string $string): bool
{
    $char_array = str_split($string);

    if ($char_array === array_reverse($char_array)) {
        return true;
    }

    return false;
}


function printBool(bool $bool)
{
    echo ($bool === true ? "True\n" : "False\n");
}


echo printBool(palindromeCheck('ababba'));
echo printBool(palindromeCheck('turtle'));
echo printBool(palindromeCheck('racecar'));
