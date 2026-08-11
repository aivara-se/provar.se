package domain

import (
	"fmt"
	"strings"
)

type DiagnosticSeverity string

const (
	SeverityError   DiagnosticSeverity = "error"
	SeverityWarning DiagnosticSeverity = "warning"
)

type DiagnosticCode string

const (
	CodeEmptyGraph        DiagnosticCode = "EMPTY_GRAPH"
	CodeMissingStart      DiagnosticCode = "MISSING_START"
	CodeDanglingEdge      DiagnosticCode = "DANGLING_EDGE"
	CodeCycleDetected     DiagnosticCode = "CYCLE_DETECTED"
	CodeDisconnectedNode  DiagnosticCode = "DISCONNECTED_NODE"
	CodeMissingActionName DiagnosticCode = "MISSING_ACTION_NAME"
)

type GraphDiagnostic struct {
	ID       string             `json:"id"`
	NodeID   string             `json:"nodeId,omitempty"`
	Severity DiagnosticSeverity `json:"severity"`
	Code     DiagnosticCode     `json:"code"`
	Message  string             `json:"message"`
}

type DiagnosticReport struct {
	IsValid  bool              `json:"isValid"`
	Errors   []GraphDiagnostic `json:"errors"`
	Warnings []GraphDiagnostic `json:"warnings"`
}

const DefaultStartID = "__start__"

// Validate runs graph health checks on f.Actions and returns a DiagnosticReport.
func (f File) Validate() *DiagnosticReport {
	errors := []GraphDiagnostic{}
	warnings := []GraphDiagnostic{}

	if len(f.Actions) == 0 {
		errors = append(errors, GraphDiagnostic{
			ID:       "empty-graph",
			Severity: SeverityError,
			Code:     CodeEmptyGraph,
			Message:  "Test file has no action nodes defined.",
		})
		return &DiagnosticReport{
			IsValid:  false,
			Errors:   errors,
			Warnings: warnings,
		}
	}

	nodesMap := make(map[string]Action, len(f.Actions))
	for _, a := range f.Actions {
		if a.ID != "" {
			nodesMap[a.ID] = a
		}
	}

	// 1. Missing start node check
	startID := f.startNodeID(nodesMap)
	if startID != "" {
		if _, ok := nodesMap[startID]; !ok && startID != DefaultStartID {
			errors = append(errors, GraphDiagnostic{
				ID:       "missing-start",
				Severity: SeverityError,
				Code:     CodeMissingStart,
				Message:  fmt.Sprintf("Graph start node %q is not defined in nodes.", startID),
			})
		}
	}

	// Build adjacency list for reachability & cycle detection
	adj := make(map[string][]string, len(nodesMap))
	for id := range nodesMap {
		adj[id] = []string{}
	}

	// 2. Edge validation (Dangling edges)
	for _, action := range f.Actions {
		if action.ID == "" {
			continue
		}
		for _, nextID := range action.Next {
			if _, exists := nodesMap[nextID]; !exists {
				errors = append(errors, GraphDiagnostic{
					ID:       fmt.Sprintf("dangling-edge-to-%s", nextID),
					NodeID:   nextID,
					Severity: SeverityError,
					Code:     CodeDanglingEdge,
					Message:  fmt.Sprintf("Edge points to non-existent node %q.", nextID),
				})
			} else {
				adj[action.ID] = append(adj[action.ID], nextID)
			}
		}
	}

	// 3. Cycle Detection (DFS)
	state := make(map[string]int, len(nodesMap)) // 0 = unvisited, 1 = visiting, 2 = visited
	cycleNodes := make(map[string]bool)

	var dfsCycle func(nodeID string, path []string)
	dfsCycle = func(nodeID string, path []string) {
		state[nodeID] = 1
		path = append(path, nodeID)

		for _, nextID := range adj[nodeID] {
			if state[nextID] == 1 {
				// Cycle found - collect nodes in cycle path
				idx := -1
				for i, pID := range path {
					if pID == nextID {
						idx = i
						break
					}
				}
				if idx >= 0 {
					for _, cNode := range path[idx:] {
						cycleNodes[cNode] = true
					}
					cycleNodes[nextID] = true
				}
			} else if state[nextID] == 0 {
				dfsCycle(nextID, path)
			}
		}

		state[nodeID] = 2
	}

	for id := range nodesMap {
		if state[id] == 0 {
			dfsCycle(id, nil)
		}
	}

	if len(cycleNodes) > 0 {
		for nodeID := range cycleNodes {
			errors = append(errors, GraphDiagnostic{
				ID:       fmt.Sprintf("cycle-%s", nodeID),
				NodeID:   nodeID,
				Severity: SeverityError,
				Code:     CodeCycleDetected,
				Message:  fmt.Sprintf("Node %q is part of a circular dependency cycle.", nodeID),
			})
		}
	}

	// 4. Reachability Analysis (BFS from start)
	entrypoint := f.Actions[0].ID
	if entrypoint != "" {
		reachable := make(map[string]bool)
		queue := []string{entrypoint}
		reachable[entrypoint] = true

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			for _, nextID := range adj[curr] {
				if !reachable[nextID] {
					reachable[nextID] = true
					queue = append(queue, nextID)
				}
			}
		}

		for id := range nodesMap {
			if !reachable[id] && id != DefaultStartID {
				warnings = append(warnings, GraphDiagnostic{
					ID:       fmt.Sprintf("disconnected-%s", id),
					NodeID:   id,
					Severity: SeverityWarning,
					Code:     CodeDisconnectedNode,
					Message:  fmt.Sprintf("Node %q is disconnected and unreachable from start.", id),
				})
			}
		}
	}

	// 5. Node Action Completeness
	for _, action := range f.Actions {
		if action.ID == DefaultStartID {
			continue
		}
		if strings.TrimSpace(action.Name) == "" {
			errors = append(errors, GraphDiagnostic{
				ID:       fmt.Sprintf("missing-name-%s", action.ID),
				NodeID:   action.ID,
				Severity: SeverityError,
				Code:     CodeMissingActionName,
				Message:  fmt.Sprintf("Action node %q has no title or action name.", action.ID),
			})
		}
	}

	return &DiagnosticReport{
		IsValid:  len(errors) == 0,
		Errors:   errors,
		Warnings: warnings,
	}
}

func (f File) startNodeID(nodesMap map[string]Action) string {
	if _, ok := nodesMap[DefaultStartID]; ok {
		return DefaultStartID
	}
	if len(f.Actions) > 0 {
		return f.Actions[0].ID
	}
	return ""
}
