package game_test

import (
	"testing"

	"github.com/lindsaygelle/nook/game"
	"github.com/lindsaygelle/nook/region"
)

func TestByKey(t *testing.T) {
	got, ok := game.ByKey(game.NewHorizons.Key)
	if !ok {
		t.Fatalf("game.ByKey(%s) not found", game.NewHorizons.Key)
	}
	if got.Key != game.NewHorizons.Key {
		t.Fatalf("game.ByKey(%s).Key = %s", game.NewHorizons.Key, got.Key)
	}
	if got.ReleaseOrder != 12 {
		t.Fatalf("game.ByKey(%s).ReleaseOrder = %d", game.NewHorizons.Key, got.ReleaseOrder)
	}
	if len(got.ReleaseDates) != 1 {
		t.Fatalf("len(game.ByKey(%s).ReleaseDates) = %d", game.NewHorizons.Key, len(got.ReleaseDates))
	}
}

func TestByKeyMissing(t *testing.T) {
	if _, ok := game.ByKey("missing"); ok {
		t.Fatal("game.ByKey(missing) unexpectedly found a game")
	}
}

func TestList(t *testing.T) {
	games := game.List()
	if len(games) != 12 {
		t.Fatalf("len(game.List()) = %d", len(games))
	}
	if games[0].Key != game.DoubutsuNoMori.Key {
		t.Fatalf("game.List()[0].Key = %s", games[0].Key)
	}
	if games[0].ReleaseOrder != 1 {
		t.Fatalf("game.List()[0].ReleaseOrder = %d", games[0].ReleaseOrder)
	}
	if games[len(games)-1].Key != game.NewHorizons.Key {
		t.Fatalf("game.List()[last].Key = %s", games[len(games)-1].Key)
	}
	if games[len(games)-1].ReleaseOrder != 12 {
		t.Fatalf("game.List()[last].ReleaseOrder = %d", games[len(games)-1].ReleaseOrder)
	}

	for i := 1; i < len(games); i++ {
		if games[i-1].ReleaseOrder >= games[i].ReleaseOrder {
			t.Fatalf("game.List() release order not increasing at %s >= %s", games[i-1].Key, games[i].Key)
		}
	}
}

func TestListReturnsCopy(t *testing.T) {
	games := game.List()
	games[0] = game.NewHorizons

	fresh := game.List()
	if fresh[0].Key != game.DoubutsuNoMori.Key {
		t.Fatalf("game.List()[0].Key after mutation = %s", fresh[0].Key)
	}
}

func TestFirstAndLastReleaseDate(t *testing.T) {
	first, ok := game.NewLeaf.FirstReleaseDate()
	if !ok {
		t.Fatal("game.NewLeaf.FirstReleaseDate() not found")
	}
	if first.Region.Key != region.Japan.Key {
		t.Fatalf("game.NewLeaf.FirstReleaseDate().Region.Key = %s", first.Region.Key)
	}
	if first.Year != 2012 || first.Month != 11 || first.Day != 8 {
		t.Fatalf("game.NewLeaf.FirstReleaseDate() = %#v", first)
	}

	last, ok := game.NewLeaf.LastReleaseDate()
	if !ok {
		t.Fatal("game.NewLeaf.LastReleaseDate() not found")
	}
	if last.Region.Key != region.Australia.Key {
		t.Fatalf("game.NewLeaf.LastReleaseDate().Region.Key = %s", last.Region.Key)
	}
	if last.Year != 2013 || last.Month != 6 || last.Day != 15 {
		t.Fatalf("game.NewLeaf.LastReleaseDate() = %#v", last)
	}
}

func TestReleaseDateByRegion(t *testing.T) {
	releaseDate, ok := game.WildWorld.ReleaseDateByRegion(region.Europe.Key)
	if !ok {
		t.Fatalf("game.WildWorld.ReleaseDateByRegion(%s) not found", region.Europe.Key)
	}
	if releaseDate.Year != 2006 || releaseDate.Month != 3 || releaseDate.Day != 31 {
		t.Fatalf("game.WildWorld.ReleaseDateByRegion(%s) = %#v", region.Europe.Key, releaseDate)
	}
}

func TestReleaseDateByRegionMissing(t *testing.T) {
	if _, ok := game.DoubutsuNoMori.ReleaseDateByRegion(region.Europe.Key); ok {
		t.Fatalf("game.DoubutsuNoMori.ReleaseDateByRegion(%s) unexpectedly found a release date", region.Europe.Key)
	}
}
