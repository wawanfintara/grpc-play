package main

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"os"

	pb "github.com/wanfintara/grpc-tkp/pb"
	grpc "google.golang.org/grpc"
)

var ctx context.Context = context.TODO()

func main() {
	// conn, err := grpc.Dial("10.255.13.244:9000", grpc.WithInsecure())
	conn, err := grpc.Dial(":9900", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := pb.NewLocalServicesClient(conn)

	checkoutValidation(client)
	// getPrice(client)
}

func createOrder(client pb.LocalServicesClient) {

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

	// *GetOrder
	// a, err = client.GetOrder(ctx, &pb.GetOrderRequest{OrderId: 1})

	if err != nil {
		fmt.Println(err)
		return
	}
	js, _ := json.Marshal(a)
	fmt.Printf("%v", string(js))
}

func checkoutValidation(client pb.LocalServicesClient) {
	// * Checkout Validation

	product := ProductFromCI{
		ID: 93,
		// ID:            10,
		Name:          "kartu nama product",
		Volume:        3000,
		TotalWeight:   1,
		TotalPrice:    185955,
		PriceOptionID: 145714,
		Qty:           1,
		QtyOption:     1,
		SpecSet:       "{\"403\":778,\"404\":782,\"405\":783,\"406\":785,\"407\":787,\"408\":788}",
		// SpecSet: "{\"41\":75,\"42\":77,\"43\":79}",
	}
	// fmt.Println(product)

	shipping := ShippingFromCI{
		ID:        1104,
		SpID:      45,
		Price:     18000,
		Insurance: 1,
		// InsurancePrice:  55,
		Weight:          1,
		OrderValue:      100000,
		OriginAddressID: 4502,
		DestAddressID:   4674761,
	}

	insurancep := (product.TotalPrice + shipping.Price) * 15 / 10000
	shipping.InsurancePrice = insurancep

	metadata := MetadataCI{
		Products: []ProductFromCI{
			product,
		},
		Shipping: shipping,
	}
	bytes, _ := json.Marshal(metadata)

	s := string(bytes)
	fmt.Println(s)

	// raw json
	rawIn, _ := json.Marshal(s)
	fmt.Println("\n\nraw:", string(rawIn))

	a, err := client.CheckoutValidation(ctx, &pb.CheckoutValidationRequest{
		UserId: 8972830, // cannot 0
		// UserId:   100, // cannot 0
		Metadata: string(bytes),
		// Metadata: "{\"products\":[{\"id\":93,\"name\":\"kartu nama product\",\"volume\":3000,\"dimension\":\"\",\"weight\":1,\"total_price\":185955,\"price_option_id\":145714,\"qty\":1,\"qty_option\":1,\"spec_set\":\"{\\\"403\\\":778,\\\"404\\\":782,\\\"405\\\":783,\\\"406\\\":785,\\\"407\\\":787,\\\"408\\\":788}\",\"notes\":\"\"}],\"shipping\":{\"shipping_id\":1104,\"sp_id\":45,\"shipping_price\":18000,\"insurance\":1,\"insurance_price\":305,\"origin_address_id\":4502,\"dest_address_id\":4674761}}",
	})
	fmt.Printf("\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	printResult(a)
}

func getPrice(c pb.LocalServicesClient) {
	r, err := c.GetProductPrice(context.TODO(), &pb.GetProductPriceRequest{
		ProductId: 45,
		SpecSet:   "{\"9\":241,\"10\":245,\"11\":250,\"18\":248,\"13\":249}",
		Area: &pb.Area{
			DistrictId: 4502,
			CityId:     100,
			ProvinceId: 200,
		},
		Quantity:  0,
		TkpUserId: 8972830,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	printResult(r)
}

func printResult(result interface{}) {
	js, _ := json.Marshal(result)
	fmt.Printf("%v", string(js))
}
