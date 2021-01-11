package main

import (
	"fmt"
	"github.com/developdeveloper/password-generator/pkg"
)

func main() {
	pwdGen := pkg.New()
	fmt.Println(pwdGen.DigitsOnly(6))
}
