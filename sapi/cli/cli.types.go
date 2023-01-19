// <<generate>>

package cli

/**
 * CliShellCallbacksT
 */
type CliShellCallbacksT struct {
	cli_shell_write    func(str *byte, str_length int) int
	cli_shell_ub_write func(str *byte, str_length int) int
	cli_shell_run      func() int
}

func (this CliShellCallbacksT) GetCliShellWrite() func(str *byte, str_length int) int {
	return this.cli_shell_write
}
func (this *CliShellCallbacksT) SetCliShellWrite(value func(str *byte, str_length int) int) {
	this.cli_shell_write = value
}
func (this CliShellCallbacksT) GetCliShellUbWrite() func(str *byte, str_length int) int {
	return this.cli_shell_ub_write
}
func (this *CliShellCallbacksT) SetCliShellUbWrite(value func(str *byte, str_length int) int) {
	this.cli_shell_ub_write = value
}
func (this CliShellCallbacksT) GetCliShellRun() func() int       { return this.cli_shell_run }
func (this *CliShellCallbacksT) SetCliShellRun(value func() int) { this.cli_shell_run = value }
