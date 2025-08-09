package migration

import "embed"

//go:embed *.sql
var Embed embed.FS
