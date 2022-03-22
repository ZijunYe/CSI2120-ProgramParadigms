sudoku([[2,1,4,3],[4,3,2,1],[1,2,3,4],[3,4,1,2]]). 
sudoku([[2,1,4,3],[4,3,2,1],[1,2,3,3],[3,4,1,2]]). 

% Question1 
% predicate different/1 that is true if all numbers in a list are different 

different([]). 
different([H|T]) :- not(member(H,T)) , different(T). 

% Question2 
% extracts the 4 columns of 4x4 mini-Sudoku 
extract([],_,[]). 
extract([H|T],I,[V|R]):- 
     extract(T,I,R),
    nth1(I,H,V).

extract4Columns(M,L):- 
    extract(M,1,A), 
    extract(M,2,B), 
    extract(M,3,C), 
    extract(M,4,D), 
    L = [A,B,C,D]. 
    

% Question3 
extract4Quadrant(M,L):- 
    % extract each value in the Sudo M 
    nth1(1,M,L1), 
    nth1(2,M,L2), 
    nth1(3,M,L3), 
    nth1(4,M,L4), 
    % set each quadrant value 
    % quadrant1 
    nth1(1,L1,A1), 
    nth1(2,L1,A2), 
    nth1(1,L2,A3), 
    nth1(2,L2,A4), 
    Q1 = [A1,A2,A3,A4], 
    % quadrant2 
    nth1(3,L1,B1), 
    nth1(4,L1,B2), 
    nth1(3,L2,B3), 
    nth1(4,L2,B4), 
    Q2 = [B1,B2,B3,B4], 
    % quadrant3 
    nth1(1,L3,C1), 
    nth1(2,L3,C2), 
    nth1(1,L4,C3), 
    nth1(2,L4,C4),
    Q3 = [C1,C2,C3,C4], 
    % quadrant4 
    nth1(3,L3,D1), 
    nth1(4,L3,D2), 
    nth1(3,L4,D3), 
    nth1(4,L4,D4), 
    Q4 = [D1,D2,D3,D4], 

    % put all together 
    L = [Q1,Q2,Q3,Q4]. 

% Question4 
% checks if each sublist of a list of contains different members 
allDifferents([]). 
allDifferents([L1|L2]) :- different(L1),allDifferents(L2). 

% Question5 

checkSudoku(L):- 
    allDifferents(L), 
    extract4Columns(L,R1), allDifferents(R1), 
    extract4Quadrant(L,R2), allDifferents(R2). 


% Check if each of the lists in the Sudoku
% allDifferents
% extract4Columns 
% extract4Quadrants 



