package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	dist "github.com/cosmos/cosmos-sdk/x/mint/types"
)

var server = ""

func queryState() error {

	arguments := os.Args
	if len(arguments) != 2 {
		panic("Need server addr!")
	}
	server = arguments[1]
	fmt.Println("Connecting to:", server)

	grpcConn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return err
	}

	// Prints the account balance of random address
	myAddress, err := sdk.AccAddressFromBech32("cosmos12adsjglpf38dyw7ecugecz8fhnrvwqg5tuw3cw")
	if err != nil {
		return err
	}
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: myAddress.String(), Denom: "uatom"},
	)
	if err != nil {
		return err
	}

	fmt.Println(bankRes.GetBalance())

	// Query mint module for annual provisions
	distClient := dist.NewQueryClient(grpcConn)
	distRes, err := distClient.AnnualProvisions(
		context.Background(),
		&dist.QueryAnnualProvisionsRequest{},
	)
	if err != nil {
		return err
	}
	fmt.Println(distRes)

	// Query mint module for misbehaviour evidience
	evidenceClient := evidence.NewQueryClient(grpcConn)
	evidenceRes, err := evidenceClient.AllEvidence(
		context.Background(),
		&evidence.QueryAllEvidenceRequest{},
	)
	if err != nil {
		return err
	}
	fmt.Println(evidenceRes)

	grpcConn.Close()
	return nil

}

func main() {
	if err := queryState(); err != nil {
		panic(err)
	}
}
