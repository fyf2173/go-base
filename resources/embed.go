package resources

import "embed"

//go:embed *.sql
var InstallationResource embed.FS
