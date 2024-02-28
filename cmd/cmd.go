package cmd

import (
	"GoMage/cmd/actions"
	"fmt"
	"github.com/spf13/cobra"
)

type Command struct {
	rootCmd *cobra.Command
}

func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:   "GoMage",
		Short: "Go Image manipulation command line.",
		Long:  "Go Image manipulation command line.",
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

	var rootCmd = &cobra.Command{

		Use:   "temp-adjust",
		Short: "Adjust image temperature, use positive number for warmer and negative number for cooler.",
		Long:  "Adjust image temperature, use positive number for warmer and negative number for cooler.",

		PreRun: func(cmd *cobra.Command, args []string) {
			//logger.WithFields(logger.Fields{"component": "command", "action": fmt.Sprintf("adjusting image %s with temp value %d", inputImage, temp)}).
			//	Infof("PreRun command done")

			fmt.Println(inputImage, outputImage, temp)

			actions.AdjustTemperature(inputImage, outputImage, temp)
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
		PostRun: func(cmd *cobra.Command, args []string) {

		},
	}

	rootCmd.Flags().StringVarP(&inputImage, "input image", "i", "", "Input image file.")
	rootCmd.Flags().StringVarP(&outputImage, "output image", "o", "", "Output image file.")
	rootCmd.Flags().Float64VarP(&temp, "temperature", "t", 0, "Temperature, use positive number for warmer and negative number for cooler.")
	rootCmd.Execute()
}
