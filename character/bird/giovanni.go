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
	// giovanniBirthday represents giovanni birthday.
	giovanniBirthday = nook.Birthday{
		Day:   6,
		Month: time.September}
)

var (
	// giovanniCode represents giovanni code.
	giovanniCode = nook.Code{
		Value: "cwa"}
)

var (
	// giovanniAmericanEnglishName represents giovanni american english name.
	giovanniAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Giovanni"}

	// giovanniCanadianFrenchName represents giovanni canadian french name.
	giovanniCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Corbac"}

	// giovanniDutchName represents giovanni dutch name.
	giovanniDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// giovanniFrenchName represents giovanni french name.
	giovanniFrenchName = nook.Name{
		Language: language.French,
		Value:    "Corbac"}

	// giovanniGermanName represents giovanni german name.
	giovanniGermanName = nook.Name{
		Language: language.German,
		Value:    "Fiete"}

	// giovanniItalianName represents giovanni italian name.
	giovanniItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Corvanni"}

	// giovanniJapaneseName represents giovanni japanese name.
	giovanniJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャンタロー"}

	// giovanniKoreanName represents giovanni korean name.
	giovanniKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캠식"}

	// giovanniLatinAmericanSpanishName represents giovanni latin american spanish name.
	giovanniLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Corvín"}

	// giovanniRussianName represents giovanni russian name.
	giovanniRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// giovanniSimplifiedChineseName represents giovanni simplified chinese name.
	giovanniSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// giovanniSpanishName represents giovanni spanish name.
	giovanniSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Corvín"}

	// giovanniTraditionalChineseName represents giovanni traditional chinese name.
	giovanniTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "露哥"}
)

var (
	// giovanniName represents giovanni name.
	giovanniName = nook.Languages{
		language.AmericanEnglish:      giovanniAmericanEnglishName,
		language.CanadianFrench:       giovanniCanadianFrenchName,
		language.Dutch:                giovanniDutchName,
		language.French:               giovanniFrenchName,
		language.German:               giovanniGermanName,
		language.Italian:              giovanniItalianName,
		language.Japanese:             giovanniJapaneseName,
		language.Korean:               giovanniKoreanName,
		language.LatinAmericanSpanish: giovanniLatinAmericanSpanishName,
		language.Russian:              giovanniRussianName,
		language.SimplifiedChinese:    giovanniSimplifiedChineseName,
		language.Spanish:              giovanniSpanishName,
		language.TraditionalChinese:   giovanniTraditionalChineseName}
)

var (
	// giovanniCharacter represents giovanni character.
	giovanniCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: giovanniBirthday,
		Code:     giovanniCode,
		Key:      character.Giovanni,
		Gender:   gender.Male,
		Name:     giovanniName,
		Special:  true}
)

var (
	// Giovanni represents giovanni.
	Giovanni = nook.Resident{
		Character: giovanniCharacter}
)
