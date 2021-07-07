package model

import (

)

// 声明一个 Customer 结构体，表示一个客户信息
type Customer struct {
	Id int
	Name string 
	Gender string 
	Age int 
	Phone string 
	Email string
}

// 使用工厂模式，返回一个 Customer 的实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string ) Customer { 
	return Customer{
		Id : id,
		Name : name, 
		Gender : gender, 
		Age : age, 
		Phone : phone, 
		Email : email,
	}
}