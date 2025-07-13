package pelican

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// peteBirthday represents pete birthday.
	peteBirthday = nook.Birthday{
		Day:   8,
		Month: time.March}
)

var (
	// peteCode represents pete code.
	peteCode = nook.Code{
		Value: "plb/plo"}
)

var (
	// peteAmericanEnglishName represents pete american english name.
	peteAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pete"}

	// peteCanadianFrenchName represents pete canadian french name.
	peteCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Antoine"}

	// peteDutchName represents pete dutch name.
	peteDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pete"}

	// peteFrenchName represents pete french name.
	peteFrenchName = nook.Name{
		Language: language.French,
		Value:    "Antoine"}

	// peteGermanName represents pete german name.
	peteGermanName = nook.Name{
		Language: language.German,
		Value:    "Peter"}

	// peteItalianName represents pete italian name.
	peteItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tino"}

	// peteJapaneseName represents pete japanese name.
	peteJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺりお"}

	// peteKoreanName represents pete korean name.
	peteKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펠리오"}

	// peteLatinAmericanSpanishName represents pete latin american spanish name.
	peteLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carturo"}

	// peteRussianName represents pete russian name.
	peteRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пит"}

	// peteSimplifiedChineseName represents pete simplified chinese name.
	peteSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "程信雄"}

	// peteSpanishName represents pete spanish name.
	peteSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carturo"}

	// peteTraditionalChineseName represents pete traditional chinese name.
	peteTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "程信雄"}
)

var (
	// peteName represents pete name.
	peteName = nook.Languages{
		language.AmericanEnglish:      peteAmericanEnglishName,
		language.CanadianFrench:       peteCanadianFrenchName,
		language.Dutch:                peteDutchName,
		language.French:               peteFrenchName,
		language.German:               peteGermanName,
		language.Italian:              peteItalianName,
		language.Japanese:             peteJapaneseName,
		language.Korean:               peteKoreanName,
		language.LatinAmericanSpanish: peteLatinAmericanSpanishName,
		language.Russian:              peteRussianName,
		language.SimplifiedChinese:    peteSimplifiedChineseName,
		language.Spanish:              peteSpanishName,
		language.TraditionalChinese:   peteTraditionalChineseName}
)

var (
	// peteCharacter represents pete character.
	peteCharacter = nook.Character{
		Animal:   animal.Pelican,
		Birthday: peteBirthday,
		Code:     peteCode,
		Key:      character.Pete,
		Gender:   gender.Male,
		Name:     peteName,
		Special:  true}
)

var (
	// Pete represents pete.
	Pete = nook.Resident{
		Character: peteCharacter}
)
