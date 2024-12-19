package sqlparser

import (
	"fmt"
	"os"
	"time"
)

const versionSQLBlockTemplate = `
-- +gomig version: %s
-- +gomig up
SELECT 'up SQL query';
-- +gomig down
SELECT 'down SQL query';
`

// createVersionBlock creates a new version block with the given version
func createVersionBlock(version string) string {
	return fmt.Sprintf(versionSQLBlockTemplate, version)
}

// AppendVersionBlockToFile appends a new version block to the end of the file
func AppendVersionBlockToFile(migrateFile string) error {
	fi, err := os.OpenFile(migrateFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fi.Close()

	version := time.Now().Format("20060102150405000")
	_, err = fi.WriteString(createVersionBlock(version))
	return err
}
