package grpc

import (
	"context"
	"fmt"
	"net"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/yura-under-review/ports-domain-service/api"
	"google.golang.org/grpc"
)

type Server struct {
	addr     string
	s        *grpc.Server
	resolver *Resolver
}

func NewServer(addr string, resolver *Resolver) *Server {
	return &Server{
		addr:     addr,
		resolver: resolver,
	}
}

func (srv *Server) Run(ctx context.Context, wg *sync.WaitGroup) error {

	l, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	srv.s = grpc.NewServer()

	wg.Add(1)
	go func() {
		defer wg.Done()

		api.RegisterPortsDomainServiceServer(srv.s, srv.resolver)

		err := srv.s.Serve(l)
		if err != nil {
			log.Errorf("failed to serve: %v", err)
			return
		}
	}()

	go func() {
		<-ctx.Done()
		srv.s.GracefulStop()
	}()

	return nil
}
