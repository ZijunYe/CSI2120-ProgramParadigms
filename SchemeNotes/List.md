# Lists
- internally a list consists two pointers 
	- the first pointer gives the address of the atom 
	- the second pointer gives the address of the next cell 

	- '()  == empty 

- List construction 
1. ```(cons obj1 obj2)``` 
	- ob1 will go in the front 
	- if obj2 is item instead if list, then it will create pair instead of list 
	- (cons element list)
2. ```(append '(1) '(2 3))```
	- both argument has to be list 
	- the first argument will be at front 
	- second argument will be at last 

3. ```(list 1 2 3)```
	- same as '(1 2 3)

EXAMPLE1
```
(cons 'a '(b c))
=> (a b c)

(cons '(a b) '(b c))
=> ((a b) b c)
```

EXAMPLE2
```
(cons 'blue empty) or (cons 'blue '())
(cons 'red (cons 'blue empty))
(cons (sqr 2) empty)
(cons(cons 3 (cons true empty)) (cons (nake-posn 1 3) empty))
```

- if the second item is not list, then it is pairs. 
```
(cons 'a 'b)
=> (a . b)
```

## List operations 
```CAR``` : always return the first item in the list, same as using ```first```
```
(car '(a b c)) 
=> a 
(car '((a b) b c))
=> (a b)
```

```CDR	```: always return the rest item in the list, same as using ```rest```
```
(cdr '(a b c))
=> (b c)

(cdr '((a b) b c))
=> (b c)

(cdr '(a (b c)))
=> ((b c))
```

- nesting list expression
	- cdr car cdr  = cd a dr = cdadr

```
(cdadr '(a (b c d) e))
=> (c d)

(cons (car '(a b c)) (cdr '(a b c)))
=> (a b c)
```









