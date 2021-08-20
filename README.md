# Nook
[![Go Reference](https://pkg.go.dev/badge/github.com/lindsaygelle/nook.svg)](https://pkg.go.dev/github.com/lindsaygelle/nook)

Nook exports basic information for characters that appear in the video game series Animal Crossing.

Specifically this package centers on the organisation and classification of residents versus villagers, as well as
defining the common attributes between them.

Residents in the context of this package are Animal Crossing characters that cannot
be invited as home owners within the players town, city, or island. Residents perform a special role within the world
and often are quest givers. An example of a Resident would be Isabelle, Tom Nook or Guliver.

Villagers are Animal Crossing characters that can be invited to stay in the players town, city or island,
and make up the larger world population. Characters such as Alfonso, Tabby and Eggbert are all examples of Villagers. Villagers
can be given gifts, have personalities and homes.

This package separates and organises Animal Crossing characters by their animal type. This is a requirement due to the large number
of characters there are in the series, and that some characters overlap with other components in name. 
An example of this is the name Petunia, who is both a Cow and a Rhinoceros. Another example is Snooty, which is the name of one of the wolves and a personality type.

The information for this package has largely been sourced from various Wiki pages across the internet.

All contributions and corrections are welcome and encouraged.

## Go
This package can be added as a dependency using Go's package manager.

### Install
Adding the package.

```sh
go get -u github.com/lindsaygelle/nook
```

## Docker
This code can be run and executed from within a Docker Container. 
Below are the instructions to use the provided Dockerfile as the development and testing environment.

### Building 
Building the Container.

```sh
docker build . -t nook
```

### Running
Developing and running Go from within the Container.

```sh
docker run -it --rm --name nook nook
```

A `docker-compose.yml` file has also been added for convience.

## Contents
Nook at the package level only exports the concrete types consumed by this package. The philosophy was to keep each package extensible, without the concern of
searching through various subpackages to find all types that make up the package. Through this, ideally, the pattern will keep the imports relatively straightforward and simple. It should also allow each child package to focus on using only the known set of building blocks to create its exports, and limit
the exposure to coding towards circular dependencies. Most of the decisions were shaped by the content of the Animal Crossing series, versus developer preference. For example, the initial concept was to have all characters live at the root package, but name conflicts, data types, and maintance became an issue. Having each character exist in its own directory was also attempted, but imports became confusing. All recommendations welcome, but for now this appears to be the best trade off.

## Usage
The nook package can be used as the starting point for your own Animal Crossing program. By default all current Animal Crossing characters are provided. These can be imported by referencing the type of animal and the characters name.

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character/cat"
)

func main() {
	for _, cat := range []nook.Villager{cat.Ankha, cat.Tabby} {
		b, err := json.Marshal(cat)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}
}

```

Characters are also exported in bulk. These can be accessed by using the animal type and the villagers or residents variable.

```go
package main

import (
	"fmt"

	"github.com/lindsaygelle/nook/character/dog"
	"golang.org/x/text/language"
)

func main() {
	for _, villager := range dog.Villagers {
		fmt.Println(fmt.Sprintf("%s is a villager in Animal Crossing", villager.Name.Must(language.AmericanEnglish).Value))
	}
	for _, resident := range dog.Residents {
		fmt.Println(fmt.Sprintf("%s is a resident in Animal Crossing", resident.Name.Must(language.AmericanEnglish).Value))
	}
}

```

Another example creating a simple API.

```go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character/cat"
	"github.com/lindsaygelle/nook/character/dog"
	"github.com/lindsaygelle/nook/character/wolf"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func createVillagersHandler(villagers nook.Villagers) http.HandlerFunc {
	b, err := json.Marshal(villagers)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/cats", createVillagersHandler(cat.Villagers))
	http.HandleFunc("/dogs", createVillagersHandler(dog.Villagers))
	http.HandleFunc("/wolves", createVillagersHandler(wolf.Villagers))
	http.ListenAndServe(":8080", nil)
}

```

## Extending
Nook was built with extensibility in mind. The simplest way to add to the functionality of the package is to import the concrete type and apply the required changes or functionality.

Below is an example of adding functionality to the `nook.Villager` type.

```go
package main

import (
	"fmt"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character/alligator"
	"golang.org/x/text/language"
)

var (
	alligators = []nook.Villager{
		alligator.Alfonso,
		alligator.Alli}
)

type Villager struct {
	nook.Villager
}

func (v Villager) Greet() string {
	name := v.Name.Must(language.AmericanEnglish).Value
	phrase := v.Phrase.Must(language.AmericanEnglish).Value
	return fmt.Sprintf("%s! My name is %s. Nice to meet you.", phrase, name)
}

func makeVillager(villager nook.Villager) Villager {
	return Villager{villager}
}

func makeVillagers(villagers ...nook.Villager) []Villager {
	v := make([]Villager, len(villagers))
	for i, villager := range villagers {
		v[i] = makeVillager(villager)
	}
	return v
}

func main() {
	villagers := makeVillagers(alligators...)
	for _, villager := range villagers {
		fmt.Println(villager.Greet())
	}
}

```

Another example is adding an outfit to one of your favourite Animal Crossing villagers. 

```go 
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character/duck"
	"golang.org/x/text/language"
)

type Outfit struct {
	Name nook.Languages
}

type Villager struct {
	nook.Villager
	Outfit Outfit
}

var (
	outfitFrogTeeAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish, Value: "Frog Tee"}
)

var (
	outfitFrogTeeName = nook.Languages{
		language.AmericanEnglish: outfitFrogTeeAmericanEnglish}
)

var (
	outfitFrogTee = Outfit{
		Name: outfitFrogTeeName}
)

var (
	Scoot = newVillager(duck.Scoot, outfitFrogTee)
)

func newVillager(villager nook.Villager, outfit Outfit) Villager {
	return Villager{
		Villager: villager,
		Outfit:   outfit}
}

func main() {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(Scoot)
	if err != nil {
		panic(err)
	}
	fmt.Println(buffer.String())
}

```
