// Package catalog assembles the full Animal Crossing character dataset into
// backend-friendly read-only registries without changing the underlying domain
// model packages. It also provides exact-match, animal-scoped lookup helpers
// over those registries.
//
// The catalog remains grouped by animal because character keys are not
// guaranteed to be globally unique across the full dataset.
package catalog
