leafNodes(T,L) :- leafNodes(T,[],L). 

% base case 
leafNodes(nil,L,L) :-!. 

leafNodes(t(Root,nil,nil),L,[Root|L]):-!. 

% K is right 
% I is left 
% what t in here mean? means a tree 
leafNodes(t(Root,Left,Right),L,K):- 
    leafNodes(Left,L,I), 
    leafNodes(Right,I,K).