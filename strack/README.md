# STRACT
## A stack-based programming language for generating messages

### Overview

in STRACT, we start each program with an empty stack of memory, an instruction pointer starting 0, and a list of instructions, either being a function name, number, or string (we'll talk more about strings later). A function will manipulate the stack in some way, and a number will push entered number to the stack. The code is a series of instruction separated by a space or newline. The program will execute instructions left to right, top to bottom, increasing the instruction pointer by 1 each time (pointing to the next  instruction) until the instruction pointer goes past the last instruction or before the first instruction, ending the program.

```
1 3 5 -> | 1 | 3 | 5 |
3 5 ADD -> | 8 |
3 5 2 MULT SUB -> | 7 |
```

When calling a function, the arguments are provided to the function in the order that they are popped off the stack. This means that if the stack `| 2 | 4 | 6 |`, the first argument would `6`, second would be `4` and third would be `2`. So `1 2 SUB` would result in a `1` on the stack, opposed to `-1`. As soon as the function is called, the cells which are arguments to the function are popped off the stack and then the function runs.

Each Cell on the stack is 4 bytes, and can be numbers from `2^31 - 1` to `-2^31`, also know as an int32. If the number goes past this range (Integer Overflow), the number should wrap around (Ex: `MAX + 1 == MIN` or `2^31 - 1 + 1 == -2^31`).

### Comments

In your code, a comment is is anything after a `#` until a newline (`\n`) character is hit. This is an example of a complete program which adds 2 numbers and prints the result
```
# My Program

1 2 ADD # The stack now has a 3 on it
INTSTRING PRINT # This should print "3"
```

Note that this program is equivalent to `1 2 ADD INTSTRING PRINT`

### Strings

In Stract, each number on the stack can be represented as 4 ascii characters (1 byte per character). A String must also end with a terminating cell where the last byte (Most Significant Byte) in the cell is a 0. String Cells can have 0s in other parts of the call that aren't the last byte which will be ignored when printing.

You could push numbers onto the stack to make your string, or you can also type out your string within double-quotes (`"`) and the correct numbers will be pushed for you. It's also important to know that strings are pushed in reverse, this helps with printing later. Characters can also be escaped with backslash `\`, the characters that can be escaped are `\n` -> newline, `\t` -> tab, `\"` -> printing `"`, and `\\` -> printing `\`.

Example pushing "Hi" to the stack:
```
Storing:
  '\0' 'i' 'H'

Make grouping divisible by 4 (padding with 0):
  '\0' 'i' 'H' '\0'

Turn characters into their ascii values:
  0 105 72 0

Turn ascii values into binary string
  00000000 01101001 1001000 00000000
  
Turn binary string into int32
  6899712
```
This also means that `6899712` and `"Hi"` are equivalent instructions in STRACT.

Pushing bigger strings will require pushing multiple numbers, let's use "STRACT" as an example:
```
Storing:
  '\0' 'T' 'C' 'A' 'R' 'T' 'S'

Make grouping divisible by 4 (padding with 0):
  '\0' 'T' 'C' 'A' 'R' 'T' 'S' '\0'

Turn characters into their ascii values:
  0 84 67 65 82 84 83 0

Turn ascii values into binary string:
  00000000 01010100 01000011 01000001
  01010010 01010100 01010011 00000000

Turn binary string into int32:
  5522241 1381257984
```

Again, that means that `5522241 1381257984` and `"STRACT"` are equivalent instructions.

To print a string, you can call the function PRINT, which will consume values on the stack until it reads a terminating cell (Last byte of cell is 0). If PRINT reads the entire stack and doesn't find a terminating cell, the program will crash.

```
"Hello, World!" -> | 2188396 (\0!dl) | 1919899424 (roW ) | 745499756 (,oll) | 1699217408 (eH\0\0) |
"Hello, World!" PRINT -> | -> output: Hello, World!
```

### CJUMP

CJUMP (Conditional Jump) is the sole way of doing both conditionals and loops. CJUMP takes 2 arguments, first being the jump offset (forwards or backwards with negative numbers), and second being the conditional value. If the conditional value is 0 no jump occurs. If the conditional value is any other value, the instruction pointer be offset by the jump offset in the first argument. If the jump offset goes outside the bounds of the program, the program will end.
```
"A" 1 2 CJUMP "B" PRINT -> | -> output: "A"
"A" 0 2 CJUMP "B" PRINT -> | 4259840 (\0A\0\0) | -> output: "B"
```

### DUP and POP

DUP and POP are the main ways of getting access to values in the stack aren't on the top of the stack. DUP takes 1 argument, which is how many cells to duplicate, and POP takes 1 argument, which is how many cells to pop off the stack. In both of these functions, the first argument is not counted to the total of cells to duplicate/pop. DUP and POP are very helpful for reading values that are anywhere in the stack as long as you know the offset from the top of the stack.

```
"H" "e" "l" "o" -> | 4718592 (\0H\0\0) | 6619136 (\0e\0\0) | 7077888 (\0l\0\0) | 7274496 (\0o\0\0) |
"H" "e" "l" "o" 4 DUP -> | 4718592 (\0H\0\0) | 6619136 (\0e\0\0) | 7077888 (\0l\0\0) | 7274496 (\0o\0\0) | 4718592 (\0H\0\0) | 6619136 (\0e\0\0) | 7077888 (\0l\0\0) | 7274496 (\0o\0\0) |
"H" "e" "l" "o" 4 DUP 3 POP -> | 4718592 (\0H\0\0) | 6619136 (\0e\0\0) | 7077888 (\0l\0\0) | 7274496 (\0o\0\0) | 4718592 (\0H\0\0) |
"H" "e" "l" "o" 4 DUP 3 POP PRINT -> | 4718592 (\0H\0\0) | 6619136 (\0e\0\0) | 7077888 (\0l\0\0) | 7274496 (\0o\0\0) | -> output: H
"H" "e" "l" "o" 4 DUP 3 POP PRINT 4 DUP 2 POP PRINT 1 POP 4 DUP 1 POP 1 DUP PRINT PRINT 2 POP 4 DUP PRINT 7 POP -> | -> output: Hello
```

### List of Functions

#### MATH
n2 n1 ADD:
Pushes `n1 + n2` to the stack

n2 n1 SUB:
Pushes `n1 - n2` to the stack

n2 n1 MULT:
Pushes `n1 * n2` to the stack

n2 n1 MOD:
Pushes `n2 modulos n1` to the stack
Note that the sign(+/-) of the result copies the sign of `n1`

n2 n1 RSFT:
Pushes `n2 >> n1` to the stack (right shift)

n2 n1 LSFT:
Pushes `n2 << n1` to the stack (left shift)

#### BITWISE

n2 n1 AND:
Pushes `n1 & n2` to the stack

n2 n1 OR:
Pushes `n1 | n2` to the stack

n2 n1 XOR:
Pushes `n1 ^ n2` to the stack

n1 INV:
Pushes `n1^0xffffffff` to the stack (flips all the bits)

#### CONDITIONALS
n2 n1 MORE:
Pushes `1` if n1 > n2, else pushes `0`

n2 n1 LESS:
Pushes `1` if n1 < n2, else pushes `0`

n2 n1 EQ:
Pushes `1` if n1 == n2, else pushes `0`

n1 NOT:
Pushes `1` if n1 == 0, else pushes `0`

offset cond CJUMP: [See Section on CJUMP](#CJUMP)

#### STRING FUNCTIONS
PRINT [See Section on Strings](#CJUMP)

n1 INTSTRING:
pushes the number n1 as a string to the stack
```
These 2 programs are equivalent
"1234567" -> | 3618357 (\0765) | 875770417 (4321) |
1234567 INTSTRING -> | 3618357 (\0765) | 875770417 (4321) |
```

#### MISC

n1 DUP [See Section on DUP and POP](#DUP-and-POP)

n1 POP [See Section on DUP and POP](#DUP-and-POP)