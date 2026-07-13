package catalog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

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
	ID        string          `json:"id"`
	Gender    LocalizedValues `json:"gender"`
	GenderKey string          `json:"gender_key"`
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
	PersonalityKey string          `json:"personality_key"`
	Personality    LocalizedValues `json:"personality"`
	Phrase         LocalizedValues `json:"phrase"`
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
		Code:      character.Code.Value,
		ID:        string(character.ID()),
		Gender:    LocalizedValuesOf(character.Gender.Name),
		GenderKey: string(character.Gender.Key),
		Key:       string(character.Key),
		Name:      LocalizedValuesOf(character.Name),
		Special:   character.Special,
	}
}

// ResidentRecordOf converts a domain resident into its API-facing record.
func ResidentRecordOf(resident nook.Resident) ResidentRecord {
	return ResidentRecord{
		CharacterRecord: CharacterRecordOf(resident.Character),
	}
}

// ResidentRecordByCode returns a transport-friendly resident record using an
// exact code match after normalization.
func ResidentRecordByCode(code string) (ResidentRecord, bool) {
	resident, ok := ResidentByCode(code)
	if !ok {
		return ResidentRecord{}, false
	}
	return ResidentRecordOf(resident), true
}

// ResidentRecordByID returns a transport-friendly resident record using an
// exact global character identifier match after normalization.
func ResidentRecordByID(id string) (ResidentRecord, bool) {
	resident, ok := ResidentByID(id)
	if !ok {
		return ResidentRecord{}, false
	}
	return ResidentRecordOf(resident), true
}

// ResidentRecordsByAnimal returns a resident bucket as deterministically sorted
// transport-friendly records ordered by animal key and then character key.
func ResidentRecordsByAnimal(animalKey nook.Key) ([]ResidentRecord, bool) {
	residents, ok := ResidentListByAnimal(animalKey)
	if !ok {
		return nil, false
	}

	records := make([]ResidentRecord, 0, len(residents))
	for _, resident := range residents {
		records = append(records, ResidentRecordOf(resident))
	}
	return records, true
}

// ResidentRecordsByBirthday returns all residents whose birthday exactly
// matches the provided month and day. Results are ordered by animal key and
// then character key.
func ResidentRecordsByBirthday(month time.Month, day uint8) []ResidentRecord {
	residents := ResidentsByBirthday(month, day)
	records := make([]ResidentRecord, 0, len(residents))
	for _, resident := range residents {
		records = append(records, ResidentRecordOf(resident))
	}
	return records
}

// ResidentRecordsByBirthMonth returns all residents whose birthday month
// exactly matches the provided month. Results are ordered by animal key and
// then character key.
func ResidentRecordsByBirthMonth(month time.Month) []ResidentRecord {
	residents := ResidentsByBirthMonth(month)
	records := make([]ResidentRecord, 0, len(residents))
	for _, resident := range residents {
		records = append(records, ResidentRecordOf(resident))
	}
	return records
}

// VillagerRecordOf converts a domain villager into its API-facing record.
func VillagerRecordOf(villager nook.Villager) VillagerRecord {
	return VillagerRecord{
		CharacterRecord: CharacterRecordOf(villager.Character),
		PersonalityKey:  string(villager.Personality.Key),
		Personality:     LocalizedValuesOf(villager.Personality.Name),
		Phrase:          LocalizedValuesOf(villager.Phrase),
	}
}

// VillagerRecordByCode returns a transport-friendly villager record using an
// exact code match after normalization.
func VillagerRecordByCode(code string) (VillagerRecord, bool) {
	villager, ok := VillagerByCode(code)
	if !ok {
		return VillagerRecord{}, false
	}
	return VillagerRecordOf(villager), true
}

// VillagerRecordByID returns a transport-friendly villager record using an
// exact global character identifier match after normalization.
func VillagerRecordByID(id string) (VillagerRecord, bool) {
	villager, ok := VillagerByID(id)
	if !ok {
		return VillagerRecord{}, false
	}
	return VillagerRecordOf(villager), true
}

// VillagerRecordsByAnimal returns a villager bucket as deterministically
// sorted transport-friendly records ordered by animal key and then character
// key.
func VillagerRecordsByAnimal(animalKey nook.Key) ([]VillagerRecord, bool) {
	villagers, ok := VillagerListByAnimal(animalKey)
	if !ok {
		return nil, false
	}

	records := make([]VillagerRecord, 0, len(villagers))
	for _, villager := range villagers {
		records = append(records, VillagerRecordOf(villager))
	}
	return records, true
}

// VillagerRecordsByBirthday returns all villagers whose birthday exactly
// matches the provided month and day. Results are ordered by animal key and
// then character key.
func VillagerRecordsByBirthday(month time.Month, day uint8) []VillagerRecord {
	villagers := VillagersByBirthday(month, day)
	records := make([]VillagerRecord, 0, len(villagers))
	for _, villager := range villagers {
		records = append(records, VillagerRecordOf(villager))
	}
	return records
}

// VillagerRecordsByPersonality returns all villagers whose localized
// personality name matches the provided value after normalization. Results are
// ordered by animal key and then character key.
func VillagerRecordsByPersonality(tag language.Tag, personality string) []VillagerRecord {
	villagers := VillagersByPersonality(tag, personality)
	records := make([]VillagerRecord, 0, len(villagers))
	for _, villager := range villagers {
		records = append(records, VillagerRecordOf(villager))
	}
	return records
}

// VillagerRecordsByBirthMonth returns all villagers whose birthday month
// exactly matches the provided month. Results are ordered by animal key and
// then character key.
func VillagerRecordsByBirthMonth(month time.Month) []VillagerRecord {
	villagers := VillagersByBirthMonth(month)
	records := make([]VillagerRecord, 0, len(villagers))
	for _, villager := range villagers {
		records = append(records, VillagerRecordOf(villager))
	}
	return records
}
