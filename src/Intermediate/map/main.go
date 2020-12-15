package main

import "fmt"

type Key struct {
	v int
}

type Value struct {
	v int
}

func main() {

	var m map[string]string
	m = make(map[string]string)

	m["abc"] = "bbb"
	fmt.Println(m["abc"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println("m1 53 = ", m1[53])
	fmt.Println("m1 55 = ", m1[55])

	m2 := make(map[int]int)

	fmt.Println("m2 30 =", m2[30])
	m2[0] = 0
	//맵의 value 가 정의되어있는지 아닌지 판별 할때 변수를 2개 써서 bool 형으로 반환
	v, ok := m2[0]
	v2, ok2 := m2[1]
	fmt.Println(v, ok, v2, ok2)

	delete(m2, 0)
	v, ok = m2[0]
	fmt.Println(v, ok, v2, ok2)

	for key, value := range m2 {
		fmt.Println(key, " = ", value)
	}

	m3 := make(map[int]bool)

	m3[4] = true
	fmt.Println(m3[4], m3[6])

	// fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	// fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	// fmt.Println("abcdf = ", dataStruct.Hash("abcdf"))
	// fmt.Println("tbcde = ", dataStruct.Hash("tbcde"))

	// m := dataStruct.CreateMap()
	// m.Add("송동명", "01020467089")
	// m.Add("송딍믱", "01023431341")
	// m.Add("송둥뭉", "01023434323")
	// m.Add("송뎅멩", "01023423412")

	// fmt.Println("송동명 = ", m.Get("송동명"))
	// fmt.Println("송딍믱 = ", m.Get("송딍믱"))
	// fmt.Println("송둥뭉 = ", m.Get("송둥뭉"))
	// fmt.Println("송뎅멩 = ", m.Get("송뎅멩"))
}
