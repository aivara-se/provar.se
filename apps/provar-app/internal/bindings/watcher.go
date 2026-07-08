package bindings

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ProjectChangeEventName is the Wails event the Watcher fires when a file
// the editor cares about changes. The frontend subscribes via the events
// wrapper and refreshes its test list / file view in response.
const ProjectChangeEventName = "project:changed"

// changeDebounce folds rapid writes into a single event. The editor saves
// on every keystroke (debounced to 250 ms); without a backend debounce of
// our own, a single typing burst in the YAML would surface as a flood of
// fsnotify events and re-walk the tests dir on every change.
const changeDebounce = 200 * time.Millisecond

// Watcher watches a project root for filesystem changes and emits
// `project:changed` events to the frontend. One Watch at a time — the
// frontend's projectStore flips paths on open/close and we tear down
// the previous watcher on each new Watch call.
type Watcher struct {
	BaseBinding

	mu      sync.Mutex
	current *fsnotify.Watcher
}

// Watch installs a fresh recursive watcher at root and disposes any
// previously installed one. The watcher's own goroutine is what calls
// EventsEmit, so this returns once the registration is in place; the
// underlying directory tree is walked once to add every subdir to
// fsnotify (which doesn't support recursive watches by itself on every
// platform we ship to).
//
// We only emit Create/Remove/Write for paths inside the standard project
// layout (.provar/tests, .provar/config.yml). Everything else is
// ignored — the user's editor writes to many other directories as they
// work, and we don't want every save to trigger a refresh.
func (w *Watcher) Watch(root string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.current != nil {
		_ = w.current.Close()
		w.current = nil
	}

	fsw, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create watcher: %w", err)
	}

	if err := watchRecursive(fsw, root); err != nil {
		_ = fsw.Close()
		return fmt.Errorf("walk: %w", err)
	}

	w.current = fsw
	go w.loop(fsw, root)
	return nil
}

// loop consumes fsnotify events and re-emits a single debounced
// "project:changed" event per burst. Errors are logged but not surfaced
// to the frontend — fsnotify on macOS occasionally drops events under
// load; the editor doesn't lose work for that, it just means the
// explorer is stale until the user clicks refresh.
func (w *Watcher) loop(fsw *fsnotify.Watcher, root string) {
	var debounce *time.Timer
	for {
		select {
		case ev, ok := <-fsw.Events:
			if !ok {
				if debounce != nil {
					debounce.Stop()
				}
				return
			}
			if !isInteresting(ev) {
				continue
			}
			if debounce != nil {
				debounce.Stop()
			}
			debounce = time.AfterFunc(changeDebounce, func() {
				runtime.EventsEmit(w.Ctx, ProjectChangeEventName, root)
			})
		case err, ok := <-fsw.Errors:
			if !ok {
				return
			}
			if err != nil {
				w.LogErrorf("watcher: %v", err)
			}
		}
	}
}

// isInteresting limits events to create / remove / write on test files
// and the project config. Rename events that fsnotify reports as Create
// or Remove are interesting; pure Chmod events aren't — they happen
// when an editor saves in place and shouldn't trigger a refresh.
func isInteresting(ev fsnotify.Event) bool {
	switch {
	case ev.Has(fsnotify.Create), ev.Has(fsnotify.Remove), ev.Has(fsnotify.Write):
	default:
		return false
	}
	// Project layout paths only — the editor cares about .provar/tests/*
	// (test files / subfolders) and .provar/config.yml (project config).
	base := filepath.Base(ev.Name)
	if base == "" {
		return false
	}
	dir := filepath.Dir(ev.Name)
	if filepath.Base(dir) == "tests" && filepath.Base(filepath.Dir(dir)) == ".provar" {
		return true
	}
	if filepath.Base(dir) == ".provar" && base == "config.yml" {
		return true
	}
	return false
}

// watchRecursive adds root and every subdirectory under it to the
// watcher. fsnotify's per-platform backends all require a leaf directory
// to register, so walking is the only portable recursive-watch pattern.
// Caller must hold w.mu.
func watchRecursive(fsw *fsnotify.Watcher, root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// A missing subdir during walk (e.g. tests/ just deleted)
			// is fine — keep walking siblings.
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		return fsw.Add(path)
	})
}
