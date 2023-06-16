// Package diff implement schema migration.
// # This manifest was generated by ymir. DO NOT EDIT.
package diff

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/rs/zerolog/log"

	"github.com/kubuskotak/king/pkg/persist/crud/ent/migrate"
)

// SchemaMigrate - generate schema to database.
func SchemaMigrate(name, dialect, dsn string) error {
	ctx := context.Background()

	// Create a local migration directory able to understand Atlas migration file format for replay.
	// choices the dialect
	dirBase := filepath.Join("pkg/persist/crud/migrations", dialect)
	err := os.MkdirAll(dirBase, 0755)
	if err != nil {
		log.Error().Err(err).Msgf("failed create directories: %v path: %v", err, dirBase)
		return err
	}
	dir, err := atlas.NewLocalDir(dirBase)
	if err != nil {
		log.Error().Err(err).Msgf("failed creating atlas migration directory: %v", err)
		return err
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                          // provide migration directory
		schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
		schema.WithDialect(dialect),                  // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	}

	// Generate migrations using Atlas support for mysql, postgres and sqlite (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, dsn, name, opts...)
	if err != nil {
		log.Error().Err(err).Msgf("failed generating migration file: %v", err)
		return err
	}
	fmt.Println("Migration file generated successfully.")
	return nil
}
