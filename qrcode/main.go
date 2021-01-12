package main

import (
	"qrcode/handler"
	pb "qrcode/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("qrcode"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterQrcodeHandler(srv.Server(), new(handler.Qrcode))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
