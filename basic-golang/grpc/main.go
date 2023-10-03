package main

import (
	"fmt"
	"log"

	pb "go-proto/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Movement",
				Price: 1000.50,
				Stock: 50,
				Category: &pb.Category{
					Id:   1,
					Name: "Silent",
				},
			},
			{
				Id:    2,
				Name:  "Continously",
				Price: 5000.50,
				Stock: 10,
				Category: &pb.Category{
					Id:   1,
					Name: "Silent",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("Marshall error: ", err)
	}

	// compact binary wire format
	fmt.Println(data)

	testProducts := &pb.Products{}
	if err = proto.Unmarshal(data, testProducts); err != nil {
		log.Fatal("Unmarshal error: ", err)
	}

	for _, product := range testProducts.GetData() {
		fmt.Println(product.GetName())
		fmt.Println(product.GetPrice())
		fmt.Println(product.GetCategory().GetName())
	}
}
