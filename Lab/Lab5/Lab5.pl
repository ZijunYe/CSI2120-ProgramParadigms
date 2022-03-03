% EXERCISE1

% departure place, arrival place, departure time, arrival time 
flight(montreal, chicoutimi, 15:30, 16:15).
flight(montreal, sherbrooke, 17:10, 17:50).
flight(montreal, sudbury, 16:40, 18:45).
flight(northbay, kenora, 13:10, 14:40).
flight(ottawa, montreal, 12:20, 13:10).
flight(ottawa, northbay, 11:25, 12:20).
flight(ottawa, thunderbay, 19:00, 20:30).
flight(ottawa, toronto, 10:30, 11:30).
flight(sherbrooke, baiecomeau, 18:40, 20:05).
flight(sudbury, kenora, 20:15, 21:55).
flight(thunderbay, kenora, 20:00, 21:55).
flight(toronto, london, 13:15, 14:05).
flight(toronto, montreal, 12:45, 14:40).
flight(windsor, toronto, 8:50, 10:10).

% Rule 1 is correct one 
on_time(H1 : _M1, D, A)  :- 
        flight(D, A, H2 : _M2, _H3 : _M3),H2 - H1 > 1.
         % CHECK the difference of hour is 1, between departure of arrival time

on_time(H1 : M1, D, A)  :- 
        flight(D, A, H2 : M2, _H3 : _M3), 
        H2 - H1 =:= 1,  % Check the hour between departured time and current time is 1
        MM is 60 - M1, % MM is the variable 
        MM + M2 >= 60. % calculate the minutues, whether has enough 60min 




       
% Exercise2 

arrival(flight(D,A),X):- flight(D,A,_,X).
% complex term 
% flight(D,A) =\= flight(D,A,H1:M1,H2:M2)


    
% Exercise3 Sum of Integer 
% sum_int(5,X). 

% if N == 1
sum_int(1,1).  % if the number is 1, then X value it just 1(base case)

% if N > 1 
sum_int(N,X):-
    N > 0,
    N1 is N - 1, % new value, decrement 
    sum_int(N1, Y), % the new value will be return in Y 
    X is Y + N. 



    