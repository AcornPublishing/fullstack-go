package dblayer

import (
	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter07/backend/src/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) error
	SignInUser(models.Customer) error
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}
