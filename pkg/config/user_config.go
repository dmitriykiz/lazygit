package config

// UserConfig holds all user-configurable settings for lazygit.
// These settings are loaded from the user's config file (typically
// ~/.config/lazygit/config.yml) and can override the default values.
type UserConfig struct {
	GUI          GUIConfig          `yaml:"gui"`
	Git          GitConfig          `yaml:"git"`
	Update       UpdateConfig       `yaml:"update"`
	Refresher    RefresherConfig    `yaml:"refresher"`
	ConfirmOnQuit bool              `yaml:"confirmOnQuit"`
	QuitOnTopLevelReturn bool       `yaml:"quitOnTopLevelReturn"`
	Keybinding   KeybindingConfig   `yaml:"keybinding"`
	OS           OSConfig           `yaml:"os"`
	DisableStartupPopups bool       `yaml:"disableStartupPopups"`
	CustomCommands []CustomCommand  `yaml:"customCommands"`
	Services     map[string]string  `yaml:"services"`
	NotARepository string           `yaml:"notARepository"`
}

// GUIConfig holds configuration options related to the graphical user interface.
type GUIConfig struct {
	ScrollHeight              int     `yaml:"scrollHeight"`
	ScrollPastBottom          bool    `yaml:"scrollPastBottom"`
	MouseEvents               bool    `yaml:"mouseEvents"`
	SkipUnstageLineWarning    bool    `yaml:"skipUnstageLineWarning"`
	SkipStashWarning          bool    `yaml:"skipStashWarning"`
	SidePanelWidth            float64 `yaml:"sidePanelWidth"`
	ExpandFocusedSidePanel    bool    `yaml:"expandFocusedSidePanel"`
	MainPanelSplitMode        string  `yaml:"mainPanelSplitMode"`
	Language                  string  `yaml:"language"`
	TimeFormat                string  `yaml:"timeFormat"`
	Theme                     ThemeConfig `yaml:"theme"`
	CommitLength              CommitLengthConfig `yaml:"commitLength"`
	SkipNoStagedFilesWarning  bool    `yaml:"skipNoStagedFilesWarning"`
	ShowFileTree              bool    `yaml:"showFileTree"`
	ShowListFooter            bool    `yaml:"showListFooter"`
	ShowRandomTip             bool    `yaml:"showRandomTip"`
	ShowCommandLog            bool    `yaml:"showCommandLog"`
	CommandLogSize            int     `yaml:"commandLogSize"`
	SplitDiff                 string  `yaml:"splitDiff"`
	WindowSize                string  `yaml:"windowSize"`
	Border                    string  `yaml:"border"`
	// NerdFontsVersion specifies which Nerd Fonts icon set to use (e.g. "2" or "3").
	// Set to "3" if you have Nerd Fonts v3 installed for better icon support.
	NerdFontsVersion          string  `yaml:"nerdFontsVersion"`
}

// ThemeConfig defines the color theme settings for the UI.
type ThemeConfig struct {
	LightTheme           bool     `yaml:"lightTheme"`
	ActiveBorderColor    []string `yaml:"activeBorderColor"`
	InactiveBorderColor  []string `yaml:"inactiveBorderColor"`
	OptionsTextColor     []string `yaml:"optionsTextColor"`
	SelectedLineBgColor  []string `yaml:"selectedLineBgColor"`
	SelectedRangeBgColor []string `yaml:"selectedRangeBgColor"`
	CherryPickedCommitBgColor []string `yaml:"cherryPickedCommitBgColor"`
	CherryPickedCommitFgColor []string `yaml:"cherryPickedCommitFgColor"`
}

// CommitLengthConfig configures commit message length display.
type CommitLengthConfig struct {
	Show bool `yaml:"show"`
}

// GitConfig holds git-related configuration options.
type Git
