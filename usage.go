package main

import (
	"fmt"

	"github.com/poroping/go-ios-xe-sdk/client"
)

func main() {
	var HostURL string = "http://172.21.21.29"
	var username string = "terraform"
	var password string = "superpassword"
	// intf := L2VlanNoId{NedVlanListNoId: NedVlanListNoId{
	// 	Name: "vd-TEST",
	// 	// ID:   2301,
	// },
	// }
	// intf := L2Vlan{NedVlanList: NedVlanList{
	// 	ID:   2302,
	// 	Name: "REST-DEVYx",
	// },
	// }
	c, err := client.NewClient(HostURL, username, password, true)
	if err != nil {
		fmt.Println("error", err)
	}
	// resp, err2 := c.ReadL3Vlan(2301)
	// resp, err2 := c.UpdateL3Vlan(2302, intf)
	// resp, err2 := c.CreateL3Vlan(2302, intf)
	// resp, err2 := c.DeleteL3Vlan(2302)
	// resp, err2 := c.ListL3Vlan()
	// resp, err2 := c.ListBgpNeighbor()
	// resp, err2 := c.ReadBgpNeighbor("10.25.0.1")
	// resp, err2 := c.ReadBgpNeighborConfig("10.25.0.1")
	// resp, err2 := c.ListBgpNeighborConfig()
	// resp, err2 := c.ListBgpNeighborPrefixList("10.25.0.1")
	// resp, err2 := c.ReadBgpNeighborPrefixList("10.25.0.1", 43892, "in")
	err2 := c.CreateL2Vlan(2305, nil)
	if err2 != nil {
		fmt.Println("error", err)
	}
	if err2 == nil {
		fmt.Println("success")
	}
	r, err3 := c.ReadL2Vlan(2305)
	if err3 != nil {
		fmt.Println("error", err)
	}
	fmt.Println(r.VlanList.ID)
	err8 := c.UpdateL2Vlan(2305, "farts")
	if err8 != nil {
		fmt.Println("error", err)
	}
	r2, err5 := c.ReadL2Vlan(2305)
	if err5 != nil {
		fmt.Println("error", err)
	}
	fmt.Println(r2.VlanList.ID)
	fmt.Println(r2.VlanList.Name)
	err4 := c.DeleteL2Vlan(2305)
	if err4 != nil {
		fmt.Println("error", err)
	}
	fmt.Println("finished")
}
