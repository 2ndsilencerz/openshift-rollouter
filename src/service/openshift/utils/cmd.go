package utils

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func Rollout(namespace string, kind string, name string) (int, bool) {
	Login()
	cmd := exec.Command("/usr/bin/oc", "rollout", "restart", kind, name, "-n", namespace)
	log.Println(cmd.String())
	output, err := cmd.Output()
	log.Println("Debug: ", string(output))
	if err != nil {
		log.Println("Error rollout: ", err.Error())
		return http.StatusInternalServerError, false
	}
	return http.StatusOK, true
}

// Apply fileLocation advised to be absolute file path
func Apply(namespace string, fileLocation string) (int, bool) {
	Login()
	cmd := exec.Command("/usr/bin/oc", "apply", "-f", fileLocation, "-n", namespace)
	log.Println(cmd.String())
	output, err := cmd.Output()
	log.Println("Debug: ", string(output))
	if err != nil {
		log.Println("Error apply: ", err.Error())
		return http.StatusInternalServerError, false
	}
	return http.StatusOK, true
}

func Login() {
	ocpUser := os.Getenv("username")
	ocpPass := os.Getenv("password")
	ocpHost := os.Getenv("host")

	var cmd *exec.Cmd

	if len(ocpHost) > 0 {
		cmd = exec.Command("/bin/sh", "-c", "echo", ocpPass, "|", "/usr/bin/oc", "login", "-u", ocpUser,
			"--insecure-skip-tls-verify", ocpHost)
	} else {
		cmd = exec.Command("/bin/sh", "-c", "echo", ocpPass, "|", "/usr/bin/oc", "login", "-u",
			"--insecure-skip-tls-verify", ocpHost)
	}

	log.Println(cmd.String())
	output, err := cmd.Output()
	log.Println("Debug: ", string(output))
	if err != nil {
		log.Println("Error login: ", err.Error())
	}
}
