package game_test

import (
	"testing"

	"github.com/lindsaygelle/nook/game"
)

func TestByKey(t *testing.T) {
	got, ok := game.ByKey(game.NewHorizons.Key)
	if !ok {
		t.Fatalf("game.ByKey(%s) not found", game.NewHorizons.Key)
	}
	if got.Key != game.NewHorizons.Key {
		t.Fatalf("game.ByKey(%s).Key = %s", game.NewHorizons.Key, got.Key)
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
	if games[len(games)-1].Key != game.NewHorizons.Key {
		t.Fatalf("game.List()[last].Key = %s", games[len(games)-1].Key)
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
