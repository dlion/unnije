package main

import (
	"bytes"

	"github.com/dlion/cue/builder"
	"github.com/dlion/cue/network"
)

func main() {
	query := builder.NewQuery(
		builder.NewHeader(22, builder.RECURSION, 1, 0, 0, 0),
		builder.NewQuestion([]byte("domenicoluciani.com"), builder.TYPE_A, builder.CLASS_IN),
	)
	client := network.NewClient("8.8.8.8", 53)
	response := client.SendQuery(query)

	reader := bytes.NewReader(response)
	responseHeader := builder.ParseHeader(reader)
	responseHeader.Print()

	for q := range responseHeader.QdCount {
		builder.ParseQuestion(reader).Print(q)
	}

	for q := range responseHeader.AnCount {
		builder.ParseRecord(reader).Print(builder.ANSWER, q)
	}

	for q := range responseHeader.NsCount {
		builder.ParseRecord(reader).Print(builder.AUTHORITIES, q)
	}

	for q := range responseHeader.ArCount {
		builder.ParseRecord(reader).Print(builder.ADDITIONALS, q)
	}
}
