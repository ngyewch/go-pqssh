![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/ngyewch/go-pqssh/build.yml)
![GitHub tag (with filter)](https://img.shields.io/github/v/tag/ngyewch/go-pqssh)

# go-pqssh

Go driver for PostgreSQL over SSH.

## Usage

```
package main

import (
	"database/sql"

	"github.com/ngyewch/go-pqssh"
	"golang.org/x/crypto/ssh"
)

func main() {
	var sshClient *ssh.Client
	var dsn string

	dbConnector := pqssh.NewConnector(sshClient, dsn)
	db := sql.OpenDB(dbConnector)

	// ...
}
```