package resources

import "embed"

//go:embed install.sql update.sql
var InstallationResource embed.FS
