package internal

type Human struct {
	name string
	age  int
}

func (a *Human) Name() string {
	return a.name
}

func (a *Human) SetName(name string) {
	a.name = name
}

func (a *Human) Age() int {
	return a.age
}

func (a *Human) SetAge(age int) {
	a.age = age
}

type Action struct {
	Human
}

func Task1() {
	action := Action{}
	action.SetAge(30)
	action.SetName("John")
	println(action.Age())
	println(action.Name())
}
