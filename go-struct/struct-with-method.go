package main

import "fmt"

type Address struct {
	Street, City, State, Zip string
	IsShippingAddress        bool
}

type Customer struct {
	FirstName, LastName, Email, Phone string
	Address                           []Address
}

func (c Customer) ToString() string {
	return fmt.Sprintf("Customer: %s %s, Email:%s", c.FirstName, c.LastName, c.Email)
}

func (c Customer) ShippingAddress() string {
	for _, v := range c.Address {
		if v.IsShippingAddress == true {
			return fmt.Sprintf("%s, %s, %s, Zip - %s", v.Street, v.City, v.State, v.Zip)
		}
	}
	return ""
}

func main() {
	c := Customer{
		FirstName: "Alex",
		LastName:  "John",
		Email:     "alex@email.com",
		Phone:     "732-757-2923",
		Address: []Address{
			Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
			Address{
				Street: "49 Stevenson Street",
				City:   "San Francisco",
				State:  "CA",
				Zip:    "94105",
			},
		},
	}
	fmt.Println(c.ToString())
	fmt.Println(c.ShippingAddress())
}
