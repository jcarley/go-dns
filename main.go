package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/dnsimple/dnsimple-go/dnsimple"
)

type Router struct {
	IP string
}

const (
	settingsFileName = "./settings.json"
)

func main() {

	// First checks for DNSIMPLE_OAUTH_TOKEN environment variable.  If the
	// variable is not found, checks the config file.
	authToken, err := GetAuthToken()
	if err != nil {
		// need better logging
		panic(err)
	}

	// Look up the router's IP
	routerIP, err := getRouterIP()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error determining router IP: \n\n%s\n", err)
	}

	fmt.Printf("Router IP: %s\n", routerIP)

	// new client
	client := dnsimple.NewClient(dnsimple.NewOauthTokenCredentials(authToken))

	// get the current authenticated account (if you don't know who you are)
	whoamiResponse, err := client.Identity.Whoami()
	if err != nil {
		fmt.Printf("Whoami() returned error: %v\n", err)
		os.Exit(1)
	}

	// either assign the account ID or fetch it from the response
	// if you are authenticated with an account token
	fmt.Printf("%#v\n", whoamiResponse.Data.Account.ID)
	accountID := strconv.Itoa(whoamiResponse.Data.Account.ID)

	// get the list of domains
	domainsResponse, err := client.Domains.ListDomains(accountID, &dnsimple.DomainListOptions{NameLike: "domain name here"})

	if err != nil {
		panic(err)
	}

	// for _, domain := range domainsResponse.Data {
	//   fmt.Printf("%#v\n", domain)
	// }

	domain := domainsResponse.Data[0]

	zoneRecordsResponse, _ := client.Zones.ListRecords(accountID, domain.Name, &dnsimple.ZoneRecordListOptions{Type: "A"})
	// zoneRecordsResponse, _ := client.Zones.ListRecords(accountID, domain.Name, nil)

	fmt.Printf("%#v\n", zoneRecordsResponse.Data)

	os.Exit(0)

	// 4) Check the current IP of each domain
	// domains := config.LoadAllDomains()
	// for _, domain := range domains {
	//   records, _ := client.Records(domain.Name, "", domain.RecordType)
	//   for _, record := range records {
	//     if record.RecordType == domain.RecordType {
	//       if routerIP == record.Content {
	//         fmt.Printf("DNS for '%s' matches\n", domain.Name)
	//       } else {
	//         // 5) If IP's don't match; update DNS record
	//         fmt.Printf("DNS for '%s' does not match\n", domain.Name)
	//         fmt.Printf("  Updating '%s' to '%s'\n", record.Content, routerIP)
	//         record.UpdateIP(client, routerIP)
	//       }
	//     }
	//   }
	// }

}

func getRouterIP() (string, error) {

	url := "https://jsonip.com"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	routerIP := Router{}

	err = json.Unmarshal(responseBytes, &routerIP)

	if err != nil {
		return "", err
	}

	return routerIP.IP, nil
}
