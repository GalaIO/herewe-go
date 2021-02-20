package grpc

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/GalaIO/herewe-go/grpc/route_guide"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"math"
	"net"
	"sync"
)

type routeGuideServerImpl struct {
	pb.UnimplementedRouteGuideServer
	savedFeatures []*pb.Feature // read-only after initialized

	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*pb.RouteNote
}

func (s *routeGuideServerImpl) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	// No feature was found, return an unnamed feature
	return &pb.Feature{Location: point}, nil
}

func (s *routeGuideServerImpl) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	for _, feature := range s.savedFeatures {
		if inRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *routeGuideServerImpl) RecordRoute(ctx context.Context, point *pb.Point) (*pb.RouteSummary, error) {
	panic("implement me")
}

func (s *routeGuideServerImpl) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	panic("implement me")
}

func (s *routeGuideServerImpl) mustEmbedUnimplementedRouteGuideServer() {
	panic("implement me")
}



func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

func newServer() *routeGuideServerImpl {
	s := &routeGuideServerImpl{routeNotes: make(map[string][]*pb.RouteNote)}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8091))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}