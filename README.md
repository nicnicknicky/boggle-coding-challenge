## Boggle Coding Challenge

Boggle is a word game that is played on a 4x4 board with 16 letter tiles. 
The goal is to find as many words as possible given a time constraint.  
For this exercise, we are making one modification.  Now it is possible for one or more of the letter tiles to be blank (denoted by *).  
When a tile is blank, it can be treated as any other letter.  Note that in one game it does not have to be the same character for each word.  
For example, if the tiles C, T, and * are adjacent.  The words cot, cat, and cut can all be used. 
You will be given a text file containing all valid English words (a dictionary).
You will also be given an initial board configuration as a text file with commas separating the letters. Use this as a guide for how 
To set up the board

For example a file may contain:

A, C, E, D, L, U, G, *, E, *, H, T, G, A, F, K

This is equivalent to the board:

A C E D
L U G *
E * H T
G A F K

Some sample words from this board are ace, dug, eight, hole, huge, hug, tide.

## Setup

In `.env` configure paths for the required files:
```
BOGGLEBOARD_PATH="challenge/TestBoard.txt"
DICTIONARY_PATH="challenge/dictionary.txt"
```

## How to Play

`!` symbolises the letters chosen on the base boggleboard to obtain your input word.

```
❯ go run main.go "TAPS"
(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧ BOGGLE ✧ﾟ･: *ヽ(◕ヮ◕ヽ)
[! ! ! !]
[E A K S]
[O B R S]
[S * X D]

❯ go run main.go "GOAT"
(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧ BOGGLE ✧ﾟ･: *ヽ(◕ヮ◕ヽ)
[! A P *]
[E ! K S]
[! B R S]
[S ! X D]
```