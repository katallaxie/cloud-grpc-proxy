package gen

// APIs ...
type APIs map[string]API

// Loader ...
type Loader interface {
	// Load ...
	Load([]string) (APIs, error)
}

// API ...
type API interface{}