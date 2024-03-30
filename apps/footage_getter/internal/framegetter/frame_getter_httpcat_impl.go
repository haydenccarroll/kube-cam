package framegetter

import (
	"context"
	"io"
	"net/http"
)

// FrameGetterHTTPCatImpl is an implementation of FrameGetter that fetches images over HTTP.
type FrameGetterHTTPCatImpl struct {
	client *http.Client
	url    string
}

// NewFrameGetterHTTPCatImpl creates a new FrameGetterHTTPCatImpl.
func NewFrameGetterHTTPCatImpl(client *http.Client, url string) *FrameGetterHTTPCatImpl {
	return &FrameGetterHTTPCatImpl{
		client: client,
		url:    url,
	}
}

// Get fetches an image (of a cat) at the URL on the FrameGetterHTTPCatImpl
func (f *FrameGetterHTTPCatImpl) Get(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", f.url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
