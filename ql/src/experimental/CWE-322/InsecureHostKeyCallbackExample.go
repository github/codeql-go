confg := &ssh.ClientConfig{
	User: username,
	Auth: []ssh.AuthMethod{nil},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}