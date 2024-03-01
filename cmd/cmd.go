package cmd

import (
	"GoMage/cmd/actions"
	"fmt"
	"github.com/spf13/cobra"
)

type Command struct {
	rootCmd *cobra.Command
}

var text = `
 ________          _____                         
 /  _____/  ____   /     \ _____     ____   ____  
/   \  ___ /  _ \ /  \ /  \\__  \   / ___\_/ __ \ 
\    \_\  (  <_> )    Y    \/ __ \_/ /_/  >  ___/ 
 \______  /\____/\____|__  (____  /\___  / \___  >
        \/               \/     \//_____/      \/ `

func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:   "GoMage",
		Short: "Go Image manipulation command line.",
		Long:  "Go Image manipulation command line.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Show display text
			fmt.Println(fmt.Sprintf(text))
		},
	}

	return &Command{
		rootCmd: rootCmd,
	}
}

// Run the all command line
func (c *Command) Run() {
	var (
		temp                    float64
		inputImage, outputImage string
	)

	var temperatureCommand = &cobra.Command{
		Use:   "temperature",
		Short: "adjust image temperature.",
		Long:  "adjust image temperature.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			actions.AdjustTemperature(inputImage, outputImage, temp)
		},
	}

	temperatureCommand.Flags().StringVarP(&inputImage, "input image", "i", "", "Input image file.")
	temperatureCommand.Flags().StringVarP(&outputImage, "output image", "o", "", "Output image file.")
	temperatureCommand.Flags().Float64VarP(&temp, "temperature", "t", 0, "Temperature, use positive number for warmer and negative number for cooler.")

	c.rootCmd.AddCommand(temperatureCommand)
	c.rootCmd.SuggestionsMinimumDistance = 1
	c.rootCmd.Execute()
}
