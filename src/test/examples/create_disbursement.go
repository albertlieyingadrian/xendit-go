package main

import (  
    "fmt"
    "../client"
)

func main () {
    disbursement := client.CreateDisbursement("demo_1475459775872", 17000, "BCA", "Bob Jones", "1231241231")

    fmt.Println("user_id:", disbursement.UserID)
    fmt.Println("external_id:", disbursement.ExternalID)
    fmt.Println("amount:", disbursement.Amount)
    fmt.Println("bank_code:", disbursement.BankCode)
    fmt.Println("account_holder_name:", disbursement.AccountHolderName)
    fmt.Println("disbursement_description:", disbursement.DisbursementDescription)
    fmt.Println("status:", disbursement.Status)
    fmt.Println("id:", disbursement.ID)
}