package nook_test

import (
	"testing"
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
)

func TestReleaseDateCompare(t *testing.T) {
	earlier := nook.ReleaseDate{
		Day:    8,
		Month:  time.November,
		Region: region.Japan,
		Year:   2012,
	}
	later := nook.ReleaseDate{
		Day:    9,
		Month:  time.June,
		Region: region.NorthAmerica,
		Year:   2013,
	}

	if got := earlier.Compare(later); got >= 0 {
		t.Fatalf("earlier.Compare(later) = %d", got)
	}
	if got := later.Compare(earlier); got <= 0 {
		t.Fatalf("later.Compare(earlier) = %d", got)
	}
	if got := earlier.Compare(earlier); got != 0 {
		t.Fatalf("earlier.Compare(earlier) = %d", got)
	}
}

func TestReleaseDateBeforeAndAfter(t *testing.T) {
	first := nook.ReleaseDate{
		Day:    21,
		Month:  time.November,
		Region: region.Worldwide,
		Year:   2017,
	}
	second := nook.ReleaseDate{
		Day:    20,
		Month:  time.March,
		Region: region.Worldwide,
		Year:   2020,
	}

	if !first.Before(second) {
		t.Fatal("first.Before(second) = false")
	}
	if first.After(second) {
		t.Fatal("first.After(second) = true")
	}
	if !second.After(first) {
		t.Fatal("second.After(first) = false")
	}
	if second.Before(first) {
		t.Fatal("second.Before(first) = true")
	}
}
