#lang racket
;Exercise1 Building list
(cons 3 (cons 4 '()))
(cons 1 (cons 2 (cons 3 '())))
; if it is char, it need ' at front 
(cons 'a (cons (cons 'b (cons 'c '())) '()))
(cons 1 '())
(cons 2(cons (cons 3 (cons (cons 4 '()) '())) '() ))

;Exercise2 List Entries
(define L '(1 2 3 4 5))

;Exercise3 Integer Range
(define (range i k)
  (cond
    ((> i k) 'range_error)
    ((= i k) (cons k '()))
    (else (cons i (range (+ i 1) k))
      )))

;Exercise4 Sum of Squares digit
; if the input number is empty return 0
; based on each digits square and plus them together 

(define (sosd n)
  (if (number? n) ;check the input is a number 
      (if (>= n 10) (+ ((lambda (x) (* x x)) (modulo n 10)) (sosd (quotient n 10)))
      ; (modulo n 10) --> get last digit of
          ;else 
      (* n n))
      ; when the number is a from 0 to 9, then just square the number 
      
      ;else, if the input is not a number, then we return a empty list 
      '()))

; Exercise5: every k-th element
; the function takes a list and an number selecting every k-th element. start counting at 1
;(define (drop l index)
  ;helper function 
  ;(define (drop l index counter result end)
    ;(if (eqv? index (- end 1)) result ;if the index reach to the end of list, we return the result
        ;else
        ;(if (eqv? index counter) (drop (rest l) index (+ counter 1) result end)
            ;else
            ;(drop l index (+ counter 1) (append result (list (list-ref l index))) end))))
    
 ;(drop l index 0 '() (length l)))
  

(define (drop L k)
  (if (not (list? L)) 'list-error
      (drophelper L k 1)))
(define (drophelper L k c)
  (cond
    ((null? L) '())
    ((= c k) (drophelper (cdr L) k 1))
    (#t (cons (car L) (drophelper (cdr L) k (+ c 1))))))



