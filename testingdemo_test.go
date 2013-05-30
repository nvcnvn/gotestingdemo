package gotestingdemo

import (
	"fmt"
	"testing"
)

func TestNextEvenNotMultipleByFive(t *testing.T) {
	if i := NextEvenNotMultipleByFive(9); i != 12 {
		t.Error("Expect 12 got", i)
	}
}

func ExampleNextEvenNotMultipleByFive() {
	fmt.Println(NextEvenNotMultipleByFive(4))
	fmt.Println(NextEvenNotMultipleByFive(7))
	fmt.Println(NextEvenNotMultipleByFive(9))
	// Output:
	// 6
	// 8
	// 12
}
