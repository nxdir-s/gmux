package tmux

const (
	SessionExists int = iota
	SessionNotExists
)

const Alias string = "tmux"

type Command string

const (
	HasSessionCmd   Command = "has-session"
	NewSessionCmd   Command = "new-session"
	NewWindowCmd    Command = "new-window"
	SelectWindowCmd Command = "select-window"
	AttachCmd       Command = "attach-session"
	SendKeysCmd     Command = "send-keys"
)

type WindowName string

const (
	EditorWindow   WindowName = "editor"
	DockerWindow   WindowName = "docker"
	DatabaseWindow WindowName = "database"
)

type Window int

const (
	_ Window = iota
	Editor
	Docker
	Database
)

func (w Window) Name() string {
	switch w {
	case Docker:
		return string(DockerWindow)
	case Database:
		return string(DatabaseWindow)
	default:
		return ""
	}
}
