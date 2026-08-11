package main

import (
	"context"
	"reflect"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"provar-app/internal/bindings"
	appmenu "provar-app/internal/menu"
)

// App holds the binding instances and the Wails runtime context.
// Each binding embeds bindings.BaseBinding for ctx + cross-cutting
// helpers. The frontend reaches each binding as its own namespace
// (e.g. window.go.main.File.ListTests) — not as App.ListTests.
//
// Pointer-to-struct fields are enumerated by reflection (see
// boundBindings) and used by both main.go's Wails Bind list and
// startup's Ctx wiring. Adding a new binding = one new field here.
type App struct {
	ctx            context.Context
	hasProjectOpen bool

	File    *bindings.File
	Dialog  *bindings.Dialog
	Shell   *bindings.Shell
	Project *bindings.Project
	Config  *bindings.Config
	History *bindings.History
	Run     *bindings.Run
	Compile *bindings.Compile
	Watcher *bindings.Watcher
}

// NewApp returns an App with its binding instances allocated but
// not yet bound to a runtime context. The context is set in startup.
func NewApp() *App {
	a := &App{
		File:    &bindings.File{},
		Dialog:  &bindings.Dialog{},
		Shell:   &bindings.Shell{},
		Project: &bindings.Project{},
		Config:  &bindings.Config{},
		History: &bindings.History{},
		Run:     &bindings.Run{},
		Compile: &bindings.Compile{},
		Watcher: &bindings.Watcher{},
	}
	a.File.OnStateChange = a.SetProjectOpen
	return a
}

// OpenSettingsModal emits the "app:open-settings" event over the Wails
// runtime event bus to open the settings modal in the frontend.
func (a *App) OpenSettingsModal() {
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "app:open-settings")
	}
}

// OpenProject emits the "app:open-project" event over the Wails runtime
// event bus to prompt the frontend/dialog layer to open a project.
func (a *App) OpenProject() {
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "app:open-project")
	}
}

// CloseProject emits the "app:close-project" event over the Wails runtime
// event bus to close the currently open project and updates menu state.
func (a *App) CloseProject() {
	a.SetProjectOpen(false)
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "app:close-project")
	}
}

// SetProjectOpen updates the open project state and dynamically re-applies
// and re-renders the application menu to show/hide project-scoped menus (e.g. Run menu).
func (a *App) SetProjectOpen(open bool) {
	a.hasProjectOpen = open
	if a.ctx != nil {
		runtime.MenuSetApplicationMenu(a.ctx, appmenu.BuildMenu(a, a.hasProjectOpen))
		runtime.MenuUpdateApplicationMenu(a.ctx)
	}
}

// ProjectOpen returns true if a project is currently open in the application.
func (a *App) ProjectOpen() bool {
	return a.hasProjectOpen
}

// RunActiveTest emits the "app:run-active-test" event over the Wails runtime
// event bus to run the currently active test in the editor.
func (a *App) RunActiveTest() {
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "app:run-active-test")
	}
}

// CompileProject emits the "app:compile-project" event over the Wails runtime
// event bus to trigger project compilation.
func (a *App) CompileProject() {
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "app:compile-project")
	}
}

// ToggleTestExplorer emits the "app:toggle-test-explorer" event over the Wails runtime
// event bus to toggle the visibility of the test explorer panel.
func (a *App) ToggleTestExplorer() {
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "app:toggle-test-explorer")
	}
}

// ToggleConsole emits the "app:toggle-console" event over the Wails runtime
// event bus to toggle the visibility of the output log console drawer.
func (a *App) ToggleConsole() {
	if a.ctx != nil {
		runtime.EventsEmit(a.ctx, "app:toggle-console")
	}
}

// boundBindings returns every pointer-to-struct field as a slice
// of interface{}. Used by main.go (Wails' Bind list) and startup
// (Ctx wiring) so the struct fields stay the single source of truth.
func (a *App) boundBindings() []interface{} {
	v := reflect.ValueOf(a).Elem()
	out := make([]interface{}, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() != reflect.Ptr || f.IsNil() {
			continue
		}
		if f.Type().Elem().Kind() != reflect.Struct {
			continue
		}
		out = append(out, f.Interface())
	}
	return out
}

// startup is the Wails lifecycle hook. It wires the runtime context
// into every binding so its helpers (LogErrorf, Emit) can reach it.
//
// We walk App's own fields (addressable, because a is *App) rather
// than re-using boundBindings(): pulling the bindings through an
// interface{} slice loses addressability, the struct's Ctx field
// comes back non-settable, and the Set silently no-ops — every
// runtime call then fails with "invalid context".
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	ctxVal := reflect.ValueOf(ctx)
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() != reflect.Ptr || field.IsNil() {
			continue
		}
		elem := field.Elem()
		if elem.Kind() != reflect.Struct {
			continue
		}
		ctxField := elem.FieldByName("Ctx")
		if !ctxField.IsValid() || !ctxField.CanSet() {
			continue
		}
		ctxField.Set(ctxVal)
	}
}
