package algoplayground

type Graph map[string][]string

func buildGraph() Graph {
	return Graph{
		"John":  {"Alice", "Suhea"},
		"Alice": {"Bob"},
		"Suhea": {"Bob", "Rohan"},
		"Bob":   {},
		"Rohan": {},
	}
}

func BFS(gr Graph, start string) {

	var queue []string
	s := gr[start]
	queue = append(queue, s...)
}
