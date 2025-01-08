package topologicalsort

import (
	"fmt"
)

type Service struct {
	Name         string
	Dependencies []string
}

type DependencyGraph struct {
	services     map[string]*Service
	dependencies map[string][]string
	inDegree     map[string]int
}

func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		services:     make(map[string]*Service),
		dependencies: make(map[string][]string),
		inDegree:     make(map[string]int),
	}
}

func (g *DependencyGraph) AddService(service *Service) {
	g.services[service.Name] = service
	g.inDegree[service.Name] = len(service.Dependencies)

	for _, dep := range service.Dependencies {
		g.dependencies[dep] = append(g.dependencies[dep], service.Name)
	}
}

func (g *DependencyGraph) detectCycle() ([]string, bool) {

	visited := make(map[string]bool)
	stack := make(map[string]bool)

	var cycle []string

	var dfs func(service string) bool
	dfs = func(service string) bool {
		if stack[service] {
			// A cycle is detected.
			cycle = append(cycle, service)
			return true
		}
		if visited[service] {
			return false
		}

		visited[service] = true
		stack[service] = true

		for _, dep := range g.dependencies[service] {
			if dfs(dep) {
				cycle = append(cycle, dep)
				return true
			}
		}
		stack[service] = false
		return false
	}

	for service := range g.services {
		if !visited[service] {
			if dfs(service) {
				for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
					cycle[i], cycle[j] = cycle[j], cycle[i]
				}
				return cycle, true
			}
		}
	}
	return nil, false
}

func (g *DependencyGraph) TopologicalSort() ([]string, error) {
	var result []string
	queue := []string{}

	for service, degree := range g.inDegree {
		if degree == 0 {
			queue = append(queue, service)
		}
	}

	for len(queue) > 0 {
		service := queue[0]
		queue = queue[1:]

		result = append(result, service)

		for _, dependent := range g.dependencies[service] {
			g.inDegree[dependent]--

			if g.inDegree[dependent] == 0 {
				queue = append(queue, dependent)
			}
		}
	}

	if len(result) != len(g.services) {
		cycle, found := g.detectCycle()
		if found {
			return nil, fmt.Errorf("cycle detected involving services: %v", cycle)
		}
		return nil, fmt.Errorf("graph contains unresolved dependencies")
	}

	return result, nil
}

func DefineServices(dependencies map[string][]string) ([]string, error) {
	graph := NewDependencyGraph()

	for service, deps := range dependencies {
		graph.AddService(&Service{
			Name:         service,
			Dependencies: deps,
		})
	}

	return graph.TopologicalSort()
}

func (g *DependencyGraph) IsDAG() (bool, error) {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var dfs func(service string) bool
	dfs = func(service string) bool {
		if recStack[service] {
			return false
		}
		if visited[service] {
			return true
		}

		visited[service] = true
		recStack[service] = true

		for _, dep := range g.dependencies[service] {
			if !dfs(dep) {
				return false
			}
		}

		recStack[service] = false
		return true
	}

	for service := range g.services {
		if !visited[service] {
			if !dfs(service) {
				return false, fmt.Errorf("graph contains a cycle")
			}
		}
	}
	return true, nil
}

func CreateAndCheckDAG(dependencies map[string][]string) (bool, error) {
	graph := NewDependencyGraph()

	for service, deps := range dependencies {
		graph.AddService(&Service{
			Name:         service,
			Dependencies: deps,
		})
	}

	return graph.IsDAG()
}
