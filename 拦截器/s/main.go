package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"mycode/src/pb"
	"net"
)

type server struct{}

func (s *server) MyTest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println(ctx.Value("niubi"))
	fmt.Println("+++++++++++++++++++++")
	st := status.New(codes.NotFound, "dsadas")
	return nil, st.Err()
	//return &pb.Response{
	//	BackJson: "hello " + in.JsonStr,
	//}, nil
}

func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	var opts []grpc.ServerOption

	// TLS认证
	//creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	//if err != nil {
	//	grpclog.Fatalf("Failed to generate credentials %v", err)
	//}
	//opts = append(opts, grpc.Creds(creds))

	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	// 创建gRPC服务器
	s := grpc.NewServer(opts...)
	// 注册服务
	pb.RegisterTesterServer(s, &server{})

	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}

// interceptor 一元拦截器
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appid string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	fmt.Println(appid)

	//fmt.Println("3213123213")
	//cc := ctx.Value("tanent_code")
	//fmt.Println(cc)
	//t := reflect.TypeOf(req)
	//
	//switch req.(type) {
	//
	//}
	//
	//
	//
	//ctx = context.WithValue(ctx, "niubi", "1-------")
	// 继续处理请求
	return handler(ctx, req)
}