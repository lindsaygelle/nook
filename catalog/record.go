package catalog

import "github.com/lindsaygelle/nook"

// LocalizedValues is an API-facing representation of localized strings keyed by
// BCP 47 language tag.
type LocalizedValues map[string]string

// BirthdayRecord is a transport-friendly birthday representation.
type BirthdayRecord struct {
	Day   uint8 `json:"day"`
	Month uint8 `json:"month"`
}

// CharacterRecord is a flattened API-facing representation of a character.
type CharacterRecord struct {
	AnimalKey string          `json:"animal_key"`
	Birthday  BirthdayRecord  `json:"birthday"`
	Code      string          `json:"code,omitempty"`
	Gender    LocalizedValues `json:"gender"`
	Key       string          `json:"key"`
	Name      LocalizedValues `json:"name"`
	Special   bool            `json:"special"`
}

// ResidentRecord is an API-facing representation of a resident.
type ResidentRecord struct {
	CharacterRecord
}

// VillagerRecord is an API-facing representation of a villager.
type VillagerRecord struct {
	CharacterRecord
	Personality LocalizedValues `json:"personality"`
	Phrase      LocalizedValues `json:"phrase"`
}

// LocalizedValuesOf converts language-tagged values into a transport-friendly
// map and omits empty values.
func LocalizedValuesOf(values nook.Languages) LocalizedValues {
	out := make(LocalizedValues, len(values))
	for tag, name := range values {
		if !name.Ok() {
			continue
		}
		out[tag.String()] = name.Value
	}
	return out
}

// CharacterRecordOf converts a domain character into its API-facing record.
func CharacterRecordOf(character nook.Character) CharacterRecord {
	return CharacterRecord{
		AnimalKey: string(character.Animal.Key),
		Birthday: BirthdayRecord{
			Day:   character.Birthday.Day,
			Month: uint8(character.Birthday.Month),
		},
		Code:    character.Code.Value,
		Gender:  LocalizedValuesOf(character.Gender.Name),
		Key:     string(character.Key),
		Name:    LocalizedValuesOf(character.Name),
		Special: character.Special,
	}
}

// ResidentRecordOf converts a domain resident into its API-facing record.
func ResidentRecordOf(resident nook.Resident) ResidentRecord {
	return ResidentRecord{
		CharacterRecord: CharacterRecordOf(resident.Character),
	}
}

// VillagerRecordOf converts a domain villager into its API-facing record.
func VillagerRecordOf(villager nook.Villager) VillagerRecord {
	return VillagerRecord{
		CharacterRecord: CharacterRecordOf(villager.Character),
		Personality:     LocalizedValuesOf(villager.Personality.Name),
		Phrase:          LocalizedValuesOf(villager.Phrase),
	}
}
