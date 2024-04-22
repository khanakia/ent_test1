package natshandlers

type Code string

func (e Code) ToString() string {
	return string(e)
}

// Common error codes
var (
	ErrCodeJSONUnmarshal            Code = "json/unmarshal"
	ErrCodeDatabaseUnexpected       Code = "db/unexpected"
	ErrCodeUnauthorized             Code = "auth/unauthorized"
	ErrCodeValidationFailed         Code = "validation/failed"
	ErrCodeInternalServerUnexpected Code = "server/unexpected"
	ErrCodeDatabaseNotFound         Code = "db/record(s)-not-found"
	ErrCodeURIParse                 Code = "url/uri-parse"
)

// Catalogue error codes
var (
	ErrCodeCatalogueNotFound      Code = "catalogue/not-found"
	ErrCodeCatalogueAlreadyExists Code = "catalogue/already-exists"
)

// Catalogue Item error codes
var (
	ErrCodeCatalogueItemNotFound Code = "catalogue-item/not-found"
)

// Country error codes
var (
	ErrCodeCountryNotFound Code = "country/not-found"
)

// Language error codes
var (
	ErrCodeLanguageNotFound Code = "language/not-found"
)
