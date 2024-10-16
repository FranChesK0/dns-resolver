package packet

import (
	"bytes"
	"fmt"
	"os"

	"github.com/FranChesK0/dns-resolver/internal/cli"
)

type DNSPacket struct {
	Header      *Header
	Questions   []*Question
	Answers     []*Record
	Additionals []*Record
	Authorities []*Record
}

func NewDNSPacket(dnsResponse []byte) *DNSPacket {
	var (
		header      *Header
		questions   []*Question
		answers     []*Record
		additionals []*Record
		authorities []*Record
	)

	reader := bytes.NewReader(dnsResponse)
	header, err := ParseHeader(reader)
	if err != nil {
		fmt.Println(cli.ErrText.Render(fmt.Sprintf("Unable to parse the response header: %v", err)))
		os.Exit(-1)
	}
	for range header.QdCount {
		questions = append(questions, ParseQuestion(reader))
	}
	for range header.AnCount {
		answers = append(answers, ParseRecord(reader))
	}
	for range header.NsCount {
		authorities = append(authorities, ParseRecord(reader))
	}
	for range header.ArCount {
		additionals = append(additionals, ParseRecord(reader))
	}

	return &DNSPacket{
		Header:      header,
		Questions:   questions,
		Answers:     answers,
		Additionals: additionals,
		Authorities: authorities,
	}
}
