package fetch

import (
	"syscall/js"

	"github.com/syumai/workers/internal/jsutil"
)

// Client is an HTTP client.
type Client struct {
	// namespace - Objects that Fetch API belongs to. Default is Global
	namespace js.Value
}

// applyOptions applies client options.
func (c *Client) applyOptions(opts []ClientOption) {
	for _, opt := range opts {
		opt(c)
	}
}

// ClientOption is a type that represents an optional function.
type ClientOption func(*Client)

// WithBinding changes the objects that Fetch API belongs to.
// This is useful for service bindings, mTLS, etc.
func WithBinding(bind js.Value) ClientOption {
	return func(c *Client) {
		c.namespace = bind
	}
}

// NewClient returns new Client
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		namespace: jsutil.Global,
	}
	c.applyOptions(opts)

	return c
}
