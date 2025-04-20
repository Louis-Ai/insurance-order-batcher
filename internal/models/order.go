package models

import "time"

type Order struct {
	CustomerID string
	Address    Address
	OrderTime  time.Time
}

type Address struct {
	AddressLineOne string
	TownCity       string
	Postcode       string
}
