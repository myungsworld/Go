package dataStruct

//Rolling Hash
func Hash(s string) int {

	h := 0
	A := 256
	B := 3571

	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}

	return h
}

//

type KeyValue struct {
	key   string
	value string
}

func CreateMap() *Map {
	return &Map{}
}

type Map struct {
	keyArray [3571][]KeyValue
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], KeyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}
	}
	return ""
}
