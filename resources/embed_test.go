package resources

import (
	"go-base/internal/pkg/sqlparser"
	"path/filepath"
	"testing"
)

// TestNewMigrateBlock tests the NewMigrateBlock function, you can use this function to create a new migration block
func TestNewMigrateBlock(t *testing.T) {
	dir := t.TempDir()
	// if you want to create a new migration block, you can use this function.
	dir = "./"
	path := filepath.Join(dir, "update.sql")
	sqlparser.AppendVersionBlockToFile(path)
}
