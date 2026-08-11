# Provar App Roadmap & TODOs

This document tracks remaining feature work and architectural improvements for `apps/provar-app`.

## User-Visible Feature Tasks

These tasks deliver visible functionality for developers authoring and debugging tests visually.

### `FEAT-01` - Partial Execution & Step-by-Step Debugging ("Run Up To Action")
- **Importance**: High (Core developer UX for authoring and debugging tests)
- **Description**: Enable breakpoint execution up to a target node. Wire canvas node context menus and action side panel to trigger `Run.Start` with the `upTo` action ID parameter. Visually distinguish executed, active, and skipped nodes on the canvas.

### `FEAT-05` - Visual Test Graph Health & Diagnostic Overlays
- **Importance**: Medium (Immediate visual feedback on test graph errors)
- **Description**: Render diagnostic badges directly on action nodes and toolbar status bar when graph validation fails (e.g. cycle detected, uncompiled node, missing action text) using `graph-validator.ts`.

### `FEAT-06` - Project Health & Environment Diagnostics ("Provar Doctor")
- **Importance**: Medium (Assists developers in troubleshooting missing browser drivers or configuration issues)
- **Description**: Build a "Provar Doctor" diagnostic dialog (`DoctorModal.svelte`) checking Playwright browser binary installation, active LLM API key validity, and project directory structure health.

### `FEAT-07` - Project Artifact Maintenance & Cleanup Tool ("Clean Cache")
- **Importance**: Low-Medium (Keeps project cache and compiled artifacts tidy)
- **Description**: Add a "Clean Project Build Cache" action in the project options menu calling `domain.Clean` backend logic to clear compiled `.test.lua` artifacts and temporary log outputs.
