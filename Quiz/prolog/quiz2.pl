gSeries(X,Y):-gSeries(X,0,Y).
gSeries(X,N,Y):-pow(X,N,Y).
gSeries(X,N,Y):- N1 is N+1, 
                gSeries(X,N1,Y).