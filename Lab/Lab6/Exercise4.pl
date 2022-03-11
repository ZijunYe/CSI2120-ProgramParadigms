% the tail of last element is empty list 
% the tail of last element is empty list 


secondLast(H,[H|[_|[]]]):- !.
           % in here the dont care term is last element 
secondLast(X,[_|Y]):- secondLast(X,Y). 


% X is the value we need to output 
% continue cut the first element in the list 