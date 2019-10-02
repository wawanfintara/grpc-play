package main

import (
	"context"
	fmt "fmt"
	"os"

	pb "github.com/wanfintara/grpc-tkp/protos"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("10.255.13.244:9000", grpc.WithInsecure())
	// conn, err := grpc.Dial("10.255.13.244:9000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := pb.NewLocalServicesClient(conn)
	ctx := context.TODO()
	// *GetCategories
	// a, err := client.GetCategories(ctx, &pb.GetCategoriesRequest{})

	// * CreateOrder
	data := pb.CreateOrderRequest{
		TxInvoice: "invoice",
		TxAmount:  999.789,
		TxStatus:  18,
		OrderDetail: &pb.OrderDetailRequest{
			VendorId:  41,
			ProductId: 99,
		},
	}
	a, err := client.CreateOrder(ctx, &data)

	// // *GetOrder
	// a, err := client.GetOrder(ctx, &pb.GetOrderRequest{OrderId: 1})

	// * UpdateOrder
	// a, err := client.UpdateOrder(ctx, &pb.UpdateOrderStatusRequest{
	// 	OrderId:  2,
	// 	TxStatus: 0,
	// })

	// a, err := client.GetProducts(ctx, &pb.GetProductsRequest{Limit: 1})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", a)
}
