package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// stuBirthday represents stu birthday.
	stuBirthday = nook.Birthday{
		Day:   20,
		Month: time.April}
)

var (
	// stuCode represents stu code.
	stuCode = nook.Code{
		Value: "bul03"}
)

var (
	// stuAmericanEnglishName represents stu american english name.
	stuAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Stu"}

	// stuCanadianFrenchName represents stu canadian french name.
	stuCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Beubeu"}

	// stuDutchName represents stu dutch name.
	stuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stu"}

	// stuFrenchName represents stu french name.
	stuFrenchName = nook.Name{
		Language: language.French,
		Value:    "Beubeu"}

	// stuGermanName represents stu german name.
	stuGermanName = nook.Name{
		Language: language.German,
		Value:    "Carlos"}

	// stuItalianName represents stu italian name.
	stuItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carlos"}

	// stuJapaneseName represents stu japanese name.
	stuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モーリス"}

	// stuKoreanName represents stu korean name.
	stuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "모리스"}

	// stuLatinAmericanSpanishName represents stu latin american spanish name.
	stuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vitorino"}

	// stuRussianName represents stu russian name.
	stuRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стью"}

	// stuSimplifiedChineseName represents stu simplified chinese name.
	stuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛乐时"}

	// stuSpanishName represents stu spanish name.
	stuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vitorino"}

	// stuTraditionalChineseName represents stu traditional chinese name.
	stuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛樂時"}
)

var (
	// stuName represents stu name.
	stuName = nook.Languages{
		language.AmericanEnglish:      stuAmericanEnglishName,
		language.CanadianFrench:       stuCanadianFrenchName,
		language.Dutch:                stuDutchName,
		language.French:               stuFrenchName,
		language.German:               stuGermanName,
		language.Italian:              stuItalianName,
		language.Japanese:             stuJapaneseName,
		language.Korean:               stuKoreanName,
		language.LatinAmericanSpanish: stuLatinAmericanSpanishName,
		language.Russian:              stuRussianName,
		language.SimplifiedChinese:    stuSimplifiedChineseName,
		language.Spanish:              stuSpanishName,
		language.TraditionalChinese:   stuTraditionalChineseName}
)

var (
	// stuCharacter represents stu character.
	stuCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: stuBirthday,
		Code:     stuCode,
		Key:      character.Stu,
		Gender:   gender.Male,
		Name:     stuName,
		Special:  false}
)

var (
	// stuAmericanEnglishPhrase represents stu american english phrase.
	stuAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "moo-dude"}

	// stuCanadianFrenchPhrase represents stu canadian french phrase.
	stuCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meuh quoi"}

	// stuDutchPhrase represents stu dutch phrase.
	stuDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boe-jend"}

	// stuFrenchPhrase represents stu french phrase.
	stuFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh quoi"}

	// stuGermanPhrase represents stu german phrase.
	stuGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muhuhu"}

	// stuItalianPhrase represents stu italian phrase.
	stuItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciaumù"}

	// stuJapanesePhrase represents stu japanese phrase.
	stuJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だなっす"}

	// stuKoreanPhrase represents stu korean phrase.
	stuKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알것소"}

	// stuLatinAmericanSpanishPhrase represents stu latin american spanish phrase.
	stuLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "olé y olé"}

	// stuRussianPhrase represents stu russian phrase.
	stuRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "му-у"}

	// stuSimplifiedChinesePhrase represents stu simplified chinese phrase.
	stuSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "奔驰"}

	// stuSpanishPhrase represents stu spanish phrase.
	stuSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "olé y olé"}

	// stuTraditionalChinesePhrase represents stu traditional chinese phrase.
	stuTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "奔馳"}
)

var (
	// stuPhrase represents stu phrase.
	stuPhrase = nook.Languages{
		language.AmericanEnglish:      stuAmericanEnglishPhrase,
		language.CanadianFrench:       stuCanadianFrenchPhrase,
		language.Dutch:                stuDutchPhrase,
		language.French:               stuFrenchPhrase,
		language.German:               stuGermanPhrase,
		language.Italian:              stuItalianPhrase,
		language.Japanese:             stuJapanesePhrase,
		language.Korean:               stuKoreanPhrase,
		language.LatinAmericanSpanish: stuLatinAmericanSpanishPhrase,
		language.Russian:              stuRussianPhrase,
		language.SimplifiedChinese:    stuSimplifiedChinesePhrase,
		language.Spanish:              stuSpanishPhrase,
		language.TraditionalChinese:   stuTraditionalChinesePhrase}
)

var (
	// Stu represents stu.
	Stu = nook.Villager{
		Character:   stuCharacter,
		Personality: personality.Lazy,
		Phrase:      stuPhrase}
)
