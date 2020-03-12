package bitbucketserver

import (
	"github.com/drone/go-scm/scm"
	"net/http"
)

type webhookService struct {
	client *wrapper
}

func (s *webhookService) Parse(req *http.Request, fn scm.SecretFunc) (scm.Webhook, error) {
	return nil, scm.ErrNotSupported
}
