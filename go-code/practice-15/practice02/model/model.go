package model

import (
	"fmt"
)

//客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}
func NewCustomer(id int, Name string,gender string ,age int,
	phone string,email string) Customer {
	return Customer{
		Id :id,
		Name:Name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,

	}
}
//返回用户信息，格式化字符串
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,
		this.Name,this.Gender,this.Age, this.Phone,this.Email)
	return info
}
func NewCustomer2( Name string,gender string ,age int,
	phone string,email string) Customer {
	return Customer{
		Name:Name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,

	}
}
