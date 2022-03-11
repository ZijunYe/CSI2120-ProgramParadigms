element(chlorine,'Cl').
element(helium,'He').
element(hydrogen,'H').
element(nitrogen,'N').
element(oxygen,'O').


lookup(S) :- 
    element(E,S), 
    write(S),write(' is the symbol for: '), writeln(E), !. 

% only will execute second lookup when the first one unable to find value 
lookup(S) :- 
    writeln('Don\'t know the symbol: '), writeln(S),!, fail. 

elements:- writeln("Elements in the Periodic Table"), 
    repeat,
    write('symbol to look-up:'), 
    read(S),
    not(lookup(S)), 
     % if it fail, then this will turning to true, and continue executing 
     % if it success, then this will turing to false, and go back to repeat 
    writeln('Exiting'), 
    !, fail. 
    % when it not(lookup(S)) is true, it will continue to executing here 
    % cut method will stop the repeat method 
    % here the fail it will stop entire program 


