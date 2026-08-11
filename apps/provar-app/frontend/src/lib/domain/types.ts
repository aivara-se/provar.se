// Canvas-local types. Shapes align strictly with the domain model (domain.Action).

export interface TestFileView {
  graph: TestFileGraph;
  order?: string[];
}

export interface TestFileGraph {
  start: string;
  nodes: Record<string, Action>;
  edges: Edge[];
}

export interface Action {
  id: string;
  name: string;
  info?: string;
  source?: string;
}

export interface Edge {
  from: string;
  to: string;
  implicit?: boolean;
}

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