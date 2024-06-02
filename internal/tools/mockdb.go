package tools

import (
	"fmt"
	"time"
)

type mockDB struct{}

// Fake data that we are using to mock the Database storing info
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

// create get user login details func
// create get coins func
// create set up database func

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call, use time.sleep
	time.Sleep(time.Second + 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		fmt.Println("Failed to get client login details")
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second + 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		fmt.Println("Failed to get client coin details")
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	fmt.Println("Set up new mocked database...")
	return nil
}
