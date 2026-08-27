package catalog_test

import (
	"testing"

	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/catalog"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/role"
)

func TestResidentsByRole(t *testing.T) {
	residents := catalog.ResidentsByRole(role.Proprietor.Key)
	if len(residents) == 0 {
		t.Fatal("catalog.ResidentsByRole(Proprietor) returned no residents")
	}

	foundTomNook := false
	for i, resident := range residents {
		if !resident.HasRole(role.Proprietor.Key) {
			t.Fatalf("catalog.ResidentsByRole(Proprietor)[%d] missing proprietor role", i)
		}
		if resident.Character.Animal.Key == animal.Raccoon.Key && resident.Character.Key == character.TomNook {
			foundTomNook = true
		}
		if i == 0 {
			continue
		}

		prev := residents[i-1]
		if resident.Character.Animal.Key < prev.Character.Animal.Key {
			t.Fatalf("catalog.ResidentsByRole(Proprietor)[%d] not sorted by animal key", i)
		}
		if resident.Character.Animal.Key == prev.Character.Animal.Key && resident.Character.Key < prev.Character.Key {
			t.Fatalf("catalog.ResidentsByRole(Proprietor)[%d] not sorted by character key", i)
		}
	}
	if !foundTomNook {
		t.Fatal("catalog.ResidentsByRole(Proprietor) missing Tom Nook")
	}
}

func TestResidentsByRoleEmptyKey(t *testing.T) {
	residents := catalog.ResidentsByRole("")
	if residents != nil {
		t.Fatalf("catalog.ResidentsByRole(\"\") = %#v", residents)
	}
}
