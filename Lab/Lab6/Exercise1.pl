
% Exercise1 
countDownR(N) :- N<0,!.
countDownR(N) :- writeln(N),
NN is N-1,
countDownR(NN).

% Test the above predicate with the query 
% ?- countDown(5).. Explain what happens and why? Find a solution to the problem that you have observed.

% countDown(N):- repeat, writeln(N), N is N-1, N<0, !. 
% repeat:  when it fail, it will brings back and repeat 