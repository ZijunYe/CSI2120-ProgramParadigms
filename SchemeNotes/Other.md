# Iteration 
- Formular: 
```
(do ((var inital_value update_value)...) (test condition) trueThen falseThen )
```
- example 
**fibonacci** 
```
(define fibonacci
	(lambda(n) 
		(if (= n 0)
			0 
			(do ((i n (- i 1))
				(a1 1 (+ a1 a2))
				(a2 0 a1))
				((= i 1) a1))
		)
	)
	)

(fibonacci 6)
=> 8 
```

# Vector 
- vector is similar as list or array 
- Basic commands: 
1. create vector with k element ```(make-vector k)```
2. create vector with initial value ```(make-vector k inital_value)```
3. test if the item is vector ```(vector? obj)```
4. make a constant vector 
	- ```#(1 2 3)```
	- ```'#(1 2 3)```
	- ```(vector 1 2 3)``` & ```(vector 'a 'b)```

5. get reference ```(vector-ref #(1 2 3) index)```
6. get vector length ```(vector-length #(1 2 3))```
7. set value in the vector ```(vector-set! #(1 2 3) 1 10)```
8. convert list into a vector ```(list->vector '(1 2 3))```
	- convert vector into list ```(vector->list '(1 2 3 4))```


# Set! 
- the function ```set!``` allows us to assign a value to a variable 
```
(set! a-num(+ 3 4))
(set! a-num (+ 1 a-num))
```
- example 
**Encapsulated set**
```
(define light-switch (let ((lit #f)) 
	(lambda() 
		(set! lit (not lit))
		(if lit 'on 'off))
	))
```

# Sort 
- ```sort``` function only accepts list 
- Merge sort: split list into two, until only one element, merge list on the way up maintaining the order 
- Quick-sort: pick a number in the list as pivot, create two list, 1 is bigger than the pivot and another on smaller than pivot 


# Tree 