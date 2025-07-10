package generate

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var CmdGenerate *cobra.Command

func Init(cmdGenerate *cobra.Command) {
	cmdGenerate.Flags().IntP("length", "l", 8, "Length of the generated Password")
	cmdGenerate.Flags().BoolP("digits", "d", false, "Include digits in the generated password")
	cmdGenerate.Flags().BoolP("special-chars", "s", false, "Include special chars in generated password")
}

func GeneratePassword(cmd *cobra.Command, args []string) {
	fmt.Println("Generating Password")

	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if isDigits {
		charset += "0123456789"
	}

	if isSpecialChars {
		charset += "!@#$%^&*()_+{}[]|;:,.<>?-="
	}

	var password string

	for length > 0 {
		password += string(charset[rand.Intn(len(charset))])
		length--
	}
	
	fmt.Println(password)
}