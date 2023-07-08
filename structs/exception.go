package structs

type Exception struct {
	Err      error
	Message  string
	ExitCode int
}
