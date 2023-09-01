package model

import "time"

type Bill struct {
	Id       	string
	BillDate 	time.Time
	EntryDate 	time.Time
	FinishDate 	*time.Time
	Employee 	Employee
	Customer 	Customer
	BillDetail  []BillDetail
}

type BillDetail struct {
	Id 				string
	BillId 			string
	Product 		Product
	ProductPrice 	int
	Qty 			int
}

type BillDetailShow struct {
	Id       	string
	BillDate 	time.Time
	EntryDate 	time.Time
	FinishDate 	*time.Time
	Employee 	Employee
	Customer 	Customer
	BillId 			string
	Product 		Product
	ProductPrice 	int
	Qty 			int
}