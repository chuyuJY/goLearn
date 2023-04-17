package servers

import (
	"context"

)

type server struct {
	.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *)
