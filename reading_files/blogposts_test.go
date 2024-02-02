package reading_files_test

import (
	"github.come/igorakimy/go_with_tests/reading_files"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := reading_files.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
