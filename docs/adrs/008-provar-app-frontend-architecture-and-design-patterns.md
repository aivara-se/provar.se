# 008 - Provar App Architecture and Design Patterns

## Context

`apps/provar-app` is the visual desktop editor for Provar built using **Wails + Go + Svelte 5 (with Runes)**. ADR-005 established Wails and Svelte 5, ADR-006 established that the desktop app is a peer of `provar-api`, and ADR-007 defined the `BaseBinding` pattern for Go backend bindings.

As features grow (visual graph editing, partial execution/debugging, live output streaming, project health diagnostics, LLM provider management), the frontend and binding interaction layer needs a clear, scalable architectural model. Without defined patterns, state management, RPC calls, event streaming, and graph validation risk becoming tightly coupled inside UI components and monolithic stores.

## Decision

We establish a **Layered Architecture** and **App-Level Design Patterns** for `apps/provar-app` that enforce strict separation of concerns, SOLID principles, zero redundancy (DRY), and high maintainability.

---

## Architecture & Layers

`apps/provar-app` is structured into 4 distinct architectural layers:

```
┌─────────────────────────────────────────────────────────┐
│ 1. UI Layer (Svelte 5 Components)                      │
│    Toolbar, TestExplorer, Canvas, SidePanels, Modals    │
└───────────────────────────┬─────────────────────────────┘
                            │ reactive state & events
┌───────────────────────────▼─────────────────────────────┐
│ 2. State Layer (Svelte 5 Runes Stores)                   │
│    editorStore, projectStore, settingsStore, uiStore    │
└───────────────────────────┬─────────────────────────────┘
                            │ delegates business logic
┌───────────────────────────▼─────────────────────────────┐
│ 3. Service Layer (Frontend Domain Services & Managers)   │
│    SettingsService, JobStreamManager, GraphValidator    │
└───────────────────────────┬─────────────────────────────┘
                            │ Wails RPC & Event Bus
┌───────────────────────────▼─────────────────────────────┐
│ 4. Binding Layer (Go Wails Bindings - ADR-007)           │
│    File, Project, Config, Run, Compile, Watcher, etc.   │
└─────────────────────────────────────────────────────────┘
```

### Directory Structure

```text
apps/provar-app/frontend/src/lib/
├── components/       # Layer 1: UI presentation (modals/, panels/, toolbar/, etc.)
├── stores/           # Layer 2: Reactive Svelte 5 Rune stores (SRP)
├── services/         # Layer 3: Frontend services & event stream managers
├── domain/           # Layer 4: Pure domain models & spec validators
└── modules/          # Encapsulated subsystems (self-contained, export via index.ts only)
    ├── canvas/       # Canvas graph renderer, viewport, & layout engine
    └── graphs/       # Pure graph AST, DAG invariants, & topological sort
```

---

## App-Level Design Patterns

### 1. Service Layer Pattern (Decoupled Binding Access)
- **Rule**: UI components MUST NOT call Wails generated bindings directly.
- **Implementation**: Create dedicated frontend service abstractions (`lib/services/`) such as `SettingsService`, `FileService`, and `ExecutionService`.
- **Benefit**: Centralizes RPC error handling, input validation, and UI toast notifications. Keeps components focused purely on presentation.

### 2. Reactive Single-Responsibility Stores (Svelte 5 Runes)
- **Rule**: Stores must adhere to the Single Responsibility Principle (SRP). Avoid monolithic stores.
- **Implementation**:
  - `projectStore`: Manages project path, test file list, `.provar/config.yml` state.
  - `editorStore`: Manages currently open test file graph, active node selection, and graph mutations.
  - `executionStore`: Manages active run/compile job state, step statuses, and live log stream output.
  - `settingsStore`: Manages global user settings (`~/.provar/settings.yml`) and setup wizard state.
  - `uiStore`: Manages modal visibility, sidebar states, and toast notifications.

### 3. Observer / Stream Manager Pattern (`JobStreamManager`)
- **Rule**: Asynchronous job streams (`compile` and `run` execution streams) must be handled by an event stream manager, separate from UI rendering logic.
- **Implementation**: `JobStreamManager` wraps Wails event subscriptions (`forJob`). It exposes typed lifecycle hooks (`onTaskStarted`, `onTaskFinished`, `onTaskFailed`, `onCompleted`) and guarantees resource cleanup upon job termination or component unmount.

### 4. Domain Specification Pattern (`GraphValidator`)
- **Rule**: Graph structural validation (DAG invariants, cycle detection, entry-point reachability, missing action text) must exist as a pure JS/TS domain service independent of rendering or persistence.
- **Implementation**: `lib/domain/graph-validator.ts` provides pure functions that evaluate graph health. `editorStore` and UI components query this service reactively to display visual diagnostic badges and gate compile/run triggers.

### 5. Transactional Auto-Save & Dirty State Queue (Unit of Work)
- **Rule**: In-memory graph mutations must be debounced and committed atomically to avoid disk watcher desynchronization or lost writes.
- **Implementation**: `editorStore` uses a transactional auto-save queue with explicit dirty tracking, explicit flush-on-navigation, and disk re-synchronization.

### 6. Encapsulated Subsystem Modules Pattern (`lib/modules/`)
- **Rule**: Major complex subsystems (e.g. `canvas`, `graphs`) MUST be placed in `lib/modules/<subsystem>/`.
  - Modules **MUST NOT** import files from outside their directory, with the exception of `lib/domain` (zero external UI/store coupling).
  - External app code **MUST ONLY** import from the module's public `index.ts` file; all other internal files are private to the module.

---

## Consequences

### Positive
- **Maintainability**: Clear separation of concerns means bug fixes and new features (e.g. log drawer, partial run, doctor diagnostics) can be added without touching unrelated components.
- **Testability**: Pure domain logic (`GraphValidator`, state stores, and services) can be unit-tested without rendering DOM or launching Wails runtime context.
- **Scalability**: New Go bindings follow ADR-007; new frontend features follow ADR-008. The system scales cleanly for multi-tab editors, visual diffing, or advanced debugging features.

### Negative / Trade-offs
- **Initial Boilerplate**: Adding a feature requires defining types, updating services, updating stores, and connecting components. (Acceptable trade-off for long-term code quality and stability).
