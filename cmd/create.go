package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/0x4c6565/gopaste/internal/pkg/util"
	"github.com/spf13/cobra"
	"gopkg.in/Luzifer/go-openssl.v1"
)

func createPasteCommand() *cobra.Command {
	// Create command
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new paste",
		RunE:  createPaste,
	}

	createCmd.Flags().StringP("content", "c", "", "Content for the paste")
	createCmd.Flags().StringP("file", "f", "", "File to read content from. Use '-' for stdin")
	createCmd.Flags().StringP("syntax", "s", "text/plain", "Syntax highlighting for the paste")
	createCmd.Flags().Int64P("expires", "e", 604800, "When the paste expires (see list-expires for options)")

	return createCmd
}

type uuidResponse struct {
	ID string `json:"id"`
}

type pasteRequest struct {
	Expires int64  `json:"expires"`
	Syntax  string `json:"syntax"`
	Content string `json:"content"`
}

func createPaste(cmd *cobra.Command, args []string) error {
	content, err := cmd.Flags().GetString("content")
	if err != nil {
		return err
	}

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	if content == "" && file == "" {
		return errors.New("either content or file must be provided")
	}

	if file != "" {
		if file == "-" {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("error reading content from stdin: %s", err)
			}
			content = string(data)
		} else {
			data, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("error reading file: %s", err)
			}
			content = string(data)
		}
	}

	syntax, err := cmd.Flags().GetString("syntax")
	if err != nil {
		return err
	}

	expires, err := cmd.Flags().GetInt64("expires")
	if err != nil {
		return err
	}

	// Generate a random password for encryption
	password := util.GeneratePassword()

	o := openssl.New()

	encrypted, err := o.EncryptString(password, content)
	if err != nil {
		return fmt.Errorf("error encrypting content: %s", err)
	}

	content = string(encrypted)

	reqBody := pasteRequest{
		Content: content,
		Syntax:  syntax,
		Expires: expires,
	}

	uuidResp, err := util.CallAPI[uuidResponse]("POST", fmt.Sprintf("%s/paste", APIURL), "application/json", reqBody)
	if err != nil {
		return fmt.Errorf("error creating paste: %s", err)
	}

	fmt.Printf("https://p.lee.io/%s#encryptionKey=%s\n", uuidResp.ID, password)
	return nil
}
