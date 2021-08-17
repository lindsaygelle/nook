package dog

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
	cookieBirthday = nook.Birthday{
		Day:   18,
		Month: time.June}
)

var (
	cookieCode = nook.Code{
		Value: "dog08"}
)

var (
	cookieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cookie"}

	cookieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cookie"}

	cookieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cookie"}

	cookieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cookie"}

	cookieGermanName = nook.Name{
		Language: language.German,
		Value:    "Rosi"}

	cookieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Briciola"}

	cookieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペリーヌ"}

	cookieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베리"}

	cookieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Purita"}

	cookieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Куки"}

	cookieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珮琳"}

	cookieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Purita"}

	cookieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "珮琳"}
)

var (
	cookieName = nook.Languages{
		language.AmericanEnglish:      cookieAmericanEnglishName,
		language.CanadianFrench:       cookieCanadianFrenchName,
		language.Dutch:                cookieDutchName,
		language.French:               cookieFrenchName,
		language.German:               cookieGermanName,
		language.Italian:              cookieItalianName,
		language.Japanese:             cookieJapaneseName,
		language.Korean:               cookieKoreanName,
		language.LatinAmericanSpanish: cookieLatinAmericanSpanishName,
		language.Russian:              cookieRussianName,
		language.SimplifiedChinese:    cookieSimplifiedChineseName,
		language.Spanish:              cookieSpanishName,
		language.TraditionalChinese:   cookieTraditionalChineseName}
)

var (
	cookieCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: cookieBirthday,
		Code:     cookieCode,
		Key:      character.Cookie,
		Gender:   gender.Female,
		Name:     cookieName}
)

var (
	cookieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "arfer"}

	cookieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "woufi"}

	cookieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwispel"}

	cookieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "woufi"}

	cookieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nuffnuff"}

	cookieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bubabù"}

	cookieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "プリリン"}

	cookieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "초롱초롱"}

	cookieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "can"}

	cookieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гав-гав"}

	cookieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "琳琳"}

	cookieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fresita"}

	cookieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "琳琳"}
)

var (
	cookiePhrase = nook.Languages{
		language.AmericanEnglish:      cookieAmericanEnglishPhrase,
		language.CanadianFrench:       cookieCanadianFrenchPhrase,
		language.Dutch:                cookieDutchPhrase,
		language.French:               cookieFrenchPhrase,
		language.German:               cookieGermanPhrase,
		language.Italian:              cookieItalianPhrase,
		language.Japanese:             cookieJapanesePhrase,
		language.Korean:               cookieKoreanPhrase,
		language.LatinAmericanSpanish: cookieLatinAmericanSpanishPhrase,
		language.Russian:              cookieRussianPhrase,
		language.SimplifiedChinese:    cookieSimplifiedChinesePhrase,
		language.Spanish:              cookieSpanishPhrase,
		language.TraditionalChinese:   cookieTraditionalChinesePhrase}
)

var (
	Cookie = nook.Villager{
		Character:   cookieCharacter,
		Personality: personality.Peppy,
		Phrase:      cookiePhrase}
)
