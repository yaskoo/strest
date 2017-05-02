package types

type Response struct {
	Body string
}

// Context hold the play data.
// This includes responses from steps, custom data, etc.
type Context struct {
	// Res holds the reponse data for the current request
	Res *Response

	// Register holds custom data registered from individual steps
	Register map[string]string
}
