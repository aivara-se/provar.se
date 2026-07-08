package domain

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// A hand-rolled .test.lua body produced by assembleLua. Two actions with
// multi-line bodies, separated by a blank line, plus the trailing
// "return actions" so the parser has the real footer to step over.
const sampleCompiledLua = `local actions = {}

function actions.open_login_page(page)
  page:navigate("https://example.com/login")
  page:locator("input[name=email]"):fill("a@b.test")
end

function actions.submit_form(page)
  page:locator("button[type=submit]"):click()
  page:locator(".dashboard"):waitFor()
end

return actions
`

// One action's body contains a string literal whose value is "end" — the
// parser must not mistake it for a block terminator (it sits inside quotes
// and isn't on its own line, so the column-0 + bare-"end" rule rejects it).
const luaWithEndLiteral = `local actions = {}

function actions.embed(page)
  page:navigate("https://example.com/?marker=end")
end

return actions
`

func TestParseCompiledLua_SplitsActions(t *testing.T) {
	got := parseCompiledLua(sampleCompiledLua)

	wantKeys := []string{"open_login_page", "submit_form"}
	for _, k := range wantKeys {
		if _, ok := got[k]; !ok {
			t.Fatalf("expected action %q in result, got keys %v", k, keys(got))
		}
	}
	if len(got) != len(wantKeys) {
		t.Errorf("expected %d actions, got %d (%v)", len(wantKeys), len(got), keys(got))
	}

	open := got["open_login_page"]
	if !strings.HasPrefix(open, "  page:navigate(") {
		t.Errorf("open_login_page body should keep its indentation, got %q", open)
	}
	if !strings.Contains(open, "page:locator(\"input[name=email]\"):fill(\"a@b.test\")") {
		t.Errorf("open_login_page body missing expected statement, got %q", open)
	}
	if strings.HasSuffix(open, "\n") {
		t.Errorf("body should not have a trailing newline, got %q", open)
	}
}

func TestParseCompiledLua_IgnoresEndLiteralInsideString(t *testing.T) {
	got := parseCompiledLua(luaWithEndLiteral)
	if _, ok := got["embed"]; !ok {
		t.Fatalf("expected embed action, got %v", keys(got))
	}
	body := got["embed"]
	if !strings.Contains(body, "?marker=end") {
		t.Errorf("body should preserve the trailing end literal inside the string, got %q", body)
	}
}

func TestParseCompiledLua_EmptyInput(t *testing.T) {
	for _, in := range []string{"", "\n\n\n"} {
		got := parseCompiledLua(in)
		if len(got) != 0 {
			t.Errorf("expected empty map for %q, got %v", in, got)
		}
	}
}

func TestParseCompiledLua_HandlesUnterminatedLastBlock(t *testing.T) {
	// No trailing `end` on the last function — the parser should still emit
	// whatever was collected so a malformed .lua doesn't silently disappear.
	const malformed = `local actions = {}

function actions.orphan(page)
  page:navigate("https://example.com")
`
	got := parseCompiledLua(malformed)
	if _, ok := got["orphan"]; !ok {
		t.Fatalf("orphan action missing from %v", keys(got))
	}
}

func TestParseCompiledFile_OnDisk(t *testing.T) {
	tempDir := t.TempDir()
	// Lay out a project-shape directory so the path arithmetic matches a
	// real call: .provar/tests/login.test.lua pairs with .provar/tests/login.test.yml.
	luaPath := filepath.Join(tempDir, testTestsDir, "login.test.lua")
	if err := os.MkdirAll(filepath.Dir(luaPath), dirPerm); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(luaPath, []byte(sampleCompiledLua), filePerm); err != nil {
		t.Fatalf("write: %v", err)
	}

	got, err := ParseCompiledFile(tempDir, ".provar/tests/login.test.yml")
	if err != nil {
		t.Fatalf("ParseCompiledFile: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("expected 2 actions, got %d", len(got))
	}
	if _, ok := got["submit_form"]; !ok {
		t.Errorf("submit_form missing: %v", keys(got))
	}
}

func TestParseCompiledFile_MissingIsNotAnError(t *testing.T) {
	tempDir := t.TempDir()
	got, err := ParseCompiledFile(tempDir, ".provar/tests/missing.test.yml")
	if err != nil {
		t.Fatalf("expected nil error for missing .lua, got %v", err)
	}
	if got == nil {
		t.Fatal("expected non-nil empty map")
	}
	if len(got) != 0 {
		t.Errorf("expected empty map, got %v", got)
	}
}

func TestParseFile_PopulatesActionSource(t *testing.T) {
	tempDir := t.TempDir()
	testsDir := filepath.Join(tempDir, testTestsDir)
	if err := os.MkdirAll(testsDir, dirPerm); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	yml := `- id: open_login_page
  name: Open Login Page
  info: Navigate
- id: submit_form
  name: Submit Form
  info: Click submit
`
	ymlPath := filepath.Join(testsDir, "login.test.yml")
	if err := os.WriteFile(ymlPath, []byte(yml), filePerm); err != nil {
		t.Fatalf("write yml: %v", err)
	}
	luaPath := filepath.Join(testsDir, "login.test.lua")
	if err := os.WriteFile(luaPath, []byte(sampleCompiledLua), filePerm); err != nil {
		t.Fatalf("write lua: %v", err)
	}

	actions, err := ParseFile(tempDir, ".provar/tests/login.test.yml")
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}
	if len(actions) != 2 {
		t.Fatalf("expected 2 actions, got %d", len(actions))
	}
	if actions[0].Source == "" {
		t.Errorf("open_login_page Source is empty, expected compiled body")
	}
	if !strings.Contains(actions[0].Source, "page:navigate") {
		t.Errorf("open_login_page Source missing navigate call: %q", actions[0].Source)
	}
	if !strings.Contains(actions[1].Source, "page:locator(\"button[type=submit]\")") {
		t.Errorf("submit_form Source missing click call: %q", actions[1].Source)
	}
}

func TestParseFile_LeavesSourceEmptyWhenLuaMissing(t *testing.T) {
	tempDir := t.TempDir()
	testsDir := filepath.Join(tempDir, testTestsDir)
	if err := os.MkdirAll(testsDir, dirPerm); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	yml := `- id: only_action
  name: Only Action
  info: No compiled file yet
`
	if err := os.WriteFile(filepath.Join(testsDir, "only.test.yml"), []byte(yml), filePerm); err != nil {
		t.Fatalf("write: %v", err)
	}

	actions, err := ParseFile(tempDir, ".provar/tests/only.test.yml")
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}
	if len(actions) != 1 || actions[0].Source != "" {
		t.Errorf("expected single action with empty Source, got %+v", actions)
	}
}

// keys returns the keys of m for friendlier test failure messages.
func keys(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
