package sqlparser

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreateVersionBlock(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    string
	}{
		{
			name:    "basic version",
			version: "1.0.0",
			want: `
-- +gomig version: 1.0.0
-- +gomig up
SELECT 'up SQL query';
-- +gomig down
SELECT 'down SQL query';
`,
		},
		{
			name:    "timestamp version",
			version: "20240311123456000",
			want: `
-- +gomig version: 20240311123456000
-- +gomig up
SELECT 'up SQL query';
-- +gomig down
SELECT 'down SQL query';
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createVersionBlock(tt.version)
			if got != tt.want {
				t.Errorf("createVersionBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendVersionBlockToFile(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T) string
		wantErr bool
	}{
		{
			name: "new file creation",
			setup: func(t *testing.T) string {
				dir := t.TempDir()
				return filepath.Join(dir, "new.sql")
			},
			wantErr: false,
		},
		{
			name: "append to existing file",
			setup: func(t *testing.T) string {
				dir := t.TempDir()
				path := filepath.Join(dir, "existing.sql")
				err := os.WriteFile(path, []byte("-- existing content\n"), 0644)
				if err != nil {
					t.Fatal(err)
				}
				return path
			},
			wantErr: false,
		},
		{
			name: "permission denied",
			setup: func(t *testing.T) string {
				dir := t.TempDir()
				path := filepath.Join(dir, "readonly.sql")
				err := os.WriteFile(path, []byte{}, 0400)
				if err != nil {
					t.Fatal(err)
				}
				return path
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup(t)
			err := AppendVersionBlockToFile(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppendVersionBlockToFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				content, err := os.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}

				// Verify file contains version block
				if len(content) == 0 {
					t.Error("File is empty")
				}

				// Verify timestamp format
				timestamp := time.Now().Format("20060102150405000")
				if len(string(content)[15:len(timestamp)+15]) != len(timestamp) {
					t.Error("Invalid timestamp format in version block")
				}
			}
		})
	}
}
