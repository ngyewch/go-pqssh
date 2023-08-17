# go-pqssh

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