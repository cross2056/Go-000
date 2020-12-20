package service

import (
	"context"
	"fmt"
	"week04/api/hello/v1"
	"week04/internal/biz"
)

type HelloServer struct {
	hello.UnimplementedGreeterServer
	msgBiz *biz.MessageBiz
}

func NewHelloServer(msgBiz *biz.MessageBiz) *HelloServer {
	s := &HelloServer{msgBiz: msgBiz}
	return s
}

func (s *HelloServer) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	msg := &biz.Message{Name: req.GetName()}
	return &hello.HelloReply{Message: fmt.Sprintf("ahaha %s", s.msgBiz.GetMore(msg))}, nil
}
