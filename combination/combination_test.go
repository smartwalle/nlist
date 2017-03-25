package combination

import (
	"testing"
	"fmt"
)

func TestCombination(t *testing.T) {
	//var l1 = []interface{}{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
	//var l2 = []interface{}{"♠", "♥", "♦", "♣"}
	var l1 = []interface{}{1, 2, 3}
	var l2 = []interface{}{"A", "B", "C", "D"}

	var p = [][]interface{}{l1, l2}

	fmt.Println(Combination(p))
}
