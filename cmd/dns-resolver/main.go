package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/FranChesK0/dns-resolver/internal/client"
	"github.com/FranChesK0/dns-resolver/internal/query"
	"github.com/FranChesK0/dns-resolver/internal/record"
)

type DNSPacket struct {
	header      *query.Header
	questions   []*query.Question
	answers     []*record.Record
	additionals []*record.Record
	authorities []*record.Record
}

func main() {
	domains := os.Args[1:]
	if len(domains) < 1 {
		fmt.Println("Usage: ./dns-resolver <domain> [<domain> ...]")
		os.Exit(0)
	}

	for _, domain := range domains {
		fmt.Println(resolve(domain, query.TYPE_A))
	}
}

func resolve(domainName string, questionType uint16) string {
	nameServer := "77.240.157.30"
	for {
		fmt.Printf("querying %s for %s\n", nameServer, domainName)
		dnsResponse := sendQuery(nameServer, domainName, questionType)
		dnsPacket := getDNSPacketFromResponse(dnsResponse)

		if ip := getAnswer(dnsPacket.answers); ip != "" {
			return ip
		}
		if nsIP := getNameServerIP(dnsPacket.additionals); nsIP != "" {
			nameServer = nsIP
			continue
		}
		if nsDomain := getNameServer(dnsPacket.authorities); nsDomain != "" {
			nameServer = resolve(nsDomain, query.TYPE_A)
		}
	}
}

func sendQuery(nameServer string, domainName string, questionType uint16) []byte {
	q := query.NewQuery(
		query.NewHeader(22, 0, 1, 0, 0, 0),
		query.NewQuestion(domainName, questionType, query.CLASS_IN),
	)
	c := client.NewClient(nameServer, 53)
	return c.SendQuery(q)
}

func getDNSPacketFromResponse(dnsResponse []byte) *DNSPacket {
	var (
		header      *query.Header
		questions   []*query.Question
		answers     []*record.Record
		additionals []*record.Record
		authorities []*record.Record
	)

	reader := bytes.NewReader(dnsResponse)
	header, err := query.ParseHeader(reader)
	if err != nil {
		fmt.Printf("unable to parse the response header: %v\n", err)
		os.Exit(-1)
	}
	for range header.QdCount {
		questions = append(questions, query.ParseQuestion(reader))
	}
	for range header.AnCount {
		answers = append(answers, record.ParseRecord(reader))
	}
	for range header.NsCount {
		authorities = append(authorities, record.ParseRecord(reader))
	}
	for range header.ArCount {
		additionals = append(additionals, record.ParseRecord(reader))
	}

	return &DNSPacket{
		header:      header,
		questions:   questions,
		answers:     answers,
		additionals: additionals,
		authorities: authorities,
	}
}

func getAnswer(answers []*record.Record) string {
	return getRecord(answers)
}

func getNameServerIP(additionals []*record.Record) string {
	return getRecord(additionals)
}

func getNameServer(authorities []*record.Record) string {
	return getRecord(authorities)
}

func getRecord(records []*record.Record) string {
	for _, record := range records {
		if record.Type == query.TYPE_A || record.Type == query.TYPE_NS {
			return record.RData
		}
	}
	return ""
}
