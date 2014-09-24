package main

import (
	"fmt"
	"github.com/alouca/gosnmp"
	"log"
	"os"
	"strconv"
)

func main() {
	s, err := gosnmp.NewGoSNMP(os.Args[1], os.Args[2], gosnmp.Version2c, 5)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.Get("1.3.6.1.4.1.17095.3.5.0")
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
