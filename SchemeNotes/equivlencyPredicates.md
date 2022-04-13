## Equivlency Predicates 
- the predicate symbol ```?``` indicates a boolean function returning #t or #f 
```
(symbol? x)
(number? x)
(eq? x y) --> true is x and y have interally the same representation
(equal? x y) --> true if x and y are identical object 
(null? x) --> true if x is empty lists
(pair? x) --> true if x is a list or pair
(procedure? x) --> true if x is a function 
(list? x) --> true if x is a list
```
- **Equality test**
1. eq? compares internal representations 
	- addresses(pointer values)
	- cannot be used to reliably compare(numbers / characters)
```
(define hello "bonjour")
(eq? hello hello)
=> #t 

(eq? "bonjour" "bonjour")
=> #f --> if the compiler in the efficient state 
```

2. eqv? 
- can used for characters and numbers
- character and numbers are compared by their *value* 
- can not be used to compare lists, strings or functions
```
(eqv? 1 1)
=> #t 

(eqv? 2 (+ 1 1))
=> #t 

(eqv? 1 1.0)
=> #f 
```

3. equal? 
- compares the structure and contents 
- works for lists, strings and functions 
```
(equal? '(a 1 2) '(a 1 2))
=> #t 

(equal? "bonjour" "bonjour") 
=> #t 

(equal? (list 1 2) '(1 2)) 
=> #t 

(equal? 'a 'a)
=> #t 

(equal? 2 2) 
=> #t 
```


## Control structures 
- control structures in scheme are simple. there are no loops. there are only functions, conditional expressions and the sequence 
- sequence start with begin 
```
(begin (display 'okay)(display '(great)))
=> okay(great)
```












