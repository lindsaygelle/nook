package catalog

import (
	"slices"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

// VillagersByPersonalityKey returns all villagers whose personality exactly
// matches the provided personality key. Results are sorted by animal key and
// then character key for deterministic backend responses.
func VillagersByPersonalityKey(personalityKey nook.Key) []nook.Villager {
	if personalityKey == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			if villager.Personality.Key != personalityKey {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}

// VillagersByPersonality returns all villagers whose localized personality name
// matches the provided value after normalization. Results are sorted by animal
// key and then character key for deterministic backend responses.
func VillagersByPersonality(tag language.Tag, personality string) []nook.Villager {
	query := NormalizeLookupValue(personality)
	if query == "" {
		return nil
	}

	villagers := make([]nook.Villager, 0)
	for _, bucket := range AllVillagers {
		for _, villager := range bucket {
			localizedPersonality, ok := villager.Personality.Name.Get(tag)
			if !ok {
				continue
			}
			if NormalizeLookupValue(localizedPersonality.Value) != query {
				continue
			}
			villagers = append(villagers, villager)
		}
	}

	slices.SortFunc(villagers, compareVillagers)
	return villagers
}
