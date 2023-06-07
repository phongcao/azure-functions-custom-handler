package main

import (
	"fmt"
)

type TaskType int

func (m TaskType) Task1() {
	fmt.Println("Invoked Task1")
}

func (m TaskType) Task2() {
	fmt.Println("Invoked Task2")
}

func (m TaskType) Task3() {
	fmt.Println("Invoked Task3")
}
