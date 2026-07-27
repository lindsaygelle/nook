package gamecategory_test

import (
	"testing"

	"github.com/lindsaygelle/nook/gamecategory"
)

func TestByKey(t *testing.T) {
	got, ok := gamecategory.ByKey(gamecategory.Mainline.Key)
	if !ok {
		t.Fatalf("gamecategory.ByKey(%s) not found", gamecategory.Mainline.Key)
	}
	if got.Key != gamecategory.Mainline.Key {
		t.Fatalf("gamecategory.ByKey(%s).Key = %s", gamecategory.Mainline.Key, got.Key)
	}
}

func TestByKeyMissing(t *testing.T) {
	if _, ok := gamecategory.ByKey("missing"); ok {
		t.Fatal("gamecategory.ByKey(missing) unexpectedly found a category")
	}
}

func TestList(t *testing.T) {
	categories := gamecategory.List()
	if len(categories) != 3 {
		t.Fatalf("len(gamecategory.List()) = %d", len(categories))
	}
	if categories[0].Key != gamecategory.Mainline.Key {
		t.Fatalf("gamecategory.List()[0].Key = %s", categories[0].Key)
	}
	if categories[len(categories)-1].Key != gamecategory.Spinoff.Key {
		t.Fatalf("gamecategory.List()[last].Key = %s", categories[len(categories)-1].Key)
	}
}

func TestListReturnsCopy(t *testing.T) {
	categories := gamecategory.List()
	categories[0] = gamecategory.Spinoff

	fresh := gamecategory.List()
	if fresh[0].Key != gamecategory.Mainline.Key {
		t.Fatalf("gamecategory.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
