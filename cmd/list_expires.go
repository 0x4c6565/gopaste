package cmd

import (
	"fmt"

	"github.com/0x4c6565/gopaste/internal/pkg/util"
	"github.com/spf13/cobra"
)

func listExpiresCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-expires",
		Short: "List available expiration options",
		RunE:  listExpires,
	}

	return cmd
}

func listExpires(cmd *cobra.Command, args []string) error {
	expires, err := util.CallAPI[[]expiresResponse]("GET", fmt.Sprintf("%s/expires", APIURL), "application/json", nil)
	if err != nil {
		return fmt.Errorf("error listing expiration options: %s", err)
	}

	fmt.Println("Available Expiration Options:")
	for _, e := range *expires {
		if e.Default {
			fmt.Printf("* %s (%d) - default\n", e.Label, e.Expires)
		} else {
			fmt.Printf("* %s (%d)\n", e.Label, e.Expires)
		}
	}

	return nil
}

type expiresResponse struct {
	Label   string `json:"label"`
	Default bool   `json:"default,omitempty"`
	Expires int64  `json:"expires"`
}
