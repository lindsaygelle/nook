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
	// cookieBirthday represents cookie birthday.
	cookieBirthday = nook.Birthday{
		Day:   18,
		Month: time.June}
)

var (
	// cookieCode represents cookie code.
	cookieCode = nook.Code{
		Value: "dog08"}
)

var (
	// cookieAmericanEnglishName represents cookie american english name.
	cookieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cookie"}

	// cookieCanadianFrenchName represents cookie canadian french name.
	cookieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cookie"}

	// cookieDutchName represents cookie dutch name.
	cookieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cookie"}

	// cookieFrenchName represents cookie french name.
	cookieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cookie"}

	// cookieGermanName represents cookie german name.
	cookieGermanName = nook.Name{
		Language: language.German,
		Value:    "Rosi"}

	// cookieItalianName represents cookie italian name.
	cookieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Briciola"}

	// cookieJapaneseName represents cookie japanese name.
	cookieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペリーヌ"}

	// cookieKoreanName represents cookie korean name.
	cookieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베리"}

	// cookieLatinAmericanSpanishName represents cookie latin american spanish name.
	cookieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Purita"}

	// cookieRussianName represents cookie russian name.
	cookieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Куки"}

	// cookieSimplifiedChineseName represents cookie simplified chinese name.
	cookieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珮琳"}

	// cookieSpanishName represents cookie spanish name.
	cookieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Purita"}

	// cookieTraditionalChineseName represents cookie traditional chinese name.
	cookieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "珮琳"}
)

var (
	// cookieName represents cookie name.
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
	// cookieCharacter represents cookie character.
	cookieCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: cookieBirthday,
		Code:     cookieCode,
		Key:      character.Cookie,
		Gender:   gender.Female,
		Name:     cookieName,
		Special:  false}
)

var (
	// cookieAmericanEnglishPhrase represents cookie american english phrase.
	cookieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "arfer"}

	// cookieCanadianFrenchPhrase represents cookie canadian french phrase.
	cookieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "woufi"}

	// cookieDutchPhrase represents cookie dutch phrase.
	cookieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwispel"}

	// cookieFrenchPhrase represents cookie french phrase.
	cookieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "woufi"}

	// cookieGermanPhrase represents cookie german phrase.
	cookieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nuffnuff"}

	// cookieItalianPhrase represents cookie italian phrase.
	cookieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bubabù"}

	// cookieJapanesePhrase represents cookie japanese phrase.
	cookieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "プリリン"}

	// cookieKoreanPhrase represents cookie korean phrase.
	cookieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "초롱초롱"}

	// cookieLatinAmericanSpanishPhrase represents cookie latin american spanish phrase.
	cookieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "can"}

	// cookieRussianPhrase represents cookie russian phrase.
	cookieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гав-гав"}

	// cookieSimplifiedChinesePhrase represents cookie simplified chinese phrase.
	cookieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "琳琳"}

	// cookieSpanishPhrase represents cookie spanish phrase.
	cookieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fresita"}

	// cookieTraditionalChinesePhrase represents cookie traditional chinese phrase.
	cookieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "琳琳"}
)

var (
	// cookiePhrase represents cookie phrase.
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
	// Cookie represents cookie.
	Cookie = nook.Villager{
		Character:   cookieCharacter,
		Personality: personality.Peppy,
		Phrase:      cookiePhrase}
)
