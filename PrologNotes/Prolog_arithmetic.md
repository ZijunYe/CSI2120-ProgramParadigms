# Prolog arithmetic (Arithmetic Expressions)

## Built-in operators
**Numbers in prolog** 
- prolog recognizes numbers: integers and floating point 
- number constants: 5,1.75,0...
- rules about arithmetics expression use:
	- number constants
	- arithmetic operators 
	- arithmetic variable 
**Built-in common operators**
```Prolog
X+Y
X-Y
X*Y
X/Y
X//Y
X mod Y
```
*Mathematical Functions* 
```
abs(X)
ln(X)
sqrt(X)
```

*Infix comparison operators*
```
X =:= Y % Equals 
X =\= Y % not Equals 
X < Y 
X =<Y
X > Y
X >= Y 
```
example:
```
?- 1+2 =:= 2+1 %true; it enforce to do the calculation first 
```
*How to evaluating arithmetic Expressions?*
"=" -> matching, unification doesn't involves any calculations 
"is" ->  after is will indicate prolog to do calculation, then do unification
Example
```prolog
?- 3 = 1 +2 %false 
?- 1+2 = 2+1 %false
?- 3 is 1+2 %true
?- X is 1+2, X is 2+1 %X = 3 
```


**Unification with numbers**

**Recursive calculations**

**Generator**




