package mock

import (
	model "github.com/brunoserralheiro/showroom-go-rest-api/main/model"
)

// ===============
// mock data @todo implement DB
func MockProducts() *[]model.Product {

	user0 := &model.User{
		ID:    "0101",
		Name:  "name",
		Email: "email",
		Rank:  "seller", //master, admin, manager, seller, buyer
	}
	user2 := &model.User{
		ID:    "0102",
		Name:  "name2",
		Email: "email2",
		Rank:  "admin", //master, admin, manager, seller, buyer
	}

	productz := []model.Product{
		model.Product{
			ID:    "001",
			Name:  "name",
			Price: "3000",
			Owner: user0,
		},
		model.Product{
			ID:    "002",
			Name:  "name2",
			Price: "price",
			Owner: user2,
		},
	}
	return &productz
}
