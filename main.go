package main

import (
	"fmt"
	. "github.com/Mathew-Estafanous/funGo/model"
	. "github.com/Mathew-Estafanous/funGo/stream"
)

type Employee struct {
	name string
	title string
	salary float32
}

func (e Employee) Equals(me Model) bool {
	empl, ok := me.(Employee)
	if !ok {
		return false
	}

	if empl.name == e.name && empl.salary == e.salary {
		return true
	}
	return false
}

func main() {
	allEmployees := getAllEmployees()

	result := NewStreamFromSlice(allEmployees).
			Map(func(m Model) Model {
				if m.(Employee).title != "Developer" {
					return m
				}
				developer := m.(Employee)
				developer.salary *= 1.5
				return developer
			}).
			Filter(func(m Model) bool {
				return m.(Employee).salary <= 100000
			}).
			Collect(ToSlice())

	fmt.Println(result)
}

func getAllEmployees() ModelSlice {
	return ModelSlice {
		Employee{
			name: "Alex",
			title: "Developer",
			salary: 72000,
		},
		Employee{
			name: "Rebecca",
			title: "Manager",
			salary: 84000,
		},
		Employee{
			name: "Joshua",
			title: "Developer",
			salary: 65000,
		},
	}
}
