package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name;
}

func (pc PhoneConnecter)Connect()  {
	fmt.Println("connect:",pc.name)
}

func main() {
	a:=PhoneConnecter{"pc"}
	a.Connect()
	Disconnet(a)

}

func Disconnet(usb interface{}){
	//fmt.Print("disconnect")
	switch v:=usb.(type) {
	case PhoneConnecter:
		fmt.Println("disconnect:",v.name)
	default:
		fmt.Print("uknown device")
	}
}
