package packet

func NewQuery(header *Header, question *Question) []byte {
	var query []byte

	query = append(query, header.ToBytes()...)
	query = append(query, question.ToBytes()...)

	return query
}
