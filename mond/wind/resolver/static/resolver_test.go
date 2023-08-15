package static

import (
	"context"
	"fmt"
	"testing"

	MetaDemo "github.com/tangbo/twatt/mond/wind/resolver/static/metaDemo"
	"github.com/tangbo/twatt/mond/wind/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

func TestNewStaticResolverBuilder(t *testing.T) {
	resolver.Register(NewStaticResolverBuilder(map[string][]string{
		"demo": []string{"localhost:8080"},
	}))
	//balancer.Register()
	conn, err := grpc.Dial("static://demo",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	utils.MustNil(err)
	client := MetaDemo.NewMetaDemoServiceClient(conn)
	res, err := client.SayHello(context.TODO(), &MetaDemo.HelloRequest{})
	if err != nil {
		fmt.Println(err)
	}
	res, err = client.SayHello(context.TODO(), &MetaDemo.HelloRequest{})
	if err != nil {
		fmt.Println(err)
	}
	//utils.MustNil(err)
	fmt.Println(res)
}
