package main

import (
	"fmt"
	"goLearn/manageUsers/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

// 1. 显示主菜单
func (cv *customerView) viewMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1. 添加用户")
		fmt.Println("                 2. 修改用户")
		fmt.Println("                 3. 删除用户")
		fmt.Println("                 4. 客户列表")
		fmt.Println("                 5. 退    出")
		fmt.Printf("请选择[1-5]:")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.addUser()
		case "2":
			//fmt.Println("2. 修改用户")
			cv.update()
		case "3":
			//fmt.Println("3. 删除用户")
			cv.delete()
		case "4":
			//fmt.Println("4. 客户列表")
			cv.viewList()
		case "5":
			cv.exit()
		default:
			fmt.Println("输入错误，请重新输入......")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("已退出，欢迎下次使用......")
}

func (cv *customerView) viewList() {
	userList := cv.customerService.List()
	fmt.Println("-----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, user := range userList {
		//fmt.Println(users.Id, "\t", users.Name...)
		fmt.Println(user.GetInfo())
	}
	fmt.Println("---------------客户列表完成---------------")
}

func (cv *customerView) addUser() {
	fmt.Println("-----------------添加客户-----------------")
	name, gender, age, phone, email := "", "", 0, "", ""
	fmt.Printf("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Printf("请输入性别：")
	fmt.Scanln(&gender)
	fmt.Printf("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Printf("请输入电话号：")
	fmt.Scanln(&phone)
	fmt.Printf("请输入邮箱：")
	fmt.Scanln(&email)
	// 系统分配 id
	if cv.customerService.Add(name, gender, age, phone, email) {
		fmt.Println("-----------------添加成功-----------------")
	} else {
		fmt.Println("-----------------添加失败-----------------")
	}
}

func (cv *customerView) delete() {
	fmt.Println("-----------------删除客户-----------------")
	fmt.Printf("请输入待删除客户编号id(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Printf("请确认是否删除(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		if cv.customerService.Delete(id) {
			fmt.Println("删除完成...")
		} else {
			fmt.Println("删除失败，该id号不存在")
		}
	}
}

func (cv *customerView) update() {
	fmt.Println("-----------------修改客户-----------------")
	id := -1
	fmt.Printf("请输入要修改的客户编号(id=-1退出)：")
	fmt.Scanln(&id)
	for {
		if id == -1 {
			fmt.Println("删除操作已退出...")
			return
		}
		index := cv.customerService.FindById(id)
		if index != -1 {
			fmt.Println("直接回车将保持原值...")
			oldInfo := cv.customerService.List()[index]
			name, gender, age, phone, email := oldInfo.Name, oldInfo.Gender, oldInfo.Age, oldInfo.Phone, oldInfo.Email
			fmt.Printf("姓名【%v】：", name)
			fmt.Scanln(&name)
			fmt.Printf("性别【%v】：", gender)
			fmt.Scanln(&gender)
			fmt.Printf("年龄【%v】：", age)
			fmt.Scanln(&age)
			fmt.Printf("电话号码【%v】：", phone)
			fmt.Scanln(&phone)
			fmt.Printf("邮箱【%v】：", email)
			fmt.Scanln(&email)
			cv.customerService.Update(index, name, gender, age, phone, email)
			fmt.Println("修改成功...")
			return
		} else {
			fmt.Printf("此用户编号不存在，请重新输入要修改的客户编号(id=-1退出)：")
			fmt.Scanln(&id)
		}
	}
}

func (cv *customerView) exit() {
	fmt.Printf("请确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "Y" || cv.key == "n" || cv.key == "N" {
			break
		} else {
			fmt.Printf("输入有误，请确认是否退出(Y/N)：")
		}
	}
	if cv.key == "y" || cv.key == "Y" {
		cv.loop = false
	}
}

func main() {
	cv := customerView{
		key:  "",
		loop: true,
	}

	cv.customerService = service.NewCustomerService()
	cv.viewMenu()
}
