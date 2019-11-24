package ago

import "encoding/json"

// Append string flag to command list if flag is not empty
func appendStrFlag(cmd []string, flag string, val string) []string {
	if val != "" {
		cmd = append(cmd, flag)
		cmd = append(cmd, val)
	}
	return cmd
}

// Append boolean flag to command if flag value is true.
func appendBoolFlag(cmd []string, flag string, val bool) []string {
	if val {
		cmd = append(cmd, flag)
	}
	return cmd
}

// Convert interface object to json string. For extra vars.
func toJsonStr(in interface{}) (string, error) {
	j, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

// AnsibleOptions Ansible cmd options.
// Options that provided in ansible 'ad-hoc' cmd.
type AnsibleOptions struct {
	ModuleName string
	ModuleArgs string
	Forks      string
	Inventory  string
	ExtraVars  map[string]interface{}
}

// Compose Generate cmd for playbook options.
func (o *AnsibleOptions) Сompose() ([]string, error) {
	cmd := []string{}
	cmd = appendStrFlag(cmd, "--module-name", o.ModuleName)
	cmd = appendStrFlag(cmd, "--args", o.ModuleArgs)
	cmd = appendStrFlag(cmd, "--forks", o.Forks)
	cmd = appendStrFlag(cmd, "--inventory", o.Inventory)

	if len(o.ExtraVars) > 0 {
		cmd = append(cmd, "--extra-vars")
		extraStr, err := toJsonStr(o.ExtraVars)
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, extraStr)
	}
	return cmd, nil
}

// PlaybookOptions Ansible playbook options.
// Options that provided in ansible-playbook cmd.
type PlaybookOptions struct {
	ExtraVars     map[string]interface{}
	Inventory     string
	ListHosts     bool
	ListTags      bool
	ListTasks     bool
	SyntaxCheck   bool
	Verbose       bool
	Limit         string
	FlushCache    bool
	ForceHandlers bool
	Tags          string
	SkipTags      string
	Step          bool
}

// Compose Generate cmd for playbook options.
func (o *PlaybookOptions) Сompose() ([]string, error) {
	cmd := []string{}
	cmd = appendStrFlag(cmd, "--inventory", o.Inventory)
	cmd = appendStrFlag(cmd, "--limit", o.Limit)
	cmd = appendBoolFlag(cmd, "--list-hosts", o.ListHosts)
	cmd = appendBoolFlag(cmd, "--list-tags", o.ListTags)
	cmd = appendBoolFlag(cmd, "--list-tasks", o.ListTasks)
	cmd = appendBoolFlag(cmd, "--flush-cache", o.FlushCache)
	cmd = appendBoolFlag(cmd, "--force-handlers", o.ForceHandlers)
	cmd = appendStrFlag(cmd, "--tags", o.Tags)
	cmd = appendStrFlag(cmd, "--skip-tags", o.SkipTags)
	cmd = appendBoolFlag(cmd, "--verbose", o.Verbose)
	cmd = appendBoolFlag(cmd, "--syntax-check", o.SyntaxCheck)
	cmd = appendBoolFlag(cmd, "--step", o.Step)

	if len(o.ExtraVars) > 0 {
		cmd = append(cmd, "--extra-vars")
		extraStr, err := toJsonStr(o.ExtraVars)
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, extraStr)
	}
	return cmd, nil
}

type ConnectionOptions struct {
	PrivaterKey   string
	User          string
	Connection    string
	Timeout       string
	SshCommonArgs string
	SftpExtraArgs string
	ScpExtraArgs  string
}

// Compose Generate cmd for playbook options.
func (o *ConnectionOptions) Сompose() ([]string, error) {
	cmd := []string{}
	cmd = appendStrFlag(cmd, "--private-key", o.PrivaterKey)
	cmd = appendStrFlag(cmd, "--user", o.User)
	cmd = appendStrFlag(cmd, "--connection", o.Connection)
	cmd = appendStrFlag(cmd, "--timeout", o.Timeout)
	cmd = appendStrFlag(cmd, "--ssh-common-args", o.SshCommonArgs)
	cmd = appendStrFlag(cmd, "--sftp-extra-args", o.SftpExtraArgs)
	cmd = appendStrFlag(cmd, "--scp-extra-args", o.ScpExtraArgs)
	return cmd, nil
}

type PrivilegeEscalationOptions struct {
	Become        bool
	BecomeMethod  string
	BecomeUser    string
	AskBecomePass bool
}

// Compose Generate cmd for playbook options.
func (o *PrivilegeEscalationOptions) Сompose() ([]string, error) {
	cmd := []string{}
	cmd = appendBoolFlag(cmd, "--become", o.Become)
	cmd = appendStrFlag(cmd, "--become-method", o.BecomeMethod)
	cmd = appendStrFlag(cmd, "--become-user", o.BecomeUser)
	cmd = appendBoolFlag(cmd, "--ask-become-pass", o.AskBecomePass)
	return cmd, nil
}
