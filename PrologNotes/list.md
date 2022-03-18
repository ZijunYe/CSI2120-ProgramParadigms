# Data structures : Lists 
- a list holds objects. elements in a list may be list themselves 
*Examples* 
- ```[1,2,3]``` list with 3 elements 
- ```[]``` empty list 
- ```[head|tail]``` return head as first element in the list, rest element return to tail 
- ```[1,2,three]``` no typing in the list 
- ```[1,2|tail]``` head may consist of multiple elements 
- ```L= [1,2,3,4]``` sign the variable to the list 

**Basic list Processing** 
1. list memebership 
```
listMember(X,[X|L]). 
listMember(X,[Y|L]):- listMember(X,L). 
```
2. determine the length of a list 
```
listLength([],0).
listLength([X|L],N):- 
    listLength(L,NN),
    N is NN+1. 
```

3. list insertion with permutations 
```
listInsert(A,L,[A|L]). 
listInsert(A,[X|L],[X|LL]):- 
    listInsert(A,L,LL). 
```

4. joining lists 
```
appendList([].Y,Y). 
appendList([A|B],Y,[A|W]):-
    appendList(B,Y,W). 
```

5. list of characters 
```
read_line(Line) :- 
    get_char(C), read_line(C,line). 

read_line('\n',[]) :- !.  % check is not the end 
read_line(C,[C| RestOfLine]) :- 
    C \= '\n',
    get_Char(NextC), 
    read_line(NextC,RestofLine). 
```

6. deletion 
*delete first occurence of an element* 
```
deleteFront(R,[R|L],L). 
deleteFront(R,[X|LL],[X|L]) :- 
    deleteFront(R,LL,L). 
```

*delete all occurences of an element* 
```
deleteAll(_,[],[]). 
deleteAll(X,[X|T],Result) :- 
    deleteAll(X,T,Result),!. 
deleteAll(X,[X|T],[H|Result]) :- 
    deleteAll(X,T,Result). 
```

7. Intersection of two list 
- using ```intersection/3```
```
intersectionList([],_,[]). 
intersection([X|Xs],Ys,Zs) :- 
    \+member(X,Ys), 
    intersectionList(Xs,Ys,Zs). 
intersection([X|Xs],Ys,[X|Zs]) :- 
    member(X,Ys), 
    intersectList(Xs,Ys,Zs). 
```


8. Quick Sorting list 

```
sortList([],[]). 
sortList([P|Q],T) :- partitionList(P,Q,G,D), 
    sortList(G,GG), 
    sortList(D,DD), 
    append(GG,[P|DD],T). 
```
partitionList code: 
```
partitionList(P,[X|L],[X|PG],PD) :- X @< P, % HERE @ is alphanumeric comparison operation, can compare both integer and letter 
    partitionList(P,L,PG,PD). 
partitionList(P,[X|L],PG,[X|PD]) :- X @>= P,
    partitionList(P,L,PG,PD). 
partitionList(P,[],[],[]). 
```


9. invert the order of list 
*Using Append* 
```
mirror([],[]).% empty list is mirrored itself 
mirror([x|L1],L2):= 
    mirror(L1,L3), 
    append(L3,[X],L2). %L3 is the list, X is the append value, L2 is the resulting list 
```
*Using accumulator- helper function/ wrapper pattern function design*
```
reverseList([],L,L) :-!. 
reverseList([H|T],L,R) :- 
    reverseList(T,[H|L],R). 
mirrorAcc(L,R) :- reverseList(L,[],R). 
```
- it works either directly calling reverseList or call mirrorAcc, just input argument different 


10. Operators on the list 
*Example1 - apply an operator to each element of a list* 
*Example2 - sum up elements of a list of numbers*


**list representation** 
1. dot operator 
- The list {pie,fruit,cream} is written (pie.(fruit.(cream.nil)))
- The variable X followed by Y can be written as (X.Y)
- {X = pie; Y = fruit.cream.nil}




**Other build-in operation**

```findall```
```bagof```: keep duplicates and not sorted 
- takes three input 
```setof```: remove all duplicates, sorted result 