package cartesian

import (
	"testing"
	"fmt"
)

func TestCartesianProduct(t *testing.T) {
	var l1 = []interface{}{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
	var l2 = []interface{}{"♠", "♥", "♦", "♣"}

	var p = [][]interface{}{l1, l2}

	fmt.Println(CartesianProduct(p))
}
