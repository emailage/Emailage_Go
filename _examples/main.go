package main

import (
	"encoding/json"
	"fmt"
	"github.com/emailage/Emailage_Go"
	"github.com/emailage/Emailage_Go/auth"
	"io/ioutil"
	"os"
	"time"
)

type config struct {
	Token         string                        `json:"token"`
	AccountSID    string                        `json:"account_sid"`
	Format        string                        `json:"format"`
	Endpoint      string                        `json:"endpoint"`
	TokenEndpoint string                        `json:"token_endpoint"`
	AuthType      emailage.AuthenticationScheme `json:"auth_type"`
}

func main() {
	f, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var c config
	if err := json.Unmarshal(b, &c); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	opts := &emailage.ClientOpts{
		Token:         c.Token,
		TokenEndpoint: c.TokenEndpoint,
		AccountSID:    c.AccountSID,
		Endpoint:      c.Endpoint,
		AuthType:      c.AuthType,
		HTTPTimeout:   3 * time.Second,
		Algorithm:     auth.HMACSHA512,
	}
	client, err := emailage.New(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.EmailOnlyScore("nigerian.prince@legit.ru", nil)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Printf("Result: %+v\n", res.Query)
}
