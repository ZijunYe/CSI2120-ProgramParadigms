indexOf(X,[]):-!,fail.
indexOf(X,[X],0).
indexOf(X,[Y|L],N):-
    indexOf(X,L,NN),
    N is NN+1.