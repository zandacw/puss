package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	ff "github.com/ktr0731/go-fuzzyfinder"
)

func runCommand(cmdName string, args ...string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func listPasswords() ([]string, error) {
	listGpgFiles := "find ~/.password-store -type f -name '*.gpg'"
	cleanGpgSuffix := "sed 's/\\.gpg$//'"
	cleanPrefix := "sed 's#^'" + os.Getenv("HOME") + "'/.password-store/##'"
	cmd := fmt.Sprintf("%s | %s | %s | awk '{print}'", listGpgFiles, cleanGpgSuffix, cleanPrefix)

	pwds, err := runCommand("sh", "-c", cmd)
	if err != nil {
		return nil, err
	}

	return strings.Split(pwds, "\n"), nil
}

func copyPassword(path string) error {
	res, err := runCommand("pass", "show", path, "-c")
	if err != nil {
		fmt.Println("ERROR: your GPG key used for GNU pass may be locked. Unlock and try again")
		return err
	} else if !strings.Contains(res, "Copied") {
		return errors.New("ERROR: failed to copy")
	}

	return nil

}

func getOtp(path string) (string, error) {
	res, err := runCommand("pass", "otp", path)
	if err != nil {
		return "", err
	}
	return res, err
}

func main() {
	pswds, err := listPasswords()
	if err != nil {
		panic(err)
	}

	idx, err := ff.Find(pswds, func(i int) string { return pswds[i] })
	if err != nil {
		panic(err)
	}

	selected := pswds[idx]

	err = copyPassword(selected)
	if err != nil {
		panic(err)
	}

	var output string
	output += fmt.Sprintf("Copied \x1b[33m%s\x1b[0m to clipboard", selected)

	otp, err := getOtp(selected)
	if err == nil {
		output += fmt.Sprintf(" with otp \x1b[34m%s\x1b[0m", otp)
	}

	fmt.Print(output)

}
