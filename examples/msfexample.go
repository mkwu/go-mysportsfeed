package main

import (
	"flag"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	msf "github.com/mkwu/go-mysportsfeed"
)

func main() {
	userPtr := flag.String("user", "", "The username to your mysportsfeed account. (https://www.mysportsfeeds.com/register)")
	passPtr := flag.String("pass", "", "The password to your mysportsfeed account. (https://www.mysportsfeeds.com/register)")
	flag.Parse()

	client := msf.NewClient(nil, *userPtr, *passPtr)
	opts := msf.CumulativePlayerStatsOptions{
		Limit: 10,
		Team:  "chi",
	}
	data, err := client.NBA.CumulativePlayerStats("2017-2018-regular", opts)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, p := range data {
			spew.Dump(p)
		}
	}
}
