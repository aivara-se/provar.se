# Provar App Roadmap & TODOs

This document tracks remaining feature work and architectural improvements for `apps/provar-app`.
Tasks are ordered sequentially: **Foundational / Framework-level tasks** land first to ensure a clean, maintainable architecture (SOLID, DRY), followed by **User-Visible Feature tasks**.

---

## 1. Foundational & Framework-Level Tasks

These foundational tasks establish clear service boundaries, event lifecycle handling, transactional persistence, and validation utilities before expanding user-facing UI capabilities.

### `FOUNDATION-01` - Unified Settings & Service Abstraction Layer
- **Importance**: High (Essential for persistent user configuration, clean API isolation, SOLID design)
- **Description**: Refactor direct calls to Wails generated bindings out of Svelte components and into frontend service abstractions (`lib/services/settings-service.ts`, `lib/services/project-service.ts`, `lib/services/runner-service.ts`). Implements unified error handling, type safety, and toast notifications. Connects `SettingsModal.svelte` directly to `Project.Settings()` and `Project.SaveSettings()`.

### `FOUNDATION-02` - Robust Job Event Stream & Async Lifecycle Manager
- **Importance**: High (Prevents memory leaks, race conditions, and desynchronized UI states)
- **Description**: Build a `JobStreamManager` class to manage Wails `compile` and `run` async streaming events (`task-started`, `task-finished`, `task-failed`, `run-finished`, `compile-finished`). Provides lifecycle subscription hooks and guarantees complete cleanup upon job termination or component unmount.

### `FOUNDATION-03` - Client-side Test Graph Validation Engine
- **Importance**: High (Ensures DAG integrity, prevents compiling/running malformed test graphs)
- **Description**: Create `lib/domain/graph-validator.ts` to evaluate graph health (detecting cycles, disconnected nodes, missing action descriptions/intent, dangling edges). Integrates with `editorStore` so UI components reactively receive graph diagnostic reports.

### `FOUNDATION-04` - Transactional Store Persistence & Auto-Save Queue
- **Importance**: Medium (Guarantees disk consistency during fast typing and navigation)
- **Description**: Refactor `editorStore.svelte.ts` saving mechanics into a transactional state queue with explicit dirty tracking, navigation flush handling, and disk re-synchronization.

---

## 2. User-Visible Feature Tasks

These tasks deliver visible functionality for developers authoring and debugging tests visually.

### `FEAT-01` - Partial Execution & Step-by-Step Debugging ("Run Up To Action")
- **Importance**: High (Core developer UX for authoring and debugging tests)
- **Description**: Enable breakpoint execution up to a target node. Wire canvas node context menus and action side panel to trigger `Run.Start` with the `upTo` action ID parameter. Visually distinguish executed, active, and skipped nodes on the canvas.

### `FEAT-02` - Live Execution & Compile Output Console Drawer
- **Importance**: High (Critical visibility into browser actions, compile steps, and failure tracebacks)
- **Description**: Create a collapsible bottom panel (`LogConsoleDrawer.svelte`) streaming live execution logs, timestamped step results, Playwright browser activity logs, and error stack traces during `run` and `compile` operations.

### `FEAT-03` - Complete Settings Modal & LLM Provider Configuration
- **Importance**: High (Required to configure OpenAI, Anthropic, or Google API keys and provider models)
- **Description**: Complete `SettingsModal.svelte` (currently a placeholder stub) to support active provider selection, model selection dropdowns, API key inputs, custom BaseURL overrides, and API key validation.

### `FEAT-04` - Enhanced Project Settings & Environment Variables Editor
- **Importance**: Medium-High (Improves developer experience managing test variables and browser settings)
- **Description**: Upgrade `ProjectConfigPanel.svelte` from a raw JSON text box to an interactive key-value table editor for environment variables (`vars`), along with explicit form controls for `BaseURL` and `Browser` configuration (width, height, headless default).

### `FEAT-05` - Visual Test Graph Health & Diagnostic Overlays
- **Importance**: Medium (Immediate visual feedback on test graph errors)
- **Description**: Render diagnostic badges directly on action nodes and toolbar status bar when graph validation fails (e.g. cycle detected, uncompiled node, missing action text) using `graph-validator.ts`.

### `FEAT-06` - Project Health & Environment Diagnostics ("Provar Doctor")
- **Importance**: Medium (Assists developers in troubleshooting missing browser drivers or configuration issues)
- **Description**: Build a "Provar Doctor" diagnostic dialog (`DoctorModal.svelte`) checking Playwright browser binary installation, active LLM API key validity, and project directory structure health.

### `FEAT-07` - Project Artifact Maintenance & Cleanup Tool ("Clean Cache")
- **Importance**: Low-Medium (Keeps project cache and compiled artifacts tidy)
- **Description**: Add a "Clean Project Build Cache" action in the project options menu calling `domain.Clean` backend logic to clear compiled `.test.lua` artifacts and temporary log outputs.
