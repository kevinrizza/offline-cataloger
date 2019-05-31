package builder

type BuildRequest struct {
	AuthorizationToken string
	Endpoint           string
	Namespace          string
	Image              string
}
