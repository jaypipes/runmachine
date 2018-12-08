package commands

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/runmachine-io/runmachine/proto"
	"github.com/spf13/cobra"
)

var objectCreateCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a new object",
	Run:   objectCreate,
}

func setupObjectCreateFlags() {
	objectCreateCommand.Flags().StringVarP(
		&cliObjectDocPath,
		"file", "f",
		"",
		"optional filepath to YAML document to send.",
	)
}

func init() {
	setupObjectCreateFlags()
}

func objectCreate(cmd *cobra.Command, args []string) {
	conn := connect()
	defer conn.Close()

	client := pb.NewRunmMetadataClient(conn)
	req := &pb.ObjectSetRequest{
		Session: getSession(),
		Format:  pb.PayloadFormat_YAML,
		Payload: readInputDocumentOrExit(),
	}

	resp, err := client.ObjectSet(context.Background(), req)
	exitIfError(err)
	obj := resp.Object
	if !quiet {
		printObject(obj)
	} else {
		fmt.Println(obj.Uuid)
	}
}
