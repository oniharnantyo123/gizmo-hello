package greetertransport

import (
	pb "gizmo-hello/pb"
	ocontext "golang.org/x/net/context"
)

// Greeting implementation of the gRPC service.
func (s *TService) Greeting(ctx ocontext.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hola Gizmo RPC " + r.Name}, nil
}
