package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/transip/gotransip"
	"github.com/transip/gotransip/domain"
)

func main() {

	home, err := os.UserHomeDir()

	req, _ := http.NewRequest("GET", "https://api.ipify.org", nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	ipv4, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(ipv4))

	c, err := gotransip.NewSOAPClient(gotransip.ClientConfig{AccountName: "bzhoek", PrivateKeyPath: fmt.Sprintf("%s/%s", home, ".ssh/transip")})
	if err != nil {
		panic(err.Error())
	}

	list, err := domain.GetInfo(c, "rangaku.com")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", list)

	entries := list.DNSEntries

	for i, v := range entries {
		if v.Name == "basmini" {
			fmt.Println(v.Name)
			entries[i].Content = string(ipv4)
		}
		fmt.Printf("%+v\n", v)
	}
	fmt.Printf("%+v\n", entries)

	err = domain.SetDNSEntries(c, "rangaku.com", list.DNSEntries)
	if err != nil {
		panic(err.Error())
	}

}
