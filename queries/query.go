package queries

import (
	"context"
	"time"

	consensus "github.com/cosmos/cosmos-sdk/x/consensus/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	mint "github.com/cosmos/cosmos-sdk/x/mint/types"
	"google.golang.org/grpc"
)

type EvidenceData struct {
	Height    int64
	Address   string
	VotePower int64
	Time      time.Time
}

type AnnualProvisions struct {
	Provisions int
}

func QueryAllEvidence(conn *grpc.ClientConn) (evidenceRes []EvidenceData, err error) {
	// Query evidence module for misbehaviour evidience
	evidenceClient := evidence.NewQueryClient(conn)
	resp, err := evidenceClient.AllEvidence(
		context.Background(),
		&evidence.QueryAllEvidenceRequest{},
	)

	for _, item := range resp.Evidence {
		result := evidence.Equivocation{}
		_ = result.Unmarshal(item.Value)
		newEvidence := EvidenceData{
			Height:    result.Height,
			Address:   result.ConsensusAddress,
			VotePower: result.Power,
			Time:      result.Time,
		}
		evidenceRes = append(evidenceRes, newEvidence)
	}
	return
}

func QueryAnnualProvisions(conn *grpc.ClientConn) (mintRes []AnnualProvisions, err error) {
	// Query mint module for annual provisions
	mintClient := mint.NewQueryClient(conn)
	resp, err := mintClient.AnnualProvisions(
		context.Background(),
		&mint.QueryAnnualProvisionsRequest{},
	)

	newMint := AnnualProvisions{
		Provisions: int(resp.AnnualProvisions.RoundInt64()),
	}
	mintRes = append(mintRes, newMint)

	return
}

func QueryConParams(conn *grpc.ClientConn) (consensusRes *consensus.QueryParamsResponse) {
	// Query consensus module for params
	consensusClient := consensus.NewQueryClient(conn)
	consensusRes, err := consensusClient.Params(
		context.Background(),
		&consensus.QueryParamsRequest{},
	)
	if err != nil {
		panic(err)
	}

	return consensusRes
}
