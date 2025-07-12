package elephant

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
	// chaiBirthday represents chai birthday.
	chaiBirthday = nook.Birthday{
		Day:   6,
		Month: time.March}
)

var (
	// chaiCode represents chai code.
	chaiCode = nook.Code{
		Value: "elp11"}
)

var (
	// chaiAmericanEnglishName represents chai american english name.
	chaiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chai"}

	// chaiCanadianFrenchName represents chai canadian french name.
	chaiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chaï"}

	// chaiDutchName represents chai dutch name.
	chaiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chai"}

	// chaiFrenchName represents chai french name.
	chaiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chaï"}

	// chaiGermanName represents chai german name.
	chaiGermanName = nook.Name{
		Language: language.German,
		Value:    "Chai"}

	// chaiItalianName represents chai italian name.
	chaiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chai"}

	// chaiJapaneseName represents chai japanese name.
	chaiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フィーカ"}

	// chaiKoreanName represents chai korean name.
	chaiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피카"}

	// chaiLatinAmericanSpanishName represents chai latin american spanish name.
	chaiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chai"}

	// chaiRussianName represents chai russian name.
	chaiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чай"}

	// chaiSimplifiedChineseName represents chai simplified chinese name.
	chaiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啡卡"}

	// chaiSpanishName represents chai spanish name.
	chaiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chai"}

	// chaiTraditionalChineseName represents chai traditional chinese name.
	chaiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啡卡"}
)

var (
	// chaiName represents chai name.
	chaiName = nook.Languages{
		language.AmericanEnglish:      chaiAmericanEnglishName,
		language.CanadianFrench:       chaiCanadianFrenchName,
		language.Dutch:                chaiDutchName,
		language.French:               chaiFrenchName,
		language.German:               chaiGermanName,
		language.Italian:              chaiItalianName,
		language.Japanese:             chaiJapaneseName,
		language.Korean:               chaiKoreanName,
		language.LatinAmericanSpanish: chaiLatinAmericanSpanishName,
		language.Russian:              chaiRussianName,
		language.SimplifiedChinese:    chaiSimplifiedChineseName,
		language.Spanish:              chaiSpanishName,
		language.TraditionalChinese:   chaiTraditionalChineseName}
)

var (
	// chaiCharacter represents chai character.
	chaiCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: chaiBirthday,
		Code:     chaiCode,
		Key:      character.Chai,
		Gender:   gender.Female,
		Name:     chaiName,
		Special:  false}
)

var (
	// chaiAmericanEnglishPhrase represents chai american english phrase.
	chaiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "flap flap"}

	// chaiCanadianFrenchPhrase represents chai canadian french phrase.
	chaiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pachipach"}

	// chaiDutchPhrase represents chai dutch phrase.
	chaiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kaneel"}

	// chaiFrenchPhrase represents chai french phrase.
	chaiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pachipach"}

	// chaiGermanPhrase represents chai german phrase.
	chaiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zimtroll"}

	// chaiItalianPhrase represents chai italian phrase.
	chaiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ruuu ruuu"}

	// chaiJapanesePhrase represents chai japanese phrase.
	chaiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ロルロル"}

	// chaiKoreanPhrase represents chai korean phrase.
	chaiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "롤롤"}

	// chaiLatinAmericanSpanishPhrase represents chai latin american spanish phrase.
	chaiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "vuelalto"}

	// chaiRussianPhrase represents chai russian phrase.
	chaiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "машу ушами"}

	// chaiSimplifiedChinesePhrase represents chai simplified chinese phrase.
	chaiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大耳"}

	// chaiSpanishPhrase represents chai spanish phrase.
	chaiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "vuelalto"}

	// chaiTraditionalChinesePhrase represents chai traditional chinese phrase.
	chaiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大耳"}
)

var (
	// chaiPhrase represents chai phrase.
	chaiPhrase = nook.Languages{
		language.AmericanEnglish:      chaiAmericanEnglishPhrase,
		language.CanadianFrench:       chaiCanadianFrenchPhrase,
		language.Dutch:                chaiDutchPhrase,
		language.French:               chaiFrenchPhrase,
		language.German:               chaiGermanPhrase,
		language.Italian:              chaiItalianPhrase,
		language.Japanese:             chaiJapanesePhrase,
		language.Korean:               chaiKoreanPhrase,
		language.LatinAmericanSpanish: chaiLatinAmericanSpanishPhrase,
		language.Russian:              chaiRussianPhrase,
		language.SimplifiedChinese:    chaiSimplifiedChinesePhrase,
		language.Spanish:              chaiSpanishPhrase,
		language.TraditionalChinese:   chaiTraditionalChinesePhrase}
)

var (
	// Chai represents chai.
	Chai = nook.Villager{
		Character:   chaiCharacter,
		Personality: personality.Peppy,
		Phrase:      chaiPhrase}
)
