package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImportData(t *testing.T) {
	// fixture
	testFilePath := "../data/riffle_entries_test.json"
	expected := []raffleEntry{
		{"001", "Alice"},
		{"002", "Bob"},
	}

	// test
	entries, err := importData(testFilePath)

	// asserts
	require.NoError(t, err)
	require.NotNil(t, entries)
	require.True(t, reflect.DeepEqual(entries, expected))
}

func TestImportData_fileDoesNotExists(t *testing.T) {
	// fixture
	testFilePath := "../data/riffle_entries_test_fail.json"

	// test
	entries, err := importData(testFilePath)

	// asserts
	require.Error(t, err)
	require.NotNil(t, err)
	require.Nil(t, entries)

}

func TestGetWinner(t *testing.T) {
	// fixture
	entries := []raffleEntry{
		{"001", "Alice"},
		{"002", "Bob"},
		{"003", "Charlie"},
	}

	// Run the test multiple times to increase the chance of catching a random failure
	for i := 0; i < 100; i++ {
		winner := getWinner(entries)
		if winner.Name == "" || winner.Id == "" {
			t.Errorf("getWinner(%v) returned invalid winner: %v", entries, winner)
		}
	}
}
