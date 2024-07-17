package cli

import (
	"fmt"
    "os"

    "github.com/spf13/cobra"

)

// cobra //نصب
//go get github.com/spf13/cobra
func TestCobra() {
	fmt.Println("----------------------------------------Start Cobra----------------------------------------")
	fmt.Println("sample 1 : go run main.go upload -f myfile.txt -s example.org -u myuser -p mypassword")
	fmt.Println("sample 1 : Uploading file 'myfile.txt' to server 'example.org' with username 'myuser' and password 'mypassword'")
	fmt.Println("sample 1 : go run main.go upload -f myfile.txt")
	fmt.Println("sample 1 : Uploading file 'myfile.txt' to server 'example.com' with username '' and password ''")
	var rootCmd = &cobra.Command{
        Use:   "myapp",
        Short: "A simple CLI application",
        Long:  `This is a sample CLI application built using the Cobra library in Go.`,
        Run: func(cmd *cobra.Command, args []string) {
            // This is the default action when no subcommands are specified
            fmt.Println("Welcome to the myapp CLI!")
        },
    }

    // Create a subcommand with flags
    var uploadCmd = &cobra.Command{
        Use:   "upload",
        Short: "Upload a file",
        Long:  `This command will upload a file to a remote server.`,
        Run: func(cmd *cobra.Command, args []string) {
            // Get the values of the flags
            filename, _ := cmd.Flags().GetString("filename")
            server, _ := cmd.Flags().GetString("server")
            username, _ := cmd.Flags().GetString("username")
            password, _ := cmd.Flags().GetString("password")

            // Perform the upload logic
            fmt.Printf("Uploading file '%s' to server '%s' with username '%s' and password '%s'\n",
                filename, server, username, password)
        },
    }

    // Define the flags for the upload command
    uploadCmd.Flags().StringP("filename", "f", "", "Name of the file to upload")
    uploadCmd.Flags().StringP("server", "s", "example.com", "Server to upload the file to")
    uploadCmd.Flags().StringP("username", "u", "", "Username for the server")
    uploadCmd.Flags().StringP("password", "p", "", "Password for the server")

    // Add the subcommand to the root command
    rootCmd.AddCommand(uploadCmd)

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	fmt.Println("----------------------------------------End Cobra  ----------------------------------------")
}
