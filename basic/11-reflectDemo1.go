package main

import (
	"fmt"
	"reflect"
)

// 1. example1
// 通过反射拿到data的type、kind、值
func reflectDemo1(data interface{}) {
	// 1. 获取reflect.Type
	refType := reflect.TypeOf(data)
	fmt.Println("refType:", refType, "refType's type:", reflect.TypeOf(refType))
	// 2. 获取reflect.Value
	refValue := reflect.ValueOf(data)
	fmt.Println("refValue:", refValue)
	// 操作refValue
	numInt64 := int64(1)
	// 不能直接操作rVal类型
	//num += refValue
	numInt64 += refValue.Int()
	fmt.Println("numInt64 =", numInt64)
	// 3. 将rVal转成interface{}
	iv := refValue.Interface()
	// 将 interface{} 通过类型断言转成需要的类型
	fmt.Println("iv type = ", reflect.TypeOf(iv))
	numInt := 1
	numInt += iv.(int)
	fmt.Println("numInt =", numInt)
}

// 2. 对结构体的reflect
type Student struct {
	Name string
	Age  int
}

func reflectDemo2(data interface{}) {
	// 1. 获取reflect.Type
	refType := reflect.TypeOf(data)
	fmt.Println("refType:", refType, "  refType's type:", reflect.TypeOf(refType))
	// 2. 获取reflect.Value
	refValue := reflect.ValueOf(data)
	fmt.Println("refValue:", refValue)
	// 3. 获取变量对应的Kind
	fmt.Printf("refType kind = %v\n", refType.Kind())
	fmt.Printf("refValue kind = %v\n", refValue.Kind())
	// 4. 转成 interface{}
	iv := refValue.Interface()
	fmt.Printf("iv = %v, iv type = %T\n", iv, iv)
	// 即使当前输出是Student类型，也不能取出数据
	// 5. 通过类型断言转换成相应类型
	stu := iv.(Student)
	fmt.Println("stu.Name =", stu.Name, ", stu.Age =", stu.Age)
}

func main() {
	// 1. example1
	//reflectDemo1(12)

	// 2. example2
	student := Student{
		Name: "Lily",
		Age:  12,
	}
	reflectDemo2(student)
}
