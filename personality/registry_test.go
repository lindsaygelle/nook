package personality_test

import (
	"testing"

	"github.com/lindsaygelle/nook/personality"
)

func TestByKey(t *testing.T) {
	got, ok := personality.ByKey(personality.Snooty.Key)
	if !ok {
		t.Fatalf("personality.ByKey(%s) not found", personality.Snooty.Key)
	}
	if got.Key != personality.Snooty.Key {
		t.Fatalf("personality.ByKey(%s).Key = %s", personality.Snooty.Key, got.Key)
	}
}

func TestByKeyMissing(t *testing.T) {
	if _, ok := personality.ByKey("missing"); ok {
		t.Fatal("personality.ByKey(missing) unexpectedly found a personality")
	}
}

func TestList(t *testing.T) {
	personalities := personality.List()
	if len(personalities) != 8 {
		t.Fatalf("len(personality.List()) = %d", len(personalities))
	}
	if personalities[0].Key != personality.BigSister.Key {
		t.Fatalf("personality.List()[0].Key = %s", personalities[0].Key)
	}
	if personalities[len(personalities)-1].Key != personality.Snooty.Key {
		t.Fatalf("personality.List()[last].Key = %s", personalities[len(personalities)-1].Key)
	}
}

func TestListReturnsCopy(t *testing.T) {
	personalities := personality.List()
	personalities[0] = personality.Snooty

	fresh := personality.List()
	if fresh[0].Key != personality.BigSister.Key {
		t.Fatalf("personality.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
