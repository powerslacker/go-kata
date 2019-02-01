# String Calculator Kata

To this kata you will need to implement functionality for a StringCalculator.
The file calc.go contains the StringCalculator struct which has a method named
Process. The Process method takes one argument, a string, which should contain
one or more arithmetic operations. For example:

* "1 + 1"
* "2 * 10"
* "7 - 5"
* "9 / 3"
* "10 * 2 + 7 - 5 / 1"

All integers in the given string will be positive whole numbers, each integer
and mathemetical operation should be separated by a single space. There will not
be any parentheses in the given string, but the Process method should still
follow the normal order of operations (PEMDAS).

In the event that the given string does not meet the format above, or if the
arithmetic does not have a valid interger answer (e.x. "0 / 1") the calculator
should return -1 to represent bad input.

The provided test suite located in calc_test.go can be used to verify whether
your implementation meets the requirements.