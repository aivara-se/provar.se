import type { Action, Edge, TestFileGraph } from '../../domain/types';

/**
 * getNextNodes reads the outgoing edges of nodeId from the graph and
 * returns their targets in the order they appear in `graph.edges`. The
 * editor stores edges as an explicit list — synthesis of "next" from
 * position is the disk-side FromActions' concern, not ours — so this is
 * just a filter+map.
 *
 * Implicit position edges (FromActions synthesises one when an action
 * has no Next and a successor exists in source order) are filtered out
 * because they don't represent a real "next". The canvas shows them as
 * connectors anyway — they're just not persisted by addNode/deleteNode.
 */
export function getNextNodes(graph: TestFileGraph, id: string): string[] {
  const out: string[] = [];
  for (const edge of graph.edges) {
    if (edge.from === id && !edge.implicit) out.push(edge.to);
  }
  return out;
}

/**
 * toEngineTasks flattens the graph into the engine's Record<id, Task>
 * shape used by buildGraphPaths. The two shapes differ — engine tasks
 * carry next inline; the canvas stores edges separately — so we
 * normalise here. Sub-graphs (a node carrying its own nested Task[]) and
 * editor-flavoured fields (data, config, source) are stripped because
 * the engine only needs id/title/info/next to enumerate paths.
 */
export function toEngineTasks(graph: TestFileGraph): Record<string, { id: string; name: string; info: string; next: string[] }> {
  const tasks: Record<string, { id: string; name: string; info: string; next: string[] }> = {};
  for (const [id, node] of Object.entries(graph.nodes)) {
    tasks[id] = {
      id,
      name: node.name,
      info: node.info ?? '',
      next: getNextNodes(graph, id),
    };
  }
  return tasks;
}

/**
 * addNodeToGraph splices a new task between fromId and toId in graph and
 * returns the new node id. Edge cases:
 *   - fromId null: prepend a new start before `toId` and point `__start__`
 *     at the new node. (The start is `__start__` by convention — see
 *     testfile.GraphStartID.)
 *   - toId null: append at the end. New node has no outgoing edge;
 *     addOrphanEndIfNeeded adds an end_<id> virtual end node only when
 *     ToActions runs on disk, the canvas renders it on the fly.
 *   - fromId has multiple successors already: the new node replaces `toId`
 *     in fromId's next array; toId still gets the other predecessors'
 *     edges untouched.
 *
 * The function mutates a deep clone of graph; callers assign the result
 * back to currentFile (Svelte 5's $state replaces, doesn't mutate).
 */
export function addNodeToGraph(
  graph: TestFileGraph,
  fromId: string | null,
  toId: string | null,
): { graph: TestFileGraph; newNodeId: string } {
  const next: TestFileGraph = {
    start: graph.start,
    nodes: { ...graph.nodes },
    edges: graph.edges.map((e) => ({ from: e.from, to: e.to, implicit: e.implicit })),
  };
  const newNodeId = generateNodeId();
  next.nodes[newNodeId] = {
    id: newNodeId,
    name: 'New Action',
    info: 'Describe what this action does…',
  };

  if (fromId === null) {
    // Prepend: the new node becomes the first user node, with __start__
    // pointing at it. Any previous incoming-from-__start__ edge is
    // re-pointed at the new node so we don't end up with two starts.
    next.edges = next.edges.filter((e) => e.from !== '__start__');
    next.edges.push({ from: '__start__', to: newNodeId });
  } else {
    // Splice: add `newNodeId` to fromId's successors. If toId was an
    // existing successor, swap it; otherwise append.
    const successors = getNextNodes(next, fromId);
    let replaced = false;
    next.edges = next.edges.map((e) => {
      if (e.from !== fromId) return e;
      if (toId !== null && e.to === toId) {
        replaced = true;
        return { from: fromId, to: newNodeId };
      }
      return e;
    });
    if (!replaced) {
      next.edges.push({ from: fromId, to: newNodeId });
    }
    // Connect newNodeId → toId only if toId was given. (When toId is
    // null, the new node terminates the chain; its End is rendered but
    // not persisted — see end_<id> synthesised by FromActions.)
    if (toId !== null) {
      next.edges.push({ from: newNodeId, to: toId });
    }
  }

  return { graph: next, newNodeId };
}

/**
 * deleteNodeFromGraph removes a node and every descendant reachable via
 * outgoing edges (recursive downstream delete), then cleans up any
 * remaining edges in/out of the deleted set. If the deleted node was the
 * start, the start pointer is cleared — the test file becomes empty
 * until the user adds a fresh entry point.
 */
export function deleteNodeFromGraph(graph: TestFileGraph, id: string): TestFileGraph {
  const deleted = new Set<string>();
  const collect = (nodeId: string): void => {
    if (deleted.has(nodeId)) return;
    deleted.add(nodeId);
    if (nodeId === '__start__') return;
    for (const edge of graph.edges) {
      if (edge.from === nodeId) collect(edge.to);
    }
  };
  collect(id);

  const nodes = { ...graph.nodes };
  for (const del of deleted) {
    if (del !== '__start__') delete nodes[del];
  }
  const edges = graph.edges.filter(
    (e) => !deleted.has(e.from) && !deleted.has(e.to),
  );

  const start = deleted.has(graph.start) ? '' : graph.start;
  return { start, nodes, edges };
}

/**
 * deleteImplicitEdges returns the input graph with any From==__start__
 * synthesised entry-point edges and Implicit==true position edges
 * dropped. Used by mutations that re-emit the file view back to disk so
 * the YAML stays clean — those edges were synthesised by FromActions on
 * load and shouldn't reappear on save.
 *
 * Edges the user added with `addNodeToGraph` carry no `implicit` flag,
 * so they're preserved here.
 */
export function dropSynthesisedEdges(graph: TestFileGraph): TestFileGraph {
  return {
    start: graph.start,
    nodes: graph.nodes,
    edges: graph.edges.filter((e) => e.from !== '__start__' && !e.implicit),
  };
}

/**
 * edgesFrom returns all edges originating at fromId, regardless of
 * `implicit`. The canvas uses this to drive the connector rendering
 * pass so the gap from `__start__` to the first action still shows up.
 */
export function edgesFrom(graph: TestFileGraph, fromId: string): Edge[] {
  const out: Edge[] = [];
  for (const edge of graph.edges) {
    if (edge.from === fromId) out.push(edge);
  }
  return out;
}

/**
 * generateNodeId produces a fresh task_<5 alphanum> id. We use the same
 * shape the Bun-era code used so existing test files / fixtures don't
 * need to be re-migrated by future tooling. Math.random is fine here —
 * collisions in a 36^5 space are vanishingly rare for graphs of the
 * size provar supports (36^5 = ~60M).
 */
export function generateNodeId(): string {
  const chars = 'abcdefghijklmnopqrstuvwxyz0123456789';
  let result = '';
  for (let i = 0; i < 5; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length));
  }
  return `task_${result}`;
}

/**
 * clones of an Action — used by editor store mutations so Svelte 5's
 * $state proxy isn't mutated in place. We do a plain shallow clone
 * because Action has no nested objects except `config` (which we treat
 * as immutable from the editor's perspective — `updateNode` rebuilds
 * the config object explicitly).
 */
export function snapshotAction(a: Action): Action {
  return {
    id: a.id,
    name: a.name,
    info: a.info,
    source: a.source,
  };
}
