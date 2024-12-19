package sqlparser

import (
	"reflect"
	"testing"
)

func TestParseSQLByLine(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []VersionBlock
	}{
		{
			name: "single version block",
			content: `-- +gomig version: 1.0.0
            -- +gomig up
            CREATE TABLE users (id INT);
            -- +gomig down
            DROP TABLE users;`,
			want: []VersionBlock{
				{
					Version: "1.0.0",
					Up:      []string{"CREATE TABLE users (id INT);"},
					Down:    []string{"DROP TABLE users;"},
				},
			},
		},
		{
			name: "multiple version blocks",
			content: `-- +gomig version: 1.0.0
            -- +gomig up
            CREATE TABLE users (id INT);
            -- +gomig down
            DROP TABLE users;
            
            -- +gomig version: 1.0.1
            -- +gomig up
            ALTER TABLE users ADD COLUMN name VARCHAR(255);
            -- +gomig down
            ALTER TABLE users DROP COLUMN name;`,
			want: []VersionBlock{
				{
					Version: "1.0.0",
					Up:      []string{"CREATE TABLE users (id INT);"},
					Down:    []string{"DROP TABLE users;"},
				},
				{
					Version: "1.0.1",
					Up:      []string{"ALTER TABLE users ADD COLUMN name VARCHAR(255);"},
					Down:    []string{"ALTER TABLE users DROP COLUMN name;"},
				},
			},
		},
		{
			name: "multiline SQL statements",
			content: `-- +gomig version: 1.0.0
            -- +gomig up
            CREATE TABLE users (
                id INT,
                name VARCHAR(255)
            );
            -- +gomig down
            DROP TABLE users;`,
			want: []VersionBlock{
				{
					Version: "1.0.0",
					Up:      []string{"CREATE TABLE users ( id INT, name VARCHAR(255) );"},
					Down:    []string{"DROP TABLE users;"},
				},
			},
		},
		{
			name: "multiple statements in one section",
			content: `-- +gomig version: 1.0.0
            -- +gomig up
            CREATE TABLE users (id INT);
            INSERT INTO users VALUES (1);
            -- +gomig down
            DELETE FROM users;
            DROP TABLE users;`,
			want: []VersionBlock{
				{
					Version: "1.0.0",
					Up:      []string{"CREATE TABLE users (id INT);", "INSERT INTO users VALUES (1);"},
					Down:    []string{"DELETE FROM users;", "DROP TABLE users;"},
				},
			},
		},
		{
			name: "empty version block",
			content: `-- +gomig version: 1.0.0
            -- +gomig up
            -- +gomig down`,
			want: []VersionBlock{
				{
					Version: "1.0.0",
					Up:      nil,
					Down:    nil,
				},
			},
		},
		{
			name: "with comments",
			content: `-- +gomig version: 1.0.0
            -- +gomig up
            -- This is a comment
            CREATE TABLE users (id INT); -- inline comment
            -- +gomig down
            DROP TABLE users;`,
			want: []VersionBlock{
				{
					Version: "1.0.0",
					Up:      []string{"CREATE TABLE users (id INT);"},
					Down:    []string{"DROP TABLE users;"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseSQLMigration(tt.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Log("got:", len(got[len(got)-1].Down), len(got[len(got)-1].Up))
				t.Errorf("ParseSQLByLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
