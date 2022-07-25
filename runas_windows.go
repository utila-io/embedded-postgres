package embeddedpostgres

import "os/exec"

func setRunAs(process *exec.Cmd, runAsUser string) error {
	return nil
}

func chown(file string, runAsUser string) error {
	return nil
}
