package nook_test

import (
	"testing"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character/gyroid"
	"github.com/lindsaygelle/nook/character/raccoon"
	"github.com/lindsaygelle/nook/character/squirrel"
	"github.com/lindsaygelle/nook/role"
)

func testResident(t *testing.T, animal nook.Key, r nook.Resident) {
	testResidentRoles(t, r)
	testResidentSpecial(t, r)
}

func testResidentRoles(t *testing.T, r nook.Resident) {
	if len(r.Roles) == 0 {
		t.Fatalf("%s.Roles is empty", r.Key)
	}

	for i, residentRole := range r.Roles {
		if residentRole.Key == "" {
			t.Fatalf("%s.Roles[%d].Key is empty", r.Key, i)
		}
	}
}

func testResidentSpecial(t *testing.T, r nook.Resident) {
	if ok := r.Special; !ok {
		t.Fatalf("%s.Special != true", r.Key)
	}
}

func TestResidentHasRole(t *testing.T) {
	if !raccoon.TomNook.HasRole(role.Proprietor.Key) {
		t.Fatalf("%s.HasRole(%s) = false", raccoon.TomNook.Key, role.Proprietor.Key)
	}
	if gyroid.Lloid.HasRole("") {
		t.Fatalf("%s.HasRole(blank) = true", gyroid.Lloid.Key)
	}
	if !gyroid.Lloid.HasRole(role.Government.Key) {
		t.Fatalf("%s.HasRole(%s) = false", gyroid.Lloid.Key, role.Government.Key)
	}
	if !gyroid.Lloid.HasRole(role.Islander.Key) {
		t.Fatalf("%s.HasRole(%s) = false", gyroid.Lloid.Key, role.Islander.Key)
	}
	if !gyroid.Lloid.HasRole(role.Proprietor.Key) {
		t.Fatalf("%s.HasRole(%s) = false", gyroid.Lloid.Key, role.Proprietor.Key)
	}
	if squirrel.Shaki.HasRole(role.Proprietor.Key) {
		t.Fatalf("%s.HasRole(%s) = true", squirrel.Shaki.Key, role.Proprietor.Key)
	}
}
