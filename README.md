# Wordle Solver

Finds all possible words given letters that are known to be at certain positions
and letters that are known to not be present in the solution.

## Installation

```
# Clone the repo
$ git clone https://github.com/jstlwy/wordle_solver.git

# Change the working directory to wordle_solver
$ cd sherlock_go

# Build
$ go build wordle_solver.go
```

## Usage

```
$ ./wordle_solver -help
Usage of ./wordle_solver:
  -debug
    	Print the dictionary map to ensure that the text file was read correctly.
  -exclude string
    	List of letters that are known to not be in the word. Separate multiple with a comma. For example: -exclude m,s,e
  -file string
    	User-specified text file from which to read in English-language words. (default "freebsd_words.txt")
  -known string
    	List of known positions (zero-indexed) and letters. Separate multiple with a comma. For example: -known 0m,1o,2u
  -length int
    	The length of the word to be found. (default 5)
  -txt
    	Write the possible solutions to a .txt file.
  -verbose
    	Show how the user's arguments were interpreted.
```

Example:
```
./wordle_solver -known 1h,3m,4e -exclude o,u,s,b,l,a,c,i
```

## Credits

Text file sources:
- [freebsd_words.txt](https://svnweb.freebsd.org/csrg/share/dict/words?revision=61569&view=markup)
- [infochimp_words.txt](https://github.com/dwyl/english-words/blob/master/words.txt)
- [infochimp_words_alpha.txt](https://github.com/dwyl/english-words/blob/master/words_alpha.txt)
