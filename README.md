# goapi
This is a practice of building go api to get someone's coin balance

## Verified through Postman
Clone the repo, under `goapi/`, run `go run cmd/api/main.go`
Go to Postman, create a new `GET` request with `http://localhost:8000/account/coins/?username=alex`
Go to `Headers` section, input the token, for example `key: Authorization, value: 123ABC`
Run Send, you will see the returned output
```
{
    "Code": 200,
    "Balance": 100
}
```
If failed, you will see:
```
{
    "Code": 400,
    "Message": "An unexpected error occurred..."
}
```