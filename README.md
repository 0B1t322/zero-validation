# Zero validation
Zero validation is a library for simple validation with 
code-gen tools integration.

Integrations presented as plugins for grpc, gqlgen and `go generate` for simple go code.

## Integration examples
1. [grpc](_example/grpc)
2. [gqlgen](_example/gqlgen)

## Concept

The concept of library is use extractors object to get object fields in runtime without reflection.
The simple extractor object looks like:
```go
type objectExtractor struct {
	ID fields.StructField[Object, uint64]
}

var ObjectExtractor = objectExtractor{
	// ... code for extractor
}
```

This allow us to use this syntax for validations:
```go
type Object struct {
	ID uint64
}

func main() {
    obj := Object{}
    err := validate.Struct(
        ctx,
        obj,
        validate.Field(ObjectExtractor.ID, rule.Required[uint64]()),
    )
    if err != nil {
        return nil, err
    }
}
```

But to not always allocate rules, you can use validators store, that allow you to reuse validation rules.
```go
type Object struct {
	ID uint64
}

type objectValidator struct {
	
}

func (objectValidator) Name() string {
	return "objectValidator"
}

func (objectValidator) Rules() []validate.FieldRules[Object] {
	return []validate.FieldRules[Object]{
	    validate.Field(ObjectExtractor.ID, rule.Required[uint64]()),	
    }
}

func main() {
	store := validators.GlobalMapStore()

	validators.InitValidatorInStore(store, objectValidator{})
	
    obj := Object{}
    err := validate.Struct(
        ctx,
        obj,
        validators.GetValidatorRulesFromStore[objectValidator]()...,
    )
    if err != nil {
        return nil, err
    }
}
```

[//]: # (TODO: WRITE ABOUT i8n)