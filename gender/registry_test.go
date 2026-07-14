package gender_test

import (
	"testing"

	"github.com/lindsaygelle/nook/gender"
)

func TestByKey(t *testing.T) {
	got, ok := gender.ByKey(gender.Female.Key)
	if !ok {
		t.Fatalf("gender.ByKey(%s) not found", gender.Female.Key)
	}
	if got.Key != gender.Female.Key {
		t.Fatalf("gender.ByKey(%s).Key = %s", gender.Female.Key, got.Key)
	}
}

func TestByKeyMissing(t *testing.T) {
	if _, ok := gender.ByKey("missing"); ok {
		t.Fatal("gender.ByKey(missing) unexpectedly found a gender")
	}
}

func TestList(t *testing.T) {
	genders := gender.List()
	if len(genders) != 2 {
		t.Fatalf("len(gender.List()) = %d", len(genders))
	}
	if genders[0].Key != gender.Female.Key {
		t.Fatalf("gender.List()[0].Key = %s", genders[0].Key)
	}
	if genders[1].Key != gender.Male.Key {
		t.Fatalf("gender.List()[1].Key = %s", genders[1].Key)
	}
}

func TestListReturnsCopy(t *testing.T) {
	genders := gender.List()
	genders[0] = gender.Male

	fresh := gender.List()
	if fresh[0].Key != gender.Female.Key {
		t.Fatalf("gender.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
