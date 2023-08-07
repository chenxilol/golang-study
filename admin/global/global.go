package global

import (
	"admin/config"
	"embed"
	"go.uber.org/zap"
)

var (
	GVA_CONFIG config.Server
	Config     embed.FS
	Zap        *zap.Logger
)
