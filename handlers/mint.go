package handlers

import (
	"fmt"

	"github.com/sryps/mismon/queries"
)

func CheckProvisions(data []queries.AnnualProvisions) (q string) {
	for _, k := range data {
		q += fmt.Sprintf("Annual_Provisions: %d\n", k.Provisions)
	}
	return
}
