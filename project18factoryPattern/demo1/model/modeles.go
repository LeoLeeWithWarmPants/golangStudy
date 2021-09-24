package model

//该结构体的作用范围仅在 package model中
type person struct {
	Id     int64
	Name   string
	Age    int8
	height int //小写字段无法在包外访问，即使在包外拿到了该struct的指针
}

// GetPerson 该函数相当于在本包中创建了一person的实例，但是返回了该实例的指针
func GetPerson(id int64, name string, age int8) *person {
	//这里是在本包使用person，不存在作用域问题
	return &person{
		Id:   id,
		Name: name,
		Age:  age,
	}
}

// GetHeight 通过方法访问私有的字段
func (p person) GetHeight() int {
	return p.height
}
func (p *person) SetHeight(height int) {
	(*p).height = height
}
