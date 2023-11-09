# Nook
![Nook](https://repository-images.githubusercontent.com/396270808/54f16924-c7d4-403d-b946-d7501d7cbbe7)

Nook exports basic information for characters that appear in the video game series [Animal Crossing](https://animal-crossing.com/).

[![Go Reference](https://pkg.go.dev/badge/github.com/lindsaygelle/nook.svg)](https://pkg.go.dev/github.com/lindsaygelle/nook)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/nook)](/LICENSE)


## Introduction
This package centers on the organisation and classification of residents versus villagers, as well as defining the common attributes between them. Residents in the context of this package are Animal Crossing characters that cannot be invited as home owners within the players town, city, or island. Residents perform a special role within the world and often are quest givers. An example of a Resident would be [Isabelle](./character/dog/isabelle.go), [Tom Nook](./character/raccoon/tomnook.go) or [Guliver](./character/seagull/gulliver.go). Villagers are Animal Crossing characters that can be invited to stay in the players town, city or island, and make up the larger world population. Characters such as [Alfonso](./character/alligator/alfonso.go), [Tabby](./character/cat/tabby.go) and [Egbert](./character/chicken/egbert.go) are all examples of Villagers. Villagers
can be given gifts, have personalities and homes.

## Design
This package separates and organises Animal Crossing characters by their animal type. This is a requirement due to the large number of characters there are in the series, and that some characters overlap with other components in name. An example of this is the name Petunia, who is both a [Cow](./character/cow/petunia.go) and a [Rhinoceros](./character/rhinoceros/petunia.go). Another example is Snooty, which is the name of one of the wolves and a [personality type](./personality/).

## Organisation
Nook at the package level only exports the concrete types consumed by this package. The philosophy was to keep each package extensible, without the concern of searching through various subpackages to find all types that make up the package. Through this, ideally, the pattern will keep the imports relatively straightforward and simple. It should also allow each child package to focus on using only the known set of building blocks to create its exports, and limit the exposure to coding towards circular dependencies. Most of the decisions were shaped by the content of the Animal Crossing series, versus developer preference. For example, the initial concept was to have all characters live at the root package, but name conflicts, data types, and maintenance became an issue. Having each character exist in its own directory was also attempted, but imports became confusing. All recommendations welcome, but for now this appears to be the best trade off.

## Disclaimer
The information for this package has largely been sourced from various Wiki pages across the internet. All contributions and corrections are welcome and encouraged. If there are any corrections, please feel free to raise them in a pull request.

## Installation
Add Nook as a dependency to your Go project using the following command:

```sh
go get -u github.com/lindsaygelle/nook
```

## Docker
You can utilize Nook within a Docker container with the provided Dockerfile. Here's how to build and run the container:

Building the Docker container.
```sh
docker build . -t nook
```

Developing and running Go from within the Docker container.
```sh
docker run -it --rm --name nook nook
```

## Docker-compose
A [`docker-compose.yml`](./docker-compose.yml) file has also been added for convenience.

# Usage
The nook package can be used as the starting point for your own Animal Crossing program. By default all current Animal Crossing characters are provided. These can be imported by referencing the type of animal and the characters name.

Accessing individual villagers:
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character/cat"
)

// main is the entry point of the program.
// It demonstrates how to use the Nook package to access information about Animal Crossing characters.
func main() {
	// Loop through an array of Animal Crossing cat villagers (Ankha and Tabby).
	for _, cat := range []nook.Villager{cat.Ankha, cat.Tabby} {
		// Marshal the current cat villager into a JSON string.
		b, err := json.Marshal(cat)
		if err != nil {
			// If there is an error during JSON marshaling, panic with the error.
			panic(err)
		}
		// Print the JSON representation of the current cat villager.
		fmt.Println(string(b))
	}
}
```

Characters are also exported in bulk. These can be accessed by using the animal type and the villagers or residents variable:
```go
package main

import (
	"fmt"

	"github.com/lindsaygelle/nook/character/dog"
	"golang.org/x/text/language"
)

// main is the entry point of the program.
// It demonstrates how to use the Nook package to access and display information about Animal Crossing dog villagers and residents.
func main() {
	// Loop through all dog villagers and print their names as villagers in Animal Crossing.
	for _, villager := range dog.Villagers {
		fmt.Println(fmt.Sprintf("%s is a villager in Animal Crossing", villager.Name.Must(language.AmericanEnglish).Value))
	}
	// Loop through all dog residents and print their names as residents in Animal Crossing.
	for _, resident := range dog.Residents {
		fmt.Println(fmt.Sprintf("%s is a resident in Animal Crossing", resident.Name.Must(language.AmericanEnglish).Value))
	}
}
```

Creating a simple API:
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

// rootHandler is an HTTP request handler for the root endpoint ("/").
// It returns a 200 OK response with an empty body.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

// createVillagersHandler creates an HTTP request handler for specific animal type villagers.
// It receives a slice of nook.Villagers for the specified animal type and returns an HTTP response with JSON data of the villagers.
func createVillagersHandler(villagers nook.Villagers) http.HandlerFunc {
	// Marshal the provided villagers into a JSON byte slice.
	b, err := json.Marshal(villagers)
	if err != nil {
		// If there is an error during JSON marshaling, panic with the error.
		panic(err)
	}
	// Return the HTTP request handler function.
	return func(w http.ResponseWriter, r *http.Request) {
		// Set the "Content-Type" header to indicate that the response contains JSON data.
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		// Write a 200 OK response with the JSON data in the response body.
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func main() {
	// Register HTTP request handlers for different endpoints.
	http.HandleFunc("/", rootHandler)              // Root endpoint
	http.HandleFunc("/cats", createVillagersHandler(cat.Villagers))   // Endpoint for cat villagers
	http.HandleFunc("/dogs", createVillagersHandler(dog.Villagers))   // Endpoint for dog villagers
	http.HandleFunc("/wolves", createVillagersHandler(wolf.Villagers)) // Endpoint for wolf villagers

	// Start the HTTP server on port 8080 and listen for incoming requests.
	http.ListenAndServe(":8080", nil)
}
```

## Extending
Nook was designed with extensibility in mind. You can easily add new functionality by importing the relevant concrete types and applying the required changes. For example, you can create your custom attributes for villagers:

Below is an example of adding functionality to the [`nook.Villager`](./villager.go) type.

```go
package main

import (
	"fmt"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/character/alligator"
	"golang.org/x/text/language"
)

// alligators is a slice containing some Animal Crossing alligator villagers.
var (
	alligators = []nook.Villager{
		alligator.Alfonso,
		alligator.Alli,
	}
)

// Villager represents an Animal Crossing villager with additional functionality.
type Villager struct {
	nook.Villager
}

// Greet returns a friendly greeting from the villager in the American English language.
func (v Villager) Greet() string {
	name := v.Name.Must(language.AmericanEnglish).Value
	phrase := v.Phrase.Must(language.AmericanEnglish).Value
	return fmt.Sprintf("%s! My name is %s. Nice to meet you.", phrase, name)
}

// makeVillager creates a new Villager from the provided nook.Villager.
func makeVillager(villager nook.Villager) Villager {
	return Villager{villager}
}

// makeVillagers creates a slice of Villagers from the provided nook.Villagers.
func makeVillagers(villagers ...nook.Villager) []Villager {
	v := make([]Villager, len(villagers))
	for i, villager := range villagers {
		v[i] = makeVillager(villager)
	}
	return v
}

func main() {
	// Create a slice of Villagers from the alligators slice.
	villagers := makeVillagers(alligators...)

	// Greet each villager and print their friendly messages.
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

// Outfit represents an outfit of an Animal Crossing villager.
type Outfit struct {
	Name nook.Languages
}

// Villager represents an Animal Crossing villager with an additional Outfit field.
type Villager struct {
	nook.Villager
	Outfit Outfit
}

// outfitFrogTeeAmericanEnglish is the American English name for an outfit called "Frog Tee".
var outfitFrogTeeAmericanEnglish = nook.Name{
	Language: language.AmericanEnglish, Value: "Frog Tee"}

// outfitFrogTeeName contains the name of the outfit "Frog Tee" in the American English language.
var outfitFrogTeeName = nook.Languages{
	language.AmericanEnglish: outfitFrogTeeAmericanEnglish}

// outfitFrogTee is an Outfit with the name "Frog Tee".
var outfitFrogTee = Outfit{
	Name: outfitFrogTeeName}

// Scoot is a Villager named "Scoot" with the outfit "Frog Tee".
var Scoot = newVillager(duck.Scoot, outfitFrogTee)

// newVillager creates a new Villager with the provided nook.Villager and Outfit.
func newVillager(villager nook.Villager, outfit Outfit) Villager {
	return Villager{
		Villager: villager,
		Outfit:   outfit,
	}
}

func main() {
	// Create a buffer to store the JSON-encoded data.
	buffer := new(bytes.Buffer)

	// Create an encoder with the buffer to encode the Villager as JSON.
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	// Encode the Scoot Villager as JSON and store it in the buffer.
	err := encoder.Encode(Scoot)
	if err != nil {
		panic(err)
	}

	// Print the JSON data stored in the buffer.
	fmt.Println(buffer.String())
}
```

## Contributing
We warmly welcome contributions to Nook. Whether you have innovative ideas, bug reports, or enhancements in mind, please share them with us by submitting GitHub issues or creating pull requests. For substantial contributions, it's a good practice to start a discussion by creating an issue to ensure alignment with the project's goals and direction. Refer to the [CONTRIBUTING](./CONTRIBUTING.md) file for comprehensive details.

## Branching
For a smooth collaboration experience, we have established branch naming conventions and guidelines. Please consult the [BRANCH_NAMING_CONVENTION](./BRANCH_NAMING_CONVENTION.md) document for comprehensive information and best practices.

## License
Nook is released under the MIT License, granting you the freedom to use, modify, and distribute the code within this repository in accordance with the terms of the license. For additional information, please review the [LICENSE](./LICENSE) file.

## Security
If you discover a security vulnerability within this project, please consult the [SECURITY](./SECURITY.md) document for information and next steps.

## Code Of Conduct
This project has adopted the [Amazon Open Source Code of Conduct](https://aws.github.io/code-of-conduct). For additional information, please review the [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md) file.
