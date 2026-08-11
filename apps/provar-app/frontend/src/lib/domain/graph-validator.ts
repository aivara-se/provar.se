import type { TestFileGraph } from './types';

export interface GraphDiagnostic {
  id: string;
  nodeId?: string;
  severity: 'error' | 'warning';
  code:
    | 'CYCLE_DETECTED'
    | 'DISCONNECTED_NODE'
    | 'DANGLING_EDGE'
    | 'MISSING_ACTION_NAME'
    | 'MISSING_START'
    | 'EMPTY_GRAPH';
  message: string;
}

export interface DiagnosticReport {
  isValid: boolean;
  errors: GraphDiagnostic[];
  warnings: GraphDiagnostic[];
}

export class GraphValidator {
  static validate(graph: TestFileGraph | null | undefined): DiagnosticReport {
    const errors: GraphDiagnostic[] = [];
    const warnings: GraphDiagnostic[] = [];

    if (!graph || !graph.nodes || Object.keys(graph.nodes).length === 0) {
      errors.push({
        id: 'empty-graph',
        severity: 'error',
        code: 'EMPTY_GRAPH',
        message: 'Test file has no action nodes defined.',
      });
      return { isValid: false, errors, warnings };
    }

    // 1. Missing start node validation
    if (!graph.start || !graph.nodes[graph.start]) {
      errors.push({
        id: 'missing-start',
        severity: 'error',
        code: 'MISSING_START',
        message: `Graph start node "${graph.start}" is not defined in nodes.`,
      });
    }

    // Build adjacency list for edges
    const adj: Record<string, string[]> = {};
    for (const nodeId of Object.keys(graph.nodes)) {
      adj[nodeId] = [];
    }

    // 2. Edge validation (Dangling edges)
    for (const edge of graph.edges || []) {
      if (!graph.nodes[edge.from]) {
        errors.push({
          id: `dangling-edge-from-${edge.from}`,
          nodeId: edge.from,
          severity: 'error',
          code: 'DANGLING_EDGE',
          message: `Edge originates from non-existent node "${edge.from}".`,
        });
      }
      if (!graph.nodes[edge.to]) {
        errors.push({
          id: `dangling-edge-to-${edge.to}`,
          nodeId: edge.to,
          severity: 'error',
          code: 'DANGLING_EDGE',
          message: `Edge points to non-existent node "${edge.to}".`,
        });
      }
      if (graph.nodes[edge.from] && graph.nodes[edge.to]) {
        adj[edge.from].push(edge.to);
      }
    }

    // 3. Cycle Detection (DFS)
    const state: Record<string, number> = {}; // 0 = unvisited, 1 = visiting, 2 = visited
    const cycleNodes = new Set<string>();

    const dfsCycle = (nodeId: string, path: string[]) => {
      state[nodeId] = 1;
      path.push(nodeId);

      for (const nextId of adj[nodeId] || []) {
        if (state[nextId] === 1) {
          // Cycle found
          const cycleStartIndex = path.indexOf(nextId);
          const cyclePath = path.slice(cycleStartIndex).concat(nextId);
          for (const cNode of cyclePath) {
            cycleNodes.add(cNode);
          }
        } else if (!state[nextId]) {
          dfsCycle(nextId, path);
        }
      }

      path.pop();
      state[nodeId] = 2;
    };

    for (const nodeId of Object.keys(graph.nodes)) {
      if (!state[nodeId]) {
        dfsCycle(nodeId, []);
      }
    }

    if (cycleNodes.size > 0) {
      for (const nodeId of cycleNodes) {
        errors.push({
          id: `cycle-${nodeId}`,
          nodeId,
          severity: 'error',
          code: 'CYCLE_DETECTED',
          message: `Node "${nodeId}" is part of a circular dependency cycle.`,
        });
      }
    }

    // 4. Reachability Analysis (BFS from start)
    if (graph.start && graph.nodes[graph.start]) {
      const reachable = new Set<string>();
      const queue = [graph.start];
      reachable.add(graph.start);

      while (queue.length > 0) {
        const current = queue.shift()!;
        for (const nextId of adj[current] || []) {
          if (!reachable.has(nextId)) {
            reachable.add(nextId);
            queue.push(nextId);
          }
        }
      }

      for (const nodeId of Object.keys(graph.nodes)) {
        if (!reachable.has(nodeId) && nodeId !== '__start__') {
          warnings.push({
            id: `disconnected-${nodeId}`,
            nodeId,
            severity: 'warning',
            code: 'DISCONNECTED_NODE',
            message: `Node "${nodeId}" is disconnected and unreachable from start.`,
          });
        }
      }
    }

    // 5. Node Action Completeness
    for (const [nodeId, node] of Object.entries(graph.nodes)) {
      if (nodeId === '__start__') continue;
      if (!node.name || node.name.trim() === '') {
        errors.push({
          id: `missing-name-${nodeId}`,
          nodeId,
          severity: 'error',
          code: 'MISSING_ACTION_NAME',
          message: `Action node "${nodeId}" has no title or action name.`,
        });
      }
    }

    return {
      isValid: errors.length === 0,
      errors,
      warnings,
    };
  }
}
