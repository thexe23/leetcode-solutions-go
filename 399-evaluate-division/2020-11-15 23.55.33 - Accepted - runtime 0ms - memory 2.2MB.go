func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := buildGraph(equations, values)
	res := make([]float64, len(queries))
	for i, v := range queries {
		res[i] = dfs(v[0], v[1], graph, map[string]bool{})
	}
	return res
}

func buildGraph(equations [][]string, values []float64) map[string]map[string]float64 {
	graph := make(map[string]map[string]float64)
	for i, e := range equations {
		if _, ok := graph[e[0]]; !ok { graph[e[0]] = make(map[string]float64) }
		if _, ok := graph[e[1]]; !ok { graph[e[1]] = make(map[string]float64) }

		graph[e[0]][e[1]] = values[i]
		graph[e[1]][e[0]] = 1 / values[i]
	}
	return graph
}

func dfs(start, end string, graph map[string]map[string]float64, visited map[string]bool) float64{
	if _, ok := graph[start]; !ok {
		return -1.0
	}

	if _, ok := graph[start][end]; ok {
		return graph[start][end]
	}
	visited[start] = true
	for k, _ := range graph[start] {
		if !visited[k] {
			weight := graph[start][k]
			path := dfs(k, end, graph, visited)
			if path != -1.0 {
				return weight * path
			}
		}
	}
	return -1.0
}