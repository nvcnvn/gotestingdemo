package quickdemo

import (
	"github.com/nvcnvn/gotestingdemo"
	"testing"
	"testing/quick"
)

func TestBlackBox(t *testing.T) {
	f := func(i int) bool {
		e := gotestingdemo.NextEvenNotMultipleByFive(i)
		return e%2 == 0 && e%5 != 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
