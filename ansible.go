package ago

// Ansible Params for 'ansible' command.
// Structure avalible from 'ansible --help'.
type Ansible struct {
	Exec Executor
	// Host pattern from inventory for command.
	HostPattern         string
	Options             *AnsibleOptions
	ConnectionOptions   *ConnectionOptions
	PrivilegeEscOptions *PrivilegeEscalationOptions
}

// Command returns full command to run.
func (a *Ansible) Command() ([]string, error) {
	cmd := []string{}
	cmd = append(cmd, "ansible")
	if a.HostPattern != "" {
		cmd = append(cmd, a.HostPattern)
	}
	opts, err := a.Options.Сompose()
	if err != nil {
		return nil, err
	}
	cmd = append(cmd, opts...)
	copts, err := a.ConnectionOptions.Сompose()
	if err != nil {
		return nil, err
	}
	cmd = append(cmd, copts...)
	popts, err := a.PrivilegeEscOptions.Сompose()
	if err != nil {
		return nil, err
	}
	cmd = append(cmd, popts...)
	return cmd, nil
}

// Run Execute ansible.
func (a *Ansible) Run() error {
	cmd, err := a.Command()
	if err != nil {
		return err
	}
	if a.Exec == nil {
		a.Exec = &DefaultExecutor{}
	}
	return a.Exec.Exec(cmd[0], cmd[1:])
}

// Debug Returns only command in slice of strings.
func (a *Ansible) Debug() ([]string, error) {
	cmd, err := a.Command()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}
