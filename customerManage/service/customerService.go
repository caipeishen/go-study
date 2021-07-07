package service

import (
	"fmt"
	model "go-study/customerManage/model"
)

type CustomerService struct { 
	Customers []model.Customer //声明一个字段，表示当前切片含有多少个客户
}

// 客户service工厂
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customer := model.NewCustomer(1, "采培珅", "男", 21, "18848848551", "peishen.cai@foxmail.com")
	customerService.Customers = append(customerService.Customers, customer)
	return customerService
}

// 添加客户
func (cs *CustomerService) AddCustomer(name string, gender string, age int, phone string, email string) {
	cs.Customers = append(cs.Customers, model.NewCustomer(len(cs.Customers)+1, name, gender, age, phone, email))
	return
}

// 返回列表数据
func (cs *CustomerService) List() []model.Customer {
	return cs.Customers
}

// 根据id获取客户
func (cs *CustomerService) GetById(id int) int {
	index := -1
	for i, v := range cs.Customers {
		if id == v.GetId() {
			index = i
			break
		}
	}
	return index
}

// 根据id删除客户
func (cs *CustomerService) DeleteId(id int) {
	index := cs.GetById(id)
	if index == -1 {
		fmt.Println("改用户不存在,无法删除")
	} else {
		cs.Customers = append(cs.Customers[:index], cs.Customers[index+1:]...)

	}
}