package hydraclient

import (
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	hydrarpc "github.com/gcetusic/hydra/service/grpc"
	"google.golang.org/grpc"
)

func Run() error {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to Hydra: %v", err)
	}

	client := hydrarpc.NewHydraClient(conn)
	stream, _ := client.Execute(context.Background())

	waitc := make(chan struct{})
	go func() {
		for {
			_, err := stream.Recv()

			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a message : %v", err)
			}
		}
	}()

	registrationEvent := &hydrarpc.Event{
		Message: &hydrarpc.Event_NodeRegistrationMessage{
			NodeRegistrationMessage: &hydrarpc.NodeRegistrationMessage{
				Runtime: hydrarpc.Runtime_GOLANG,
			},
		},
	}

	if err := stream.Send(registrationEvent); err != nil {
		log.Fatalf("Failed to send a registration request: %v", err)
	}

	runningStatusEvent := &hydrarpc.Event{
		Message: &hydrarpc.Event_NodeStatusMessage{
			NodeStatusMessage: &hydrarpc.NodeStatusMessage{
				Status: hydrarpc.Status_RUNNING,
			},
		},
	}
	stream.Send(runningStatusEvent)
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)

	readyStatusEvent := &hydrarpc.Event{
		Message: &hydrarpc.Event_NodeStatusMessage{
			NodeStatusMessage: &hydrarpc.NodeStatusMessage{
				Status: hydrarpc.Status_READY,
			},
		},
	}
	stream.Send(readyStatusEvent)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	stream.CloseSend()
	<-waitc

	defer conn.Close()
	return nil
}
