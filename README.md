# Wordle Solver

Finds all possible words given letters that are known to be in the solution
(positions can be specified or unspecified)
and letters that are known to not be in the solution.

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
  -dict string
    	User-specified text file from which to read in English-language words. (default "freebsd_words.txt")
  -exclude string
    	List of letters that are known to not be in the word. Separate multiple with a comma. For example: -exclude m,s,e
  -include string
    	List of letters that are known to be in the word but whose positions are unknown. Separate multiple with a comma. For example: -include m,s,e
  -known string
    	List of known positions and letters. Separate multiple with a comma. For example: -known 1m,2o,3u
  -length int
    	The length of the word to be found. (default 5)
  -save
    	Save the potential solutions in a .txt file.
  -verbose
    	Show how the user's arguments were interpreted.
```

Example:
```
./wordle_solver -exclude m,o,a,c -include u -known 1s,5e -dict infochimp_words_alpha.txt
```

## Credits

Text file sources:
- [freebsd_words.txt](https://svnweb.freebsd.org/csrg/share/dict/words?revision=61569&view=markup)
- [infochimp_words.txt](https://github.com/dwyl/english-words/blob/master/words.txt)
- [infochimp_words_alpha.txt](https://github.com/dwyl/english-words/blob/master/words_alpha.txt)
