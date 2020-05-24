package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	pb "github.com/zhang0911/leanMicro/proto/greeter"
	"log"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(micro.Name("greeter"))
	service.Init()
	pb.RegisterGreeterHandler(service.Server(), &Greeter{})
	if err := service.Run(); err != nil {
		log.Fatalf("run error :%v", err)
	}
}
