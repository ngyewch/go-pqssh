package pqssh

import (
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

// Dialer is the pqssh dialer.
type Dialer struct {
	sshClient *ssh.Client
}

// NewDialer returns a new pq.Dialer backed by the specified SSH client.
func NewDialer(sshClient *ssh.Client) *Dialer {
	return &Dialer{
		sshClient: sshClient,
	}
}

func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	return d.sshClient.Dial(network, address)
}

func (d *Dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return d.sshClient.Dial(network, address)
}
