% Maximum and Minimum elements 

maxmin([H|T],Max,Min):- 
    maxmin(T,H,H,Max,Min).  % here we setup curr max and min to first index value 

% 4 cases 
% 1 Base case empty list 
maxmin([],Max,Min,Max,Min). 

% 2 when the number bigger than the currMax, change currMax 
maxmin([H|T],MX,MN,Max,Min):- 
    H > MX,
    maxmin(T,H,MN,Max,Min). 

% 3 when the number smaller than the currMin, change currMin 
maxmin([H|T],MX,MN,Max,Min):- 
    H < MN, 
    maxmin(T,MX,H,Max,Min). 

% 4 when the number neither 2,3cases, just move on to next index 
maxmin([_|T],MX,MN,Max,Min) :- 
    maxmin(T,MX,MN,Max,Min). 
