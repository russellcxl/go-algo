package race_conditions

import (
	"fmt"
	"testing"
)

func TestRace(t *testing.T) {
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(race())
	}

}