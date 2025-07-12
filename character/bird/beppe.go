package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// beppeBirthday represents beppe birthday.
	beppeBirthday = nook.Birthday{
		Day:   18,
		Month: time.July}
)

var (
	// beppeCode represents beppe code.
	beppeCode = nook.Code{
		Value: "cwc"}
)

var (
	// beppeAmericanEnglishName represents beppe american english name.
	beppeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Beppe"}

	// beppeCanadianFrenchName represents beppe canadian french name.
	beppeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Choucac"}

	// beppeDutchName represents beppe dutch name.
	beppeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// beppeFrenchName represents beppe french name.
	beppeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Choucac"}

	// beppeGermanName represents beppe german name.
	beppeGermanName = nook.Name{
		Language: language.German,
		Value:    "Schorsch"}

	// beppeItalianName represents beppe italian name.
	beppeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amper"}

	// beppeJapaneseName represents beppe japanese name.
	beppeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピンちゃん"}

	// beppeKoreanName represents beppe korean name.
	beppeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑군"}

	// beppeLatinAmericanSpanishName represents beppe latin american spanish name.
	beppeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Grajardo"}

	// beppeRussianName represents beppe russian name.
	beppeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// beppeSimplifiedChineseName represents beppe simplified chinese name.
	beppeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// beppeSpanishName represents beppe spanish name.
	beppeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Grajardo"}

	// beppeTraditionalChineseName represents beppe traditional chinese name.
	beppeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "車仔"}
)

var (
	// beppeName represents beppe name.
	beppeName = nook.Languages{
		language.AmericanEnglish:      beppeAmericanEnglishName,
		language.CanadianFrench:       beppeCanadianFrenchName,
		language.Dutch:                beppeDutchName,
		language.French:               beppeFrenchName,
		language.German:               beppeGermanName,
		language.Italian:              beppeItalianName,
		language.Japanese:             beppeJapaneseName,
		language.Korean:               beppeKoreanName,
		language.LatinAmericanSpanish: beppeLatinAmericanSpanishName,
		language.Russian:              beppeRussianName,
		language.SimplifiedChinese:    beppeSimplifiedChineseName,
		language.Spanish:              beppeSpanishName,
		language.TraditionalChinese:   beppeTraditionalChineseName}
)

var (
	// beppeCharacter represents beppe character.
	beppeCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: beppeBirthday,
		Code:     beppeCode,
		Key:      character.Beppe,
		Gender:   gender.Male,
		Name:     beppeName,
		Special:  true}
)

var (
	// Beppe represents beppe.
	Beppe = nook.Resident{
		Character: beppeCharacter}
)
