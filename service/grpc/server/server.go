package hydraserver

import (
	"fmt"
	"io"
	"log"

	hydrarpc "github.com/gcetusic/hydra/service/grpc"
	"github.com/satori/go.uuid"
)

type HydraService struct{}

type Nodes map[string]*hydrarpc.Node

var nodes = make(Nodes)

func (srv *HydraService) ListNodes(node *hydrarpc.NodeListRequest, stream hydrarpc.Hydra_ListNodesServer) error {
	return nil
}

func (srv *HydraService) Execute(stream hydrarpc.Hydra_ExecuteServer) error {

	var nodeID uuid.UUID
	for {
		event, err := stream.Recv()
		log.Printf("INCOMING REQUEST")

		if err == io.EOF {
			delete(nodes, nodeID.String())
			log.Printf("Closing connection for node %s: %d nodes running\n", nodeID, len(nodes))

			return nil
		}
		if err != nil {
			return err
		}

		switch event.GetRequest().(type) {
		case *hydrarpc.Event_NodeRegistrationRequest:
			nodeID, _ = uuid.NewV4()
			node := &hydrarpc.Node{
				Id:      nodeID.String(),
				Runtime: event.GetNodeRegistrationRequest().GetRuntime(),
				Status:  event.GetNodeRegistrationRequest().GetStatus(),
			}

			nodes[node.Id] = node
			log.Printf(fmt.Sprintf("Node: %s", nodeID))
			log.Printf(fmt.Sprintf("Runtime: %s", node.Runtime))
			log.Printf(fmt.Sprintf("Status: %s", node.Status))
			log.Printf(fmt.Sprintf("Number of nodes: %d", len(nodes)))

		case *hydrarpc.Event_NodeStatusRequest:
			nodes[nodeID.String()].Status = event.GetNodeStatusRequest().GetStatus()
			log.Printf("Node Status: %s", nodes[nodeID.String()].Status)
		}
	}
}
