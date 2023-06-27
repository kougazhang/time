package time

import (
	"fmt"
	"testing"
)

func TestUnitDigitOfMinute(t *testing.T) {
	for i := 0; i < 60; i++ {
		fmt.Println(i, unitDigit(i))
	}
}
