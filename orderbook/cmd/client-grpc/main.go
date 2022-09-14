package main

func main() {
	// get configuration
	/*
		port := flag.String("server", "8080", "gRPC server in format host:port")
		flag.Parse()

		address := "localhost:" + *port
		cl := grpc.NewOrderBookClient(address)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		limitOrderReq := &pb.LimitOrderRequest{
			Side:     int32(1),
			OrderId:  "my-order-1",
			Quantity: "5.2",
			Price:    "10.00",
		}

		lr, err := cl.ProcessLimitOrder(ctx, limitOrderReq)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("limit orders: %v\n", lr)

		depth, err := cl.Depth(ctx, &pb.Empty{})
		if err != nil {
			log.Println("depth error ", err)
		}

		fmt.Printf("depth: %v\n", depth)

	*/

}
