# Recursion Example 

1. Inverting of list (slide#21)
2. list membership (slide#22)
3. size(length) of a list (Slide#23)
4. equivalence of two lists(Slide#26)
5. removing duplicates (Slide#27)
6. Minimal Element (Slide#29)
7. Traversal --> accept function as argument 
	```
		(define (apply-list fct L)  
			(if (null? L ) '() 
				(cons (fct (car L)) 
						(apply-list fct (cdr L)))))
	```
	- simliar as ```(map function list)```

8. adding a prefix to the elements (Slide#33)
9. Reduction of list to value (Slide#35) --> similar to foldl & foldr 