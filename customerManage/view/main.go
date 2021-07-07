package main

import(
	"fmt"
	service "go-study/customerManage/service"
)


type customerView struct {
	// 定义必要字段
	key string // 接收用户输入...
	loop bool	// 表示是否循环的显示主菜单
	customerService *service.CustomerService // 增删改查service
}


// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("客户信息管理软件")
		fmt.Println("1.添加客户")
		fmt.Println("2.修改客户")
		fmt.Println("3.删除客户")
		fmt.Println("4.客户列表")
		fmt.Println("5.退出")
		fmt.Print("请选择(1-5)：")
		
		fmt.Scanln(&this.key)

		switch this.key {
			case "1":
				fmt.Println("添加客户")
			case "2":	
				fmt.Println("修改客户")
			case "3":
				fmt.Println("删除客户")
			case "4":
				fmt.Println("客户列表")
			case "5":
				this.loop = false
			default :
				fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop { 
			break
		}
		fmt.Println()
	}
	fmt.Println("你退出了客户关系管理系统...")
}
func main() {
	// 在main 函数中，创建一个 customerView,并运行显示主菜单.. 
	customerView := customerView {
		key : "", 
		loop : true,
	}

	//这里完成对 customerView 结构体的 customerService 字段的初始化
	customerView.customerService = service.NewCustomerService()

	// 显示主菜单.. 
	customerView.mainMenu()
}