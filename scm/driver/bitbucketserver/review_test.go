package bitbucketserver

import (
	"context"
	"github.com/drone/go-scm/scm"
	"testing"
)

func TestReviewFind(t *testing.T) {
	_, _, err := NewDefault().Reviews.Find(context.Background(), "", 0, 0)
	if err != scm.ErrNotSupported {
		t.Errorf("Expect Not Supported error")
	}
}

func TestReviewList(t *testing.T) {
	_, _, err := NewDefault().Reviews.List(context.Background(), "", 0, scm.ListOptions{})
	if err != scm.ErrNotSupported {
		t.Errorf("Expect Not Supported error")
	}
}

func TestReviewCreate(t *testing.T) {
	_, _, err := NewDefault().Reviews.Create(context.Background(), "", 0, &scm.ReviewInput{})
	if err != scm.ErrNotSupported {
		t.Errorf("Expect Not Supported error")
	}
}

func TestReviewDelete(t *testing.T) {
	_, err := NewDefault().Reviews.Delete(context.Background(), "", 0, 0)
	if err != scm.ErrNotSupported {
		t.Errorf("Expect Not Supported error")
	}
}
