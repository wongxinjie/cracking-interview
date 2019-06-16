/**
* @File: graph.go
* @Author: wongxinjie
* @Date: 2019/6/9
*/
package graph

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	UnVisit = iota
	Visiting
	Visited
)

type Vertex struct {
	Value int64
	VisitStatus int
}

type Graph struct {
	Vertexes []*Vertex
	Edges map[Vertex] []*Vertex
	Count int64
}

func New() *Graph{
	vertexes := make([]*Vertex, 0)
	edges := make(map[Vertex][]*Vertex)
	return &Graph{
		Vertexes: vertexes,
		Edges: edges,
		Count: 0,
	}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.Vertexes = append(g.Vertexes, v)
}

func (g *Graph) AddEdge(from, to *Vertex) {
	g.Edges[*from] = append(g.Edges[*from], to)
}

func (g *Graph) String() string {
	s := ""
	for v, edges := range g.Edges {
		s += fmt.Sprintf("%v -> ", v.Value)
		values := make([]string, 0)
		for _, e := range edges {
			values = append(values, strconv.FormatInt(e.Value, 10))
		}
		s += fmt.Sprintf("%v\n", strings.Join(values, ","))
	}

	return s
}

func CreateGraphFromMap(m map[int64][]int64) *Graph {
	g := &Graph{
		Vertexes: make([]*Vertex, 0),
		Edges: make(map[Vertex] []*Vertex),
		Count: 0,
	}

	valueVertexMap := make(map[int64] *Vertex)

	for key := range m {
		v := &Vertex{
			key,
			UnVisit,
		}
		g.Vertexes = append(g.Vertexes, v)
		g.Count += 1

		valueVertexMap[key] = v
	}

	fmt.Println(valueVertexMap)

	for key, values := range m {
		from := valueVertexMap[key]
		for _, n := range values {
			to := valueVertexMap[n]
			g.AddEdge(from, to)
		}
	}

	return g
}
