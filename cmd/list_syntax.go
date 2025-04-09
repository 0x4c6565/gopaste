package cmd

import (
	"fmt"

	"github.com/0x4c6565/gopaste/internal/pkg/util"
	"github.com/spf13/cobra"
)

func listSyntaxCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-syntax",
		Short: "List available syntax highlighting options",
		RunE:  listSyntax,
	}

	return cmd
}

func listSyntax(cmd *cobra.Command, args []string) error {
	syntax, err := util.CallAPI[[]syntaxResponse]("GET", fmt.Sprintf("%s/syntax", APIURL), "application/json", nil)
	if err != nil {
		return fmt.Errorf("error listing syntax options: %s", err)
	}

	fmt.Println("Available Syntax Highlighting Options:")
	for _, s := range *syntax {
		if s.Default {
			fmt.Printf("* %s (%s) - default\n", s.Label, s.Syntax)
		} else {
			fmt.Printf("* %s (%s)\n", s.Label, s.Syntax)
		}
		if len(s.Aliases) > 0 {
			fmt.Printf("  Aliases: %v\n", s.Aliases)
		}
	}

	return nil
}

type syntaxResponse struct {
	Label   string   `json:"label"`
	Syntax  string   `json:"syntax"`
	Default bool     `json:"default,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
}
