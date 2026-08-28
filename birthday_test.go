package nook_test

import (
	"testing"
	"time"

	"github.com/lindsaygelle/nook"
)

func TestBirthdayZodiacSignKey(t *testing.T) {
	tests := []struct {
		name     string
		birthday nook.Birthday
		want     nook.Key
	}{
		{
			name: "capricorn december boundary",
			birthday: nook.Birthday{
				Day:   22,
				Month: time.December,
			},
			want: "Capricorn",
		},
		{
			name: "aquarius january boundary",
			birthday: nook.Birthday{
				Day:   20,
				Month: time.January,
			},
			want: "Aquarius",
		},
		{
			name: "virgo september final day",
			birthday: nook.Birthday{
				Day:   22,
				Month: time.September,
			},
			want: "Virgo",
		},
		{
			name: "libra september boundary",
			birthday: nook.Birthday{
				Day:   23,
				Month: time.September,
			},
			want: "Libra",
		},
	}

	for _, tt := range tests {
		got, ok := tt.birthday.ZodiacSignKey()
		if !ok {
			t.Fatalf("%s: ZodiacSignKey() not found", tt.name)
		}
		if got != tt.want {
			t.Fatalf("%s: ZodiacSignKey() = %s", tt.name, got)
		}
	}
}

func TestBirthdayZodiacSignKeyInvalidBirthday(t *testing.T) {
	if _, ok := (nook.Birthday{}).ZodiacSignKey(); ok {
		t.Fatal("Birthday{}.ZodiacSignKey() unexpectedly found a zodiac sign")
	}
}
