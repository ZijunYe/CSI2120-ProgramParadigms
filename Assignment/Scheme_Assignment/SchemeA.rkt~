#lang racket

; ---------Question1--------------
; a function lowest-exponent
; need 1 local variable to keep track the exponent
; using loop for increase the exponent     
(define (lowest-exponent x y)
  ;; helper
  (define (lowest-exponent e x)
    (cond ((>= (expt x e) y) e)
          (else (lowest-exponent (+ e 1) x))))
  ;; call helper
  (lowest-exponent 0 x))


; -----------Question2------------
; find abundant taking 1 positive argument(AS a bound)
; return a list all abundant number
; looping through all the number smaller than the bound
; x check 2*n
; y check all sum of divisors of n
; if x > y, input n to the list
; list has odered from largest to the smallest 

(define (find-abundant x)
  ;; helper
  
  (define (find-abundant v l)
    (if (eqv? v 0) l
     ;else
        (if (> (findDivisorSum v) (* 2 v))  (find-abundant(- v 1) (append l (list v)))
            (find-abundant (- v 1) l))))
  
  ;; call helper
  (find-abundant (- x 1) '() ))

; helper function to do calculates sum of all divisor, --> number mod d is 0
(define (sumDivisors n i)
  (cond ((= 1 i) 1)
        ((= (remainder n i) 0)
         (+ i (sumDivisors n (- i 1))))
        (else (sumDivisors n (- i 1)))))

(define (findDivisorSum n)
  (sumDivisors n n))


; --------------Question3----------------
; take a n natural number
; return list of strings


; 3 conditions
; if the number = 0 , append "Finished", return the list
; if the number = 1, append "second"
; if the number > 1, append "seconds" 

(define (make-string-list x)
  ;; helper
  
  (define (make-string-list v l)
    (cond ((eqv? v 0) (append l '("Finished")))
          ((eqv? v 1) (make-string-list (- v 1) (append l (list (string-append (number->string v) " second")))))
          (else (make-string-list (- v 1)(append l (list (string-append (number->string v) " seconds")))))))
  
  ;; call helper
  (make-string-list x '() ))


