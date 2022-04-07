import:-
    csv_read_file('partition65.csv', Data65, [functor(partition)]),maplist(assert, Data65),
    csv_read_file('partition74.csv', Data74, [functor(partition)]),maplist(assert, Data74),
    csv_read_file('partition75.csv', Data75, [functor(partition)]),maplist(assert, Data75),
    csv_read_file('partition76.csv', Data76, [functor(partition)]),maplist(assert, Data76),
    csv_read_file('partition84.csv', Data84, [functor(partition)]),maplist(assert, Data84),
    csv_read_file('partition85.csv', Data85, [functor(partition)]),maplist(assert, Data85),
    csv_read_file('partition86.csv', Data86, [functor(partition)]),maplist(assert, Data86),listing(partition).

% PROBLEM RN: 
% the breakpoint in checkpoint 
% change the label for return 

mergeClusters(L):- 
    findall([D,X,Y,C],partition(_,D,X,Y,C),PL), % we got global list 
    process(PL,L). 

% check base case, clusterlist is empty 

process(PL,NL):- 
    process(PL,0,[],[],NL).
	% 5 argument: global list, check if it is first cluster, append list, cluster id list, result list 
process([],_,R,_,R). 
process([H|T],N,_,CID,R):- 
    N = 0, 
    nth1(4,H,ID), % get current cluster ID
    findall([D,X,Y,ID],partition(_,D,X,Y,ID),CCL), %get cluster of the ID
    process(T,1,CCL,[ID|CID],R). % append cluster to the cluster list, and collect it cluster id 

process([H|T],N,I,CID,R):- 
    N = 1, 
    % find other cluster and check ID 
    nth1(4,H,ID),
    findall([D,X,Y,ID],partition(_,D,X,Y,ID),CCL),% this is second cluster 
    % passing current cluster ID, current cluster List, overall cluster list, result 
    check(ID,CID,CCL,I,NCID,BACK), % to append the new cluster into cluster 
    process(T,1,BACK,NCID,R). 
   
    % first check the cluster id is not in the cluster id list 
    % need to compare the point id check if it is in the current overall cluster 
    % if it is different with all ID --> append 
    % once find the same --> relabel all in the overall cluster 

check(0,_,CID,_,OC,CID,OC). % return back, cuz the cluster already in the cluster list 
check(1,ID,CID,_,OC,[ID|CID],OC). 

check(ID,CID,_,OC,NCID,R) :- 
    member(ID,CID), % if it is true, then we just return back
    check(0,_,CID,_,OC,NCID,R).

check(ID,CID,CC,OC,NCID,R) :- 
    % passing current cluster ID, current cluster List, current cluster,overall cluster list, result 
    not(member(ID,CID)), % if it is true, then we just return back
     % add the id to the collection of cluster id 
    checkPointID(CC,OC,NOC), 
    check(1,ID,CID,CC,NOC,NCID,R). 
   
% we compare each of point in current cluster with overall cluster 

checkPointID([],OC,OC). 
% when we looping through all the point in the current cluster 

checkPointID([H|T],OC,R):- 
    intersect(H,OC,OC,I), 
    checkPointID(T,I,R). 
	% passing current cluster point, cluster list for checking,mutable cluster list, result 

intersect(POINT,[],OC,R):- 
	append(OC,[POINT],R). 
	% we compared all the point in the cluster list with the point
	% no same point id exists in the cluster list 
	% append the point into the overall cluster list 
intersect(0,_,OC,OC).
intersect(POINT,[H|T],OC,R):- 
    nth1(1,H,HID),
    nth1(1,POINT,PID),
    nth1(4,H,OCID),
    nth1(4,POINT,NCID), 
    HID =:= PID, 
    relabel(OCID,NCID,OC,B),% return back a new list with changed cluster id 
    intersect(0,T,B,R). % continue compare the point with point in the clusters list

intersect(POINT,[H|T],OC,R):- 
    nth1(1,H,HID),
    nth1(1,POINT,PID),
    HID =\= PID, % then we check next point in cluster list 
    intersect(POINT,T,OC,R). 

% cluster ID are different, check Point ID 
	% point ID are same, relabel 
	% point ID are different, append this to the list 

% Point ID are the same, relabel 
% Point ID are different, append to the list 

% ---------Relabel part---------------------
% relabel takes 4 argument 
% odd cluster ID, new cluster id, overall cluster list, result 
deleteLastElement([_], []).
deleteLastElement([Head, Next|Tail], [Head|NTail]):-
  deleteLastElement([Next|Tail], NTail).

relabel(OID,NID,[H|T],R):- 
    relabel(OID,NID,[H|T],[],R). 

relabel(_,_,[],I,I). 
    
% parameter:odd cluster id, new cluster id, cluster list, changable cluster list, result     
relabel(OID,NID,[H|T],I,R):- 
    nth1(4,H,CID), 
    CID =:= OID, 
    deleteLastElement(H,NEW),
    append(NEW,[NID],N), % create a new list
    append(I,[N],CL),
    relabel(OID,NID,T,CL,R). 

    
relabel(OID,NID,[H|T],I,R):- 
    nth1(4,H,CID), 
    CID =\= OID, 
    append(I,[H],CL),
    relabel(OID,NID,T,CL,R).
    
test(relabel):- 
    write('relabel(33, 77,
[[1,2.2,3.1,33], [2,2.1,3.1,22], [3,2.5,3.1,33], [4,2.1,4.1,33],[5,4.1,3.1,30]],Result)'),nl,
relabel(33, 77,
[[1,2.2,3.1,33], [2,2.1,3.1,22], [3,2.5,3.1,33], [4,2.1,4.1,33], [5,4.1,3.1,30]],Result),
write(Result).