import { describe, expect, test } from 'bun:test';
import { GraphValidator } from './graph-validator';
import type { TestFileGraph } from './types';

describe('GraphValidator', () => {
  test('validates a healthy graph', () => {
    const graph: TestFileGraph = {
      start: '__start__',
      nodes: {
        __start__: { id: '__start__', name: 'Start' },
        node1: { id: 'node1', name: 'Click Button', info: 'Click login button' },
        node2: { id: 'node2', name: 'Verify Text', info: 'Check dashboard text' },
      },
      edges: [
        { from: '__start__', to: 'node1' },
        { from: 'node1', to: 'node2' },
      ],
    };

    const report = GraphValidator.validate(graph);
    expect(report.isValid).toBe(true);
    expect(report.errors).toHaveLength(0);
    expect(report.warnings).toHaveLength(0);
  });

  test('detects cycles in graph', () => {
    const graph: TestFileGraph = {
      start: '__start__',
      nodes: {
        __start__: { id: '__start__', name: 'Start' },
        node1: { id: 'node1', name: 'Action 1' },
        node2: { id: 'node2', name: 'Action 2' },
      },
      edges: [
        { from: '__start__', to: 'node1' },
        { from: 'node1', to: 'node2' },
        { from: 'node2', to: 'node1' },
      ],
    };

    const report = GraphValidator.validate(graph);
    expect(report.isValid).toBe(false);
    expect(report.errors.some((e) => e.code === 'CYCLE_DETECTED')).toBe(true);
  });

  test('detects disconnected nodes', () => {
    const graph: TestFileGraph = {
      start: '__start__',
      nodes: {
        __start__: { id: '__start__', name: 'Start' },
        node1: { id: 'node1', name: 'Connected Action' },
        orphan: { id: 'orphan', name: 'Orphaned Action' },
      },
      edges: [{ from: '__start__', to: 'node1' }],
    };

    const report = GraphValidator.validate(graph);
    expect(report.isValid).toBe(true); // Warnings don't invalidate graph, but surface in warnings array
    expect(report.warnings.some((w) => w.code === 'DISCONNECTED_NODE')).toBe(true);
  });

  test('detects missing action names', () => {
    const graph: TestFileGraph = {
      start: '__start__',
      nodes: {
        __start__: { id: '__start__', name: 'Start' },
        unnamed: { id: 'unnamed', name: '   ' },
      },
      edges: [{ from: '__start__', to: 'unnamed' }],
    };

    const report = GraphValidator.validate(graph);
    expect(report.isValid).toBe(false);
    expect(report.errors.some((e) => e.code === 'MISSING_ACTION_NAME')).toBe(true);
  });

  test('handles empty graph', () => {
    const report = GraphValidator.validate(null);
    expect(report.isValid).toBe(false);
    expect(report.errors.some((e) => e.code === 'EMPTY_GRAPH')).toBe(true);
  });
});
