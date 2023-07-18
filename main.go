package main

import (
	"grpc-client/common/pb/hello"

	"github.com/maybgit/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.Dial("localhost:6666", grpc.WithTransportCredentials(insecure.NewCredentials(), insecure.NewCredentials()))
	if err != nil {
		glog.Warning("客户端发起连接失败：", err.Error())
	}

	defer clientConn.Close()
	//grpc.Dial
	//greeterClient:=hello.NewGreeterClient(cc grpc.ClientConnInterface)
	
	hello.NewGreeterClient(cc grpc.ClientConnInterface)

}
