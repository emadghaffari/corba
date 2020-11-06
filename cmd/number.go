package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var (
	min int
	max int
)

func init() {
	rootCmd.AddCommand(numberCmd)
	numberCmd.Flags().IntVarP(&min, "min", "", 1, "min for random int")
	numberCmd.Flags().IntVarP(&max, "max", "", 1, "max for random int")
	numberCmd.MarkFlagRequired("min")
	numberCmd.MarkFlagRequired("max")
}

var numberCmd = &cobra.Command{
	Use:   "random",
	Short: "generate a random of a length",
	Run: func(cmd *cobra.Command, args []string) {
		if min < max {
			rand.Seed(time.Now().UnixNano())

			fmt.Printf("%d \n", rand.Intn(max-min+1)+min)
		} else {
			fmt.Printf("min is bigger than max number min: %d and max: %d \n", min, max)
		}

	},
}
