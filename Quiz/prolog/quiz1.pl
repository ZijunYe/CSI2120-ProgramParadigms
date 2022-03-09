

% Create the inOrder/3 predicate 
% returning true if the three numbers are in order( increasing or descending)

increasing(A,B):- A >= B. 
descending(A,B):- A =< B.
inOrder(X,Y,Z):- 
    increasing(X,Y),increasing(Y,Z);
    descending(X,Y),descending(Y,Z). 



% Notes 
% Prolog do comparison can only between two variable 