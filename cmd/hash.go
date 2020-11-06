package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var (
	lenght  int
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-@%$&()_")
)

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.Flags().IntVarP(&lenght, "lenght", "l", 1, "lenght for random hash string")
	hashCmd.MarkFlagRequired("lenght")
}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "generate a string of a fixed length",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		b := make([]rune, lenght)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		fmt.Printf("%s \n", string(b))
	},
}
