package main

import (
	"bytes"
	"fmt"

	"github.com/dlion/cue/builder"
	"github.com/dlion/cue/network"
)

func main() {
	client := network.NewClient("8.8.8.8", 53)
	header := builder.NewHeader(22, builder.RECURSION, 1, 0, 0, 0)
	question := builder.NewQuestion([]byte("domenicoluciani.com"), builder.TYPE_A, builder.CLASS_IN)
	query := builder.NewQuery(header, question)
	fmt.Printf("Query: %X\n", query)
	response := client.SendQuery(query)
	fmt.Printf("Response: %X\n", response)
	reader := bytes.NewReader(response)

	responseHeader := builder.ParseHeader(reader)
	fmt.Println("--- HEADER ---")
	fmt.Printf("ID: %v\n", responseHeader.Id)
	fmt.Printf("AnCount: %v\n", responseHeader.AnCount)
	fmt.Printf("ArCount: %v\n", responseHeader.ArCount)
	fmt.Printf("NsCount: %v\n", responseHeader.NsCount)
	fmt.Printf("QdCount: %v\n", responseHeader.QdCount)
	fmt.Printf("Flags: 0x%X\n", responseHeader.Flags)

	for q := range responseHeader.QdCount {
		responseQuestion := builder.ParseQuestion(reader)
		fmt.Printf("--- QUESTION %d ---\n", q)
		fmt.Printf("Name: %s\n", responseQuestion.QName)
		fmt.Printf("Type: 0x%X\n", responseQuestion.QType)
		fmt.Printf("Class: 0x%X\n", responseQuestion.QClass)
	}

	for q := range responseHeader.AnCount {
		responseRecord := builder.ParseRecord(reader)
		fmt.Printf("--- ANSWER %d ---\n", q)
		fmt.Printf("Name: %s\n", responseRecord.Name)
		fmt.Printf("Type: 0x%X\n", responseRecord.Type)
		fmt.Printf("Class: 0x%X\n", responseRecord.Class)
		fmt.Printf("TTL: %d\n", responseRecord.TTL)
		fmt.Printf("Length: %d\n", responseRecord.RdLength)
		fmt.Printf("Data: %s\n", responseRecord.Rdata)
	}

	for q := range responseHeader.NsCount {
		responseRecord := builder.ParseRecord(reader)
		fmt.Printf("--- AUTHORITIES %d ---\n", q)
		fmt.Printf("Name: %s\n", responseRecord.Name)
		fmt.Printf("Type: 0x%X\n", responseRecord.Type)
		fmt.Printf("Class: 0x%X\n", responseRecord.Class)
		fmt.Printf("TTL: %d\n", responseRecord.TTL)
		fmt.Printf("Length: %d\n", responseRecord.RdLength)
		fmt.Printf("Data: %s\n", responseRecord.Rdata)
	}

	for q := range responseHeader.ArCount {
		responseRecord := builder.ParseRecord(reader)
		fmt.Printf("---  ADDITIONALS %d ---\n", q)
		fmt.Printf("Name: %s\n", responseRecord.Name)
		fmt.Printf("Type: 0x%X\n", responseRecord.Type)
		fmt.Printf("Class: 0x%X\n", responseRecord.Class)
		fmt.Printf("TTL: %d\n", responseRecord.TTL)
		fmt.Printf("Length: %d\n", responseRecord.RdLength)
		fmt.Printf("Data: %s\n", responseRecord.Rdata)
	}
}
