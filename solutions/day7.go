package solutions

import (
	utils "adventofcode/utils"
	"errors"
	"strings"
)

var (
	ErrVertexDoesNotExist  = errors.New("Vertex does not exist")
	ErrVertexAlreadyExists = errors.New("Vertex already exists")
)

type Vertex struct {
	name         string
	pointsTo     []string
	pointsFrom   []string
	containsGold bool
}

func NewVertex(name string) (n *Vertex) {

	n = new(Vertex)
	n.name = name
	n.pointsTo = make([]string, 0)
	n.pointsFrom = make([]string, 0)
	return
}

type Graph struct {
	vertices     map[string]*Vertex
	edgeWeight   map[string]int
	containsGold map[string]bool
	bagCount     int
}

func NewGraph() (g Graph) {
	g.vertices = make(map[string]*Vertex)
	g.edgeWeight = make(map[string]int)
	g.containsGold = make(map[string]bool)
	g.bagCount = 0
	return
}

func (g *Graph) AddVertex(name string) error {
	_, ok := g.vertices[name]

	if ok {
		return ErrVertexAlreadyExists
	}

	g.vertices[name] = NewVertex(name)
	return nil
}

func (g *Graph) AddEdge(name, containsName string, weight int) error {
	vertex, ok := g.vertices[name]

	if !ok {
		return ErrVertexDoesNotExist
	}

	vertexTo, ok := g.vertices[containsName]

	if !ok {
		return ErrVertexDoesNotExist
	}

	vertex.pointsTo = append(vertex.pointsTo, containsName)
	vertexTo.pointsFrom = append(vertexTo.pointsFrom, name)
	g.edgeWeight[name+containsName] = weight
	return nil
}

func (g *Graph) MarkVertexAndParents(name string) error {

	vertex, ok := g.vertices[name]
	if !ok {
		return ErrVertexDoesNotExist
	}

	_, ok = g.containsGold[name]
	if ok {
		return nil
	}

	g.containsGold[name] = true

	for _, parent := range vertex.pointsFrom {
		err := g.MarkVertexAndParents(parent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Graph) followGold(name string) int {

	vertex, _ := g.vertices[name]

	totalBags := 1

	for _, child := range vertex.pointsTo {

		w, _ := g.edgeWeight[name+child]
		bags := g.followGold(child)
		totalBags += w * bags

	}

	return totalBags
}

func lineAsTopology(line string) (vertex string, contains []string, weights []int) {

	sections := strings.Split(line, ",")

	words := strings.Fields(sections[0])
	vertex = words[0] + words[1]

	if len(words) == 7 {
		return
	}

	contains = append(contains, words[5]+words[6])
	weights = append(weights, utils.AtoiPanic(words[4]))

	for i, bag := range sections {
		if i == 0 {
			continue
		}

		things := strings.Fields(bag)
		contains = append(contains, things[1]+things[2])
		weights = append(weights, utils.AtoiPanic(things[0]))
	}
	return
}

func fileAsGraph() Graph {

	g := NewGraph()

	//g.AddVertex("shinygold")

	inputs := utils.FileAsLines("input/day7.txt")

	for _, input := range inputs {
		vertex, edges, weights := lineAsTopology(input)

		g.AddVertex(vertex)
		for i, containsVertex := range edges {
			g.AddVertex(containsVertex)
			g.AddEdge(vertex, containsVertex, weights[i])
		}
	}

	return g

}

func SolveDay7() (string, string) {

	part1 := 0
	part2 := 0

	graph := fileAsGraph()
	graph.MarkVertexAndParents("shinygold")
	part2 = graph.followGold("shinygold") - 1

	part1 = len(graph.containsGold) - 1

	return utils.ResultStrings(part1, part2)
}
