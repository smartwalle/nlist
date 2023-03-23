package nlist_test

import (
	"github.com/smartwalle/nlist"
	"testing"
)

func TestCombination(t *testing.T) {
	var l1 = []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
	var l2 = []string{"♠", "♥", "♦", "♣"}

	var p = [][]string{l1, l2}

	t.Log(nlist.Combination[string](p))
}

func TestCombination2(t *testing.T) {
	var l1 = []string{"1", "2", "3", "4"}
	var l2 = []string{"A", "B", "C", "D"}
	var l3 = []string{"★", "☆"}

	var p = [][]string{l1, l2, l3}

	t.Log(nlist.Combination[string](p))
}
