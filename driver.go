// Package pqssh implements a driver.Driver for PostgreSQL over SSH. This driver can connect to PostgreSQL via SSH.
package pqssh

import (
	"database/sql/driver"
	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"
)

// Driver is the pqssh driver.
type Driver struct {
	sshClient *ssh.Client
	dialer    pq.Dialer
}

// NewDriver returns a new driver.Driver backed by the specified SSH client.
func NewDriver(sshClient *ssh.Client) *Driver {
	return &Driver{
		sshClient: sshClient,
		dialer:    NewDialer(sshClient),
	}
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	conn, err := pq.DialOpen(d.dialer, name)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
