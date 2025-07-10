package cmd

import (
	"fmt"
	"math/rand"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "Generate random passwords",
	Long: `Generate random passwords with customizable options.
		For example:
		dcxcli generate -l 12 -d -s`,
	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("length", "l", 8, "Length of the generated Password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in the generated password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special chars in generated password")
}

func generatePassword(cmd *cobra.Command, args []string) {
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
	log.Info("Generating password")
	// log.Info(("Generating password")
	// fmt.Println("Generating password")
	fmt.Println(password)
}