package bitbucketserver

import (
	"context"
	"github.com/drone/go-scm/scm"
)

type linker struct {
	base string
}

// Resource returns a link to the resource.
func (l *linker) Resource(ctx context.Context, repo string, ref scm.Reference) (string, error) {
	return "", scm.ErrNotSupported
}

// Diff returns a link to the diff.
func (l *linker) Diff(ctx context.Context, repo string, source, target scm.Reference) (string, error) {
	return "", scm.ErrNotSupported
}