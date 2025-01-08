# Topological Sort in Go

This package provides a utility for performing **Topological Sorting** of directed acyclic graphs (DAGs). It allows users to add services with dependencies and then obtain the topologically sorted order of services. It also includes functionality for detecting cycles in the graph to ensure that the graph is a valid DAG.

## Features
- **Topological Sort**: Perform topological sorting on a graph of services with dependencies.
- **Cycle Detection**: Detect cycles in the graph to verify whether the graph is a Directed Acyclic Graph (DAG).
- **Error Handling**: Provide detailed error messages when the graph contains unresolved dependencies or cycles.

## Installation

To use this package, include it in your Go project by importing it as follows:

```go
import "github.com/MahshidJabari/dependit"
