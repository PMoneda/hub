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
else
elif
listen
```

## Rules

```
program => : program 
           | statement
program => END_PROG

statement =>  : statements 
              | print expression
              | get string function
              | post string function
              | read variable
              | if expressions { statements } ELSEIF
              | if expressions { statements } 
              | for expressions { statements }
              | import string
ELSIF => : 0
         | elif expressions { statements }
         | else { statements }
statements => : statements
              | statement
   
```