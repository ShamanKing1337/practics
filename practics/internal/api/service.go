package service

import "practics/internal/pkg"

type Implementation struct {
	usersService  pkg.UsersService
	itemsService pkg.ItemsService
	purchasesService pkg.PurchasesService
}

func NewPracticsApiService(
	usersService  pkg.UsersService,
	itemsService pkg.ItemsService,
	purchasesService pkg.PurchasesService,
) *Implementation {
	return &Implementation{
		usersService:  usersService,
		itemsService: itemsService,
		purchasesService: purchasesService,
	}
}


