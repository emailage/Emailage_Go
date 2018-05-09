package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/emailage/Emailage_Go"
)

type config struct {
	Token      string `json:"token"`
	AccountSID string `json:"account_sid"`
	Format     string `json:"format"`
	Endpoint   string `json:"endpoint"`
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
		Format:     c.Format,
		Token:      c.Token,
		AccountSID: c.AccountSID,
		Endpoint:   c.Endpoint,
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
