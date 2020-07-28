package main

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestSetupBoggleBoard(t *testing.T) {
	// Given
	testBBString := "T, A, P, *, E, A, K, S, O, B, R, S, S, *, X, D"

	// When
	gotBB, err := SetupBoggleBoard(testBBString)

	// Then
	expectedBB := BoggleBoard{
		Grid: [4][4]string{
			{"T", "A", "P", "*"},
			{"E", "A", "K", "S"},
			{"O", "B", "R", "S"},
			{"S", "*", "X", "D"},
		},
	}

	require.Equal(t, expectedBB, gotBB)
	require.NoError(t, err)
}

func TestGetSelectableTiles(t *testing.T) {
	// Given
	givenBB := BoggleBoard{
		Grid: [4][4]string{
			{"T", "A", "P", "*"},
			{"E", "A", "K", "S"},
			{"O", "B", "R", "S"},
			{"S", "*", "X", "D"},
		},
	}

	givenTile := Tile{
		Row:    2,
		Col:    2,
		Letter: "R",
	}

	// When
	gotTiles := GetSelectableTiles(givenBB, givenTile, "D")

	// Then
	expectedTiles := []Tile{
		{3, 1, "*"},
		{3, 3, "D"},
	}

	require.Equal(t, expectedTiles, gotTiles)
}

func TestSelectTiles(t *testing.T) {
	// Given
	baseBB := BoggleBoard{
		Grid: [4][4]string{
			{"T", "A", "P", "*"},
			{"E", "A", "K", "S"},
			{"O", "B", "R", "S"},
			{"S", "*", "X", "D"},
		},
	}

	testCases := []struct {
		givenBB             BoggleBoard
		givenTiles          []Tile
		givenWord           string
		expectedBB          BoggleBoard
		expectedRemaingWord string
	}{
		{
			// start with exact first letter
			givenBB:    baseBB,
			givenTiles: []Tile{{0, 0, "T"}},
			givenWord:  "TARS",
			expectedBB: BoggleBoard{
				Grid: [4][4]string{
					{"!", "A", "P", "*"},
					{"E", "!", "K", "!"},
					{"O", "B", "!", "S"},
					{"S", "*", "X", "D"},
				}},
			expectedRemaingWord: "",
		},
		{
			// start with *
			givenBB:    baseBB,
			givenTiles: []Tile{{3, 1, "*"}},
			givenWord:  "GOAT",
			expectedBB: BoggleBoard{
				Grid: [4][4]string{
					{"!", "A", "P", "*"},
					{"E", "!", "K", "S"},
					{"!", "B", "R", "S"},
					{"S", "!", "X", "D"},
				}},
			expectedRemaingWord: "",
		},
		{
			// cannot boggle midway
			givenBB:    baseBB,
			givenTiles: []Tile{{0, 0, "T"}},
			givenWord:  "TAN",
			expectedBB: BoggleBoard{
				Grid: [4][4]string{
					{"!", "A", "P", "*"},
					{"E", "A", "K", "S"},
					{"O", "B", "R", "S"},
					{"S", "*", "X", "D"},
				}},
			// A has not been selected because there is no way to select N next
			expectedRemaingWord: "AN",
		},
		{
			// wrong givenTiles
			givenBB:    baseBB,
			givenTiles: []Tile{{1, 1, "T"}},
			givenWord:  "TAPE",
			expectedBB: BoggleBoard{
				Grid: [4][4]string{
					{"T", "A", "P", "*"},
					{"E", "A", "K", "S"},
					{"O", "B", "R", "S"},
					{"S", "*", "X", "D"},
				}},
			expectedRemaingWord: "TAPE",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotBB, gotRemainingWord := SelectTiles(tc.givenBB, tc.givenTiles, tc.givenWord)
			for i, expBBRow := range tc.expectedBB.Grid {
				if got, want := gotBB.Grid[i], expBBRow; !cmp.Equal(want, got) {
					t.Errorf("boggleboard mismatch (-want, +got):\n%s", cmp.Diff(want, got))
				}
			}
			require.Equal(t, tc.expectedRemaingWord, gotRemainingWord)
		})
	}
}

func TestIsWordInDictionary(t *testing.T) {
	// Given
	givenDictFilePath := "challenge/dictionary.txt"

	// When
	got, err := IsWordInDictionary(givenDictFilePath, "abandon")

	// Then
	require.True(t, got)
	require.NoError(t, err)
}

func TestGetStartingTiles(t *testing.T) {
	// Given
	givenBB := BoggleBoard{
		Grid: [4][4]string{
			{"T", "A", "P", "*"},
			{"E", "A", "K", "S"},
			{"O", "B", "R", "S"},
			{"S", "*", "X", "D"},
		},
	}

	// When
	gotTiles := GetStartingTiles(givenBB, "A")

	// Then
	expectedTiles := []Tile{
		{0, 1, "A"},
		{0, 3, "*"},
		{1, 1, "A"},
		{3, 1, "*"},
	}

	require.Equal(t, expectedTiles, gotTiles)
}
