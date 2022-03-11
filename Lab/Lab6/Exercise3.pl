canalOpen( saturday ).
canalOpen( monday ).
canalOpen( tuesday ).

raining( saturday ).

melting( saturday ).
melting( sunday ).
melting( monday ).


% canal is open 
% not raining 
% not melting 

% winterlude(X). 
% X = tuesday 

weather( X ) :- melting( X ),
!, fail.
weather( X ) :- raining( X ),
!, fail.
weather( _ ).


winterlude( X ) :- 
    canalOpen( X ),% if the canal is open true, then it then take look the next predicate 
    weather( X ). % there has three weather predicate, first two both need to be fail then it will reach the last one back to true 

% weather(_) % always will return true 
% winterlude(tudesday) --> return ture or false 
% winterlude(X) --> return the Variable value 