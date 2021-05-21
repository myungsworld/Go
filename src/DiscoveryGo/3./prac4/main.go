package main

func main() {

	var totalSlice []int
	var willBeDeleted []int

	// slice 의 삭제
	for i := 0; i < len(totalSlice); i++ {
		if len(TotalSlice) == 0
		for j := 0; j < len(willBeDeleted); j++ {
			if totalSlice[i] == willBeDeleted[j] {
				totalSlice = append(totalSlice[:i],totalSlice[i+1:]...)
				i--
				continue	
			}
		}
	}
}
