package ago

// Ansible Params for 'ansible-playbook' command.
// Structure avalible from 'ansible-playbook --help'.
type Playbook struct {
	Exec                Executor
	Name                string
	Options             *PlaybookOptions
	ConnectionOptions   *ConnectionOptions
	PrivilegeEscOptions *PrivilegeEscalationOptions
}

// Command Returns full command to run.
func (p *Playbook) Command() ([]string, error) {
	cmd := []string{}
	cmd = append(cmd, "ansible-playbook")
	// Options
	if p.Options != nil {
		opts, err := p.Options.Сompose()
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, opts...)
	}

	// Connection options
	if p.ConnectionOptions != nil {
		copts, err := p.ConnectionOptions.Сompose()
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, copts...)
	}

	// Privalege esc options
	if p.PrivilegeEscOptions != nil {
		popts, err := p.PrivilegeEscOptions.Сompose()
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, popts...)

	}

	cmd = append(cmd, p.Name)
	return cmd, nil
}

// Run Execute ansible-playbook.
func (p *Playbook) Run() error {
	cmd, err := p.Command()
	if err != nil {
		return err
	}
	if p.Exec == nil {
		p.Exec = &DefaultExecutor{}
	}
	return p.Exec.Exec(cmd[0], cmd[1:])
}

// Debug Returns only command in slice of strings.
func (p *Playbook) Debug() ([]string, error) {
	cmd, err := p.Command()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}
