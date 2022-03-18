% reverse the order of the elements 
% remove every even index number 


reverseDrop(L,R):- reverseDrop(L,1,[],R),!.  % start with odd index 
reverseDrop([],_,R,R):- !. 

reverseDrop([HL|TL],N,I,R) :- % when the index is odd, add index value 
    N = 1, 
    reverseDrop(TL,0,[HL|I],R). 

reverseDrop([_|TL],N,I,R) :- % when the index is even, ignore the value 
    N = 0, 
    reverseDrop(TL,1,I,R). 
