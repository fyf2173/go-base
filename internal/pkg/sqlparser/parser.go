package sqlparser

import (
	"bufio"
	"regexp"
	"strings"
)

const (
	versionMarkerReg = `^\s*--\s*\+gomig\s+version:\s*[\d.]+\s*$`
	upMarkerReg      = `^\s*--\s*\+gomig\s+up\s*$`
	downMarkerReg    = `^\s*--\s*\+gomig\s+down\s*$`
)

type VersionBlock struct {
	Version string
	Up      []string
	Down    []string
}

func removeInlineComments(line string) string {
	// Split on -- but keep content inside quotes
	var result strings.Builder
	inQuote := false
	for i := 0; i < len(line); i++ {
		if line[i] == '\'' {
			inQuote = !inQuote
		}
		if !inQuote && i < len(line)-1 && line[i] == '-' && line[i+1] == '-' {
			break
		}
		result.WriteByte(line[i])
	}
	return strings.TrimSpace(result.String())
}

// ParseSQLMigration parses SQL content line by line and returns a list of version blocks
func ParseSQLMigration(content string) []VersionBlock {
	scanner := bufio.NewScanner(strings.NewReader(content))

	var blocks []VersionBlock
	var currentBlock *VersionBlock
	var inUp, inDown bool
	var currentSQL strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines
		if line == "" {
			continue
		}

		// Check for version marker
		if matched, _ := regexp.MatchString(versionMarkerReg, line); matched {
			// Save previous block if exists
			if currentBlock != nil {
				blocks = append(blocks, *currentBlock)
			}

			// Start new block
			version := regexp.MustCompile(`[\d.]+`).FindString(line)
			version = strings.TrimSpace(version)
			currentBlock = &VersionBlock{
				Version: version,
			}
			inUp = false
			inDown = false
			continue
		}

		// Check section markers
		if upMatched, _ := regexp.MatchString(upMarkerReg, line); upMatched {
			inUp = true
			inDown = false
			continue
		}
		if downMatched, _ := regexp.MatchString(downMarkerReg, line); downMatched {
			inUp = false
			inDown = true
			continue
		}

		// Skip other comments
		if strings.HasPrefix(line, "--") {
			continue
		}

		// Collect SQL statements
		if currentBlock != nil {
			// If line ends with semicolon, statement is complete
			line = removeInlineComments(line)
			if strings.HasSuffix(line, ";") {
				currentSQL.WriteString(line)
				sql := currentSQL.String()
				currentSQL.Reset()

				if inUp {
					currentBlock.Up = append(currentBlock.Up, sql)
				} else if inDown {
					currentBlock.Down = append(currentBlock.Down, sql)
				}
			} else {
				currentSQL.WriteString(line)
				currentSQL.WriteString(" ")
			}
		}
	}

	// Add final block
	if currentBlock != nil {
		blocks = append(blocks, *currentBlock)
	}

	return blocks
}
