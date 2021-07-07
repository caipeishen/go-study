package service

import (
	model "go-study/customerManage/model"
)

type CustomerService struct { 
	customers []model.Customer //声明一个字段，表示当前切片含有多少个客户
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customer := model.NewCustomer(1, "张三", "男", 20, "18848848551", "peishen.cai@foxmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}