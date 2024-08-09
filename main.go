package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jamous/pingsweep"
)

func main() {
	//Check if address was passed
	var args string
	if len(os.Args) > 1 {
		args = os.Args[1]
	}

	//Build PSconfig. Include arg pos 1 if provided.
	psconfig := pingsweep.PSconfig{UseDefaultNetwork: false, MaxSubnetSize: 21, CustomSubnet: args}

	//Call pingDriver
	subnetList, err := pingsweep.PingDriver(psconfig)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	fmt.Println(subnetList)
}
