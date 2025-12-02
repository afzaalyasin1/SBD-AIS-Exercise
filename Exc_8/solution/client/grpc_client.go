package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	// todo
	// 1. List drinks
	fmt.Println("Requesting drinks 🍹🍺☕")
	drinkList, err := c.client.GetDrinks(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}

	fmt.Println("Available drinks:")
	for _, d := range drinkList.Drinks {
		fmt.Printf("\t> id:%d  name:%q  price:%.0f  description:%q\n", d.Id, d.Name, d.Price, d.Description)
	}

	// 2. Order a few drinks
	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	for _, d := range drinkList.Drinks {
		count := 2
		fmt.Printf("\t> Ordering: %d x %s\n", count, d.Name)
		for i := 0; i < count; i++ {
			_, err := c.client.OrderDrink(context.Background(), d)
			if err != nil {
				return err
			}
		}
	}
	// 3. Order more drinks
	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	for _, d := range drinkList.Drinks {
		count := 6
		fmt.Printf("\t> Ordering: %d x %s\n", count, d.Name)
		for i := 0; i < count; i++ {
			_, err := c.client.OrderDrink(context.Background(), d)
			if err != nil {
				return err
			}
		}
	}
	// 4. Get order total
	fmt.Println("Getting the bill 💹💹💹")
	orders, err := c.client.GetOrders(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}

	totals := make(map[string]int)
	for _, o := range orders.Orders {
		if o.Drink != nil {
			totals[o.Drink.Name]++
		}
	}

	for name, count := range totals {
		fmt.Printf("\t> Total: %d x %s\n", count, name)
	}
	// print responses after each call
	fmt.Println("Orders complete!")
	return nil
}
