package main
/*
State поведенческий паттерн 
позволяет управлять поведением объекта в зависимости от состояния.
*/

import "fmt"

type Activity interface {
	justDoIt()
}

type Coding struct{}

func (coding *Coding) justDoIt() {
	fmt.Println("Writing code...")
}

type Reading struct{}

func (reading *Reading) justDoIt() {
	fmt.Println("Reading book...")
}

type Sleeping struct{}

func (sleeping *Sleeping) justDoIt() {
	fmt.Println("Sleeping...")
}

type Training struct{}

func (training *Training) justDoIt() {
	fmt.Println("Training...")
}

type Developer struct {
	activity Activity
}

func (developer *Developer) setActivity(activity Activity) {
	developer.activity = activity
}

func (developer *Developer) changeActivity() {
	switch developer.activity.(type) {
	case *Sleeping:
		developer.setActivity(&Training{})
	case *Training:
		developer.setActivity(&Coding{})
	case *Coding:
		developer.setActivity(&Reading{})
	case *Reading:
		developer.setActivity(&Sleeping{})
	}
}

func (developer *Developer) justDoIt() {
	developer.activity.justDoIt()
}

func main() {
	activity := Sleeping{}
	developer := Developer{}
	developer.setActivity(&activity)
	for i := 0; i < 10; i++ {
		developer.justDoIt()
		developer.changeActivity()
	}
}
