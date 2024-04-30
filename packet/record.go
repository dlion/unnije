package packet

import "fmt"

type Record struct {
	Name     []byte
	Type     uint16
	Class    uint16
	TTL      uint32
	RdLength uint16
	Rdata    string
}

const (
	QUESTION = iota
	ANSWER
	AUTHORITIES
	ADDITIONALS
)

func (r *Record) Print(t int, n uint16) {
	switch t {
	case ANSWER:
		fmt.Printf("--- ANSWER %d ---\n", n)
	case AUTHORITIES:
		fmt.Printf("--- AUTHORITIES %d ---\n", n)
	case ADDITIONALS:
		fmt.Printf("---  ADDITIONALS %d ---\n", n)
	}
	fmt.Printf("Name: %s\n", r.Name)
	fmt.Printf("Type: 0x%X\n", r.Type)
	fmt.Printf("Class: 0x%X\n", r.Class)
	fmt.Printf("TTL: %d\n", r.TTL)
	fmt.Printf("Length: %d\n", r.RdLength)
	fmt.Printf("Data: %s\n", r.Rdata)
}
