package utils

func TopoSort(graph map[int][]int, seq []int) []int {
	visited := make(map[int]bool)
	cycleCheck := make(map[int]bool)
	result := make([]int, 0)
	input := make(map[int]bool)

	for _, node := range seq {
		input[node] = true
	}

	var visit func(node int) bool
	visit = func(node int) bool {
		if cycleCheck[node] {
			return false
		}
		if visited[node] {
			return true
		}
		cycleCheck[node] = true

		for _, neighbor := range graph[node] {
			if input[neighbor] {
				if !visit(neighbor) {
					return false // Cycle
				}
			}
		}

		cycleCheck[node] = false
		visited[node] = true
		result = append(result, node)
		return true
	}

	for _, node := range seq {
		if !visited[node] {
			if !visit(node) {
				return nil // Cycle
			}
		}
	}

	// Reverse
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
