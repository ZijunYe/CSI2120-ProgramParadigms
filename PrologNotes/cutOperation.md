# CUT ! 
- eliminates the backtracking, no looking for more solution 

- reason for using the cut: 
    - faster execution 
    - more efficient use of memory 
- syntax: ```!```

*Example*
```
p(X,Y):- q(X),r(X,Y)!. %1 solution 
p(X,Y):- q(x),!,r(X,Y). %2 solution 
p(X,Y):- !,q(X),r(X,Y). %5 solution 

p(c,c1).
q(a).
q(b).
r(a,a1).
r(a,a2).
r(b,b1).
r(b,b2).
r(b,b3). 
```


**if-the-else with the cut** 
1. explicit test 
```
max(X,Y,X) :- X >= Y. 
max(X,Y,Y) :- X < Y. 
```

2. cut for efficiency 
```
max(X,Y,X) :- X >= Y, !. 
max(_,Y,X). 
```

**Predicate fail** 
- the fail predicate always fails 
- combined with the cut we will get negation(cut-fail)

```
is_false(P) :- P,!, fail. 
is_false(P). 
```
    - if P in the first predicate fail, then it will move on to the next predicate to success
    - if P in the first predicate true, then the whole statement will be fail 




**Negation: not or \+** 
- the negation predicate not(F) with F a term can also written with the operator \+F 
```
\+(2=4). 
true. 


not(2=4). 
true. 
```

 **Term type testing** 
 ```
 var(X) %an uninstatiated variable 
 nonvar(X) %not a variable, or is an instantiated variable 
 compound(X) %a compound term, not atomic 
 atom %an atom, or bound to an atom 
 integer(X) %a floating point number 
 number(X) %a number, a bound to number(integer or float)
 atomic(X) %a number, string or atom or bound to one
 ```