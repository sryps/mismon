package queries

import (
	"context"

	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	mint "github.com/cosmos/cosmos-sdk/x/mint/types"
	"google.golang.org/grpc"
)

func QueryEvidence(conn *grpc.ClientConn) (evidenceRes *evidence.QueryAllEvidenceResponse) {
	// Query evidence module for misbehaviour evidience
	evidenceClient := evidence.NewQueryClient(conn)
	evidenceRes, err := evidenceClient.AllEvidence(
		context.Background(),
		&evidence.QueryAllEvidenceRequest{},
	)
	if err != nil {
		panic(err)
	}
	//fmt.Println(evidenceRes)
	return evidenceRes
}

func QueryAnnualProvisions(conn *grpc.ClientConn) (mintRes *mint.QueryAnnualProvisionsResponse) {
	// Query mint module for annual provisions
	mintClient := mint.NewQueryClient(conn)
	mintRes, err := mintClient.AnnualProvisions(
		context.Background(),
		&mint.QueryAnnualProvisionsRequest{},
	)
	if err != nil {
		panic(err)
	}
	//fmt.Println(mintRes)
	return mintRes
}
