package main

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/sryps/mismon/handlers"
	"github.com/sryps/mismon/queries"
)

func queryClient(conn *grpc.ClientConn, module string) (q string, err error) {

	switch module {
	case "evidence":
		data, err := queries.QueryAllEvidence(conn)
		q := handlers.CheckEvidence(data)
		return q, err

	case "provisions":
		data, err := queries.QueryAnnualProvisions(conn)
		q := handlers.CheckProvisions(data)
		return q, err

	case "consensus":
		q = queries.QueryConParams(conn).String()

	default:
		err = fmt.Errorf("invalid module: %s not found", module)
	}

	return
}

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("USEAGE: <binary> <server IP:PORT> <evidence | provisions | consensus>")
		panic("Need server address and queries!")
	}
	server := arguments[1]
	modules := arguments[2:]

	fmt.Printf("Connecting to: %s\n\n", server)

	conn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		panic(err)
	}

	for _, module := range modules {
		result, err := queryClient(conn, module)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		fmt.Println(result)
	}

	conn.Close()
}
