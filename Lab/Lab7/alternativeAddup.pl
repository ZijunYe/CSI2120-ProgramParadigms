% alternatively add number 


if the index is odd add 
if the index is even subtract 

addAlternate(L,S) :- addAlternate(L,0,p,S). 

% base case 
addAlternate([],A,_,A). 

addAlternate([H|T],A,p,S) :- !, 
    AA is A + H, 
    addAlternate(T, AA, m , S). 

addAlternate([H|T], A, m,S) :- !, 
    AA is A - H, 
    addAlternate(T,AA,p,S). 

