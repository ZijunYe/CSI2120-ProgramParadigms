
% Exercise1 


countDownR(N) :- N<0,!.
countDownR(N) :- writeln(N),
NN is N-1,
countDownR(NN).