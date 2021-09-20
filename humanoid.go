package main

import (
	"fmt"
	"math"
	"os"
)

type Persone struct {
	Name    string
	Age     int
	Health  int
	Wallet  int
	IceBox  int
	WorkDay int
	Cost    int
	Energy  int
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

func (chel *Persone) Eat() {
	if chel.Health > 100 {
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

func (chel *Persone) Work() {
	chel.Health -= 30
	chel.Wallet += chel.WorkDay
}

func (chel *Persone) Pay() {
	if chel.Wallet >= chel.Cost {
		chel.Wallet -= chel.Cost
		chel.IceBox += 4
		chel.Health -= 10
	}
}

func do(p PersoneDo, a *Persone) {
	damage := 0.0
	days := 0
	i := float64(a.Age)
	for a.Health > 0 {
		days += 1
		if days > 360 {
			p.Olding()
			days = 0
			i += 1
		}
		//		p.Olding()
		damage = math.Pow(0.04, -i/100)
		a.Health -= int(damage)
		if a.Health <= 0 {
			break
		}
		p.Eat()
		if a.Health > 100 {
			a.Health = 100
		}
		p.Work()
		if a.Health <= 0 {
			break
		}
		p.Pay()
		if a.Health <= 0 {
			break
		}
		p.Eat()
		if a.Health > 100 {
			a.Health = 100
		}
		p.Sleep()
		if a.Health > 100 {
			a.Health = 100
		}
	}
}

func main() {
	var name string
	fmt.Print("inter name")
	fmt.Fscan(os.Stdin, &name)
	persone := &Persone{
		Health:  100,
		IceBox:  4,
		Age:     33,
		Wallet:  20000,
		WorkDay: 300000,
		Cost:    2000,
		Energy:  2,
		Name:    name,
	}
	do(persone, persone)
	fmt.Println(persone.Age)
}
