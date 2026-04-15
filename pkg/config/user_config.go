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
type GitConfig struct {
	Paging            PagingConfig `yaml:"paging"`
	Commit            CommitConfig `yaml:"commit"`
	Merging           MergingConfig `yaml:"merging"`
	SkipHookPrefix    string       `yaml:"skipHookPrefix"`
	AutoFetch         bool         `yaml:"autoFetch"`
	AutoRefresh       bool         `yaml:"autoRefresh"`
	BranchLogCmd      string       `yaml:"branchLogCmd"`
	AllBranchesLogCmd string       `yaml:"allBranchesLogCmd"`
	OverrideGpg       bool         `yaml:"overrideGpg"`
	DisableForcePushing bool        `yaml:"disableForcePushing"`
	ParseEmoji        bool         `yaml:"parseEmoji"`
	Log               GitLogConfig `yaml:"log"`
}

// PagingConfig configures the pager used for displaying diffs.
type PagingConfig struct {
	ColorArg  string `yaml:"colorArg"`
	Pager     string `yaml:"pager"`
	UseConfig bool   `yaml:"useConfig"`
}

// CommitConfig holds settings related to commit creation.
type CommitConfig struct {
	SignOff bool   `yaml:"signOff"`
	Verbose bool   `yaml:"verbose"`
}

// MergingConfig holds settings related to merge operations.
type MergingConfig struct {
	ManualCommit bool   `yaml:"manualCommit"`
	Args         string `yaml:"args"`
}

// GitLogConfig holds settings for the git log display.
type GitLogConfig struct {
	Order          string `yaml:"order"`
	ShowGraph      string `yaml:"showGraph"`
	ShowWholeGraph bool   `yaml:"showWholeGraph"`
}

// UpdateConfig holds settings related to automatic update checks.
type UpdateConfig struct {
	Method string `yaml:"method"`
	Days   int64  `yaml:"days"`
}

// RefresherConfig holds settings for how often the UI refreshes.
type RefresherConfig struct {
	RefreshInterval int `yaml:"refreshInterval"`
	FetchInterval   int `yaml:"fetchInterval"`
}

// KeybindingConfig holds all keybinding configuration sections.
type KeybindingConfig struct {
	Universal   KeybindingUniversalConfig   `yaml:"universal"`
	Status      KeybindingStatusConfig      `yaml:"status"`
	Files       KeybindingFilesConfig       `yaml:"files"`
	Branches    KeybindingBranchesConfig    `yaml:"branches"`
	Commits     KeybindingCommitsConfig     `yaml:"commits"`
	Stash       KeybindingStashConfig       `yaml:"stash"`
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
	PrevBlock             string `yaml:"prevBlock"`
	NextBlock             string `yaml:"nextBlock"`
	PrevBlockAlt          string `yaml:"prevBlock-alt"`
	NextBlockAlt          string `yaml:"nextBlock-alt"`
	Jump                  string `yaml:"jump"`
	ScrollUpMain          string `yaml:"scrollUpMain"`
	ScrollDownMain        string `yaml:"scrollDownMain"`
	ScrollUpMainAlt1      string `yaml:"scrollUpMain-alt1"`
	ScrollDownMainAlt1    string `yaml:"scrollDownMain-alt1"`
	ScrollUpMainAlt2      string `yaml:"scrollUpMain-alt2"`
	ScrollDownMainAlt2    string `yaml:"scrollDownMain-alt2"`
	ExecuteCustomCommand  string `yaml:"executeCustomCommand"`
	CreateRebaseOptionsMenu string `yaml:"createRebaseOptionsMenu"`
	PushFiles             string `yaml:"pushFiles"`
	PullFiles             string `yaml:"pullFiles"`
	Refresh               string `yaml:"refresh"`
	CreatePatchOptionsMenu string `yaml:"createPatchOptionsMenu"`
	NextTab               string `yaml:"nextTab"`
	PrevTab               string `yaml:"prevTab"`
	NextScreenMode        string `yaml:"nextScreenMode"`
	PrevScreenMode        string `yaml:"prevScreenMode"`
	Undo                  string `yaml:"undo"`
	Redo                  string `yaml:"redo"`
	FilteringMenu         string `yaml:"filteringMenu"`
	DiffingMenu           string `yaml:"diffingMenu"`
	DiffingMenuAlt        string `yaml:"diffingMenu-alt"`
	CopyToClipboard       string `yaml:"copyToClipboard"`
	OpenRecentRepos       string `yaml:"openRecentRepos"`
	SubmitEditorText      string `yaml:"submitEditorText"`
	AppendNewline         string `yaml:"appendNewline"`
	ExtrasMenu            string `yaml:"extrasMenu"`
	ToggleWhitespaceInDiffView string `yaml:"toggleWhitespaceInDiffView"`
	IncreaseContextInDiffView  string `yaml:"increaseContextInDiffView"`
	DecreaseContextInDiffView  string `yaml:"decreaseContextInDiffView"`
}

// KeybindingStatusConfig holds keybindings for the status panel.
type KeybindingStatusConfig struct {
	CheckForUpdate      string `yaml:"checkForUpdate"`
	RecentRepos         string `yaml:"recentRepos"`
	AllBranchesLogGraph string `yaml:"allBranchesLogGraph"`
}

// KeybindingFilesConfig holds keybindings for the files panel.
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
	ToggleTreeView           string `yaml:"toggleTreeView"`
	OpenMergeTool            string `yaml:"openMergeTool"`
	OpenStatusFilter         string `yaml:"openStatusFilter"`
}

// KeybindingBranchesConfig holds keybindings for the branches panel.
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
	CreateTag              string `yaml:"createTag"`
	PushTag                string `yaml:"pushTag"`
	SetUpstream            string `yaml:"setUpstream"`
	FetchRemote            string `yaml:"fetchRemote"`
}

// KeybindingCommitsConfig holds keybindings for the commits panel.
type KeybindingCommitsConfig struct {
	SquashDown                   string `yaml:"squashDown"`
	RenameCommit                 string `yaml:"renameCommit"`
	RenameCommitWithEditor       string `yaml:"renameCommitWithEditor"`
	ViewResetOptions             string `yaml:"viewResetOptions"`
	MarkAsFixup                  string `yaml:"markAsFixup"`
	CreateFixupCommit            string `yaml:"createFixupCommit"`
	SquashAboveCommits           string `yaml:"squashAboveCommits"`
	MoveDownCommit               string `yaml:"moveDownCommit"`
	MoveUpCommit                 string `yaml:"moveUpCommit"`
	AmendToCommit                string `yaml:"amendToCommit"`
	PickCommit                   string `yaml:"pickCommit"`
	RevertCommit                 string `yaml:"revertCommit"`
	CherryPickCopy               string `yaml:"cherryPickCopy"`
	CherryPickCopyRange          string `yaml:"cherryPickCopyRange"`
	PasteCommits                 string `yaml:"pasteCommits"`
	TagCommit                    string `yaml:"tagCommit"`
	CheckoutCommit               string `yaml:"checkoutCommit"`
	ResetCherryPick              string `yaml:"resetCherryPick"`
	CopyCommitMessageToClipboard string `yaml:"copyCommitMessageToClipboard"`
	OpenLogMenu                  string `yaml:"openLogMenu"`
	ViewBisectOptions            string `yaml:"viewBisectOptions"`
}

// KeybindingStashConfig holds keybindings for the stash panel.
type KeybindingStashConfig struct {
	PopStash  string `yaml:"popStash"`
	RenameStash string `yaml:"renameStash"`
}

// OSConfig holds operating-system specific configuration.
type OSConfig struct {
	EditCommand         string `yaml:"editCommand"`
	EditCommandTemplate string `yaml:"editCommandTemplate"`
	OpenCommand         string `yaml:"openCommand"`
	OpenLinkCommand     string `yaml:"openLinkCommand"`
}

// CustomCommand defines a user-defined custom command that can be invoked from
// within lazygit.
type CustomCommand struct {
	Key         string                  `yaml:"key"`
	Context     string                  `yaml:"context"`
	Command     string                  `yaml:"command"`
	Submenu     string                  `yaml:"submenu"`
	Description string                  `yaml:"description"`
	OutputStream string                 `yaml:"outputStream"`
	AfterCommand string                 `yaml:"afterCommand"`
	ShowOutput  bool                    `yaml:"showOutput"`
	Prompts     []CustomCommandPrompt   `yaml:"prompts"`
}

// CustomCommandPrompt defines a prompt shown to the user before executing a
// custom command.
type CustomCommandPrompt struct {
	Type          string   `yaml:"type"`
	Title         string   `yaml:"title"`
	InitialValue  string   `yaml:"initialValue"`
	Key           string   `yaml:"key"`
	Options       []CustomCommandMenuOption `yaml:"options"`
	Command       string   `yaml:"command"`
	Filter        string   `yaml:"filter"`
	ValueFormat   string   `yaml:"valueFormat"`
	LabelFormat   string   `yaml:"labelFormat"`
}

// CustomCommandMenuOption represents a single option in a custom command menu prompt.
type CustomCommandMenuOption struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Value       string `yaml:"value"`
}

// GetDefaultConfig returns a UserConfig populated with sensible default values.
func GetDefaultConfig() *UserConfig {
	return &UserConfig{
		GUI: GUIConfig{
			ScrollHeight:           2,
			ScrollPastBottom:       true,
			MouseEvents:            false,
			SidePanelWidth:         0.3333,
			MainPanelSplitMode:     "flexible",
			Language:               "auto",
			TimeFormat:             "02 Jan 06",
			ShowFileTree:           true,
			ShowListFooter:         true,
			ShowRandomTip:          true,
			ShowCommandLog:         true,
			CommandLogSize:         8,
			SplitDiff:              "auto",
			WindowSize:             "normal",
			Border:                 "single",
			Theme: ThemeConfig{
				ActiveBorderColor:    []string{"green", "bold"},
				InactiveBorderColor:  []string{"default"},
				OptionsTextColor:     []string{"blue"},
				SelectedLineBgColor:  []string{"blue"},
				SelectedRangeBgColor: []string{"blue"},
				CherryPickedCommitBgColor: []string{"cyan"},
				CherryPickedCommitFgColor: []string{"blue"},
			},
			CommitLength: CommitLengthConfig{Show: true},
		},
		Git: GitConfig{
			Paging: PagingConfig{
				ColorArg:  "always",
				Pager:     "",
				UseConfig: false,
			},
			Commit: CommitConfig{
				SignOff: false,
				Verbose: false,
			},
			Merging: MergingConfig{
				ManualCommit: false,
				Args:         "",
			},
			SkipHookPrefix:    "WIP",
			AutoFetch:         true,
			AutoRefresh:       true,
			BranchLogCmd:      "git log --graph --color=always --abbrev-commit --decorate --date=relative --pretty=medium {{branchName}} --",
			AllBranchesLogCmd: "git log --graph --all --color=always --abbrev-commit --decorate --date=relative --pretty=medium",
			OverrideGpg:       false,
			DisableForcePushing: false,
			ParseEmoji:        false,
			Log: GitLogConfig{
				Order:          "topo-order",
				ShowGraph:      "when-maximised",
				ShowWholeGraph: false,
			},
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
		DisableStartupPopups: false,
		NotARepository:       "prompt",
		OS:                   GetPlatformDefaultConfig(),
	}
}
