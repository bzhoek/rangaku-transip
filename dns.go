package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/transip/gotransip"
	"github.com/transip/gotransip/domain"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("Usage:", os.Args[0], "username", "private_key", "domainName", "hostName")
		return
	}

	username := os.Args[1]
	private_key := os.Args[2]
	domainName := os.Args[3]
	hostName := os.Args[4]

	req, _ := http.NewRequest("GET", "https://api.ipify.org", nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	ipv4, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("Update %s on %s in account %s with external IPv4 %s\n", hostName, domainName, username, string(ipv4))

	c, err := gotransip.NewSOAPClient(gotransip.ClientConfig{AccountName: username, PrivateKeyPath: private_key})
	if err != nil {
		panic(err.Error())
	}

	list, err := domain.GetInfo(c, domainName)
	if err != nil {
		panic(err.Error())
	}

	entries := updateEntries(list.DNSEntries, hostName, ipv4)

	fmt.Printf("Updated DNS entries %+v\n", entries)

	err = domain.SetDNSEntries(c, domainName, entries)
	if err != nil {
		panic(err.Error())
	}

}

func updateEntries(entries []domain.DNSEntry, hostname string, ipv4 []byte) []domain.DNSEntry {

	entries = filter(entries, hostname)
	entries = filter(entries, fmt.Sprintf("_%s", hostname))

	entries = append(entries, domain.DNSEntry{hostname, 60 * 60 * 24, domain.DNSEntryTypeA, string(ipv4)})
	entries = append(entries, domain.DNSEntry{fmt.Sprintf("_%s", hostname), 60 * 60 * 24, domain.DNSEntryTypeTXT, time.Now().String()})

	return entries
}

func filter(entries []domain.DNSEntry, name string) []domain.DNSEntry {
	n := 0
	for _, v := range entries {
		if v.Name != name {
			entries[n] = v
			n++
		}
	}
	return entries[:n]
}
