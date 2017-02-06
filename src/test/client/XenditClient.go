package client

import (  
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "os"

    "../config"
)

const SERVER_DOMAIN = "https://api.xendit.co"
const SECRET_API_KEY = config.SECRET_API_KEY

type Disbursement struct {
    UserID string `json:"user_id"`
    ExternalID string `json:"external_id"`
    Amount int `json:"amount"`
    BankCode string `json:"bank_code"`
    AccountHolderName string `json:"account_holder_name"`
    DisbursementDescription string `json:"disbursement_description"`
    Status string `json:"status"`
    ID string `json:"id"`
}

func CreateDisbursement (externalID string, amount int, bankCode string, accountHolderName string, accountNumber string) Disbursement {
    endPoint := SERVER_DOMAIN + "/disbursements"

    type DisbursementData struct {
        ExternalID string `json:"external_id"`
        Amount int `json:"amount"`
        BankCode string `json:"bank_code"`
        AccountHolderName string `json:"account_holder_name"`
        AccountNumber string `json:"account_number"`
    }

    disbursementData := DisbursementData{
        ExternalID: externalID,
        Amount: amount,
        BankCode: bankCode,
        AccountHolderName: accountHolderName,
        AccountNumber: accountNumber,
    }

    disbursementDataInBytes := new(bytes.Buffer)
    json.NewEncoder(disbursementDataInBytes).Encode(disbursementData)

    req, err := http.NewRequest("POST", endPoint, disbursementDataInBytes)
    req.SetBasicAuth(SECRET_API_KEY, "")
    req.Header.Set("Content-Type", "application/json")

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    client := &http.Client{}
    res, err := client.Do(req)

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    defer res.Body.Close()

    disbursement := Disbursement{}

    data, err := ioutil.ReadAll(res.Body)
    if err == nil && data != nil {
        err = json.Unmarshal(data, &disbursement)
    }

    return disbursement
}