package poki_test

import (
	"context"
	"testing"

	"github.com/pokstad/poki"
)

var mdP1 = poki.Post{
	Raw: []byte(``),
	Meta: poki.MetaData{
		Title: "mdP1",
	},
}

func TestMemoryStorage(t *testing.T) {
	ctx := context.Background()
	var sto poki.Storage = poki.NewMemoryStorage()

	rev, err := sto.Create(ctx, mdP1)
	if err != nil {
		t.Fatalf("Could not create new post from %+v: %s", mdP1, err)
	}
	if rev.RevisionID == "" {
		t.Fatalf("Revision ID for new post must be non-zero value string")
	}
}
