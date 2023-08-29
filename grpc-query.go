package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/codec"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	mint "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func queryState(conn *grpc.ClientConn, module string) error {

	switch m := module; m {
	case "evidence":
		queryEvidence(conn)
	case "provisions":
		queryAnnualProvisions(conn)
	default:
		fmt.Println("Unknown Module....Please try again.")
	}

	return nil
}

func queryEvidence(conn *grpc.ClientConn) {
	// Query mint module for misbehaviour evidience
	evidenceClient := evidence.NewQueryClient(conn)
	evidenceRes, err := evidenceClient.AllEvidence(
		context.Background(),
		&evidence.QueryAllEvidenceRequest{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(evidenceRes)
}

func queryAnnualProvisions(conn *grpc.ClientConn) {
	// Query mint module for annual provisions
	mintClient := mint.NewQueryClient(conn)
	mintRes, err := mintClient.AnnualProvisions(
		context.Background(),
		&mint.QueryAnnualProvisionsRequest{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(mintRes)
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		panic("Need server address and module!")
	}
	server := arguments[1]
	module := arguments[2]

	fmt.Println("Connecting to:", server)

	conn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		panic(err)
	}

	if err := queryState(conn, module); err != nil {
		panic(err)
	}

	conn.Close()
}
