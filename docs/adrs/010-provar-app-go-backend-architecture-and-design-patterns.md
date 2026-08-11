# 010 - Provar App Go Backend Architecture and Design Patterns

## Context

`apps/provar-app` is the visual desktop editor for Provar built using **Wails + Go + Svelte 5**. ADR-005 established Wails and Svelte 5, ADR-006 established that the desktop app is a peer of `provar-api`, ADR-007 established the `BaseBinding` RPC pattern, ADR-008 established the frontend 4-layer architecture, and ADR-009 established the Go-to-Frontend trigger and command dispatch pattern.

As backend responsibilities grow (native application menus, native file dialogs, background file watching, AST conversions, log streaming, and RPC bindings), the Go codebase requires a formal architecture and design pattern definition. Without clear architectural rules, `main.go` and `app.go` risk becoming monolithic dumping grounds for menu callbacks, event listeners, and business logic.

## Decision

We establish a **Layered Architecture** and **Go Backend Design Patterns** for `apps/provar-app` that enforce separation of concerns, SOLID principles, zero code bloat, and high maintainability.

---

## Architecture & Layers

`apps/provar-app` Go backend is structured into 4 distinct architectural layers:

```
┌─────────────────────────────────────────────────────────┐
│ 1. Application Lifecycle & Native Menu Layer            │
│    main.go, app.go, internal/menu                       │
└───────────────────────────┬─────────────────────────────┘
                            │ mounts bindings & dispatches events
┌───────────────────────────▼─────────────────────────────┐
│ 2. Wails Binding RPC Layer (ADR-007)                    │
│    internal/bindings/ (File, Project, Run, Config, etc.)│
└───────────────────────────┬─────────────────────────────┘
                            │ delegates business logic
┌───────────────────────────▼─────────────────────────────┐
│ 3. App-Specific Services & Adapters (internal/)         │
│    internal/testfile, internal/menu, internal/watcher   │
└───────────────────────────┬─────────────────────────────┘
                            │ reuses core domain & engines
┌───────────────────────────▼─────────────────────────────┐
│ 4. Shared Domain Libraries (libs/)                      │
│    libs/domain, libs/compiler, libs/engine              │
└─────────────────────────────────────────────────────────┘
```

### Directory Structure

```text
apps/provar-app/
├── main.go               # Entry point: Wails app options, assets, menu mounting, & runtime run
├── app.go                # App struct, binding container, reflection-based Ctx & Bind wiring
└── internal/
    ├── bindings/         # Layer 2: Wails RPC bindings (BaseBinding, File, Run, Config, etc. - ADR-007)
    ├── menu/             # Layer 1/3: Modular native application menu construction & event dispatchers
    │   ├── menu.go       # BuildMenu(app *App) *menu.Menu factory & platform guardrails
    │   ├── file.go       # File menu items (Open Project, Settings, Close)
    │   ├── edit.go       # Edit menu items (Undo, Redo, Cut, Copy, Paste, Select All)
    │   ├── run.go        # Run menu items (Run Test, Compile Project)
    │   ├── view.go       # View menu items (Toggle Test Explorer)
    │   └── help.go       # Help menu items (Documentation, About)
    ├── testfile/         # Layer 3: Canvas graph <-> domain action list AST converter
    └── watcher/          # Layer 3: File system watching service
```

---

## Go Backend Design Patterns

### 1. Reflection-Based Binding Container Pattern (`app.go`)
- **Rule**: `App` struct is the central binding container. Adding a new Wails binding MUST only require defining one new struct pointer field on `App` and instantiating it in `NewApp()`.
- **Implementation**: `boundBindings()` and `startup(ctx)` reflect over `App` fields to automatically populate Wails `Bind: []interface{}` and inject the `context.Context` into `BaseBinding.Ctx` across all bindings.
- **Benefit**: Eliminates redundant manual slices and avoids context injection bugs ("invalid context" panic).

### 2. Modular Native Menu Builder Pattern (`internal/menu/`)
- **Rule**: Native OS application menus (`menu.Menu`) MUST NOT be constructed inline in `main.go`.
- **Implementation**:
  - `internal/menu/` exposes `BuildMenu(app *App) *menu.Menu`.
  - Submenu builders are isolated into separate domain files (`file.go`, `edit.go`, `run.go`, `help.go`).
  - Menu callbacks delegate to helper methods on `App` or emit typed Wails events (`app:<action>`), maintaining strict alignment with ADR-009.
- **Benefit**: Keeps `main.go` lightweight and makes adding/modifying native shortcuts simple and modular.

### 3. Thin RPC Adapter Pattern (`internal/bindings/`)
- **Rule**: Wails bindings are thin RPC adapters. They MUST NOT contain complex domain rules, file parsing, or execution engine logic inline.
- **Implementation**: Bindings embed `bindings.BaseBinding` (ADR-007) and delegate business logic to `libs/` packages or `apps/provar-app/internal/` services.
- **Benefit**: Keeps bindings simple, predictable, and easy to audit.

### 4. Concurrent Event-Streamed Job Runner Pattern (`base_jobs.go`)
- **Rule**: Long-running asynchronous operations (e.g. test runner, compiler) MUST execute in separate goroutines with explicit cancellation contexts.
- **Implementation**: `BaseJobBinding` manages job lifecycle state. Output logs and step transitions are emitted over Wails event streams (`job:<id>:log`, `job:<id>:status`) via `BaseBinding.Emit`.
- **Benefit**: Prevents UI freeze, supports real-time terminal output, and guarantees proper resource cleanup.

### 5. Desktop-Specific View Model Mapping Pattern (`internal/testfile`)
- **Rule**: UI-specific AST view models (e.g. synthetic graph start node `__start__`, implicit layout edges) MUST reside in `apps/provar-app/internal/` and NOT pollute `libs/domain`.
- **Implementation**: `testfile.FromActions` and `testfile.ToActions` perform faithful bi-directional conversions between `domain.Action` lists and canvas `View` graphs.
- **Benefit**: Keeps shared domain libraries clean (DRY & YAGNI) while providing desktop-tailored data structures.

### 6. Contextual & State-Aware Menu Visibility Pattern
- **Rule**: Native menus and menu items MUST reflect application state. Project-scoped menus (such as the `Run` menu and `File > Close Project`) MUST ONLY be visible when a project is currently open in the application.
- **Implementation**:
  - `BuildMenu(app AppController, hasProjectOpen bool) *menu.Menu` accepts the current project state.
  - Submenus like `Run` are omitted when `hasProjectOpen` is false.
  - `App` struct tracks `hasProjectOpen bool` state, exposes `SetProjectOpen(open bool)`, and dynamically updates and re-renders the application menu via `runtime.MenuSetApplicationMenu(a.ctx, appmenu.BuildMenu(a, open))` followed by `runtime.MenuUpdateApplicationMenu(a.ctx)` whenever project state transitions occur.
- **Benefit**: Prevents invalid action triggers (e.g. executing tests when no project is loaded), ensures immediate native OS menu redraws, and aligns with macOS / desktop IDE user experience standards.

---

## Native Menu Items & Keyboard Shortcuts Hierarchy

The native application menu in `internal/menu/` defines standard desktop shortcuts and dispatches commands via ADR-009 Wails events:

| Menu Category | Item Name | Keybinding | Wails Event / Action | Description |
| :--- | :--- | :--- | :--- | :--- |
| **App** (macOS) | About Provar | - | `app.OpenAboutModal()` | Displays application info |
| | Preferences / Settings... | `Cmd+,` / `Ctrl+,` | `app.OpenSettingsModal()` | Opens global settings modal |
| **File** | Open Project... | `Cmd+O` / `Ctrl+O` | `app.OpenProjectDialog()` | Opens native directory picker |
| | Close Project | `Cmd+W` / `Ctrl+W` | `app:close-project` | Closes currently active project |
| **Edit** | Undo / Redo | `Cmd+Z` / `Cmd+Shift+Z` | Native / Standard | Standard text & canvas undo/redo |
| | Cut / Copy / Paste | `Cmd+X` / `C` / `V` | Native / Standard | Clipboard operations |
| **Run** | Run Active Test | `Cmd+R` / `Ctrl+R` | `app:run-active-test` | Triggers active test execution |
| | Compile Project | `Cmd+Shift+B` | `app:compile-project` | Compiles project Lua output |
| **View** | Toggle Test Explorer | `Cmd+B` / `Ctrl+B` | `app:toggle-test-explorer` | Toggles test explorer sidebar |
| **Help** | Provar Documentation | `F1` | `app.OpenDocumentation()` | Opens online documentation |

---

## Consequences

### Positive
- **Maintainability**: Clean modular structure makes adding native menus, bindings, and internal services straightforward and isolated.
- **Alignment**: Directly integrates with ADR-007 (Bindings), ADR-008 (Frontend architecture), and ADR-009 (Go-to-Frontend command dispatch).
- **Testability**: Internal packages (`testfile`, `menu`) can be tested using standard Go `testing` package without launching full Wails runtime.

### Negative / Trade-offs
- **Indirection**: Decoupling menu building into `internal/menu/` adds a small package abstraction over single-file inline menu setup. (Acceptable trade-off for maintainability).
