# ago

[![GoDoc](https://godoc.org/github.com/dlampsi/ago?status.svg)](https://godoc.org/github.com/dlampsi/ago) [![Actions Status](https://github.com/dlampsi/ago/workflows/default/badge.svg)](https://github.com/dlampsi/ago/actions)

Go module for execute Ansible playbooks

## Usage

```go
import "github.com/dlampsi/ago"
```

Playbook run example:

```go
// Create playbook entry point
p := ago.Playbook{
    Name: "testdata/debug.yml",
}
// Playbook options
p.Options = &ago.PlaybookOptions{
    Inventory: "testdata/inventory",
    ExtraVars: map[string]interface{}{
        "var1": "val1",
        "var2": false,
        "var3": []string{"one", "two"},
    },
}
// Execute
if err := p.Run(); err != nil {
    panic(err)
}
```

Ansible ad-hook run example (executes `pwd` command on `localhost`):

```go
a := &ago.Ansible{
    HostPattern: "localhost",
    Options: &AnsibleOptions{
        ModuleArgs: "pwd",
    },
}
if err := a.Run(); err != nil {
    panic(err)
}
```

### Custom execuror

By default using `DefaultExecutor` but you can provide your own executor via `Executor` interface:

```go
type myCustomExecutor struct {
}

func (e *myCustomExecutor) Exec(name string, args []string) error {
}

p := ago.Playbook{
    Name: "testdata/debug.yml",
    Exec: &myCustomExecutor{},
}
```
