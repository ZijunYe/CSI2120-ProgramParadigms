#lang racket
;Exercise1
;Calculate the sum of the squares for the numbers 1 to 4. Hint: Your result should be 30.
(+ (* 1 1) (* 2 2) (* 3 3) (* 4 4))

;Calculate both solutions to 2x^2 + 5x - 3 = 0 using function equation
; d = b^2 - 4ac
;( -b + / - sqrt(d) )/2a

; a = 2 , b = 5 , c= -3
(/ (+ (- 5)(sqrt (- (* 5 5) (* 4 2 (- 3))))) (* 2 2))
(/ (- (- 5)(sqrt (- (* 5 5) (* 4 2 (- 3))))) (* 2 2))



;Calculate the result of the sin(pi/4) cos(pi/3) + cos(pi/4) sin(pi/3)
(+ (*(sin (/ pi 4)) (cos (/ pi 3))) (*(cos (/ pi 4)) (sin (/ pi 3))) )

;Exercise2 Local and Global definition 
;global definition use define and local definition use let binding
;define a global constant (define my favorite 43) --> using constant
;define locals with let (let ((x 1)(y 2)) (+ x y)) 
(let ((x (/ pi 4)) (y (/ pi 3))) (+ (*(sin x) (cos y)) (*(cos x) (sin y)) ) )

;Exercise3 Procedures
; lambda expressions create local procedures
; example: ((lambda (x y) (+ x y)) 1 2) , they code will take argument 1 and 2
; (define foo (lambda (x y) (+ x y))) define global procedures

; Use a lambda with the angles as arguments to calculate sin(pi/4) cos(pi/3) + cos(pi/4) sin(pi/3) again.
((lambda (x y) (+ (*(sin x) (cos y)) (*(cos x) (sin y)) )) (/ pi 4)(/ pi 3))

;Define a global function to calculate sin(alpha) cos(beta) + cos(alpha) sin(beta) with alpha and beta as parameters.
(define f (lambda(x y) (+ (*(sin x) (cos y)) (*(cos x) (sin y)) )) )
(f (/ pi 4) (/ pi 3)) ;making function call


;Exercise4 Quadratic equation
; return them as a list example:
; (define foo (lambda ( x y)(list (+ x y)(- x y))))
;using the function equation b^2 - 4ac = d 
; (-b +/- d )/ 2a 
(define q (lambda(a b c)(let ((d (sqrt(- (* b b)(* 4 a c)))   ))
            (list (/ (+ (- b) d) (* a 2))
                  (/ (- (- b) d) (* a 2))))))
(q 2 5 -3) 
            


    