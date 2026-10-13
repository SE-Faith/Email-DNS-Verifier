package main

import (
	"bufio"
	"fmt"

	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an email address: ")
	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}
	email := scanner.Text()
	if !isValidEmail(email) {
		fmt.Println("Invalid email address")
		return
	}

	fmt.Println("Valid email address")
	parts := strings.Split(email, "@")
	domain := parts[1]
	checkIP(domain)
	checkMX(domain)
	checkSPF(domain)
	checkDMARC(domain)

}

func isValidEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	if strings.ContainsAny(email, " \t\n\r") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return false
	}

	local := parts[0]
	if len(local) < 1 || len(local) > 64 {
		return false
	}

	if strings.HasPrefix(local, ".") ||
		strings.HasSuffix(local, ".") ||
		strings.Contains(local, "..") {
		return false
	}
	domain := parts[1]
	if len(domain) < 1 || len(domain) > 253 {
		return false
	}

	if strings.HasPrefix(domain, ".") ||
		strings.HasSuffix(domain, ".") ||
		strings.Contains(domain, "..") ||
		!strings.Contains(domain, ".") {
		return false
	}

	return true
}

func checkIP(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else if len(ips) > 0 {
		fmt.Printf("Domain IP Lookup: Found %d IP address(es)\n", len(ips))
	} else {
		fmt.Println("Domain IP Lookup: No IP addresses found")
	}
}

func checkMX(domain string) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf("Error: MX lookup failed: %v\n", err)
	} else if len(mxRecords) > 0 {
		fmt.Printf("Domain MX Lookup: Found %d MX records\n", len(mxRecords))
	} else {
		fmt.Println("Domain MX Lookup: No MX records found")
	}
}

func checkSPF(domain string) {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("Error: SPF lookup failed: %v\n", err)
		return 
	} 
	for _, record := range txtRecords {
		if strings.HasPrefix(strings.ToLower(record), "v=spf1") {
			fmt.Printf("Domain SPF Lookup: Found SPF record: %s\n", record)
			return
		}
	}
	fmt.Println("Domain SPF Lookup: No SPF record found")
}

func checkDMARC(domain string) {
	dmarcDomain := "_dmarc." + domain
	txtRecords, err := net.LookupTXT(dmarcDomain)
	if err != nil {
		fmt.Printf("Error: DMARC lookup failed: %v\n", err)
		return
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(strings.ToLower(record), "v=dmarc1") {
			fmt.Printf("Domain DMARC Lookup: Found DMARC record: %s\n", record)
			return
		}
	}
	fmt.Println("Domain DMARC Lookup: No DMARC record found")
}