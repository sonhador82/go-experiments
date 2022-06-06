package main

import (
	"fmt"
	"net"
	"os"
	"sandbox/powerdns/pdns"
	"sandbox/powerdns/selectel"
	"sync"
)


func main() {
	s := selectel.NewSelectel(&selectel.Config{
		Username:          os.Getenv("OS_USERNAME"),
		Password:          os.Getenv("OS_PASSWORD"),
		ProjectDomainName: os.Getenv("OS_PROJECT_DOMAIN_NAME"),
		ProjectName:       os.Getenv("OS_PROJECT_NAME"),
	})

	_, netBlock, err := net.ParseCIDR("192.168.0.0/24")
	if err != nil {
		panic(err)
	}

	servers := s.GetServerList(*netBlock)

	var wg sync.WaitGroup
	lock := &sync.Mutex{}


	for _, srv := range servers {
		wg.Add(1)
		go func(addrPair selectel.ServerPair, lock *sync.Mutex) {
			fmt.Println(addrPair)
			// добавиьт в API)
			config := pdns.PDnsConfig{
				Url:    "http://localhost:8081",
				Domain: "sandbox",
				ApiKey: "secret",
				Vhost:  "localhost",
			}
			p := pdns.NewPowerDns(config)
			lock.Lock()
			p.AddOrUpdateRecord("sandbox", fmt.Sprintf("%s.%s", addrPair.Name, "sandbox"), 3600, []string{addrPair.IP.String()})
			lock.Unlock()
			wg.Done()
		}(srv, lock)
	}
	wg.Wait()
}
