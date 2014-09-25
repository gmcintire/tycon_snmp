package main

import (
	"fmt"
	"github.com/alouca/gosnmp"
	"log"
	"os"
	"strconv"
)

func main() {
	var oid string

	s, err := gosnmp.NewGoSNMP(os.Args[2], os.Args[3], gosnmp.Version2c, 5)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "relay1":
		oid = "1.3.6.1.4.1.17095.3.1.0"
	case "relay2":
		oid = "1.3.6.1.4.1.17095.3.2.0"
	case "relay3":
		oid = "1.3.6.1.4.1.17095.3.3.0"
	case "relay4":
		oid = "1.3.6.1.4.1.17095.3.4.0"
	case "voltage1":
		oid = "1.3.6.1.4.1.17095.3.5.0"
	case "voltage2":
		oid = "1.3.6.1.4.1.17095.3.6.0"
	case "voltage3":
		oid = "1.3.6.1.4.1.17095.3.7.0"
	case "voltage4":
		oid = "1.3.6.1.4.1.17095.3.8.0"
	case "amp1":
		oid = "1.3.6.1.4.1.17095.3.9.0"
	case "amp2":
		oid = "1.3.6.1.4.1.17095.3.10.0"
	case "amp3":
		oid = "1.3.6.1.4.1.17095.3.11.0"
	case "amp4":
		oid = "1.3.6.1.4.1.17095.3.12.0"
	case "temp1":
		oid = "1.3.6.1.4.1.17095.3.13.0"
	case "temp2":
		oid = "1.3.6.1.4.1.17095.3.14.0"
	}

	resp, err := s.Get(oid)
	if err == nil {
		for _, v := range resp.Variables {
			switch v.Type {
			case gosnmp.OctetString:
				value, _ := strconv.ParseFloat(v.Value.(string), 64)
				if value > 20 {
					fmt.Printf("OK %s\n", v.Value.(string))
					os.Exit(0)
				} else {
					fmt.Printf("CRITICAL: %s\n", v.Value.(string))
					os.Exit(2)
				}
			}
		}
	}
}
