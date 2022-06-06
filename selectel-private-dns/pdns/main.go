package pdns

import (
	"github.com/joeig/go-powerdns/v2"
	"log"
)

type pDnsClient struct {
	client *powerdns.Client
}

type PDnsConfig struct {
	Url string
	Domain string
	ApiKey string
	Vhost string
}


func NewPowerDns(config PDnsConfig) *pDnsClient {
	client := powerdns.NewClient(config.Url, config.Vhost, map[string]string{"X-API-Key": config.ApiKey}, nil)
	c := new(pDnsClient)
	c.client = client
	return c
}


func (c *pDnsClient) AddOrUpdateRecord(domain string, name string, ttl uint32, values []string ) {
	err := c.client.Records.Add(domain, name, powerdns.RRTypeA, ttl, values)
	if err != nil {
		log.Fatalf(err.Error())
	}
}


	// zones, err := pdns.Zones.List()
   //zone, err := pdns.Zones.AddNative(domain, false, "", false, "", "", false, []string{"ns1.sandbox."})
	//if err != nil {
	//	log.Panicln(err)
	//}

