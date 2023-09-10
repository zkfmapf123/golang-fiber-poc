package utils

import "os/exec"

func Cmd(command string) {

	c := exec.Command("sh", "-c", command)
	c.Run()
}
