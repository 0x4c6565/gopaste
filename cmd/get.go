package cmd

import (
	"fmt"

	"github.com/0x4c6565/gopaste/internal/pkg/util"
	"github.com/spf13/cobra"
	"gopkg.in/Luzifer/go-openssl.v1"
)

func getPasteCommand() *cobra.Command {
	// Get command
	getCmd := &cobra.Command{
		Use:   "get <paste: id>",
		Short: "Get a paste by ID",
		Args:  cobra.ExactArgs(1),
		RunE:  getPaste,
	}

	getCmd.Flags().StringP("password", "p", "", "Password for client-side decryption")
	getCmd.Flags().BoolP("raw", "r", false, "Output raw content without metadata")

	return getCmd
}

type pasteResponse struct {
	Expires int64  `json:"expires"`
	Syntax  string `json:"syntax"`
	Content string `json:"content"`
	Burnt   bool   `json:"burnt"`
}

func getPaste(cmd *cobra.Command, args []string) error {
	id := args[0]
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return err
	}

	raw, err := cmd.Flags().GetBool("raw")
	if err != nil {
		return err
	}

	paste, err := util.CallAPI[pasteResponse]("GET", fmt.Sprintf("%s/paste/%s", APIURL, id), "application/json", nil)
	if err != nil {
		return fmt.Errorf("error fetching paste: %s", err)
	}

	content := paste.Content

	o := openssl.New()

	decrypted, err := o.DecryptString(password, content)
	if err != nil {
		return fmt.Errorf("error decrypting content: %s", err)
	}

	content = string(decrypted)

	if raw {
		fmt.Print(content)
	} else {
		fmt.Println("Paste Details:")
		fmt.Println("Syntax:", paste.Syntax)
		if paste.Expires == -1 {
			fmt.Println("Expires: Never")
		} else if paste.Expires == -2 {
			fmt.Println("Expires: Burn after reading")
			if paste.Burnt {
				fmt.Println("Status: Burnt (this was the only time this paste could be viewed)")
			}
		} else {
			fmt.Println("Expires:", paste.Expires)
		}
		fmt.Println("\nContent:")
		fmt.Println(content)
	}

	return nil
}
