# Missandei
A simple interpreter built with Go, for an expression-oriented and lexically-scoped "Let" language.


The Let language BNF grammar rules:
1. _Expression_ ::= _Number_
2. _Expression_ ::= `minus` ( _Expression_, _Expression_ )
3. _Expression_ ::= `iszero` ( _Expression_ )
4. _Expression_ ::= `if` _Expression_ `then` _Expression_ `else` _Expression_
5. _Expression_ ::= _Identifier_
6. _Expression_ ::= `let` _Identifier_ = _Expression_ `in` _Expression_

- The interpreter is comprised of three parts: the scanner, the parser, and the evaluator (Abstract Syntax Tree).
- Each tree node is an environment (key, value pair) that holds the values of the variables; the evaluator maintains these bindings with a stack and lookup.
- The interpreter output is simply the resulting evaluation of the program.

## To Run:
`run main.go`, with the supplied example programs: `program0` and `program1`.

``` 
Example:

let x = 7                             // 1st outer declaration of x
in let y = 2                          // 1st outer declaration of y
   in let y = let x = minus(x, 1)     // nested declarations, which shadow outer declarations
              in minus(x,y)           // inner declaration used for calculation where inner x scope ends and y acquires value
      in minus(minus(x,8), y)         // scope of inner y; evals and returns to outer scope twice for final evaluation
```




