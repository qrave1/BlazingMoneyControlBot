package storage

import "embed"

//go:embed migrations/*.sql
var EmbedMigrations embed.FS
