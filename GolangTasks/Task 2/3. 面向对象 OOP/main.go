package main

import "fmt"

// 题目 1：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。


type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Length + r.Width) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 3.14 * 2 * c.Radius
}


// 题目 2：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工姓名:%v\t 年龄:%v\t 员工编号:%v\n", e.Name, e.Age, e.EmployeeID)
}



func main() {
	// Question 1: 接口 + 两个结构体实现
	fmt.Println("[OOP-1] Shape 接口 + Rectangle/Circle 实现")

	var rect = Rectangle{Width: 6, Length: 4}
	var circle = Circle{Radius: 5}

	var a Shape = rect
	var b Shape = circle

	fmt.Printf("矩形面积:%v\t 矩形周长:%v\n", a.Area(), a.Perimeter())
	fmt.Printf("园面积:5%f\t 园周长:5%f\n", b.Area(), b.Perimeter())

	// Question 2:
	var p1 = Person{"韩德水烧鸡", 190}
	var e1 = Employee{p1, "E-1001"}

	var e2 = Employee{
		Person:     Person{Name: "吊炉盐烧饼", Age: 98},
		EmployeeID: "E-1002",
	}

	e1.PrintInfo()
	e2.PrintInfo()

}
