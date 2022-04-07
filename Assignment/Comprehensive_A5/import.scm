#lang racket
(define (readlist filename)
 (call-with-input-file filename
  (lambda (in)
    (read in))))
(define (import)
  (let ((p65 (readlist "partition65.scm"))
        (p74 (readlist "partition74.scm")) 
        (p75 (readlist "partition74.scm")))
    (append p65 p74 p75)))

; get global list (import does that?)
; go to first elem in the global list
; findout their cluster id, then do extractCluster
; check if the result list is empty
    ; if the list is empty, append entire current cluster element to the cluster list, append current cluster id into cluster id collector 
    ; if the list is not empty
        ; if cluster id is in the cluster id collector, skip the cluster, extract another cluster from global list
        ; if cluster id is not in the cluster id collector, add to the cluster id
             ; go to the first point in the current cluster
                 ;check intersection yes --> relabel, no--> continue check till end of the cluster list, then append 


; length of global list (keep tracking we loop all the point)
; cluster id collector (to check if we see the cluster or not)
; result list --> clusters list
(define (mergeClusters L) ;input a global list
                      
   (define (mergeClusters gl counter collector cl)                    
         ;condition1: if the counter reach to zero
         (cond ((eqv? counter 0) cl)
               ; condition2: still has point have loop yet
               ; cluster id: (list-ref (first gl) 5)
               ; find out their cluster extractCluster((list-ref (first gl) 5) gl) 
               (else
                (if (eqv? (length collector) 0) (mergeClusters (rest gl) (- counter 1) (cons (list-ref (first gl) 4) collector)  (extractCluster (list-ref (first gl) 4) gl))
                    ;else
                    (if (eqv? (member (list-ref (first gl) 4) collector ) #f) ;current cluster is not in the list
                        ; add the cluster id to the collector: (cons (list-ref (first gl) 5)) collector)
                        ; counter decrease by 1 : (- counter 1)
                        ; get the current cluster by it id: extractCluster((list-ref (first gl) 5) gl)
                        ; check if it has intersection --> here will return me back a new list
                        (mergeClusters (rest gl) (- counter 1) (cons (list-ref (first gl) 4) collector) (intersection (extractCluster(list-ref (first gl) 4) gl) cl))
                        ; else current cluster already checked, so we just skip the cluster
                        (mergeClusters (rest gl) (- counter 1) collector cl)) )
                )))
               
            
   (mergeClusters L (length L) '() '()))


; ---------------- extract cluster from main list and remove the first point --------
; the main list comes from the import

(define (extractCluster cid l)
  ; go through the list
  ; if the cid is same
  ; append to the list
  ; if counter reach the end
  ; return the list

  ;;helper function
  (define (extractCluster cid l result counter)
    (if (eqv? counter 0) result
        ;else 
        (if (eqv? cid (list-ref (first l) 4)) (extractCluster cid (rest l) (append result (list(rest(first l)))) (- counter 1))
            ; else
            (extractCluster cid (rest l) result (- counter 1)))))
    
 (extractCluster cid l '() (length l))) 


; ----------------------------------------------------- CHECKINTERSECTION ----------------------------------------------------------------------
; passing cluster, current overall cluster list
; return list of cluster ID that have a point in common

(define (intersection cc cl)
  (define (intersection cc cl counter)
    ; current cluster, clusters list, result list, inside counter, outside counter
    ; condition1: if the counter = 0, then return the new cluster list
    ; condition2: counter != 0, then calling for next point in the current cluster
    (cond ((eqv? counter 0) cl)
          (else
           (intersection (rest cc) (checkPointID (first cc) cl) (- counter 1)))))
(intersection cc cl (length cc)))


;checkPointID --> return a new cluster list 

; compare 1 point with the cluster list
  ; passing: point & clusters list 
(define (checkPointID point cl)
  
  (define (checkPointID p l same cl counter)
    ; condition1: if the counter = 0, return result 
    ; condition2: if the pointID is same as one point in the result, append to the result
    (cond ((eqv? counter 0)
           (cond ((eqv? (length same) 0) (append cl (list p))) ;return a new cluster list with point append
               ;else
               (else (relabeling (list-ref p 3) same cl)   )));it will return back a new clusters list

          
          (else
              (if (eqv? (list-ref p 0) (list-ref (first l) 0)) (checkPointID p (rest l) (cons (list-ref (first l) 3) same ) cl (- counter 1))
                  ;else
                  (checkPointID p (rest l) same cl (- counter 1)))) ))


  (checkPointID point cl '() cl  (length cl))) 




; ----------------------------------------------------------RELABEL PART----------------------------------------------------------------------------------
;relabeling
;passing new cluster id, list cluster id intersect, overall cluster list


 (define (relabeling nid li cl)
   (define (relabeling nid li cl counter)
     ;condition1: counter=0 return result
     ; else passing each one to the relabel
     (cond ((eqv? counter 0) cl)
           (else (relabeling nid (rest li) (relabel (first li) nid  cl) (- counter 1)))))
      
   (relabeling nid li cl  (length li))) 

; need one list contain the all the cluster ID has same point
;looping through the cluster
  ; then looping through the overall cluster list


    ; if the counter for the cluster list is not zero 
             ; ------- compare point ID -------   
             ; if the point ID are same, append the cluster id, continue looping move to the next point
             ; if the point ID not same, move on to the next point in the cluster list
    ; else the counter reach to the zero, passing the next point in current cluster

; --------- Relabel part ----------
; parameter: new Cluster ID, odd Cluster ID, overall cluster list

; condition1: if the Cluster ID match with odd Cluster ID, construct a new list and append to the overall cluster list 
; condition2: if the Cluster ID doesn't match with odd cluster ID, just append the list
; condition3: counter = (length l), if the counter reach to 0, then return the final list 
; how do you know the list already done

(define (relabel o n l)
  ;; helper
  
  (define (relabel o n l nl c)
    (cond ((eqv? c 0) nl) ; we already loop through the all the element in the overall list 
          ((eqv? (list-ref (first l) 3) o) ; Cluster id are matched, then we have to create a new list
           (relabel o n (rest l) (append nl (list (list (first (first l)) (list-ref (first l) 1) (list-ref (first l) 2) n)) ) (- c 1)) )
           ; new list and append to the n l 
          (else (relabel o n (rest l) (append nl (list (first l)) ) (- c 1))))) ; no need to be change just append the original list into the nl
 
  ;; call helper
  (relabel o n l '() (length l)))