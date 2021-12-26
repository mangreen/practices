package main

import (
    "context"
    "fmt"
    pb "grpc/protobuf"
    "google.golang.org/grpc"
    "log"
    "time"
)

func main() {
    addr := "127.0.0.1:9999"
    conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Can not connect to gRPC server: %v", err)
    }
    defer conn.Close()

    c := pb.NewHelloServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.SayHello(ctx, &pb.HelloRequest{Greeting: "Moto"})
    if err != nil {
        log.Fatalf("Could not get nonce: %v", err)
    }
    fmt.Println("Response:", r.GetReply())
}