package files

import "embed"

//go:embed static
var Static embed.FS

//go:embed template
var Template embed.FS
