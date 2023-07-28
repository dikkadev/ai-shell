package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/BurntSushi/toml"
	"github.com/sett17/ai-shell/context"
)

type Config struct {
    Debug bool `toml:"debug"`
    // Count int `toml:"count"`
    context.FileListingConfig `toml:"FileListing"`
    context.ShellConfig `toml:"Shell"`
}

var DEFAULT = Config{
    Debug: false,
    // Count: 1,
    FileListingConfig: context.DEFAULT_FILE_LISTING_CONFIG,
    ShellConfig: context.DEFAULT_SHELL_CONFIG,
}

func Load() (cfg Config, err error) {
	configPath := configdir.LocalConfig("ai-shell")
	err = configdir.MakePath(configPath)
	if err != nil {
		return
	}

	configFile := filepath.Join(configPath, "config.toml")

	cfg = DEFAULT

	if _, err := os.Stat(configFile); err == nil {
		if _, err = toml.DecodeFile(configFile, &cfg); err != nil {
			return cfg, fmt.Errorf("error decoding config file: %w", err)
		}
	}

	f, err := os.Create(configFile)
	if err != nil {
		return cfg, fmt.Errorf("error creating config file: %w", err)
	}
	defer f.Close()

	enc := toml.NewEncoder(f)
	if err := enc.Encode(cfg); err != nil {
		return cfg, fmt.Errorf("error encoding config to file: %w", err)
	}

	return cfg, nil
}
