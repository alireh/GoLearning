package microservicesTools

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/v5/server"
)

//installation command
// go get github.com/smallnest/rpcx/v5
//go get github.com/smallnest/rpcx/v5/codec@v5.7.9
//go get github.com/smallnest/rpcx/v5/log@v5.7.9
//go get github.com/smallnest/rpcx/v5/server@v5.7.9

func TestRpcx() {
	fmt.Println("----------------------------------------Start Rpcx----------------------------------------")

	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", ":8972")

	fmt.Println("----------------------------------------End Rpcx  ----------------------------------------")
}

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith struct{}

func (a *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}
