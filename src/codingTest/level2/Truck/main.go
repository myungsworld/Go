package main

func solution(bridge_length int, weight int, truck_weights []int) int {
	cnt := 0
	unpassedT := truck_weights
	passedT := []int{}
	handle := 0
	for len(unpassedT) > 0 {
		for i := 0; i < len(unpassedT); i++ {
			handle += unpassedT[i]

			if handle > weight {
				handle -= unpassedT[i]
				unpassedT = unpassedT[i-1:]
				break
			}
			passedT = append(passedT, unpassedT[i])
			cnt++
		}
		cnt += bridge_length
	}

	return cnt
}

func main() {
	bridge_length := 2
	weight := 10

	truck_weights := []int{7, 4, 5, 6}
	solution(bridge_length, weight, truck_weights)
}
