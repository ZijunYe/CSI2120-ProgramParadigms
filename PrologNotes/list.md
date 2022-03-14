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

**list representation** 
1. dot operator 
- The list {pie,fruit,cream} is written (pie.(fruit.(cream.nil)))
- The variable X followed by Y can be written as (X.Y)
- {X = pie; Y = fruit.cream.nil}




