package main

type TaskType int

func (m TaskType) Task1() string {
	return "Invoked Task1"
}

func (m TaskType) Task2() string {
	return "Invoked Task2"
}

func (m TaskType) Task3() string {
	return "Invoked Task3"
}
