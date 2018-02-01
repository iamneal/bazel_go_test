package main

import (
	"bazel_go_test/lib"
	"fmt"
	"path"

	"net"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Main() {
	myToml := make(map[string]interface{})
	md, err := toml.DecodeFile(path.Clean("./Gopkg.lock"), &myToml)
	if err != nil {
		panic(err)
	}
	logrus.Printf("myToml: %+v\n", myToml)
	logrus.Printf("metadata: %+v\n", md)
	logrus.Printf("%s world\n", lib.GetString())
	logrus.Printf("%+v user\n", lib.GetUser())
	service := lib.NewUBuilder().
		WithDefaultQueryHandlers().
		WithSpannerURI(context.Background(), "projects/test/instances/test/databases/test").
		MustBuild()
	server := grpc.NewServer()
	lib.RegisterUServer(server, service)
	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(list); err != nil {
		fmt.Printf("error erving: %v\n", err)
	}
}

func main() {
	Main()
}
