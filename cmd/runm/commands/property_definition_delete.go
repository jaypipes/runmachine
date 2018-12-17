package commands

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/runmachine-io/runmachine/proto"
	"github.com/spf13/cobra"
)

var propertyDefinitionDeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete property definitions matching one or more filters",
	Run:   propertyDefinitionDelete,
}

func setupPropertyDefinitionDeleteFlags() {
	propertyDefinitionDeleteCommand.Flags().StringArrayVarP(
		&cliPropertyDefinitionFilters, // defined in cmd/runm/commands/property_definition_list.go
		"filter", "f",
		nil,
		usagePropertyDefinitionFilterOption, // defined in cmd/runm/commands/property_definition_list.go
	)
}

func init() {
	setupPropertyDefinitionDeleteFlags()
}

func propertyDefinitionDelete(cmd *cobra.Command, args []string) {
	conn := connect()
	defer conn.Close()

	client := pb.NewRunmMetadataClient(conn)
	req := &pb.PropertyDefinitionDeleteRequest{
		Session: getSession(),
		Any:     buildPropertyDefinitionFilters(),
	}
	resp, err := client.PropertyDefinitionDelete(context.Background(), req)
	if s, ok := status.FromError(err); ok {
		if s.Code() != codes.OK {
			fmt.Fprintf(os.Stderr, "Error: %s\n", s.Message())
			if resp != nil && len(resp.Errors) > 0 {
				fmt.Fprintf(os.Stderr, "Details:\n")
				for x, errText := range resp.Errors {
					fmt.Fprintf(os.Stderr, "%d: %s\n", x, errText)
				}
			}
			os.Exit(int(s.Code()))
		}
	}
	if !quiet {
		if verbose {
			fmt.Fprintf(os.Stdout, "deleted %d property definition(s)\n", resp.NumDeleted)
		} else {
			fmt.Fprintf(os.Stdout, "ok\n")
		}
	}
}
