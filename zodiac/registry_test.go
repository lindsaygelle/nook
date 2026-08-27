package zodiac_test

import (
	"testing"

	"github.com/lindsaygelle/nook/zodiac"
)

func TestByKey(t *testing.T) {
	found, ok := zodiac.ByKey(zodiac.Virgo.Key)
	if !ok {
		t.Fatalf("zodiac.ByKey(%s) not found", zodiac.Virgo.Key)
	}
	if found.Key != zodiac.Virgo.Key {
		t.Fatalf("zodiac.ByKey(%s).Key = %s", zodiac.Virgo.Key, found.Key)
	}

	if _, ok := zodiac.ByKey("missing"); ok {
		t.Fatal("zodiac.ByKey(missing) unexpectedly found a zodiac sign")
	}
}

func TestList(t *testing.T) {
	zodiacSigns := zodiac.List()
	if len(zodiacSigns) != 12 {
		t.Fatalf("len(zodiac.List()) = %d", len(zodiacSigns))
	}
	if zodiacSigns[0].Key != zodiac.Aquarius.Key {
		t.Fatalf("zodiac.List()[0].Key = %s", zodiacSigns[0].Key)
	}
	if zodiacSigns[len(zodiacSigns)-1].Key != zodiac.Virgo.Key {
		t.Fatalf("zodiac.List()[last].Key = %s", zodiacSigns[len(zodiacSigns)-1].Key)
	}

	zodiacSigns[0] = zodiac.Virgo

	fresh := zodiac.List()
	if fresh[0].Key != zodiac.Aquarius.Key {
		t.Fatalf("zodiac.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}
