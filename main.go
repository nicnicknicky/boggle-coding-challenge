package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// BoggleBoard ...
type BoggleBoard struct {
	Grid [4][4]string
}

// SetupBoggleBoard ...
func SetupBoggleBoard(bbString string) (BoggleBoard, error) {
	bb := BoggleBoard{}

	// boggle board has a 4x4 grid
	bbLetters := strings.Split(bbString, ",")
	if len(bbLetters) != 16 {
		return bb, errors.New("invalid number of boggleBoard elements")
	}

	for idx, letter := range bbLetters {
		row := idx / 4
		col := idx % 4

		bb.Grid[row][col] = strings.TrimSpace(letter)
	}

	return bb, nil
}

// Tile ...
type Tile struct {
	Row    int
	Col    int
	Letter string
}

// GetSelectableTiles ...
func GetSelectableTiles(bb BoggleBoard, subjectTile Tile, targetLetter string) []Tile {
	var selectableTiles []Tile

	possibleAdjRows := []int{subjectTile.Row - 1, subjectTile.Row, subjectTile.Row + 1}
	for _, r := range possibleAdjRows {
		if r >= 0 && r <= 3 {
			for c, l := range bb.Grid[r] {
				switch c {
				case subjectTile.Col:
					// exclude the subjectTile
					if r == subjectTile.Row {
						continue
					}
					fallthrough
				case subjectTile.Col - 1, subjectTile.Col + 1:
					// exclude failed tiles
					if l == "!" {
						continue
					}
					if l == strings.ToUpper(targetLetter) || l == "*" {
						selectableTiles = append(selectableTiles, Tile{r, c, l})
					}
				}
			}
		}
	}

	return selectableTiles
}

// SelectTiles ...
func SelectTiles(bb BoggleBoard, selectableTiles []Tile, remainingWord string) (BoggleBoard, string) {
	var nextSelectableTiles []Tile
	nextBB := bb
	// done: no remaining choices / no remaining processing
	if len(selectableTiles) == 0 || len(remainingWord) == 0 {
		return bb, remainingWord
	}
	for _, tile := range selectableTiles {
		// checks: tile data matches boggleboard, tile letter matches currently processed rune or special char *
		if tile.Letter == bb.Grid[tile.Row][tile.Col] && (tile.Letter == string([]rune(remainingWord)[0]) || tile.Letter == "*") {
			maybeNextBB := nextBB
			maybeNextBB.Grid[tile.Row][tile.Col] = "!" // prevents reverse selection
			// prevent index out of range
			if len(remainingWord) > 1 {
				// check if the next letter has selectable tiles, otherwise current selection is a dead end
				maybeNextSelectableTiles := GetSelectableTiles(maybeNextBB, tile, string([]rune(remainingWord)[1]))
				if len(maybeNextSelectableTiles) == 0 {
					continue
				}
				// has selectable tiles, can 'commit'
				nextSelectableTiles = maybeNextSelectableTiles
				remainingWord = remainingWord[1:]
				nextBB = maybeNextBB
			} else {
				// last letter selected
				remainingWord = ""
				nextBB = maybeNextBB
				break
			}
		}
	}
	return SelectTiles(nextBB, nextSelectableTiles, remainingWord)
}

// IsWordInDictionary ...
func IsWordInDictionary(dictFilePath string, word string) (bool, error) {
	dictFile, err := os.Open(dictFilePath)
	if err != nil {
		return false, err
	}

	// scan by line
	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		if scanner.Text() == word {
			return true, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

// GetStartingTiles ...
func GetStartingTiles(bb BoggleBoard, letter string) []Tile {
	var tiles []Tile

	for r, bbRow := range bb.Grid {
		for c, l := range bbRow {
			if l == letter || l == "*" {
				tiles = append(tiles, Tile{r, c, l})
			}
		}
	}

	return tiles
}

func main() {
	// ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	boggleBoardFilePath := os.Getenv("BOGGLEBOARD_PATH")
	dictionaryFilePath := os.Getenv("DICTIONARY_PATH")
	targetWord := os.Args[1]
	targetWordUpper := strings.ToUpper(targetWord)

	// check valid word
	readyToBoggle, err := IsWordInDictionary(dictionaryFilePath, strings.ToLower(targetWord))
	if err != nil {
		log.Fatal(err.Error())
	}
	if !readyToBoggle {
		fmt.Printf("%s is not a valid word in the dictionary, can't play boggle\n", targetWord)
		os.Exit(0)
	}

	// setup boggle board
	bbData, err := ioutil.ReadFile(boggleBoardFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	boggleBoard, err := SetupBoggleBoard(string(bbData))
	if err != nil {
		log.Fatal(err.Error())
	}

	// play boggle
	resultBoggleBoard := boggleBoard
	resultWord := targetWordUpper
	startingTiles := GetStartingTiles(boggleBoard, string([]rune(targetWordUpper)[0]))
	// startingTile must be passed in one at a time to allow 'reset' of the boggleboard
	for _, startingTile := range startingTiles {
		bb, word := SelectTiles(boggleBoard, []Tile{startingTile}, targetWordUpper)
		if len(word) == 0 {
			resultBoggleBoard = bb
			resultWord = word
			break
		}
	}

	if len(resultWord) == 0 {
		fmt.Println("(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧ BOGGLE ✧ﾟ･: *ヽ(◕ヮ◕ヽ)")
		for _, row := range resultBoggleBoard.Grid {
			fmt.Println(row)
		}
		os.Exit(0)
	}

	fmt.Println("(╯°□°）╯︵ ┻━┻ NO BOGGLE")
}
