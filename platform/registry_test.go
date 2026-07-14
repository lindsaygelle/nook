package platform_test

import (
	"testing"

	"github.com/lindsaygelle/nook/platform"
)

func TestByKey(t *testing.T) {
	got, ok := platform.ByKey(platform.NintendoSwitch.Key)
	if !ok {
		t.Fatalf("platform.ByKey(%s) not found", platform.NintendoSwitch.Key)
	}
	if got.Key != platform.NintendoSwitch.Key {
		t.Fatalf("platform.ByKey(%s).Key = %s", platform.NintendoSwitch.Key, got.Key)
	}
}

func TestByKeyMissing(t *testing.T) {
	if _, ok := platform.ByKey("missing"); ok {
		t.Fatal("platform.ByKey(missing) unexpectedly found a platform")
	}
}

func TestList(t *testing.T) {
	platforms := platform.List()
	if len(platforms) != 10 {
		t.Fatalf("len(platform.List()) = %d", len(platforms))
	}
	if platforms[0].Key != platform.Android.Key {
		t.Fatalf("platform.List()[0].Key = %s", platforms[0].Key)
	}
	if platforms[len(platforms)-1].Key != platform.WiiU.Key {
		t.Fatalf("platform.List()[last].Key = %s", platforms[len(platforms)-1].Key)
	}
}

func TestListReturnsCopy(t *testing.T) {
	platforms := platform.List()
	platforms[0] = platform.WiiU

	fresh := platform.List()
	if fresh[0].Key != platform.Android.Key {
		t.Fatalf("platform.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
