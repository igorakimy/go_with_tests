package reading_files_test

import (
	"github.come/igorakimy/go_with_tests/reading_files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker`
	)

	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := reading_files.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], reading_files.Post{
		Title:       "Post 2",
		Description: "Description 2",
		Tags:        []string{"rust", "borrow-checker"},
	})
}

func assertPost(t *testing.T, got, want reading_files.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
