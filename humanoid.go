package main

import (
	"fmt"
	"math"
)

type Persone struct {
	Name string
	Age int
	Health int
	Wallet int
	IceBox int
	WorkDay int
	Cost int
	Energy int
}

type PersoneDo interface {
	Eat()
	Sleep()
	Work()
	Pay()
	Olding()
}

func (chel *Persone) Olding() {
	chel.Age += 1
}

func (chel *Persone) Eat(){
	if chel.Health > 100{
		chel.Health = 100
	}
	if chel.IceBox >= chel.Energy {
		chel.Health += 10
		chel.IceBox -= chel.Energy
	}
}
func (chel *Persone) Sleep() {
	chel.Health += 40
}

func (chel *Persone) Work(){
	chel.Health -= 50
	chel.Wallet += chel.WorkDay
}

func (chel *Persone)  Pay(){
	if chel.Wallet >= chel.Cost {
		chel.Wallet -= chel.Cost
		chel.IceBox += 4
	}
}

func do(p PersoneDo,a *Persone) {
	damage := 0.0
	i := float64(a.Age)
	for a.Health > 0{
		p.Olding()
		i += 1
		damage = math.Pow(0.04, -i / 100)
		a.Health -= int(damage)
		if a.Health <= 0{
			break
		}
		p.Eat()
		p.Work()
		if a.Health <= 0{
			break
		}
		p.Pay()
		if a.Health <= 0{
			break
		}
		p.Eat()
		p.Sleep()
		fmt.Println(a.Age)
	}
}

func main()  {
	persone := &Persone{
		Health:  100,
		IceBox: 4,
		Age: 33,
		Wallet: 20000,
		WorkDay: 1000,
		Cost: 2000,
		Energy: 2,
	}
	do(persone, persone)
	fmt.Println(persone.Age)
}
