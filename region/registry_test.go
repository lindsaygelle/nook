package region_test

import (
	"testing"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
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

func TestRegionLocalizedNames(t *testing.T) {
	tests := []struct {
		name   string
		region nook.Region
		tag    language.Tag
		want   string
	}{
		{
			name:   "australia traditional chinese",
			region: region.Australia,
			tag:    language.TraditionalChinese,
			want:   "澳洲",
		},
		{
			name:   "japan italian",
			region: region.Japan,
			tag:    language.Italian,
			want:   "Giappone",
		},
		{
			name:   "north america japanese",
			region: region.NorthAmerica,
			tag:    language.Japanese,
			want:   "北アメリカ大陸",
		},
		{
			name:   "worldwide canadian french",
			region: region.Worldwide,
			tag:    language.CanadianFrench,
			want:   "Monde",
		},
	}

	for _, tt := range tests {
		got, ok := tt.region.Name.Get(tt.tag)
		if !ok {
			t.Fatalf("%s: %s.Name.Get(%s) not found", tt.name, tt.region.Key, tt.tag)
		}
		if got.Value != tt.want {
			t.Fatalf("%s: %s.Name.Get(%s) = %s", tt.name, tt.region.Key, tt.tag, got.Value)
		}
	}
}

func TestRegionLocalizedNamesComplete(t *testing.T) {
	for _, regionValue := range region.List() {
		if !regionValue.Name.Complete() {
			t.Fatalf("%s.Name.Complete() = false; missing %v", regionValue.Key, regionValue.Name.MissingSupportedTags())
		}
	}
}
