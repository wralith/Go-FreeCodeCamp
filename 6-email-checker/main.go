package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	// SPF Records
	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Print(domain, " ", hasMX, " ", hasSPF, " ", spfRecord, " ", hasDMARC, " ", dmarcRecord)

}
