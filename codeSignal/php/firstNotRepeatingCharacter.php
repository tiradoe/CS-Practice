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

    foreach ($tempArray as $letter => $letterCount) {
        if ($letterCount === 1) return $letter;
    }

    return '_';
}

echo firstNotRepeatingCharacter('zzabaccxccyyzz');
