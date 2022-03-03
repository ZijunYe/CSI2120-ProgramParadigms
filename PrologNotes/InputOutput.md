# Prolog Input & Output 

## Reading and writing to console 
- write to the screen or to a file 
- read from the keyboard or from a file 

**Write** 
- what you write is what you get 
```
- write(X) %X can be anything you wanna to input 

-  write(X) nl %nl new line command 

- tab(N) %output N spaces 

```
Example: 
```
write(1+2)  output 1+2 
```

What's the difference between ```Write``` & ```display```? 
- write displays exactly same as user input 
- display turn user input into predicate format 
```
write(3+4), nl, display(3+4), nl. 
```

output: 
```
3+4 
+(3+4)
```

**Read** 
- input from currently open stream ```read```
- the prompt: ```:```
```
?- read(X).
|: a(1,2).
X = a(1,2). 
```

**Interactive example**
```
age(X,Y):- 
	write('Give the age of'),
	write(X),write(': '),
	read(Y). 



?- age(teddy,Z).
Give the age of teddy: 22. 
Z = 22
```

## Reading and writing to file 

## Character i/o 