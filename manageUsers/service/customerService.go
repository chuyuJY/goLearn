package service

import "goLearn/manageUsers/model"

// CustomerService 完成对Customer的crud
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段，表示当前切片含有客户数量
	customerNum int
}

func NewCustomerService() *CustomerService {
	customer := model.NewCustomer(1, "chuyu", "男", 23, "123455", "554321@qq.com")
	return &CustomerService{
		customers:   []model.Customer{customer},
		customerNum: 1,
	}

}

// List 返回客户列表
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// Add 添加客户
func (cs *CustomerService) Add(name string, gender string, age int, phone string, email string) bool {
	cs.customerNum++
	newCustomer := model.NewCustomer(cs.customerNum, name, gender, age, phone, email)
	cs.customers = append(cs.customers, newCustomer)
	return true
}

// Delete 删除用户
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}
	// 从切片中删除
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

// Update 更新用户
func (cs *CustomerService) Update(index int, name string, gender string, age int, phone string, email string) {
	cs.customers[index].Name = name
	cs.customers[index].Gender = gender
	cs.customers[index].Age = age
	cs.customers[index].Phone = phone
	cs.customers[index].Email = email
}

// 查找Slice下标
func (cs *CustomerService) FindById(id int) int {
	index := -1
	for i, user := range cs.customers {
		if user.Id == id {
			index = i
			break
		}
	}
	return index
}
