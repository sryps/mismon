package queries

import (
	"context"
	"fmt"
	"time"

	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
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
	Provisions any
}

func QueryAllEvidence(conn *grpc.ClientConn) (evidenceRes []EvidenceData, err error) {
	// Query evidence module for misbehaviour evidience

	evidenceClient := evidence.NewQueryClient(conn)
	resp, err := evidenceClient.AllEvidence(
		context.Background(),
		&evidence.QueryAllEvidenceRequest{},
	)
	if err != nil {
		return nil, err
	}

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
	if err != nil {
		return nil, err
	}

	newMint := AnnualProvisions{
		Provisions: resp.AnnualProvisions,
	}
	mintRes = append(mintRes, newMint)

	return
}

func QueryBank(conn *grpc.ClientConn) (consensusRes *bank.QueryDenomsMetadataResponse) {
	// Query consensus module for params
	bankClient := bank.NewQueryClient(conn)
	consensusRes, err := bankClient.DenomsMetadata(
		context.Background(),
		&bank.QueryDenomsMetadataRequest{},
	)
	if err != nil {
		fmt.Println(err)
	}

	return
}
