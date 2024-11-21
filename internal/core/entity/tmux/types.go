package tmux

const (
	Alias string = "tmux"
)

const (
	SessionExists int = iota
	SessionNotExists
)

type Command string

const (
	EnterCmd        Command = "C-m"
	HasSessionCmd   Command = "has-session"
	NewSessionCmd   Command = "new-session"
	NewWindowCmd    Command = "new-window"
	SelectWindowCmd Command = "select-window"
	AttachCmd       Command = "attach-session"
	SendKeysCmd     Command = "send-keys"
)
