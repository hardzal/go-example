package services

import (
	"context"
	"go-grpc/pb/pagination"
	productPb "go-grpc/pb/product"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
	productPb.UnimplementedProductServiceServer
}

func (p *ProductService) GetProducts(context.Context, *productPb.Empty) (*productPb.Products, error) {

	// products :=
	// 	&productPb.Products{
	// 		Pagination: &pagingPb.Pagination{
	// 			Total:       10,
	// 			PerPage:     5,
	// 			CurrentPage: 1,
	// 			LastPage:    2,
	// 		},
	// 		Data: []*productPb.Product{
	// 			{
	// 				Id:    1,
	// 				Name:  "Metallica T-shirt",
	// 				Price: 1000000.0,
	// 				Stock: 15,
	// 				Category: &productPb.Category{
	// 					Id:   1,
	// 					Name: "Shirt",
	// 				},
	// 			},
	// 			{
	// 				Id:    1,
	// 				Name:  "White T-shirt",
	// 				Price: 500000.0,
	// 				Stock: 30,
	// 				Category: &productPb.Category{
	// 					Id:   1,
	// 					Name: "Shirt",
	// 				},
	// 			},
	// 		},
	// 	}

	var products []*productPb.Product

	rows, err := p.DB.Table("products AS p").
		Joins("LEFT JOIN categories AS c ON c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id as category_id", "c.name as category_name").
		Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var product productPb.Product
		var category productPb.Category

		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
			log.Fatalf("Gagal mengambil row data %v", err.Error())
		}

		product.Category = &category
		products = append(products, &product)
	}

	response := &productPb.Products{
		Pagination: &pagination.Pagination{
			Total:       2,
			PerPage:     1,
			CurrentPage: 1,
			LastPage:    1,
		},
		Data: products,
	}

	return response, nil
}
