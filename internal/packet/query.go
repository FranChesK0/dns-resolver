package packet

func NewQuery(header *Header, question *Question) []byte {
	query := make([]byte, 0)
	query = append(query, header.ToBytes()...)
	query = append(query, question.ToBytes()...)
	return query
}
