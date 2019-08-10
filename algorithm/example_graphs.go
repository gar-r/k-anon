package algorithm

import (
	"gonum.org/v1/gonum/graph/simple"
)

// GetUndirectedTestGraph1 returns an example undirected graph as follows:
//                   ---------- 0 ---------------------------------
//                   |                      |          |           |
//         --------- 1 ---------          - 5 -      - 6 -       - 7 -
//         |         |          |        |     |     |     |    |     |
//       - 2 -     - 3 -      - 4 -     14    15    16     17  18     19
//      |     |   |     |    |     |
//      8     9  10     11  12     13
func GetUndirectedTestGraph1() *simple.UndirectedGraph {
	g := CreateNodesUndirected(20)
	AddEdge(g, 0, 1)
	AddEdge(g, 1, 2)
	AddEdge(g, 1, 3)
	AddEdge(g, 1, 4)
	AddEdge(g, 2, 8)
	AddEdge(g, 2, 9)
	AddEdge(g, 3, 10)
	AddEdge(g, 3, 11)
	AddEdge(g, 4, 12)
	AddEdge(g, 4, 13)
	AddEdge(g, 0, 5)
	AddEdge(g, 0, 6)
	AddEdge(g, 0, 7)
	AddEdge(g, 5, 14)
	AddEdge(g, 5, 15)
	AddEdge(g, 6, 16)
	AddEdge(g, 6, 17)
	AddEdge(g, 7, 18)
	AddEdge(g, 7, 19)
	return g
}

// GetUndirectedTestGraph2 returns an example undirected graph as follows:
//            ------- 0 ------
//           |                |
//      ---- 1 ----           6
//     |     |     |          |
//     2     3     4          7
//           |
//           5
//
func GetUndirectedTestGraph2() *simple.UndirectedGraph {
	g := CreateNodesUndirected(8)
	AddEdge(g, 0, 1)
	AddEdge(g, 1, 2)
	AddEdge(g, 1, 3)
	AddEdge(g, 1, 4)
	AddEdge(g, 3, 5)
	AddEdge(g, 0, 6)
	AddEdge(g, 6, 7)
	return g
}

// GetUndirectedTestGraph3 returns an example undirected graph as follows:
//         ------ 0 -------
//        |    |      |    |
//        1    3      5    7
//        |    |      |    |
//        2    4      6    8
//
func GetUndirectedTestGraph3() *simple.UndirectedGraph {
	g := CreateNodesUndirected(9)
	AddEdge(g, 0, 1)
	AddEdge(g, 0, 3)
	AddEdge(g, 0, 5)
	AddEdge(g, 0, 7)
	AddEdge(g, 1, 2)
	AddEdge(g, 3, 4)
	AddEdge(g, 5, 6)
	AddEdge(g, 7, 8)
	return g
}
