package storage_test

import (
	"context"
	"testing"

	"reflect"

	"github.com/pokstad/poki"
	"github.com/pokstad/poki/storage"
	"github.com/pokstad/poki/storage/memory"
)

var mdP1 = poki.Post{
	Raw: []byte(`# This is a title
And this is the start of a paragraph.

And this is the second paragraph.

- This is a list.
	- This is indented once
		This is indented more`),
	Meta: poki.MetaData{
		Title: "mdP1",
	},
}

var mdP2 = poki.Post{
	Raw: []byte(`This is a simple one paragraph.`),
	Meta: poki.MetaData{
		Title: "mdP1",
	},
}

func TestMemoryStorage(t *testing.T) {
	ctx := context.Background()

	tableTests := []struct {
		Storage storage.Storage
	}{
		{
			Storage: memory.NewMemoryStorage(),
		},
	}

	for _, row := range tableTests {

		// creating new post should return a rev with non zero string
		rev, err := row.Storage.Create(ctx, mdP1)
		if err != nil {
			t.Fatalf("Could not create new post from %+v: %s", mdP1, err)
		}
		if rev.RevisionID == "" {
			t.Fatalf("Revision ID for new post must be non-zero value string")
		}

		// trying to create something at the same path should cause
		// an error:
		_, err = row.Storage.Create(ctx, mdP1)
		if err == nil {
			t.Fatalf("Expected an error from creating the same post twice")
		}

		// try to read that post
		rev, err = row.Storage.Read(ctx, mdP1.Path)
		if err != nil {
			t.Fatalf("Could not read post from path %s", mdP1.Path)
		}

		if !reflect.DeepEqual(rev.Post, mdP1) {
			t.Fatalf("Retrieved post and reference post do not match: %+v %+v",
				rev.Post, mdP1)
		}

		// try to update post revision with new post version
		newRev, err := row.Storage.Update(ctx, mdP2, rev)
		if newRev.RevisionID == rev.RevisionID {
			t.Fatalf("New revision ID is the same as before: %+v vs %+v",
				newRev, rev)
		}
	}
}
