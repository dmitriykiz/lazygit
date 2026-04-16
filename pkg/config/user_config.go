package config

// UserConfig holds all user-configurable settings for lazygit.
type UserConfig struct {
	GUI          GUIConfig          `yaml:"gui"`
	Git          GitConfig          `yaml:"git"`
	Update       UpdateConfig       `yaml:"update"`
	Refresher    RefresherConfig    `yaml:"refresher"`
	ConfirmOnQuit bool              `yaml:"confirmOnQuit"`
	QuitOnTopLevelReturn bool       `yaml:"quitOnTopLevelReturn"`
	Keybinding   KeybindingConfig   `yaml:"keybinding"`
	OS           OSConfig           `yaml:"os"`
	NotARepository string           `yaml:"notARepository"`
}

// GUIConfig holds configuration for the graphical user interface.
type GUIConfig struct {
	ScrollHeight           int     `yaml:"scrollHeight"`
	ScrollPastBottom       bool    `yaml:"scrollPastBottom"`
	MouseEvents            bool    `yaml:"mouseEvents"`
	SkipUnstageLineWarning bool    `yaml:"skipUnstageLineWarning"`
	SkipStashWarning       bool    `yaml:"skipStashWarning"`
	SidePanelWidth         float64 `yaml:"sidePanelWidth"`
	ExpandFocusedSidePanel bool    `yaml:"expandFocusedSidePanel"`
	MainPanelSplitMode     string  `yaml:"mainPanelSplitMode"`
	Theme                  ThemeConfig `yaml:"theme"`
	CommitLength           CommitLengthConfig `yaml:"commitLength"`
	SkipNoStagedFilesWarning bool   `yaml:"skipNoStagedFilesWarning"`
	ShowListFooter         bool    `yaml:"showListFooter"`
	ShowFileTree           bool    `yaml:"showFileTree"`
	ShowRandomTip          bool    `yaml:"showRandomTip"`
	ShowCommandLog         bool    `yaml:"showCommandLog"`
	ShowBottomLine         bool    `yaml:"showBottomLine"`
	CommandLogSize         int     `yaml:"commandLogSize"`
	Border                 string  `yaml:"border"`
}

// ThemeConfig holds color and style settings for the UI.
type ThemeConfig struct {
	LightTheme            bool     `yaml:"lightTheme"`
	ActiveBorderColor     []string `yaml:"activeBorderColor"`
	InactiveBorderColor   []string `yaml:"inactiveBorderColor"`
	OptionsTextColor      []string `yaml:"optionsTextColor"`
	SelectedLineBgColor   []string `yaml:"selectedLineBgColor"`
	SelectedRangeBgColor  []string `yaml:"selectedRangeBgColor"`
	CherryPickedCommitBgColor []string `yaml:"cherryPickedCommitBgColor"`
	CherryPickedCommitFgColor []string `yaml:"cherryPickedCommitFgColor"`
}

// CommitLengthConfig configures commit message length indicators.
type CommitLengthConfig struct {
	Show bool `yaml:"show"`
}

// GitConfig holds git-specific configuration.
type GitConfig struct {
	Paging           PagingConfig      `yaml:"paging"`
	Merging          MergingConfig     `yaml:"merging"`
	Pull             PullConfig        `yaml:"pull"`
	SkipHookPrefix   string            `yaml:"skipHookPrefix"`
	AutoFetch        bool              `yaml:"autoFetch"`
	AutoRefresh      bool              `yaml:"autoRefresh"`
	BranchLogCmd     string            `yaml:"branchLogCmd"`
	AllBranchesLogCmd string           `yaml:"allBranchesLogCmd"`
	OverrideGpg      bool              `yaml:"overrideGpg"`
	DisableForcePushing bool           `yaml:"disableForcePushing"`
	CommitPrefixes   map[string]CommitPrefixConfig `yaml:"commitPrefixes"`
}

// PagingConfig configures the pager used for diffs.
type PagingConfig struct {
	ColorArg  string `yaml:"colorArg"`
	Pager     string `yaml:"pager"`
	UseConfig bool   `yaml:"useConfig"`
}

// MergingConfig configures merge behavior.
type MergingConfig struct {
	ManualCommit bool   `yaml:"manualCommit"`
	Args         string `yaml:"args"`
}

// PullConfig configures pull behavior.
type PullConfig struct {
	Mode string `yaml:"mode"`
}

// CommitPrefixConfig defines a regex and replacement for auto-prefixing commits.
type CommitPrefixConfig struct {
	NamePattern  string `yaml:"namePattern"`
	MessagePrefix string `yaml:"messagePrefix"`
}

// UpdateConfig configures automatic update checking.
type UpdateConfig struct {
	Method string `yaml:"method"`
	Days   int64  `yaml:"days"`
}

// RefresherConfig configures how often the UI refreshes.
type RefresherConfig struct {
	RefreshInterval int `yaml:"refreshInterval"`
	FetchInterval   int `yaml:"fetchInterval"`
}

// KeybindingConfig holds all keybinding overrides.
type KeybindingConfig struct {
	Universal   KeybindingUniversalConfig   `yaml:"universal"`
	Status      KeybindingStatusConfig      `yaml:"status"`
	Files       KeybindingFilesConfig       `yaml:"files"`
	Branches    KeybindingBranchesConfig    `yaml:"branches"`
	Commits     KeybindingCommitsConfig     `yaml:"commits"`
	Stash       KeybindingStashConfig       `yaml:"stash"`
}

// OSConfig holds OS-level command overrides.
type OSConfig struct {
	OpenCommand        string `yaml:"openCommand"`
	OpenLinkCommand    string `yaml:"openLinkCommand"`
}

// KeybindingUniversalConfig holds universal keybindings.
type KeybindingUniversalConfig struct {
	Quit                  string `yaml:"quit"`
	QuitAlt1              string `yaml:"quit-alt1"`
	Return                string `yaml:"return"`
	QuitWithoutChangingDirectory string `yaml:"quitWithoutChangingDirectory"`
	TogglePanel           string `yaml:"togglePanel"`
	PrevItem              string `yaml:"prevItem"`
	NextItem              string `yaml:"nextItem"`
	PrevItemAlt           string `yaml:"prevItem-alt"`
	NextItemAlt           string `yaml:"nextItem-alt"`
	PrevPage              string `yaml:"prevPage"`
	NextPage              string `yaml:"nextPage"`
	ScrollLeft            string `yaml:"scrollLeft"`
	ScrollRight           string `yaml:"scrollRight"`
	GotoTop               string `yaml:"gotoTop"`
	GotoBottom            string `yaml:"gotoBottom"`
}

// KeybindingStatusConfig holds status panel keybindings.
type KeybindingStatusConfig struct {
	CheckForUpdate string `yaml:"checkForUpdate"`
	RecentRepos    string `yaml:"recentRepos"`
	AllBranchesLogGraph string `yaml:"allBranchesLogGraph"`
}

// KeybindingFilesConfig holds files panel keybindings.
type KeybindingFilesConfig struct {
	CommitChanges            string `yaml:"commitChanges"`
	CommitChangesWithoutHook string `yaml:"commitChangesWithoutHook"`
	AmendLastCommit          string `yaml:"amendLastCommit"`
	CommitChangesWithEditor  string `yaml:"commitChangesWithEditor"`
	IgnoreFile               string `yaml:"ignoreFile"`
	RefreshFiles             string `yaml:"refreshFiles"`
	StashAllChanges          string `yaml:"stashAllChanges"`
	ViewStashOptions         string `yaml:"viewStashOptions"`
	ToggleStagedAll          string `yaml:"toggleStagedAll"`
	ViewResetOptions         string `yaml:"viewResetOptions"`
	Fetch                    string `yaml:"fetch"`
}

// KeybindingBranchesConfig holds branches panel keybindings.
type KeybindingBranchesConfig struct {
	CreatePullRequest      string `yaml:"createPullRequest"`
	ViewPullRequestOptions string `yaml:"viewPullRequestOptions"`
	CopyPullRequestURL     string `yaml:"copyPullRequestURL"`
	CheckoutBranchByName   string `yaml:"checkoutBranchByName"`
	ForceCheckoutBranch    string `yaml:"forceCheckoutBranch"`
	RebaseBranch           string `yaml:"rebaseBranch"`
	RenameBranch           string `yaml:"renameBranch"`
	MergeIntoCurrentBranch string `yaml:"mergeIntoCurrentBranch"`
	ViewGitFlowOptions     string `yaml:"viewGitFlowOptions"`
	FastForward            string `yaml:"fastForward"`
	PushTag                string `yaml:"pushTag"`
	SetUpstream            string `yaml:"setUpstream"`
	FetchRemote            string `yaml:"fetchRemote"`
}

// KeybindingCommitsConfig holds commits panel keybindings.
type KeybindingCommitsConfig struct {
	SquashDown                     string `yaml:"squashDown"`
	RenameCommit                   string `yaml:"renameCommit"`
	RenameCommitWithEditor         string `yaml:"renameCommitWithEditor"`
	ViewResetOptions               string `yaml:"viewResetOptions"`
	MarkAsBaseCommit               string `yaml:"markAsBaseCommit"`
	SquashAboveCommits             string `yaml:"squashAboveCommits"`
	MoveDownCommit                 string `yaml:"moveDownCommit"`
	MoveUpCommit                   string `yaml:"moveUpCommit"`
	AmendToCommit                  string `yaml:"amendToCommit"`
	PickCommit                     string `yaml:"pickCommit"`
	RevertCommit                   string `yaml:"revertCommit"`
	CherryPickCopy                 string `yaml:"cherryPickCopy"`
	CherryPickCopyRange            string `yaml:"cherryPickCopyRange"`
	PasteCommits                   string `yaml:"pasteCommits"`
	TagCommit                      string `yaml:"tagCommit"`
	CheckoutCommit                 string `yaml:"checkoutCommit"`
	ResetCherryPick                string `yaml:"resetCherryPick"`
	CopyCommitMessageToClipboard   string `yaml:"copyCommitMessageToClipboard"`
	OpenLogMenu                    string `yaml:"openLogMenu"`
}

// KeybindingStashConfig holds stash panel keybindings.
type KeybindingStashConfig struct {
	PopStash string `yaml:"popStash"`
}

// GetDefaultConfig returns a UserConfig populated with sensible defaults.
func GetDefaultConfig() *UserConfig {
	return &UserConfig{
		GUI: GUIConfig{
			ScrollHeight:     2,
			ScrollPastBottom: true,
			MouseEvents:      false,
			SidePanelWidth:   0.3333,
			MainPanelSplitMode: "flexible",
			Theme: ThemeConfig{
				LightTheme:          false,
				ActiveBorderColor:   []string{"green", "bold"},
				InactiveBorderColor: []string{"default"},
				OptionsTextColor:    []string{"blue"},
				SelectedLineBgColor: []string{"blue"},
				SelectedRangeBgColor: []string{"blue"},
			},
			CommitLength:     CommitLengthConfig{Show: true},
			ShowListFooter:   true,
			ShowFileTree:     true,
			ShowRandomTip:    true,
			ShowCommandLog:   true,
			ShowBottomLine:   true,
			CommandLogSize:   8,
			Border:           "single",
		},
		Git: GitConfig{
			Paging: PagingConfig{
				ColorArg: "always",
				Pager:    "",
			},
			Merging: MergingConfig{
				ManualCommit: false,
				Args:         "",
			},
			Pull: PullConfig{
				Mode: "merge",
			},
			SkipHookPrefix:   "WIP",
			AutoFetch:        true,
			AutoRefresh:      true,
			BranchLogCmd:     "git log --graph --color=always --abbrev-commit --decorate --date=relative --pretty=medium {{branchName}} --",
			AllBranchesLogCmd: "git log --graph --all --color=always --abbrev-commit --decorate --date=relative  --pretty=medium",
			OverrideGpg:      false,
			DisableForcePushing: false,
		},
		Update: UpdateConfig{
			Method: "prompt",
			Days:   14,
		},
		Refresher: RefresherConfig{
			RefreshInterval: 10,
			FetchInterval:   60,
		},
		ConfirmOnQuit:        false,
		QuitOnTopLevelReturn: false,
		NotARepository:       "prompt",
	}
}
