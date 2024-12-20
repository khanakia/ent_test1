package entsaasmedia

import (
	"encoding/json"
	"go/token"
	"strings"
	"unicode"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema"
	"github.com/go-openapi/inflect"
)

func ruleset() *inflect.Ruleset {
	rules := inflect.NewDefaultRuleset()
	// Add common initialism from golint and more.
	for _, w := range []string{
		"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID",
		"HCL", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC",
		"MB", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO",
		"TCP", "TLS", "TTL", "UDP", "UI", "UID", "URI", "URL", "UTF8", "UUID",
		"VM", "XML", "XMPP", "XSRF", "XSS",
	} {
		acronyms[w] = struct{}{}
		rules.AddAcronym(w)
	}
	return rules
}

var (
	rules = ruleset()
)

type FieldType struct {
	Name string
}

type Field struct {
	Name            string
	Type            string
	Unique          bool
	BuilderField    string
	MutationAdd     string
	MutationAdded   string
	MutationSet     string
	MutationRemove  string
	MutationRemoved string
	MutationClear   string
	MutationCleared string
	MutationReset   string
	StructField     string
	Tag             string

	GqlCreate string
	GqlAdd    string
	GqlRemove string
	GqlClear  string
}

var (
	camel    = gen.Funcs["camel"].(func(string) string)
	pascal   = gen.Funcs["pascal"].(func(string) string)
	plural   = gen.Funcs["plural"].(func(string) string)
	singular = gen.Funcs["singular"].(func(string) string)
	snake    = gen.Funcs["snake"].(func(string) string)
)

func (a *Field) make() {
	a.Unique = false
	a.Type = "string"
	a.BuilderField = builderField(a.Name)
	a.MutationAdd = a.GetMutationAdd()
	a.MutationAdded = a.GetMutationAdded()
	a.MutationSet = a.GetMutationSet()
	a.MutationRemove = a.GetMutationRemove()
	a.MutationRemoved = a.GetMutationRemoved()
	a.MutationClear = a.GetMutationClear()
	a.MutationCleared = a.GetMutationCleared()
	a.MutationReset = a.GetMutationReset()
	a.StructField = pascal(a.Name)

	a.GqlCreate = camel(singular(a.Name)) + "IDs"
	a.GqlAdd = "add" + pascal(singular(a.Name)) + "IDs"
	a.GqlRemove = "remove" + pascal(singular(a.Name)) + "IDs"
	a.GqlClear = camel(snake(a.MutationClear))
}

// MutationAdd returns the method name for adding edge ids.
func (e Field) GetMutationAdd() string {
	return "Add" + pascal(rules.Singularize(e.Name)) + "IDs"
}

func (e Field) GetMutationAdded() string {
	return pascal(rules.Singularize(e.Name)) + "IDs"
}

func (e Field) GetMutationRemoved() string {
	return "Removed" + pascal(rules.Singularize(e.Name)) + "IDs"
}

// var mutMethods = func() map[string]bool {
// 	names := map[string]bool{"Client": true, "Tx": true, "Where": true, "SetOp": true}
// 	t := reflect.TypeOf(new(ent.Mutation)).Elem()
// 	for i := 0; i < t.NumMethod(); i++ {
// 		names[t.Method(i).Name] = true
// 	}
// 	return names
// }()

// MutationClear returns the method name for clearing the edge value.
// The default name is "Clear<EdgeName>". If the method conflicts
// with the mutation methods, suffix the method with "Edge".
func (e Field) GetMutationClear() string {
	name := "Clear" + pascal(e.Name)
	// if _, ok := mutMethods[name]; ok {
	// name += "Field"
	// }
	return name
}

// MutationRemove returns the method name for removing edge ids.
func (e Field) GetMutationRemove() string {
	return "Remove" + pascal(rules.Singularize(e.Name)) + "IDs"
}

// MutationSet returns the method name for setting the edge id.
func (e Field) GetMutationSet() string {
	return "Set" + pascal(e.Name) + "ID"
}

// MutationCleared returns the method name for indicating if the edge
// was cleared in the mutation. The default name is "<EdgeName>Cleared".
// If the method conflicts with the mutation methods, add "Edge" the
// after the edge name.
func (e Field) GetMutationCleared() string {
	name := pascal(e.Name) + "Cleared"
	// if _, ok := mutMethods[name]; ok {
	// 	return pascal(e.Name) + "EdgeCleared"
	// }
	return name
}

// MutationReset returns the method name for resetting the edge value.
// The default name is "Reset<EdgeName>". If the method conflicts
// with the mutation methods, suffix the method with "Edge".
func (e Field) GetMutationReset() string {
	name := "Reset" + pascal(e.Name)
	// if _, ok := mutMethods[name]; ok {
	// 	name += "Edge"
	// }
	return name
}

type Annotation struct {
	Fields []*Field
}

// Name implements ent.Annotation interface.
func (Annotation) Name() string {
	return "SaasMedia"
}

func (a *Annotation) Decode(annotation interface{}) error {
	buf, err := json.Marshal(annotation)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, a)
}

func ExtractAnnotation(ants gen.Annotations) (*Annotation, error) {
	ant := &Annotation{}
	if ants != nil && ants[ant.Name()] != nil {
		if err := ant.Decode(ants[ant.Name()]); err != nil {
			return nil, err
		}
	}
	return ant, nil
}

func (a Annotation) Merge(other schema.Annotation) schema.Annotation {

	var ant Annotation
	switch other := other.(type) {
	case Annotation:
		ant = other
	case *Annotation:
		if other != nil {
			ant = *other
		}
	default:
		return a
	}

	if len(ant.Fields) > 0 {
		a.Fields = append(a.Fields, ant.Fields...)
	}

	for _, field := range a.Fields {
		field.Unique = true
	}

	// if a.StructTag == nil && len(ant.StructTag) > 0 {
	// 	a.StructTag = make(map[string]string, len(ant.StructTag))
	// }
	// for k, v := range ant.StructTag {
	// 	a.StructTag[k] = v
	// }
	// if len(ant.ID) > 0 {
	// 	a.ID = ant.ID
	// }
	return a
}

type FieldOption struct {
	Name string
	Tag  string
}

func MediaFields(inputs ...FieldOption) Annotation {
	if len(inputs) == 0 {
		inputs = []FieldOption{}
	}

	a := []*Field{}
	for _, f := range inputs {

		if f.Tag == "" {
			f.Tag = strings.ToLower(f.Name)
		}

		field := &Field{
			Name: f.Name + "_medias",
			Tag:  f.Tag,
		}
		field.make()

		a = append(a, field)
	}
	return Annotation{Fields: a}
}

// func (e MediaInputOption) BuilderField() string {
// 	return e.Name
// 	// return builderField(e.Name)
// }

func names(ids ...string) map[string]struct{} {
	m := make(map[string]struct{})
	for i := range ids {
		m[ids[i]] = struct{}{}
	}
	return m
}

var (
	privateField = names(
		"config",
		"ctx",
		"done",
		"hooks",
		"inters",
		"limit",
		"mutation",
		"offset",
		"oldValue",
		"order",
		"op",
		"path",
		"predicates",
		"typ",
		"unique",
		"driver",
	)
)

// builderField returns the struct field for the given name
// and ensures it doesn't conflict with Go keywords and other
// builder fields, and it is not exported.
func builderField(name string) string {
	_, ok := privateField[name]
	if ok || token.Lookup(name).IsKeyword() || strings.ToUpper(name[:1]) == name[:1] {
		return "_" + name
	}
	return name
}

var _ interface {
	schema.Annotation
	schema.Merger
} = (*Annotation)(nil)

var (
	acronyms = make(map[string]struct{})
)

// pascal converts the given name into a PascalCase.
//
//	user_info 	=> UserInfo
//	full_name 	=> FullName
//	user_id   	=> UserID
//	full-admin	=> FullAdmin
// func pascal(s string) string {
// 	words := strings.FieldsFunc(s, isSeparator)
// 	return pascalWords(words)
// }

func isSeparator(r rune) bool {
	return r == '_' || r == '-' || unicode.IsSpace(r)
}

func pascalWords(words []string) string {
	for i, w := range words {
		upper := strings.ToUpper(w)
		if _, ok := acronyms[upper]; ok {
			words[i] = upper
		} else {
			words[i] = rules.Capitalize(w)
		}
	}
	return strings.Join(words, "")
}
