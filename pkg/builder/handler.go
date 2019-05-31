package builder

func NewHandler() Handler {
	return &handler{
		downloader:   NewDownloader(),
		imageBuilder: NewImageBuilder(),
	}
}

type Handler interface {
	Handle(request *BuildRequest) error
}

type handler struct {
	downloader   Downloader
	imageBuilder ImageBuilder
}

func (h *handler) Handle(request *BuildRequest) error {
	// download yaml
	h.downloader.GetManifests(request)

	// parse yaml for additional images

	// push yaml to temp directory

	// create dockerfile pointing to the yaml

	// docker build
	h.imageBuilder.Build(request.Image)

	// pull additional images to local registry

	return nil
}
