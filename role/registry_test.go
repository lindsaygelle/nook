package role_test

import (
	"testing"

	"github.com/lindsaygelle/nook/role"
)

func TestByKey(t *testing.T) {
	found, ok := role.ByKey(role.Proprietor.Key)
	if !ok {
		t.Fatalf("role.ByKey(%s) not found", role.Proprietor.Key)
	}
	if found.Key != role.Proprietor.Key {
		t.Fatalf("role.ByKey(%s).Key = %s", role.Proprietor.Key, found.Key)
	}

	if _, ok := role.ByKey("missing"); ok {
		t.Fatal("role.ByKey(missing) unexpectedly found a role")
	}
}

func TestList(t *testing.T) {
	roles := role.List()
	if len(roles) != 6 {
		t.Fatalf("len(role.List()) = %d", len(roles))
	}
	if roles[0].Key != role.Government.Key {
		t.Fatalf("role.List()[0].Key = %s", roles[0].Key)
	}
	if roles[len(roles)-1].Key != role.Unused.Key {
		t.Fatalf("role.List()[last].Key = %s", roles[len(roles)-1].Key)
	}

	roles[0] = role.Unused

	fresh := role.List()
	if fresh[0].Key != role.Government.Key {
		t.Fatalf("role.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
