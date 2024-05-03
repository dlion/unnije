package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/dlion/runne/network"
	"github.com/dlion/runne/packet"
	"github.com/dlion/runne/response"
)

type DNSPacket struct {
	header      *packet.Header
	questions   []*packet.Question
	answers     []*packet.Record
	authorities []*packet.Record
	additionals []*packet.Record
}

func main() {
	fmt.Println(resolve("domenicoluciani.com", packet.TYPE_A))
}

func resolve(domainName string, type_ uint16) string {
	nameServer := "198.41.0.4"
	for {
		fmt.Printf("Querying %s for %s\n", nameServer, domainName)
		dnsResponse := sendQuery(nameServer, domainName, type_)
		dnsPacket := getDnsPacketFromResponse(dnsResponse)

		if ip := getAnswer(dnsPacket.answers); ip != "" {
			return ip
		}

		if nsIp := getNameServerIp(dnsPacket.additionals); nsIp != "" {
			nameServer = nsIp
			continue
		}

		if nsDomain := getNameServer(dnsPacket.authorities); nsDomain != "" {
			nameServer = resolve(nsDomain, packet.TYPE_A)
		}
	}
}

func getNameServer(authorities []*packet.Record) string {
	for _, authority := range authorities {
		if authority.Type == packet.TYPE_NS {
			return authority.Rdata
		}
	}
	return ""
}

func getNameServerIp(additionals []*packet.Record) string {
	for _, additional := range additionals {
		if additional.Type == packet.TYPE_A {
			return additional.Rdata
		}
	}
	return ""
}

func getAnswer(answers []*packet.Record) string {
	for _, answer := range answers {
		if answer.Type == packet.TYPE_A {
			return answer.Rdata
		}
	}
	return ""
}

func sendQuery(nameServer, domainName string, type_ uint16) []byte {
	query := packet.NewQuery(
		packet.NewHeader(22, 0, 1, 0, 0, 0),
		packet.NewQuestion(domainName, type_, packet.CLASS_IN),
	)

	client := network.NewClient(nameServer, 53)
	return client.SendQuery(query)
}

func getDnsPacketFromResponse(dnsResponse []byte) *DNSPacket {
	var (
		header      *packet.Header
		questions   []*packet.Question
		answers     []*packet.Record
		authorities []*packet.Record
		additionals []*packet.Record
	)

	reader := bytes.NewReader(dnsResponse)
	header, err := packet.ParseHeader(reader)
	if err != nil {
		fmt.Printf("Can't parse the response header: %v", err)
		os.Exit(-1)
	}
	for range header.QdCount {
		questions = append(questions, response.ParseQuestion(reader))
	}

	for range header.AnCount {
		answers = append(answers, response.ParseRecord(reader))
	}

	for range header.NsCount {
		authorities = append(authorities, response.ParseRecord(reader))
	}

	for range header.ArCount {
		additionals = append(additionals, response.ParseRecord(reader))
	}

	return &DNSPacket{header: header, questions: questions, answers: answers, authorities: authorities, additionals: additionals}
}
