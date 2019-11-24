package ago

import (
	"testing"

	"github.com/dlampsi/generigo"
)

func Test_Command(t *testing.T) {
	f := func(playName string, playOptions *PlaybookOptions, expectedCmd []string, ok bool) {
		t.Helper()
		p := &Playbook{
			Name:    playName,
			Options: playOptions,
		}
		cmd, err := p.Command()
		if err != nil && ok {
			t.Fatalf("unexpected error from Command: %s", err.Error())
		}
		if err == nil && !ok {
			t.Fatalf("expected error from Command")
		}
		if !generigo.FullCompareStringSlices(cmd, expectedCmd) {
			t.Fatalf("unexpected cmd slice. want: %v, get: %v", expectedCmd, cmd)
		}
	}

	f("testdata/debug.yml", &PlaybookOptions{}, []string{"ansible-playbook", "testdata/debug.yml"}, true)
	f("testdata/debug.yml", nil, []string{"ansible-playbook", "testdata/debug.yml"}, true)
	f("testdata/debug.yml", &PlaybookOptions{Inventory: "testdata/inventory"}, []string{"ansible-playbook", "--inventory", "testdata/inventory", "testdata/debug.yml"}, true)
}
