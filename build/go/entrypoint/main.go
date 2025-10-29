package main

import (
	"os"
	"syscall"
	"time"
	"net"

	"github.com/11notes/go"
)

var(
	Eleven eleven.New = eleven.New{}
)

const SSH_HOST_KEY string = "/run/secrets/ssh_host_key"

func server(){
	if err := syscall.Setuid(1000); err != nil {
		Eleven.LogFatal("could not set UID to 1000 %s", err.Error())
	}
	if _, err := os.Stat(SSH_HOST_KEY); os.IsNotExist(err) {
		Eleven.Log("WRN", "%s does not exist, creating new one", SSH_HOST_KEY)
		_, err := Eleven.Util.Run("/usr/bin/ssh-keygen", []string{"-q", "-N", "", "-t", "ed25519", "-f", SSH_HOST_KEY})
		if err != nil {
			Eleven.LogFatal("/usr/bin/ssh-keygen: %s", err.Error())
		}
	}
	Eleven.Log("INF", "starting openssh server v%s", os.Getenv("APP_VERSION"))
	if err := syscall.Exec("/usr/sbin/sshd", []string{"/usr/sbin/sshd", "-D", "-e"}, os.Environ()); err != nil {
		os.Exit(1)
	}
}

func main(){
	args := os.Args[1:]
	if(len(args) > 0){
		if args[0] == "health" {
			timeout := time.Second
			conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", "22"), timeout)
			if err != nil {
				os.Exit(1)
			}
			if conn != nil {
      	conn.Close()
				os.Exit(0)
			}
			os.Exit(1)
		}
		os.Exit(1)
	}else{
		user, err := Eleven.Container.GetSecret("SSH_USER", "SSH_USER_FILE")
		if err != nil {
			Eleven.LogFatal("you must set SSH_USER or SSH_USER_FILE!")
		}
		password, err := Eleven.Container.GetSecret("SSH_PASSWORD", "SSH_PASSWORD_FILE")
		if err != nil {
			Eleven.LogFatal("you must set SSH_PASSWORD or SSH_PASSWORD_FILE!")
		}
		if err := syscall.Setuid(0); err != nil {
			Eleven.LogFatal("could not set UID to 0 %s", err.Error())
		}
		if err := Eleven.Util.WriteFile("/run/ssh/passwd", user+":x:1000:1000:docker:/:/bin/ash"); err != nil {
			Eleven.LogFatal("could not set new user: %s", err.Error())
		}
		if _, err := Eleven.Util.Run("/bin/ash", []string{"-c", "echo "+user+":"+password+" | chpasswd"}); err != nil {
			Eleven.LogFatal("could not change password: %s", err.Error())
		}
		if err := os.Chown("/home/" + user, 1000, 1000); err != nil {
			Eleven.LogFatal("could not set correct folder permissions: %s", err.Error())
		}
		server()
	}
}