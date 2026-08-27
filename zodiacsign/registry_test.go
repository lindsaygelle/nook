package zodiacsign_test

import (
	"testing"

	"github.com/lindsaygelle/nook/zodiacsign"
)

func TestByKey(t *testing.T) {
	found, ok := zodiacsign.ByKey(zodiacsign.Virgo.Key)
	if !ok {
		t.Fatalf("zodiacsign.ByKey(%s) not found", zodiacsign.Virgo.Key)
	}
	if found.Key != zodiacsign.Virgo.Key {
		t.Fatalf("zodiacsign.ByKey(%s).Key = %s", zodiacsign.Virgo.Key, found.Key)
	}

	if _, ok := zodiacsign.ByKey("missing"); ok {
		t.Fatal("zodiacsign.ByKey(missing) unexpectedly found a zodiac sign")
	}
}

func TestList(t *testing.T) {
	zodiacSigns := zodiacsign.List()
	if len(zodiacSigns) != 12 {
		t.Fatalf("len(zodiacsign.List()) = %d", len(zodiacSigns))
	}
	if zodiacSigns[0].Key != zodiacsign.Aquarius.Key {
		t.Fatalf("zodiacsign.List()[0].Key = %s", zodiacSigns[0].Key)
	}
	if zodiacSigns[len(zodiacSigns)-1].Key != zodiacsign.Virgo.Key {
		t.Fatalf("zodiacsign.List()[last].Key = %s", zodiacSigns[len(zodiacSigns)-1].Key)
	}

	zodiacSigns[0] = zodiacsign.Virgo

	fresh := zodiacsign.List()
	if fresh[0].Key != zodiacsign.Aquarius.Key {
		t.Fatalf("zodiacsign.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
