package main

import (
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/sryps/mismon/queries"
)

func queryClient(conn *grpc.ClientConn, module string) error {
	s := strings.Split(module, " ")

	for i := 0; i < len(s); i++ {
		fail := true
		fmt.Println("\n\nRunning", s[i], "module query...")
		if s[i] == "evidence" {
			q := queries.QueryEvidence(conn)
			fmt.Println(q)
			fail = false
		}
		if s[i] == "provisions" {
			q := queries.QueryAnnualProvisions(conn)
			fmt.Println(q)
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
