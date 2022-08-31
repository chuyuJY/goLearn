package model

import "fmt"

// Customer 客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式， 返回Customer实例
func NewCustomer(id int, name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 格式化用户信息
func (c Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
}
