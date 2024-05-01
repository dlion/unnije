package main

import (
	"bytes"

	"github.com/dlion/runne/network"
	"github.com/dlion/runne/packet"
	response "github.com/dlion/runne/response"
)

func main() {
	query := packet.NewQuery(
		packet.NewHeader(22, packet.RECURSION_FLAG, 1, 0, 0, 0),
		packet.NewQuestion("domenicoluciani.com", packet.TYPE_A, packet.CLASS_IN),
	)
	client := network.NewClient("8.8.8.8", 53)
	dnsResponse := client.SendQuery(query)

	reader := bytes.NewReader(dnsResponse)
	responseHeader, _ := packet.ParseHeader(reader)
	responseHeader.Print()

	for q := range responseHeader.QdCount {
		response.ParseQuestion(reader).Print(q)
	}

	for q := range responseHeader.AnCount {
		response.ParseRecord(reader).Print(packet.ANSWER, q)
	}

	for q := range responseHeader.NsCount {
		response.ParseRecord(reader).Print(packet.AUTHORITIES, q)
	}

	for q := range responseHeader.ArCount {
		response.ParseRecord(reader).Print(packet.ADDITIONALS, q)
	}
}
