package model

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