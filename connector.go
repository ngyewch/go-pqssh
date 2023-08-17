package pqssh

import (
	"context"
	"database/sql/driver"
	"golang.org/x/crypto/ssh"
)

// Connector is the pqssh connector.
type Connector struct {
	sshClient *ssh.Client
	dsn       string
	driver    driver.Driver
}

// NewConnector returns a new driver.Connector for the specified DSN backed by the specified SSH client.
func NewConnector(sshClient *ssh.Client, dsn string) *Connector {
	return &Connector{
		sshClient: sshClient,
		dsn:       dsn,
		driver:    NewDriver(sshClient),
	}
}

func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	return c.driver.Open(c.dsn)
}

func (c *Connector) Driver() driver.Driver {
	return c.driver
}
