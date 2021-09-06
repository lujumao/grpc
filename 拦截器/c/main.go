package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mycode/src/pb"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(unaryInterceptor))
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	// 连接服务器
	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()

	// 新建一个客户端
	c := pb.NewTesterClient(conn)
	// 调用服务端函数
	cxt := context.Background()
	cxt = context.WithValue(cxt, "tanent_code", "111111111111111")

	cc := cxt.Value("tanent_code")

	fmt.Println(cc)

	r, err := c.MyTest(cxt, &pb.Request{JsonStr: "222222"})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Printf("调用成功: %s", r.BackJson)
}

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = context.WithValue(ctx, "niuhi", "32151232")
	return nil
}

// unaryInterceptor 一个简单的 unary interceptor 示例。
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// pre-processing
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...) // invoking RPC method
	// post-processing
	end := time.Now()
	println("RPC: %s, req:%v start time: %s, end time: %s, err: %v", method, req, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return err
}

// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return false
}