package service

import "github.com/go-training/go-code/practice-14/practice03/01"

type CustomerService struct {
	customers []model.Customer
	customerNum int
}