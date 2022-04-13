# Let 
- define a list of local variables for a list of expressions 
- ```(let ((a 3 ) (b 4))  (+ a b))``` 
	- variable definition : local variable a and b 
	- the expression for bound 


**Let vs Let* **
- let is not allow to have two variable definition dependent to each other 
- let* allows for sequential definition 
```
(let* ((x 1) (y (+ x 1))) (list x y)) 
```

**Let vs letrec** 
- letrec permits the recursive definitions of functions 
```
(letrec ((fact (lambda(n)  
		(if (= n 1)
			1
			(* n (fact (- n 1))))))) (fact5))
```

- same function use let and letrec (factorial example)
	- name in the let expression ```(let name ((var val))  exp1 exp2)```
	- name in the letrec express ```(letrec((name (lambda (var...) exp1 exp2 ))(name val)...)```
```
(let ft ((k 5))  
	(if (<= k 0) 1
		(* k (ft (- k 1)))))
```

```
(letrec ((ft (lambda (k) 
			(if (<= k 0) 1
				(* k (ft (- k 1))))))) (ft 5))
```


**named let-bound** 
```
(define divisors
	(lambda(n) 
		(let f ((i 2)) 
			(cond        
				((>= i n) '())
				((integer? (/ n i)) 
					(cons i (f (+ i 1)))) 
				(else (f (+ i 1)))))))
```

