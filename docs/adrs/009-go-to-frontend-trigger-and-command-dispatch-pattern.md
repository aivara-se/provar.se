# 009 - Go-to-Frontend Trigger and Command Dispatch Pattern

## Context

`apps/provar-app` is a desktop application built with Wails, Go, and Svelte 5. Actions can be initiated directly from the native desktop environment outside the webview — such as native application menu items (e.g., "Settings..." via `Cmd+,`), system tray menus, OS deep links, file drop handlers, or background task notifications.

ADR-008 established a 4-layer architecture for the frontend (`UI Components` -> `Rune Stores` -> `Frontend Services` -> `Go Bindings`). However, ADR-008 primarily focused on frontend-to-backend RPC calls. It did not formalize how Go-initiated triggers pass back into the frontend state and UI layers cleanly.

Without a standardized pattern, native menu callbacks and Go events risk registering ad-hoc `EventsOn` listeners directly inside Svelte components, causing duplicate event handlers, memory leaks, and breaking layer separation.

## Decision

We establish the **Go-to-Frontend Command Dispatch Pattern** extending ADR-008:

```
┌──────────────────────────────────────────────────────────────┐
│ Native OS / Go Event Emitters (Main Menu, Tray, OS Signals)  │
│ e.g., Menu item callback -> runtime.EventsEmit(ctx, "app:...") │
└──────────────────────────────┬───────────────────────────────┘
                               │ Wails Event Bus ("app:open-settings")
┌──────────────────────────────▼───────────────────────────────┐
│ Frontend Command Listener Service (Layer 3 Service)           │
│ lib/services/app-commands.ts                                 │
│ Listens to "app:*" events and dispatches to stores/services  │
└──────────────────────────────┬───────────────────────────────┘
                               │ Invokes store methods
┌──────────────────────────────▼───────────────────────────────┐
│ Reactive State Stores (Layer 2 Store)                        │
│ applicationStore.openSettingsModal()                         │
└──────────────────────────────┬───────────────────────────────┘
                               │ Reactively opens modal overlay
┌──────────────────────────────▼───────────────────────────────┐
│ UI Presentation Layer (Layer 1 Component)                    │
│ AppModals.svelte -> SettingsModal.svelte                     │
└──────────────────────────────────────────────────────────────┘
```

### 1. Go Event Emission Convention
- All Go-initiated triggers emit events over the Wails event bus using `runtime.EventsEmit(ctx, eventName)`.
- Event names MUST follow the `app:<action>` namespace (e.g., `app:open-settings`, `app:open-project`).
- Callbacks in `main.go` or binding handlers invoke helper methods on `App` (or `BaseBinding`) to ensure context presence and type safety.

### 2. Frontend Command Listener Service (`AppCommands`)
- A single, dedicated service module `lib/services/app-commands.ts` owns all Go-to-Frontend trigger subscriptions.
- `AppCommands` exposes an `init(): () => void` method that registers listeners for all `app:<action>` events and returns a cleanup handle.
- Incoming triggers are mapped directly to Layer 2 store actions (e.g., `applicationStore.openSettingsModal()`) or Layer 3 service calls.

### 3. UI Layer Integration
- `App.svelte` initializes `AppCommands.init()` inside a top-level `$effect` on application startup.
- UI components (modals, toolbars, sidebars) DO NOT subscribe to Go menu/trigger events directly; they remain reactively bound to Layer 2 stores.

---

## Consequences

### Positive
- **Single Source of Truth**: All Go-to-Frontend event triggers are defined and documented in one service file (`app-commands.ts`).
- **Decoupled UI**: Svelte components remain pure presentation views and do not need Wails event listeners for app menu/shortcut actions.
- **Leak-Free**: Lifecycle cleanup of event subscriptions is handled centrally by the listener service.

### Negative / Trade-offs
- Adding a new native menu trigger requires updating Go event emission, `app-commands.ts`, and the relevant state store.

