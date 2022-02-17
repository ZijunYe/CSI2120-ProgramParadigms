# Introduction To Prolog
```ls.``` : show list of directory

```cd('directoryName').```: move into one specific directory 

```pwd.```: check which directory you at 

```['fileName.pl'].```: consult file content 

```X```: any cap in the (), then it will return answer 

```_```: means don't care 


## Output& input quries 

*output1:*
```
?- book(X,_,'Springer',_,_).
X = 'Programming in Prolog'.
```

```
?- book(Y,_,_,X,_),X >1990.
Y = 'Programming in Prolog',
X = 2003 ;
Y = 'Prolog programming for artificial intelligence',
X = 2001.
```

```
?- reader(X,'James Brown'),checkedOut(X,Y),book(T,_,_,_,Y). 
X = 1234567,
Y = 'QA 76.73 .P76 C57 2003',
T = 'Programming in Prolog'.
```

*output2:*
```
1 ?- gmp(emma,zach).
true.

2 ?- gmp(marie,paul).
false.

3 ?- gmp(X,tim).
X = josee ;
false.

4 ?- gmp(X,anne).
false.
```