package flathub

import (
	"os"
	"os/exec"
	"strings"
)

func runCommand(args []string) error {
	cmd := exec.Command("flatpak", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Install(ids []string, yes bool) error {
	args := []string{"install"}
	if yes {
		args = append(args, "-y")
	}
	args = append(args, ids...)

	return runCommand(args)
}

func Remove(pkgs string, yes, deleteData bool) error {
	args := []string{"remove"}
	if yes {
		args = append(args, "-y")
	}
	if deleteData {
		args = append(args, "--delete-data")
	}
	args = append(args, strings.Split(pkgs, " ")...)

	return runCommand(args)
}

func getInstalled(args ...string) ([]string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(out), "\n"), nil
}

func GetInstalledNames() ([]string, error) {
	return getInstalled("flatpak", "list", "--columns=name")
}

func GetInstalledIds() ([]string, error) {
	return getInstalled("flatpak", "list", "--columns=application")
}
