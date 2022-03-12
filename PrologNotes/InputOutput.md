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
```open(X,Y,Z)``` need 3 arguments 
- argument1: Filename 
- argument2: file mode{ write, append, read} 
- argument3: instantiated to the file handle (the output); use to manipulate the stream status(close, set-input..) 

- current input and output stream can be set, affecting all input and output commands 
```set_input(X)``` : Query with current_input(X) 
```set_output(X)```: Query with current_output(X) 

- all the read and write predicates can take an extra parameter for the file handle 
```write(X,Y)```: X is the file handle
```
   read(X,Y)
   get(X,Y) 
   get0(X,Y) 
 ```
 
 **Example: write to file** 
 ```
 writeFile(x):- 
 	open('text.txt',append,F),
	write(F,X),nl(F), 
	close(F). 
```
- alternative ways to set the current input and output stream 
``` see(Filename)```: filename becomes are current output stream
```seen.``` : closes current output stream and reverts back to the console 
```tell(Filename)```: filename becomes the current input stream 
``` told```: closes current input stream and reverts back to the keyboard 

## Character i/o 
```put_char(Character)```: puts a character code into the current stream 
- character can either be an integer(e.g ASCII_code) or character (e.g 'a') 
- ```put(ASCII_Code)``` also exists as a non-ISO primitive 

```get_char(Character)```: gets a character into current stream 
- non-iso primitives 
	- ```get0(X)```: unifies the variable X with ASCII code character entered 
	- ```get(X)```: same as get0 but skips spaces 



## Repeat 
- The built-in predicate repeat is a way to generate multiple solution through backtracking 
- Definition: 
```
repeat 
repeat :- repeat
```
**Example for repeat** 
```
test:- repeat,
	write('Answer to everything?(num)'), % tells user to input 
	read(X), 
	(X=:= 42). 
```

















