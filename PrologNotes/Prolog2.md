# Predicate Calculus 


## Predicated in prolog 
- B <- A1 AND A2 AND A3: all true Ai then implies B is true
- B <- : This is facts because truth is always implied 
- <- A1 AND A2...AND An : without head, correctness still need to be proven

## Horn clauses 
- facts and rules are horn clauses 
```
F :- F1,F2...Fn
```
	-  F is head of clauses 
	- Fi are terms / literals 

Example: 

```prolog
salary(X):- employed(Y,X). 
```
```
salary(X):- employed(_,X). 
```
	- anoymous variable 

## Unification 
- to make two formulas identical by giving values to the variable they contain (matching)
- Predicate name need to be match 
- the arguments number need to be match
Example
![uni EXAMPLE](https://github.com/ZijunYe/CSI2120-ProgramParadigms/blob/main/PrologNotes/sample/Unification_example1.png)
![uni EXAMPLE](https://github.com/ZijunYe/CSI2120-ProgramParadigms/blob/main/PrologNotes/sample/Unification_example2.png)


**Demonstration Steps**


## Resolution 
- resolution is a technique of producing a new clause by resolving two clauses that 
	- contain complimentary literal 
	- produces proof by contraction 

Example: 
![RESOLUTION EXAMPLE](https://github.com/ZijunYe/CSI2120-ProgramParadigms/blob/main/PrologNotes/sample/resolution_example.png) 


