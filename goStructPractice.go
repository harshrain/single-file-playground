package main

import (
	"fmt"
	"strconv"
)

// A :		Function to be used in struct
type FormatTransactionMessage func(tx *Transaction) string

//Declaring a struct
type Transaction struct {
	id     string
	from   string
	to     string
	amount int
	//Function inside struct -------------------- A
	getTransactionDetails FormatTransactionMessage
}

//Calling a function for a created instance of the structure
func (tx *Transaction) validate() bool {
	if tx.id == "" {
		return false
	}
	return true
}

func main() {
	// var dummyTransaction Transaction = Transaction{"tr1", "harsh@rain.bh", "piyush@rain.bh", 1000}

	//Creating instance of the Transaction Structure with valid value in id field and passing anonymous function in same layout as declared @A and used in Transaction struct
	anotherDummyTransaction := Transaction{id: "tr2", amount: 2000, from: "harsh@rain.bh", to: "piyush@rain.bh", getTransactionDetails: func(tx *Transaction) string {
		return "From : " + tx.from + " | " + "To : " + tx.to + " | " + "Amount : " + strconv.Itoa(tx.amount) + " | " + "ID : " + tx.id
	}}

	//Creating instance of the Transaction Structure with invalid/null/empty value in id field and passing anonymous function in same layout as declared @A and used in Transaction struct
	withoutID := Transaction{amount: 2000, from: "harsh@rain.bh", to: "piyush@rain.bh", getTransactionDetails: func(tx *Transaction) string {
		return "From : " + tx.from + " | " + "To : " + tx.to + " | " + "Amount : " + strconv.Itoa(tx.amount)
	}}

	// fmt.Println(dummyTransaction)
	fmt.Println(anotherDummyTransaction)
	fmt.Println("In format : ", anotherDummyTransaction.getTransactionDetails(&anotherDummyTransaction), "\n")
	fmt.Println(withoutID)
	fmt.Println("In format : ", withoutID.getTransactionDetails(&withoutID), "\n\n")

	fmt.Println("Validating withoutID.......")
	fmt.Println(withoutID.validate(), "\n")
	fmt.Println("Validating anotherDummyTransaction........")
	fmt.Println(anotherDummyTransaction.validate())
}
