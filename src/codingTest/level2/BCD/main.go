package main

func solution(arr []int) int {

	newarray := []int{}
	m := 1
	for len(arr) > 0 {
		for i := 1; i < len(arr); i++ {

			if arr[i]%arr[0] == 0 {
				arr[i] = arr[i] / arr[0]
			}
		}
		newarray = append(newarray, arr[0])
		arr = arr[1:]

	}

	for i := 0; i < len(newarray); i++ {
		m *= newarray[i]
	}
	return m
}

func main() {
	arr := []int{2, 6, 8, 14}
	solution(arr)
}
