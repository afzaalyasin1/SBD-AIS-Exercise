package server

import (
	"context"
	"exc8/pb"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
	drinks *pb.Drinks
	orders *pb.Orders
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// Implement functions

func (s *GRPCService) GetDrinks(ctx context.Context, empty *emptypb.Empty) (*pb.Drinks, error) {
	// If drinks are not initialized yet, let's create them now
	if s.drinks == nil {
		s.drinks = &pb.Drinks{
			Drinks: []*pb.Drink{
				{Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
				{Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
				{Id: 3, Name: "Coffee", Description: "Mifare isn't that secure"},
			},
		}
	}
	return s.drinks, nil
}

func (s *GRPCService) OrderDrink(ctx context.Context, drink *pb.Drink) (*emptypb.Empty, error) {
	// If orders list is nil, initialize it
	if s.orders == nil {
		s.orders = &pb.Orders{
			Orders: []*pb.Order{},
		}
	}

	newOrder := &pb.Order{Drink: drink}
	s.orders.Orders = append(s.orders.Orders, newOrder)
	return &emptypb.Empty{}, nil
}

func (s *GRPCService) GetOrders(ctx context.Context, empty *emptypb.Empty) (*pb.Orders, error) {
	// If orders list is nil, initialize it
	if s.orders == nil {
		s.orders = &pb.Orders{
			Orders: []*pb.Order{},
		}
	}
	return s.orders, nil
}
