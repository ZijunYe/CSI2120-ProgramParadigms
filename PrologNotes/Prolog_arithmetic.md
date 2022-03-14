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
- unification of 1+2 and 3 fails, because 1+2 is term, 3 is number 

**Recursive calculations**
1. Min predicate 
```
min(X,Y,X) :- X < Y 
min(X,Y,Y) :- X >= Y 
```

2. Power 
```
pow(X,1,X).
pow(X,Y,Z) :- Y > 1, 
	Y1 is Y -1, 
	pow(X,Y1,Z1),
	Z is X* Z1 
```

3. gcd 
```
gcd(U,0,U). 
gcd(U,V,W) :- V>0, R is U mod V, gcd(V,R,W). 
```

4. fibonacci 
```
fib(N,F):- N>1 
	N1 is N -1, 
	N2 is N -2, 
	fib(N1,N2),
	fib(N2,F2), 
	F is F1 + F2. 
```

5. Check even or odd 
```
even(0). 
odd(N) :- N > 0, 
	M is N - 1, 
	even(M). 

even(N) :- N > 0, 
	M is N - 1, 
	odd(M). 
```



