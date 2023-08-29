package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/codec"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	mint "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func queryClient(conn *grpc.ClientConn, module string) error {
	s := strings.Split(module, " ")

	for i := 0; i < len(s); i++ {
		fail := true
		fmt.Println("\n\nRunning", s[i], "module query...")
		if s[i] == "evidence" {
			queryEvidence(conn)
			fail = false
		}
		if s[i] == "provisions" {
			queryAnnualProvisions(conn)
			fail = false
		}
		if fail {
			fmt.Println("\n\n", s[i], "module query is not available")
			fmt.Println("Available queries: evidence, provisions, ")
			panic("unknown module used")
		}
	}

	return nil
}

func queryEvidence(conn *grpc.ClientConn) {
	// Query evidence module for misbehaviour evidience
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
	if len(arguments) < 3 {
		panic("Need server address and module!")
	}
	server := arguments[1]
	module := strings.Join(arguments[2:], " ")

	fmt.Println("Connecting to:", server)

	conn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		panic(err)
	}

	if err := queryClient(conn, module); err != nil {
		panic(err)
	}

	conn.Close()
}
