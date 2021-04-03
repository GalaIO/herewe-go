package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)
import pb "github.com/GalaIO/herewe-go/grpc/route_guide"

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8091", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	feature, err := client.GetFeature(context.Background(), &pb.Point{Latitude: 409146138, Longitude: -746188906})
	if err != nil {
		panic(err)
	}
	fmt.Println(feature.String())

	rect := &pb.Rectangle{} // initialize a pb.Rectangle
	stream, err := client.ListFeatures(context.Background(), rect)
	if err != nil {
		panic(err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
}
