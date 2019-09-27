package main

import (
	"context"
	fmt "fmt"
	"os"

	pb "github.com/wanfintara/grpc-tkp/protos"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := pb.NewLocalServicesClient(conn)

	// a, err := client.GetCategories(context.TODO(), &pb.GetCategoriesRequest{})
	// data := pb.CreateOrderRequest{
	// 	TxInvoice: "invoice123",
	// 	TxAmount:  999.789,
	// 	TxStatus:  18,
	// 	OrderDetail: &pb.OrderDetailRequest{
	// 		VendorId:  41,
	// 		ProductId: 99,
	// 	},
	// }
	a, err := client.GetOrder(context.TODO(), &pb.GetOrderRequest{OrderId: 20})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}
