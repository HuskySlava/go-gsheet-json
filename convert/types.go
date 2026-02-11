package convert

//type RawRow []interface{}

type Row struct {
	Key   string
	Value interface{}
}

func (r Row) ToSlice() []interface{} {
	return []interface{}{r.Key, r.Value}
}

func RowsToSlices(rows []Row) [][]interface{} {
	raw := make([][]interface{}, len(rows))
	for i, r := range rows {
		raw[i] = r.ToSlice()
	}
	return raw
}
