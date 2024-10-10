package resolver

import (
	"fmt"

	"github.com/FranChesK0/dns-resolver/internal/client"
	"github.com/FranChesK0/dns-resolver/internal/packet"
)

type Resolver struct {
	nameServer string
}

func NewResolver(nameServer string) *Resolver {
	return &Resolver{nameServer: nameServer}
}

func (r *Resolver) Resolve(domainName string, questionType uint16) string {
	nameServer := r.nameServer
	for {
		fmt.Printf("querying %s for %s\n", nameServer, domainName)
		dnsResponse := sendQuery(nameServer, domainName, questionType)
		dnsPacket := packet.NewDNSPacket(dnsResponse)

		if ip := getAnswer(dnsPacket.Answers); ip != "" {
			return ip
		}
		if nsIP := getNameServerIP(dnsPacket.Additionals); nsIP != "" {
			nameServer = nsIP
			continue
		}
		if nsDomain := getNameServer(dnsPacket.Authorities); nsDomain != "" {
			nameServer = r.Resolve(nsDomain, packet.TYPE_A)
		}
	}
}

func sendQuery(nameServer string, domainName string, questionType uint16) []byte {
	q := packet.NewQuery(
		packet.NewHeader(22, 0, 1, 0, 0, 0),
		packet.NewQuestion(domainName, questionType, packet.CLASS_IN),
	)
	c := client.NewClient(nameServer, 53)
	return c.SendQuery(q)
}

func getAnswer(answers []*packet.Record) string {
	return getRecord(answers)
}

func getNameServerIP(additionals []*packet.Record) string {
	return getRecord(additionals)
}

func getNameServer(authorities []*packet.Record) string {
	return getRecord(authorities)
}

func getRecord(records []*packet.Record) string {
	for _, record := range records {
		if record.Type == packet.TYPE_A || record.Type == packet.TYPE_NS {
			return record.RData
		}
	}
	return ""
}
