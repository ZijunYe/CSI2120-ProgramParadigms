# Character 
- character constants: 
``` #\a #\A #\( #\space #\newline```

- predicates for character 
```
(char? obj)
(char-alphabetic? char)
(char-numeric? char)
(char-whitespace? char)
(char-upper-case? char)
(char-lower-case? char)
```

- comparison between two characters 
```
(char=? char1 char2)
(char<? char1 char2)
(char>? char1 char2)
(char<=? char1 char2)
(char>=? char1 char2)
```

- corresponding case insensitive 
```
(char=? #\a #\A)
=> #f 

(char-ci=? #\a #\A)
=> #t 
```

- character conversion 
- character to ascii ```(char->integer #\a)```
- ascii to character ```(integer->char 97)```

# String 
- string constants written in double quotation marks ``` "Hello" ```
- comparison between string 
```
(string=? string1 string2)
(string<? string1 string2)
(string>? string1 string2)
(string<=? string1 string2)
(string>=? string1 string2)
(string-ci=? string1 string2) ; insensitive 
```

- more string functions 
```
(string-length "Hello")
(string->list "Hello")
(substring "computer" 3 6)
```

-caesar cipher: ```((charNum - base) +shiftNum) mod totalNum + base ```

# Input & Output 
- How to print something in the prolog 
```(display "hello world")```
- How to read from terminal 
```(read)``` then it will pop up the window to type, it will return exactly the same value as you type in 
- combine read and write 
```(display (read))```

- for formatted output 
```newline``` 

- example 
** number entry ** 
```
(define (ask-number) 
	(display "Enter a number:")
	(let ((n (read)) 
		(if (number? n) n (ask-number))))
)

==> (ask-number)  
```


**Reading a file into a list** 
```
(let ((p (open-input-file "short.scm")))
	(let f ((x (read p))) 
		(if (eof-object? x)
		 (begin 
		 	(close-input-port p)
		 	'())
		 (cons x (f (read p)))
		)
	)
)
```





