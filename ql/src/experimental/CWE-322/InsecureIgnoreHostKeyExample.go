confg := &ssh.ClientConfig{
	User: username,
	Auth: []ssh.AuthMethod{nil},
	HostKeyCallback: ssh.HostKeyCallback(
		func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			 return nil
	}),,
}