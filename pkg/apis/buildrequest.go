package apis

// BuildRequest is a struct to describe the API used by
// the command line package to make requests to the builder
// handler.
type BuildRequest struct {
	AuthorizationToken string
	Endpoint           string
	Namespace          string
	Image              string
	ImageBuildArgs     string
}
