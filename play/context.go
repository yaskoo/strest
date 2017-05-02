package types

type Response struct {
	Body string
}

// Context hold the play data.
// This includes responses from steps, custom data, etc.
type Context struct {
	// Res holds the reponse data for the current request
	Res *Response

	// Register is string/string map holding custom data
	Register map[string]string
}
