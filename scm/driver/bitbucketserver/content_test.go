package bitbucketserver

import (
	"context"
	"encoding/json"
	"github.com/drone/go-scm/scm"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"testing"

	"github.com/h2non/gock"
)

func TestContentFind(t *testing.T) {
	defer gock.Off()

	gock.New("https://bitbucket.server").
		Get("/rest/api/1.0/projects/atlassian/repos/atlaskit/browse/README").
		MatchParam("at", "dev").
		Reply(200).
		Type("application/json").
		File("testdata/content.json")

	client, _ := New("https://bitbucket.server")
	got, _, err := client.Contents.Find(context.Background(), "atlassian/atlaskit", "README", "dev")
	if err != nil {
		t.Error(err)
	}

	want := new(scm.Content)
	raw, _ := ioutil.ReadFile("testdata/content.json.golden")
	json.Unmarshal(raw, want)

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}

func TestContentCreate(t *testing.T) {
	defer gock.Off()

	gock.New("https://bitbucket.server").
		Put("/rest/api/1.0/projects/atlassian/repos/atlaskit/browse/README").
		Reply(200)

	params := &scm.ContentParams{
		Message: "test",
		Branch: "master",
		Signature: scm.Signature{
			Name:  "Monalisa Octocat",
		},
		Data: []byte("create data"),
	}

	client, _ := New("https://bitbucket.server")
	res, err := client.Contents.Create(context.Background(), "atlassian/atlaskit", "README", params)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Status != 200 {
		t.Errorf("Unexpected Results")
	}
}
