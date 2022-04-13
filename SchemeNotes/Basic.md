# Function applications in Racket 
- in a racket function application, parentheses are put around the function name and the arguments 
```
g(3,1) becomes (g 3,1) 
g(f(2),g(3,1)) becomes (g (f 2) (g 3 1)) 
```

- ```(trace functionName)``` : return each the procedure happen in function 

## Built-in functions 
```
(quotient 75 7)
(remainder 75 7)
(fliter even? '(1 2 3 4)) => '(2 4) 
```

## Defining functions in Racket 
- ```define``` is a special form
- when you define a varaible in definition window, you can't change it; but you can do in execution part 
```
(define n 7)
(define n 9)--> giving error 
```
```
F(X) = X^2 
(define (f x)(* x x))
		header  body 

G(x,y) = x - y 
(define(g x y)(- x y)) --> takes two variable

```

- a function definition consists: 
	1. a name for the function 
	2. a list of parameters 
	3. a single "body" expression 
- How to do function call?
EXAMPLE: 
```
(define (showIt pts)
	(cond ((or (<= pts 3)(>= pts 65)) 0)
	((<= 4 pts 6) 0.5)
	((<= 7 pts 12) 1.0)
	((<= 13 pts 15) 1.5)
	((<= 16 pts 18) 1.6 )
	(else 2.0)
	)
)

(showIt 25) <-- function call
```

- function definition with lambda 
	- a definition associates a function expression to a name: 
	```
	(define square (lambda(x) (* x x)))
	```
	- or equivalently(and shorter): 
	```
	(define (square x)(* x x))
	```
	- use of define, here procedure square: 
	```
	(square 2)
	=> 4 
	```



## Types of comparisons 
- A comparison is a function that consumes two numbers and produces a boolean value 
```
(= x y)
(< x y)
(> x y)
(<= x y)
(>= x y)
string = ?, string <? --> using to compare string 
```
- complex relationships : AND, OR, NOT 
```
the statement 3<= x < 7
(and (<= 3 x)( < x 7))
```

## Predicates 
- a predicate is a function that produces a boolean result: true if data is of a particular form, and false otherwise 
- built-in predicates: ```even?``` ```negative``` ```zero``` ```string?```
- User-defined: 
```
(define (between? low high nbr))
	(and(< low nbr)(< nbr heigh))
```

```
(define (can-drink? age))
	(>= age 19)
```

## Special syntactic form 
1. **if statement**
- it contains three value 
	1. expression evaluate first 
	2. if it is true , second agrument evaluate 
	3. if it is false, third argument evaluate 
```
(if (= x 0) "Infinity" (/ 1 x)) 
```

2. **Conditional expression**
- using special form ```cond```
- each argument of cond is a question/answer pair 
	- QUESTION is boolean expression 
	- ANSWER is a possible value of conditional expression 
EXAMPLE: 
```
(cond
	[< x 0] (-x)]
	[(>= x 0 ) x])
```

```
(define n 5)
(cond [(even? n) "even"] [(odd? n) "odd"])
```
- square brackets used by convention, for readability 

FORMATE: (similar to switch)
```
(cond
[question1 answer1]
[question2 answer2]
....
[questionk answerk]
)
```


3. **Creating local scope**
```
(let ((a 3)(b 4)) (* a b))
=> 12 
```

4. **Quotations**
- the quote function ensures that an argument list is not evaluated 
```
(quote (1 2 3))
=> (1 2 3)
```

- the quotation is necessary here, otherwise the first will be evaluate 
```
'(+ 3 4) => (+ 3 4)
(+ 3 4) => 7
```

- the quotation function can simply written with a ```'```
```
'(1 2 3)
=> ( 1 2 3)
```

- Example(using let and quotation together)
```
(let ((a '(1 2 3))(b '(3 4 5))) (cons a b)) == (cons '(1 2 3) '(3 4 5))
=> ((1 2 3) 3 4 5)

```
- the function cons is the dot operator for list(construction function)
	- (cons item list)
	- the list(1 2 3) becomes the first item 

- Building a list with the function list 
```
(list 'a 'b 'c)
=> (a b c)

(list '(a b c)) 
=> ((a b c))
```

5.  **Lambdas**
- lambda expressions are 'local functions'
```
(lambda (var1, var2,...) exp1 exp2 ) --> returns a function applies it to the variable 
```

```
((lambda(x)(* x x)) 3)
=> 3 
```

- multiple variabkes and multiple expression(result of last expression is the answer)
```
((lambda (f x y) (f x x) (f x y) (f y y)) + 2 3) 
=> 6 
only (f y y ) will be apply here 
```


- Lambda expression and let-bound variable 
EXAMPLE: 
```
(let ((x 2) (y 3)) (+ x y)) 
is equivalent to 
((lambda(x y) (+ x y)) 2 3)
```
General format: 
```
(let (var value) (expression)) = ((lambda(var1 var2)(expression)) value1 value2)
```

- top define combined with a lambda 
```
(define fct(lambda (f x)(f x x))) --> here define a function with 2 parameter, first is the operator, second is value 
(fct + 13) --> function call 
(fct * 4)
```




















