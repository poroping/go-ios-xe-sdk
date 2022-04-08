## IOSXE Client (Go)

```go
package main

import (
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/client"
	"github.com/poroping/go-ios-xe-sdk/config"
	"github.com/poroping/go-ios-xe-sdk/models"
)

func main() {
	username := "cisco"
	password := "cisco"
	host := "192.168.1.1"
	insecure := true
	userAgent := "go-sdk-iosxe"
	cfg := config.Config{
		Username:  username,
		Password:  password,
		Host:      host,
		Insecure:  insecure,
		UserAgent: userAgent,
	}
	c, err := client.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	m := &models.BgpNeighbor{}
	m.Neighbor.ASN = 65421
	i := 64442
	m.Neighbor.RemoteAs = &i
	m.Neighbor.ID = "8.8.8.1"

	m2 := models.BgpNeighborConfig{}
	m2.NeighborConfig.Activate = &models.CiscoEnabled
	m2.NeighborConfig.DefaultOriginate = &struct{}{}
	m2.NeighborConfig.ASN = m.Neighbor.ASN
	m2.NeighborConfig.ID = m.Neighbor.ID
	m2.NeighborConfig.AddressFamilyType = "ipv4"

	err = c.CreateBgpNeighbor(*m)
	if err != nil {
		log.Fatal(err)
	}

	err = c.CreateBgpNeighborConfig(m2)
	if err != nil {
		log.Fatal(err)
	}

	err = c.DeleteBgpNeighbor(*m)
	if err != nil {
		log.Fatal(err)
	}

	m3 := &models.BgpNeighborConfig{}
	m3.NeighborConfig.ASN = 65421
	m3.NeighborConfig.Activate = &models.CiscoEnabled
	m3.NeighborConfig.ID = "8.8.8.8"
	m3.NeighborConfig.AddressFamilyType = "ipv4"
	i2 := 64442
	m3.NeighborConfig.RemoteAs = &i2
	vrf := "FLAB"
	m3.NeighborConfig.Vrf = &vrf

	err = c.CreateBgpNeighborConfig(*m3)
	if err != nil {
		log.Fatal(err)
	}

	read := &models.BgpNeighborConfig{}
	read.NeighborConfig.ID = "8.8.8.8"
	read.NeighborConfig.Vrf = &vrf
	read.NeighborConfig.ASN = 65421
	read.NeighborConfig.AddressFamilyType = "ipv4"

	r, err := c.ReadBgpNeighborConfig(*read)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)

	err = c.DeleteBgpNeighborConfig(*read)
	if err != nil {
		log.Fatal(err)
	}
}
```
