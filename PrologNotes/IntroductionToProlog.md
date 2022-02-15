# Prolog Introduction 

- logic programming 
- prolog is descriptive( as opposed to prescriptive)
	- **Descriptive**: describing known facts and relationships(or rules) 
	- **Prescriptive**: sequence of steps taken by a computer to solve a specific problem

- steps in Prolog 
	1. specific facts 
	2. Define rules(establish new facts)
	3. Start queries 


**First order logic**
- predicate symbols 
- equality 
- negation 
- logic binary connections
- quantifiers 'for all' & 'there exists' 

**Commutation in Prolog** 
- partly by logical declaration semantics of prolog 
- partly by what new facts prolog can infer from given 
- partly explicit control information supplies by the programmer


## Facts 

```like(dogs,cats).```
- lower case for both individuals and relationships
	- is the argument is Upper case, then it is variable
- relationship(or predicate) is written first 
- individuals (or arguments) are written in parenthesis
- end with a dot 
order of arguments is important 

## Rules 
- rules need go to the pl file, not ?- part 
- a prolog rule has a head and body, separated by ```:-``` pronounced "if"
- head + body (first satify body, then satify head) 
- the rule also end in "."
- if there no match with rule, it will return ```false```

```
likes(horses,X) :- like(X,hay). 
```
	- horses like X if X like hay
```
like(horses,X) :- like(X,hay),like(X,mice). 
```
	- body can be conjunction 
	



## constants or Atoms
```
domestic(cows). 
faster(horses,cows). 
take(cats,milk,cows). 
```
	- exmaple  cows. horses,milk
	- symbolic: small caps letter followed by letters and numbers 
	- Numbers : integer and float 

## Variable 
- More interesting question the type: "Do cats like X"
	- we want prolog to tell What X could stand for 
```
?- like(cats,X)
```
- variables start with uppercase letters 


**how prolog answes?**
- when prolog is first ask question, variable X is initially not instantiated 
- prolog seraches through the database, looking for a facts that unifies with the question(same structure)
- when the facts is found, then X become instantiated with the argument that is exist in the database 
- Prolog search the facts in order(top to bottom)
- Prolog will marks the place in the database where the unifier is found 
*IF THERE HAVE MULTIPLE ANSWERS*
- when entering ```;``` we ask prolog to re-satisfy the goal or search to another solution 
- prolog resume its search, starting from where it left the place-maker 
- If it directly jump to ```?-``` --> means there is no more answer 

## Conjunctions
```,``` represents "and"
- can have any number of question separated by comma and end with dot 
```
?- like(cats,dogs), like(dogs,cats).
```
	- do cats and dogs like each other 

```
?- like(cows,X),like(horses,X). 
```
	- is there anything that horses and cows both like 
	- steps
		1. find if there is some X that cows like 
		2. find if horse like what ever X is 



## Queries or Questions 
- question are about indivduals and their relationships 
```
?- eat(cats,mice). 
```
	- means "do cats eat mice? "


## Database 
- sequence of facts 
```
like(horses,fish).
like(dogs,cats).
like(cats,mice).
like(dogs,mice).
like(horses,racing).
like(cats,horses).
like(tigers,cats).
like(cats,hay).
like(cows,grass).
like(cows,hay).
like(horses,hay).
```

- simple queries 
```
?- like(dogs,bones).
?- like(cats,dogs).
?- like(cats,hay).
?- enjoy(horses,racing
```


```%``` : comments 













