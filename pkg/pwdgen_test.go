package pkg

import (
	"fmt"
	"testing"
)

func TestPasswordGenerator_DigitsOnly(t *testing.T) {
	pwdGen := New()
	fmt.Println(pwdGen.DigitsOnly(6))
}

func TestPasswordGenerator_CharsOnly(t *testing.T) {
	pwdGen := New()
	fmt.Println(pwdGen.CharsOnly(6))
}
