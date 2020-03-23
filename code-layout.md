## Go-scm file tree

This document describes the directory structure of the go-scm repository.

├──transport // Package transport provides facilities for setting up authenticated http.RoundTripper given credentials and base RoundTripper.
├──webhook.go // Defindes all struct of interface and data.
├──token.go
├──git.go
├──org.go
├──user.go
├──review.go
├──client.go
├──util.go
├──repo.go
├──linker.go
├──pr.go
├──content.go
├──const.go
├──issue.go
├──driver // Defindes dirvers of implement scm interface, contain bitbucket cloud, bitbucket server, github, gitlab etc.
├── ├──bitbucketserver // bitbucket server driver.
├── ├── ├──webhook.go
├── ├── ├──testdata
├── ├── ├──git.go
├── ├── ├──org.go
├── ├── ├──user.go
├── ├── ├──review.go
├── ├── ├──bitbucketserver.go
├── ├── ├──repo.go
├── ├── ├──linker.go
├── ├── ├──pr.go
├── ├── ├──content.go
├── ├── ├──issue.go
├── ├──gitlab // gitlab driver.
├── ├── ├──webhook.go
├── ├── ├──testdata
├── ├── ├──git.go
├── ├── ├──org.go
├── ├── ├──user.go
├── ├── ├──review.go
├── ├── ├──util.go
├── ├── ├──gitlab.go
├── ├── ├──repo.go
├── ├── ├──linker.go
├── ├── ├──pr.go
├── ├── ├──content.go
├── ├── ├──issue.go
├── ├──gitea // gitea driver
├── ├── ├──webhook.go
├── ├── ├──testdata
├── ├── ├──git.go
├── ├── ├──org.go
├── ├── ├──user.go
├── ├── ├──review.go
├── ├── ├──repo.go
├── ├── ├──linker.go
├── ├── ├──pr.go
├── ├── ├──content.go
├── ├── ├──gitea.go
├── ├── ├──issue.go
├── ├──github // github driver
├── ├── ├──webhook.go
├── ├── ├──testdata
├── ├── ├──git.go
├── ├── ├──org.go
├── ├── ├──user.go
├── ├── ├──review.go
├── ├── ├──util.go
├── ├── ├──repo.go
├── ├── ├──linker.go
├── ├── ├──pr.go
├── ├── ├──content.go
├── ├── ├──issue.go
├── ├── ├──github.go
├── ├──stash // stash driver
├── ├── ├──webhook.go
├── ├── ├──testdata
├── ├── ├──git.go
├── ├── ├──org.go
├── ├── ├──user.go
├── ├── ├──review.go
├── ├── ├──util_test.go
├── ├── ├──stash.go
├── ├── ├──util.go
├── ├── ├──repo.go
├── ├── ├──linker.go
├── ├── ├──pr.go
├── ├── ├──content.go
├── ├── ├──issue.go
├── ├──bitbucket // bitbucket cloud driver
├── ├── ├──webhook.go
├── ├── ├──testdata
├── ├── ├──git.go
├── ├── ├──org.go
├── ├── ├──user.go
├── ├── ├──bitbucket.go
├── ├── ├──review.go
├── ├── ├──util.go
├── ├── ├──repo.go
├── ├── ├──linker.go
├── ├── ├──pr.go
├── ├── ├──content.go
├── ├── ├──issue.go
├── ├──gogs // gogs driver
├── ├── ├──webhook.go
├── ├── ├──testdata
├── ├── ├──git.go
├── ├── ├──org.go
├── ├── ├──user.go
├── ├── ├──gogs.go
├── ├── ├──review.go
├── ├── ├──repo.go
├── ├── ├──linker.go
├── ├── ├──pr.go
├── ├── ├──content.go
├── ├── ├──webhook_test.go
├── ├── ├──issue.go