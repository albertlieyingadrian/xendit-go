# xendit-go
Example of using Xendit API in Golang

# Installation & Configuration
Ensure you have installed:
1. Golang (https://golang.org/doc/install)
2. Golang package management (https://github.com/Masterminds/glide)

How to run:
1. Run `glide install` to install all your dependencies
2. Go to examples folder
`cd src/test/examples`
3. Create Disbursement
`go run create_disbursement.go`
4. Disbursement Callback Server
`go run disbursement_callback.go`
Your server will run in port 3000
5. Hit localhost:3000/disbursement_callback_url with
	- method: POST
	- header: 'Content-Type: application/json'
	- body: 
	```
	{
	  "user_id": "5785e6334d7b410667d355c4",
	  "external_id": "12345",
	  "amount": 1000,
	  "bank_code": "BCA",
	  "account_holder_name": "RAIDY WIJAYA",
	  "disbursement_description": "Refunds for shoes",
	  "status": "PENDING",
	  "id": "57f1ce05bb1a631a65eee662"
	}
	```