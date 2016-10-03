# Grammar for the Hub lang

This will describe all grammar rules and keywords for the language


## keywords

```
datasource
get
post
if
return
for
defun
mod
and
or
not
import
case
switch
print
read
else
elif
listen
nand
nor
xor
```

## Rules

```
program => : program 
           | statement
program => END_PROG

statement =>  : statements 
              | print expression
              | read ID
              | get string function
              | post string function
              | read variable
              | if expressions { statements } ELSEIF
              | for expressions { statements }
              | import string
              | listen number ( statements )
              | type ID { IDS }
              | var ID = expressions
              | const ID = expressions
              | defun ID (args) { statements }


expressions => : expressions
               | expression
               | (expression)

expression => : ID
              | ID OP expressions

args => : 0
        | arg

arg  => : args
        | IDS

OP => : +
      | -
      | *
      | /
      | mod
      | and
      | or
      | not
      | nand
      | nor
      | xor
ELSIF => : 0
         | elif expressions { statements }
         | else { statements }

ID  => : name

IDS => : IDS,
       | ID
number => integer

statements => : 0
              | statements
              | statement
   
```