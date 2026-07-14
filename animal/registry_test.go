package animal_test

import (
	"testing"

	"github.com/lindsaygelle/nook/animal"
)

func TestByKey(t *testing.T) {
	got, ok := animal.ByKey(animal.FrillneckedLizard.Key)
	if !ok {
		t.Fatalf("animal.ByKey(%s) not found", animal.FrillneckedLizard.Key)
	}
	if got.Key != animal.FrillneckedLizard.Key {
		t.Fatalf("animal.ByKey(%s).Key = %s", animal.FrillneckedLizard.Key, got.Key)
	}
}

func TestByKeyMissing(t *testing.T) {
	if _, ok := animal.ByKey("missing"); ok {
		t.Fatal("animal.ByKey(missing) unexpectedly found an animal")
	}
}

func TestList(t *testing.T) {
	animals := animal.List()
	if len(animals) != 65 {
		t.Fatalf("len(animal.List()) = %d", len(animals))
	}
	if animals[0].Key != animal.Alligator.Key {
		t.Fatalf("animal.List()[0].Key = %s", animals[0].Key)
	}
	if animals[len(animals)-1].Key != animal.Wolf.Key {
		t.Fatalf("animal.List()[last].Key = %s", animals[len(animals)-1].Key)
	}
}

func TestListReturnsCopy(t *testing.T) {
	animals := animal.List()
	animals[0] = animal.Wolf

	fresh := animal.List()
	if fresh[0].Key != animal.Alligator.Key {
		t.Fatalf("animal.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
