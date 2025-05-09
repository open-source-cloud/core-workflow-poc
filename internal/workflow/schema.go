package workflow

import "github.com/open-source-cloud/fuse/pkg/graph"

// Schema describes a workflow schema for execution
type Schema interface {
	ID() string
	RootNode() graph.Node
}

type schema struct {
	id    string
	graph graph.Graph
}

// LoadSchema creates a new Schema object
func LoadSchema(id string, graph graph.Graph) Schema {
	return &schema{
		id:    id,
		graph: graph,
	}
}

func (s *schema) ID() string {
	return s.id
}

func (s *schema) RootNode() graph.Node {
	return s.graph.Root()
}
