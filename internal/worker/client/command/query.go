package command

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gaurav36/systemd-wrapper/internal/worker/proto"
	"google.golang.org/grpc"
)

type QueryCommand struct {
	client proto.WorkerServiceClient
}

func NewQueryCommand(client proto.WorkerServiceClient) Runner {
	return &QueryCommand{
		client: client,
	}
}

func (c *QueryCommand) Run(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass an argument")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	command := proto.QueryRequest{
		JobID: args[0],
	}
	res, err := c.client.Query(ctx, &command, grpc.WaitForReady(true))
	if err != nil {
		return err
	}
	os.Stdout.WriteString(fmt.Sprintf("Pid: %v Exit code: %v Exited: %v\n", res.Pid, res.ExitCode, res.Exited))
	return nil
}
