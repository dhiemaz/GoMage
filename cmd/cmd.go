package cmd

import (
	"GoMage/internal/logger"
	"fmt"
	"github.com/spf13/cobra"
)

type Command struct {
	rootCmd *cobra.Command
}

var text = `GoMage`

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
		temp      uint32
		imagePath string
	)

	var rootCommands = []*cobra.Command{
		{
			Use:   "temp-adjust",
			Short: "Adjust image temperature, use positive number for warmer and negative number for cooler.",
			Long:  "Adjust image temperature, use positive number for warmer and negative number for cooler.",
			PreRun: func(cmd *cobra.Command, args []string) {

				// Show display text
				fmt.Println(fmt.Sprintf(text))

				logger.WithFields(logger.Fields{"component": "command", "action": "adjust image temp"}).
					Infof("PreRun command done")
			},
			Run: func(cmd *cobra.Command, args []string) {

			},
			PostRun: func(cmd *cobra.Command, args []string) {

			},
		},
	}

	for _, command := range rootCommands {
		c.rootCmd.SetHelpCommand(command)
		c.rootCmd.AddCommand(command)
	}

	c.rootCmd.Flags().StringVarP(&imagePath, "image path", "u", "", "Image path.")
	c.rootCmd.Flags().Uint32VarP(&temp, "temperature", "t", 0, "Temperature, use positive number for warmer and negative number for cooler.")
	c.rootCmd.Execute()
}

// GetRoot the command line service
func (c *Command) GetRoot() *cobra.Command {
	return c.rootCmd
}
