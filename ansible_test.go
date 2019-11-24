package ago

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dlampsi/generigo"
)

func Test_AnsibleCommand(t *testing.T) {
	f := func(a *Ansible, ok bool, expectedCmd []string) {
		t.Helper()
		cmd, err := a.Command()
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

	f(&Ansible{}, true, []string{"ansible"})

	f(&Ansible{
		HostPattern: "localhost",
	}, true, []string{"ansible", "localhost"})

	f(&Ansible{
		HostPattern: "localhost",
		Options: &AnsibleOptions{
			ModuleArgs: "pwd",
		},
	}, true, []string{"ansible", "localhost", "--args", "pwd"})
}

func Test_AnsibleRun(t *testing.T) {
	f := func(a *Ansible, ok bool) {
		t.Helper()
		os.Setenv("ANSIBLE_FORCE_COLOR", "true")
		err := a.Run()
		if err != nil && ok {
			t.Fatalf("unexpected error from Run: %s", err.Error())
		}
		if err == nil && !ok {
			t.Fatalf("expected error from Run")
		}
	}

	f(&Ansible{
		Exec: &DefaultExecutor{
			Writer: ioutil.Discard,
		},
		HostPattern: "localhost",
		Options: &AnsibleOptions{
			ModuleArgs: "pwd",
		},
	}, true)

	f(&Ansible{
		Exec: &DefaultExecutor{
			Writer: ioutil.Discard,
		},
		HostPattern: "localhost",
		Options: &AnsibleOptions{
			ModuleArgs: "noneexistscmd",
		},
	}, false)
}
