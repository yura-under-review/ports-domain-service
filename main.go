package main

import (
	"context"
	"github.com/yura-under-review/ports-domain-service/opts"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/yura-under-review/ports-domain-service/grpc"
	"github.com/yura-under-review/ports-domain-service/repository"
)

const (
	INFO  = "INFO"
	ERROR = "ERROR"
	WARN  = "WARN"

	// DEBUG = "DEBUG"
)

func main() {

	config := opts.LoadConfigFromEnv()

	initLogger(config.LogLevel, config.PrettyLogs)

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	setupGracefulShutdown(cancel)

	// TODO: replace with envs
	repo := repository.New(config.Postgres)
	if err := repo.Init(ctx); err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}
	defer repo.Close()

	// initializing gRPC
	resolver := grpc.NewResolver(repo)
	server := grpc.NewServer(config.GRPC, resolver)

	if err := server.Run(ctx, &wg); err != nil {
		log.Fatalf("failed to run grpc servier: %v", err)
	}

	log.Info("service runs")

	wg.Wait()
}

func initLogger(lvl string, prettyLogs bool) {

	switch strings.ToUpper(lvl) {
	case INFO:
		log.SetLevel(log.InfoLevel)

	case WARN:
		log.SetLevel(log.WarnLevel)

	case ERROR:
		log.SetLevel(log.ErrorLevel)

	default:
		log.SetLevel(log.DebugLevel)
	}

	log.SetFormatter(&log.JSONFormatter{PrettyPrint: prettyLogs})
	log.SetOutput(os.Stderr)
}

func setupGracefulShutdown(cancel context.CancelFunc) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Info("interrupt signal")
		cancel()
	}()
}
