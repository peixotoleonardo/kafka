package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/peixotoleonardo/kafka/golang/internal/kafka"
	"github.com/peixotoleonardo/kafka/golang/internal/logger"
)

var log = logger.WithPrefix(context.TODO(), "producer")

func main() {
	producer, err := kafka.NewProducer()

	if err != nil {
		log.WithError(err).Fatal("failed to create producer")
	}

	log.Info("enviando evento")

	producer.Send("Test")
	log.Info("evento enviado")

	awaitForInterruptSignal()

	if err = producer.Close(); err != nil {
		log.WithError(err).Fatal("producer forced to shutdown")
	}

	log.Info("producer exiting")
}

func awaitForInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
