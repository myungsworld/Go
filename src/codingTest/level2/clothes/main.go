package main

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func solution(clothes [][]string) []int {
	// nameofc := 0
	typeofc := []string{}
	numofc := []int{}
	//delete := []string{}
	typecount := 1
	// cnt := 0

	for _, array := range clothes {
		typeofc = append(typeofc, array[1])
	}
	for len(typeofc) > 0 {
		if len(typeofc) == 1 {
			numofc = append(numofc, typecount)
			break
		}
		for i := 1; i < len(typeofc); i++ {

			if typeofc[0] == typeofc[i] {
				typecount++
				remove(typeofc, i)
				break
			}

		}
		remove(typeofc, 0)
		numofc = append(numofc, typecount)
		typecount = 1
	}
	return numofc
}

func main() {
	var multiArray = [][]string{
		{"yellow_hat", "headgear"},
		{"blue_sunglasses", "eyewear"},
		{"green_turban", "headgear"},
	}
	solution(multiArray)
}
