//go:build embed

package global

import (
	"embed"
)

func init() {
	Config = configs
}

var (
	//go:embed config/*.yaml
	configs embed.FS
)
