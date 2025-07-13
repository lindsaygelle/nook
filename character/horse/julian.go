package horse

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
	// julianBirthday represents julian birthday.
	julianBirthday = nook.Birthday{
		Day:   15,
		Month: time.March}
)

var (
	// julianCode represents julian code.
	julianCode = nook.Code{
		Value: "hrs13"}
)

var (
	// julianAmericanEnglishName represents julian american english name.
	julianAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Julian"}

	// julianCanadianFrenchName represents julian canadian french name.
	julianCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lico"}

	// julianDutchName represents julian dutch name.
	julianDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Julian"}

	// julianFrenchName represents julian french name.
	julianFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lico"}

	// julianGermanName represents julian german name.
	julianGermanName = nook.Name{
		Language: language.German,
		Value:    "Jimmy"}

	// julianItalianName represents julian italian name.
	julianItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giuliano"}

	// julianJapaneseName represents julian japanese name.
	julianJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュリー"}

	// julianKoreanName represents julian korean name.
	julianKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "유니오"}

	// julianLatinAmericanSpanishName represents julian latin american spanish name.
	julianLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Azulino"}

	// julianRussianName represents julian russian name.
	julianRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джулиан"}

	// julianSimplifiedChineseName represents julian simplified chinese name.
	julianSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱黎"}

	// julianSpanishName represents julian spanish name.
	julianSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Azulino"}

	// julianTraditionalChineseName represents julian traditional chinese name.
	julianTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱黎"}
)

var (
	// julianName represents julian name.
	julianName = nook.Languages{
		language.AmericanEnglish:      julianAmericanEnglishName,
		language.CanadianFrench:       julianCanadianFrenchName,
		language.Dutch:                julianDutchName,
		language.French:               julianFrenchName,
		language.German:               julianGermanName,
		language.Italian:              julianItalianName,
		language.Japanese:             julianJapaneseName,
		language.Korean:               julianKoreanName,
		language.LatinAmericanSpanish: julianLatinAmericanSpanishName,
		language.Russian:              julianRussianName,
		language.SimplifiedChinese:    julianSimplifiedChineseName,
		language.Spanish:              julianSpanishName,
		language.TraditionalChinese:   julianTraditionalChineseName}
)

var (
	// julianCharacter represents julian character.
	julianCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: julianBirthday,
		Code:     julianCode,
		Key:      character.Julian,
		Gender:   gender.Male,
		Name:     julianName,
		Special:  false}
)

var (
	// julianAmericanEnglishPhrase represents julian american english phrase.
	julianAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "glitter"}

	// julianCanadianFrenchPhrase represents julian canadian french phrase.
	julianCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "western"}

	// julianDutchPhrase represents julian dutch phrase.
	julianDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "glitter"}

	// julianFrenchPhrase represents julian french phrase.
	julianFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trottoune"}

	// julianGermanPhrase represents julian german phrase.
	julianGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wihihihi"}

	// julianItalianPhrase represents julian italian phrase.
	julianItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vamos"}

	// julianJapanesePhrase represents julian japanese phrase.
	julianJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ね、キミ"}

	// julianKoreanPhrase represents julian korean phrase.
	julianKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그대여"}

	// julianLatinAmericanSpanishPhrase represents julian latin american spanish phrase.
	julianLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trotót"}

	// julianRussianPhrase represents julian russian phrase.
	julianRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "блеск"}

	// julianSimplifiedChinesePhrase represents julian simplified chinese phrase.
	julianSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喂！你"}

	// julianSpanishPhrase represents julian spanish phrase.
	julianSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pinchito"}

	// julianTraditionalChinesePhrase represents julian traditional chinese phrase.
	julianTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喂！你"}
)

var (
	// julianPhrase represents julian phrase.
	julianPhrase = nook.Languages{
		language.AmericanEnglish:      julianAmericanEnglishPhrase,
		language.CanadianFrench:       julianCanadianFrenchPhrase,
		language.Dutch:                julianDutchPhrase,
		language.French:               julianFrenchPhrase,
		language.German:               julianGermanPhrase,
		language.Italian:              julianItalianPhrase,
		language.Japanese:             julianJapanesePhrase,
		language.Korean:               julianKoreanPhrase,
		language.LatinAmericanSpanish: julianLatinAmericanSpanishPhrase,
		language.Russian:              julianRussianPhrase,
		language.SimplifiedChinese:    julianSimplifiedChinesePhrase,
		language.Spanish:              julianSpanishPhrase,
		language.TraditionalChinese:   julianTraditionalChinesePhrase}
)

var (
	// Julian represents julian.
	Julian = nook.Villager{
		Character:   julianCharacter,
		Personality: personality.Smug,
		Phrase:      julianPhrase}
)
