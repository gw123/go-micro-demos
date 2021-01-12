package handler

import (
	"context"
	"fmt"

	log "github.com/micro/micro/v3/service/logger"

	qrcode "qrcode/proto"
)

type Qrcode struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Qrcode) GetQrcode(ctx context.Context, req *qrcode.Request, rsp *qrcode.Response) error {
	log.Info("Received Qrcode.Call request")
	rsp.Msg = fmt.Sprintf("content: %s, size: %d, format: %d",  req.Content, req.Size, req.Format)
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Qrcode) Stream(ctx context.Context, req *qrcode.StreamingRequest, stream qrcode.Qrcode_StreamStream) error {
	log.Infof("Received Qrcode.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&qrcode.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Qrcode) PingPong(ctx context.Context, stream qrcode.Qrcode_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&qrcode.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
