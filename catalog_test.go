package nook_test

import "testing"

func TestCatalogBucketsNonEmpty(t *testing.T) {
	for animal, residents := range residents {
		if len(residents) == 0 {
			t.Fatalf("resident catalog bucket %s is empty", animal)
		}
	}

	for animal, villagers := range villagers {
		if len(villagers) == 0 {
			t.Fatalf("villager catalog bucket %s is empty", animal)
		}
	}
}
