package handlers

import (
	"fmt"

	"github.com/sryps/mismon/queries"
)

func CheckEvidence(data []queries.EvidenceData) (q string) {
	for _, k := range data {
		q += fmt.Sprintf("Height: %d, Address: %s, VotePower: %d, Time: %s\n", k.Height, k.Address, k.VotePower, k.Time)
	}
	return
}
