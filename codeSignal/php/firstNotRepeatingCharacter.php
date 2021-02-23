<?php
function firstNotRepeatingCharacter(string $string): string
{
    $tempArray = [];

    for ($i = 0; $i < strlen($string); $i++) {
        if (array_key_exists($string[$i], $tempArray)) {
            $tempArray[$string[$i]] = $tempArray[$string[$i]] + 1;
        } else {
            $tempArray[$string[$i]] = 1;
        }
    }

    foreach ($tempArray as $key => $letterCount) {
        if ($letterCount === 1) return $key;
    }

    return '_';
}

echo firstNotRepeatingCharacter('zzabaccxccyyzz');
