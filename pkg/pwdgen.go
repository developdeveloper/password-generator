package pkg

import (
	"bytes"
	"github.com/developdeveloper/password-generator/internal"
)

type PasswordGenerator struct{}

//New Create password generator
func New() *PasswordGenerator {
	return &PasswordGenerator{}
}

//DigitsOnly generate password only contains digits
func (pg *PasswordGenerator) DigitsOnly(length int) string {
	var buf bytes.Buffer
	digits := "0123456789"

	for i := 0; i < length; i++ {
		buf.WriteByte(digits[internal.RandNumber(10)])
	}

	return buf.String()
}

//CharsOnly generate password only contains chars
func (pg *PasswordGenerator) CharsOnly(length int) string {
	var buf bytes.Buffer
	chars := "abcdefghijklmnopkrstuvwxyz"

	for i := 0; i < length; i++ {
		buf.WriteByte(chars[internal.RandNumber(26)])
	}

	return buf.String()
}
