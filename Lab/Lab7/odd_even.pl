% Odd and Even 
% check the number
% using mod 2 

% how to append number to the list? 

% base case 
oddEven([],[]). 
oddEven([H|T],[even|L]):- 
    0 is H mod 2,!, % checking the curr head is even number 
    oddEven(T,L). 


oddEven([H|T],[odd|L]) :- 
    1 is H mod 2 ,!, 
    oddEven(T,L). 

