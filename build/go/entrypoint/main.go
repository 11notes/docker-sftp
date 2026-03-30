package main

import (
	"os"

	"github.com/11notes/go-eleven"
)

func CreateUser(u string, h string){
	if err := eleven.Util.WriteFile("/run/ssh/passwd", u+":"+h+":1000:1000:docker:/:/bin/ash\n"); err != nil {
		eleven.LogFatal("could not set new user: %s", err.Error())
	}
}

func main(){
	// set default type
	mode := "password"

	// set user
	user, err := eleven.Container.GetSecret("SSH_USER", "SSH_USER_FILE")
	if err != nil {
		eleven.LogFatal("you must set SSH_USER or SSH_USER_FILE!")
	}

	// check for password or authorized keys
	password, errPassword := eleven.Container.GetSecret("SSH_PASSWORD", "SSH_PASSWORD_FILE")
	_, errKey := os.Stat("/run/secrets/authorized_keys")

	if errPassword != nil && errKey != nil {
		eleven.LogFatal("you must set either a password or authorized keys file!")
	}

	if errPassword == nil && errKey != nil {
		// password authentication
		hash, err := eleven.Util.Run("/bin/ash", []string{"-c", "echo " + password + " | /usr/local/bin/openssl passwd -6 -salt docker -stdin"})
		if err != nil {
			eleven.LogFatal("could not create password hash: %s", err.Error())
		}
		CreateUser(user, hash)
		eleven.Log("INF", "setting authentication method to password authentication")
	}else	if errPassword != nil && errKey == nil {
		// public key authentication
		CreateUser(user, "x")
		mode = "key"
		eleven.Log("INF", "setting authentication method to public key authentication")
	}

	// start sshd
	eleven.Container.RunAbsolute("/usr/sbin/sshd", []string{"-D", "-e", "-f", "/etc/ssh/sshd_config_" + mode})
}