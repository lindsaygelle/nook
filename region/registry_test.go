package region_test

import (
	"testing"

	"github.com/lindsaygelle/nook/region"
)

func TestByKey(t *testing.T) {
	got, ok := region.ByKey(region.NorthAmerica.Key)
	if !ok {
		t.Fatalf("region.ByKey(%s) not found", region.NorthAmerica.Key)
	}
	if got.Key != region.NorthAmerica.Key {
		t.Fatalf("region.ByKey(%s).Key = %s", region.NorthAmerica.Key, got.Key)
	}
}

func TestByKeyMissing(t *testing.T) {
	if _, ok := region.ByKey("missing"); ok {
		t.Fatal("region.ByKey(missing) unexpectedly found a region")
	}
}

func TestList(t *testing.T) {
	regions := region.List()
	if len(regions) != 6 {
		t.Fatalf("len(region.List()) = %d", len(regions))
	}
	if regions[0].Key != region.Australia.Key {
		t.Fatalf("region.List()[0].Key = %s", regions[0].Key)
	}
	if regions[len(regions)-1].Key != region.Worldwide.Key {
		t.Fatalf("region.List()[last].Key = %s", regions[len(regions)-1].Key)
	}
}

func TestListReturnsCopy(t *testing.T) {
	regions := region.List()
	regions[0] = region.Worldwide

	fresh := region.List()
	if fresh[0].Key != region.Australia.Key {
		t.Fatalf("region.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
