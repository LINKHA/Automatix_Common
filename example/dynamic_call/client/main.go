package main

import (
	"fmt"

	dynamic_call "github.com/LINKHA/automatix-common/dynamic_call"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("start...")
	defer fmt.Println("end...")

	dynamic_call.SetProtoSetFiles("dynamic_call.protoset")
	err := dynamic_call.InitDescSource()

	if err != nil {
		panic(err.Error())
	}

	var handler = DefaultEventHandler{}
	var sendBody string

	grpcEnter, err := dynamic_call.New(
		dynamic_call.SetHookHandler(&handler),
	)
	grpcEnter.Init()

	sendBody = `{"name": "golang world"}`
	res, err := grpcEnter.Call("127.0.0.1:50051", "apigrpc.Greeter", "SayHello", sendBody)
	fmt.Printf("%+v \n", res)
	fmt.Println("req/reply return err: ", err)

}

type DefaultEventHandler struct {
	sendChan chan []byte
}

func (h *DefaultEventHandler) OnReceiveData(md metadata.MD, resp string, respErr error) {
}

func (h *DefaultEventHandler) OnReceiveTrailers(stat *status.Status, md metadata.MD) {
}
