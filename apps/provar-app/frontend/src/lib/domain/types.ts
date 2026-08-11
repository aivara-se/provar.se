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