package main

func main() {

	var totalSlice []int
	var willBeDeleted []int

	// slice 의 삭제
	for i := 0; i < len(totalSlice); i++ {
		if len(totalSlice) == 0 {
			break
		}
		for j := 0; j < len(willBeDeleted); j++ {
			
			if len(willBeDeleted) == 0 {
				break
			}

			if totalSlice[i] == willBeDeleted[j] {
				totalSlice = append(totalSlice[:i],totalSlice[i+1:]...)
				i--
				continue	
			}
		}
	}
}
