package SfTech

import "strings"

func hasCycle(graph string) bool {
	links := strings.Split(graph, ",")
	m := make(map[string]string)
	for _, l := range links {
		nodes := strings.Split(l, "->")
		m[nodes[0]] = nodes[1]
	}

	for k := range m {
		start := k
		next, exist := m[start]
		for exist {
			next, exist = m[next]
			if next == start {
				return true
			}
		}
	}

	return false
}
