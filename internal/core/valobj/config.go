package valobj

type Config struct {
	Session    string
	StartIndex int
	Windows    []Window
}

type Window struct {
	Name string
	Cmd  []string
}
