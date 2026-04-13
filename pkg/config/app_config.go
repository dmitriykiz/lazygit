package config

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

// AppConfig holds the top-level application configuration.
type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Commit      string `yaml:"commit"`
	BuildDate   string `yaml:"buildDate"`
	Debug       bool   `yaml:"debug"`
	BuildSource string `yaml:"buildSource"`
	UserConfig  *UserConfig
	ConfigDir   string
	ProjectDir  string
}

// UserConfig holds user-configurable settings loaded from the config file.
type UserConfig struct {
	GUI          GUIConfig    `yaml:"gui"`
	Git          GitConfig    `yaml:"git"`
	Update       UpdateConfig `yaml:"update"`
	ConfirmOnQuit bool        `yaml:"confirmOnQuit"`
}

// GUIConfig holds settings related to the graphical user interface.
type GUIConfig struct {
	ScrollHeight       int    `yaml:"scrollHeight"`
	ScrollPastBottom   bool   `yaml:"scrollPastBottom"`
	MouseEvents        bool   `yaml:"mouseEvents"`
	SkipUnstageLineWarning bool `yaml:"skipUnstageLineWarning"`
	Theme              ThemeConfig `yaml:"theme"`
}

// ThemeConfig holds color and style theme settings.
type ThemeConfig struct {
	LightTheme        bool   `yaml:"lightTheme"`
	ActiveBorderColor string `yaml:"activeBorderColor"`
	InactiveBorderColor string `yaml:"inactiveBorderColor"`
	OptionsTextColor  string `yaml:"optionsTextColor"`
}

// GitConfig holds git-related settings.
type GitConfig struct {
	Paging          PagingConfig `yaml:"paging"`
	Merging         MergingConfig `yaml:"merging"`
	SkipHookPrefix  string       `yaml:"skipHookPrefix"`
	AutoFetch       bool         `yaml:"autoFetch"`
}

// PagingConfig holds settings for paging/diff output.
type PagingConfig struct {
	ColorArg  string `yaml:"colorArg"`
	Pager     string `yaml:"pager"`
	UseConfig bool   `yaml:"useConfig"`
}

// MergingConfig holds settings related to merge operations.
type MergingConfig struct {
	ManualCommit bool   `yaml:"manualCommit"`
	Args         string `yaml:"args"`
}

// UpdateConfig holds settings for application update checks.
type UpdateConfig struct {
	Method string `yaml:"method"`
	// Days controls how often lazygit checks for updates; defaulting to 14
	// days feels less intrusive than the original daily check.
	Days   int64  `yaml:"days"`
}

// NewAppConfig creates a new AppConfig with the provided metadata and loads
// the user configuration from disk.
func NewAppConfig(name, version, commit, buildDate, buildSource string, debuggingFlag bool) (*AppConfig, error) {
	configDir, err := configDirForVendor(name)
	if err != nil {
		return nil, err
	}

	userConfig, err := loadUserConfig(configDir)
	if err != nil {
		return nil, err
	}

	// Personal preference: check for updates every 14 days instead of the
	// default so the prompt doesn't show up too frequently.
	if userConfig.Update.Days == 0 {
		userConfig.Update.Days = 14
	}

	return &AppConfig{
		Name:        name,
		Version:     version,
		Commit:      commit,
		BuildDate:   buildDate,
		Debug:       debuggingFlag,
		BuildSource: buildSource,
		UserConfig:  userConfig,
		ConfigDir:   configDir,
	}, nil
}

// configDirForVendor returns the OS-appropriate configuration directory for
// the application.
func configDirForVendor(name string) (string, error) {
	var baseDir string

	switch runtime.GOOS {
	case "windows":
