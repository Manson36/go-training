package service

import (
	"github.com/go-training/go-code/practice-26/practice02"
)

//对客户信息的增删改查操作
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前有多少客户
	customerNum int

}
//为了看到有客户在切片中，初始化一个客户
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"张三","男",22,
		"1212","souhu@souhu.com")
	customerService.customers = append(customerService.customers,customer)
	return customerService
}
//返回客户切片
func (this *CustomerService)List() []model.Customer {
	return this.customers
}