# Scheme
- a scheme interpreter session uses the interactive READ-EVAL-PRINT programming model (REPL: read,evaluate,print,loop)
- need to have space between identifier and constant
- representation of numbers: most math operations return exact numbers  
    ```
    (/ 50.0 6)
    8.3333333334 

    conversion from inaccurate to exact: 
    (inexact > exact(/ 50.0 6))
    8 187649984473771/562949953421312

    TRUNCATE will return the integer that precedes 
    (truncate(/10 3)) 
    3 

    (truncate 3.7)
    3.0

    ```

## Pure Functional Programming 
- a program corresponds to a function call 
- a function is composition of functions 
- no variables, no assignments 
- no loops, no control statement 

## Basic Concepts 
- list is the fundamental data structure 
```
'(1 2 3) --> return list 
(list 1 2 3) --> return list '(1 2 3)
```
- atom: a number, a character string or a symbol(all data types are equal)
- expression: an atom or a list 
    - list: a series of expressions in parentheses ```(+(sqrt 3)6)```
    - including the empty list,()Nile, both list and atom

- a function is a first-class data object that can be created, assigned to variables, passed as a parameter or returned as a value 

## Evaluations of Expressions 
- constants are evaluated to what they are 
- identifiers are evaluated at the value commonly assigned to them 
- list are evaluated by first evaluating the first expression that composes them 
    - the value of the expression must be function 
    - the arguments of the function are values obtained by the evaluation of the expressions contained the rest of the list 



