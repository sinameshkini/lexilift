package core

import (
	"fmt"
	"testing"
)

func Test_proficiencyScore(t *testing.T) {

	for i := -20; i <= 20; i++ {
		fmt.Println(i, " = ", proficiencyScore(i))
	}

}

func Test_timeScore(t *testing.T) {

	//for i := 0; i <= 13; i++ {
	//	fmt.Println(i, " = ", timeScore(i, 10))
	//}

	for i := 0; i <= 25; i++ {
		fmt.Println(i, " = ", timeScore(i, 20))
	}

}
