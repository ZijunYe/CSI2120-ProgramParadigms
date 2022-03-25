bunkbeds(L):- L=[[N1,C1],[N2,C2],[kayla,C3],[N4,C4],[N5,C5],[N6,C6]],
    ((N1=reeva,N2=haley);(N2=reeva,N1=haley)),
    % reeva and haley are closest to the door 
    % kayla is in the middle bed 
    member(C1,[brown,black,blue]),
    member(C3,[brown,black,blue]), 
    member(C5,[brown,black,blue]), 
    % color of blankets on the bottom bunks all begin with same letter 
    member([beth,C],L),member(C,[yellow,green,red]), 
    % the color of Beth's blanket doesn't begin with same letter as her name 
    member([blue,red],[[C1,C2],[C3,C4],[C5,C6]]),
    % blue blanket is under the red blanket 
    member(liza,[N1,N5]), 
    % liza like lower bunks 
    member(zoe,[N1,N2,N5,N6]),
    % zoe doesn't share a bunk with kayla, so No N4 
    brown = C5, 
    member([black,yellow],[[C1,C2],[C3,C4],[C5,C6]]),
    member(green,[C1,C2,C3,C4,C5,C6]). 
