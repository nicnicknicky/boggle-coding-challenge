package main

import (
	"testing"

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
	givenBB := BoggleBoard{
		Grid: [4][4]string{
			{"T", "A", "P", "*"},
			{"E", "A", "K", "S"},
			{"O", "B", "R", "S"},
			{"S", "*", "X", "D"},
		},
	}

	givenTiles := []Tile{
		{0, 0, "T"},
	}

	givenWord := "TARS"

	// When
	expectedBB := BoggleBoard{
		Grid: [4][4]string{
			{"!", "A", "P", "*"},
			{"E", "!", "K", "!"},
			{"O", "B", "!", "S"},
			{"S", "*", "X", "D"},
		},
	}
	gotBB, gotRemainingWord := SelectTiles(givenBB, givenTiles, givenWord)
	// Then
	require.Equal(t, expectedBB, gotBB)
	require.Zero(t, gotRemainingWord)
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
