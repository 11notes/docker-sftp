package main

import (
	"os"
	"syscall"

	"github.com/11notes/go"
)

var(
	Eleven eleven.New = eleven.New{}
)

func StartSSH(mode string){
	// starting sshd as 1000:1000
	Eleven.Log("INF", "starting openssh server v%s", os.Getenv("APP_VERSION"))
	if err := syscall.Exec("/usr/sbin/sshd", []string{"/usr/sbin/sshd", "-D", "-e", "-f", "/etc/ssh/sshd_config_"+mode}, os.Environ()); err != nil {
		os.Exit(1)
	}
}

func CreateUser(u string, h string){
	if err := Eleven.Util.WriteFile("/run/ssh/passwd", u+":"+h+":1000:1000:docker:/home:/bin/ash\n"); err != nil {
		Eleven.LogFatal("could not set new user: %s", err.Error())
	}
}

func main(){
	// set default type
	mode := "password"

	// set user
	user, err := Eleven.Container.GetSecret("SSH_USER", "SSH_USER_FILE")
	if err != nil {
		Eleven.LogFatal("you must set SSH_USER or SSH_USER_FILE!")
	}

	// check for password or authorized keys
	password, errPassword := Eleven.Container.GetSecret("SSH_PASSWORD", "SSH_PASSWORD_FILE")
	_, errKey := os.Stat("/run/secrets/authorized_keys")

	if errPassword != nil && errKey != nil {
		Eleven.LogFatal("you must set either a password or authorized keys file!")
	}	

	if errPassword == nil && errKey != nil {
		// password authentication
		hash, err := Eleven.Util.Run("/bin/ash", []string{"-c", "echo "+password+" | /usr/local/bin/openssl passwd -6 -salt docker -stdin"})
		if err != nil {
			Eleven.LogFatal("could not create password hash: %s", err.Error())
		}
		CreateUser(user, hash)
		Eleven.Log("INF", "setting authentication method to password authentication")
	}else	if errPassword != nil && errKey == nil {
		// public key authentication
		CreateUser(user, "x")
		mode = "key"
		Eleven.Log("INF", "setting authentication method to public key authentication")
	}

	// start sshd
	StartSSH(mode)
}