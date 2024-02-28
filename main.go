package main

import command "GoMage/cmd"

func main() {
	// This is the entry point for the application
	// Create a new command
	cmd := command.NewCommand()

	// Run the command
	cmd.Run()
}
