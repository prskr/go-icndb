package assets

import "embed"

var (
	//go:embed swagger-ui api
	FS embed.FS
)
