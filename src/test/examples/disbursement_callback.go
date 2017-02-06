package main

import (  
    "encoding/json"
    "fmt"
    "net/http"

    "../../../vendor/github.com/julienschmidt/httprouter"
)

func RunServer () {
    router := httprouter.New()

    router.POST("/disbursement_callback_url", func (responseWriter http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

        decoder := json.NewDecoder(req.Body)

        disbursementData := Disbursement{}

        err := decoder.Decode(&disbursementData)

        if err != nil {
            panic(err)
        }

        defer req.Body.Close()

        disbursement, _ := json.Marshal(disbursementData)

        responseWriter.Header().Set("Content-Type", "application/json")
        responseWriter.WriteHeader(200)
        fmt.Fprintf(responseWriter, "%s", disbursement)
    })

    fmt.Println("Your server is running on port 3000")
    http.ListenAndServe("localhost:3000", router)
}

func main() {  
    RunServer()
}