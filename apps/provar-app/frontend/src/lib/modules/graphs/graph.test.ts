import { describe, expect, test } from 'bun:test';
import {
  addNodeToGraph,
  deleteNodeFromGraph,
  dropSynthesisedEdges,
  generateNodeId,
  getNextNodes,
  toEngineTasks,
} from './graph';
import type { TestFileGraph } from '../../domain/types';

function makeGraph(
  nodes: Record<string, { name?: string }>,
  edges: { from: string; to: string; implicit?: boolean }[] = [],
  start: string = '__start__',
): TestFileGraph {
  const nodeMap: TestFileGraph['nodes'] = {};
  for (const [id, n] of Object.entries(nodes)) {
    nodeMap[id] = { id, name: n.name ?? id, info: '' };
  }
  return {
    start,
    nodes: nodeMap,
    edges: edges.map((e) => ({ from: e.from, to: e.to, implicit: e.implicit })),
  };
}

describe('graph utilities', () => {
  test('getNextNodes reads outgoing edges in order, ignores implicit', () => {
    const g = makeGraph(
      { task_a: {}, task_b: {}, task_c: {} },
      [
        { from: '__start__', to: 'task_a' },
        { from: 'task_a', to: 'task_b', implicit: true },
        { from: 'task_a', to: 'task_c' },
      ],
    );
    expect(getNextNodes(g, 'task_a')).toEqual(['task_c']);
  });

  test('addNodeToGraph prepends when fromId is null', () => {
    // graph.start is the sentinel '__start__' constant — only the edges
    // out of it change when we prepend. The new node becomes the only
    // entry point.
    const g = makeGraph(
      { task_a: {}, task_b: {} },
      [
        { from: '__start__', to: 'task_a' },
        { from: 'task_a', to: 'task_b' },
      ],
    );
    const { graph, newNodeId } = addNodeToGraph(g, null, null);
    expect(graph.start).toBe('__start__');
    expect(Object.keys(graph.nodes)).toContain(newNodeId);
    // __start__ edge is re-pointed at the new node; old __start__ edge gone.
    expect(graph.edges).toContainEqual({ from: '__start__', to: newNodeId });
    expect(graph.edges.filter((e) => e.from === '__start__')).toHaveLength(1);
  });

  test('addNodeToGraph splices between from and to', () => {
    const g = makeGraph({ task_a: {}, task_b: {} }, [
      { from: '__start__', to: 'task_a' },
      { from: 'task_a', to: 'task_b' },
    ]);
    const { graph, newNodeId } = addNodeToGraph(g, 'task_a', 'task_b');
    expect(graph.edges).toContainEqual({ from: 'task_a', to: newNodeId });
    expect(graph.edges).toContainEqual({ from: newNodeId, to: 'task_b' });
    // No duplicate task_a → task_b edge
    expect(graph.edges.filter((e) => e.from === 'task_a' && e.to === 'task_b')).toHaveLength(0);
  });

  test('addNodeToGraph with toId null appends at end', () => {
    const g = makeGraph({ task_a: {}, task_b: {} }, [
      { from: 'task_a', to: 'task_b' },
    ]);
    const { graph, newNodeId } = addNodeToGraph(g, 'task_b', null);
    expect(graph.nodes[newNodeId]).toBeDefined();
    expect(getNextNodes(graph, 'task_b')).toContain(newNodeId);
    expect(getNextNodes(graph, newNodeId)).toHaveLength(0);
  });

  test('deleteNodeFromGraph removes downstream descendants and orphans edges', () => {
    const g = makeGraph(
      { task_a: {}, task_b: {}, task_c: {} },
      [
        { from: '__start__', to: 'task_a' },
        { from: 'task_a', to: 'task_b' },
        { from: 'task_b', to: 'task_c' },
      ],
    );
    const after = deleteNodeFromGraph(g, 'task_a');
    // task_a, task_b, task_c all gone; nothing referencing them.
    expect(Object.keys(after.nodes)).not.toContain('task_a');
    expect(Object.keys(after.nodes)).not.toContain('task_b');
    expect(Object.keys(after.nodes)).not.toContain('task_c');
    // Edges that touched deleted nodes should be gone.
    for (const e of after.edges) {
      expect(['task_a', 'task_b', 'task_c']).not.toContain(e.from);
      expect(['task_a', 'task_b', 'task_c']).not.toContain(e.to);
    }
  });

  test('deleteNodeFromGraph leaves start sentinel alone', () => {
    // The graph start is a sentinel constant — deleting a user-action
    // never touches it (the canvas re-renders entry-point for whatever
    // node has an incoming edge from __start__ after the delete).
    const g = makeGraph(
      { task_a: {}, task_b: {} },
      [
        { from: '__start__', to: 'task_a' },
        { from: 'task_a', to: 'task_b' },
      ],
    );
    const after = deleteNodeFromGraph(g, 'task_a');
    expect(after.start).toBe('__start__');
  });

  test('dropSynthesisedEdges strips __start__ and implicit edges', () => {
    const g = makeGraph(
      { task_a: {}, task_b: {} },
      [
        { from: '__start__', to: 'task_a' },
        { from: 'task_a', to: 'task_b', implicit: true },
        { from: 'task_a', to: 'task_b' },
      ],
    );
    const out = dropSynthesisedEdges(g);
    expect(out.edges).toEqual([{ from: 'task_a', to: 'task_b' }]);
  });

  test('toEngineTasks flattens to engine shape', () => {
    const g = makeGraph(
      { task_a: {}, task_b: {}, task_c: {} },
      [
        { from: 'task_a', to: 'task_b' },
        { from: 'task_a', to: 'task_c' },
      ],
    );
    const tasks = toEngineTasks(g);
    expect(tasks['task_a'].next).toEqual(['task_b', 'task_c']);
    expect(tasks['task_b'].next).toEqual([]);
  });

  test('generateNodeId matches the task_<5 alnum> shape', () => {
    const id = generateNodeId();
    expect(id).toMatch(/^task_[a-z0-9]{5}$/);
  });
});
