package schema

/* cookie-cutter standard interface stuff */

func (typeBase) _Type()                    {}
func (t typeBase) TypeSystem() *TypeSystem { return t.universe }
func (t typeBase) Name() TypeName          { return t.name }

func (TypeBool) Kind() Kind   { return Kind_Bool }
func (TypeString) Kind() Kind { return Kind_String }
func (TypeBytes) Kind() Kind  { return Kind_Bytes }
func (TypeInt) Kind() Kind    { return Kind_Int }
func (TypeFloat) Kind() Kind  { return Kind_Float }
func (TypeMap) Kind() Kind    { return Kind_Map }
func (TypeList) Kind() Kind   { return Kind_List }
func (TypeLink) Kind() Kind   { return Kind_Link }
func (TypeUnion) Kind() Kind  { return Kind_Union }
func (TypeStruct) Kind() Kind { return Kind_Struct }
func (TypeEnum) Kind() Kind   { return Kind_Enum }

/* interesting methods per Type type */

// IsAnonymous is returns true if the type was unnamed.  Unnamed types will
// claim to have a Name property like `{Foo:Bar}`, and this is not guaranteed
// to be a unique string for all types in the universe.
func (t TypeMap) IsAnonymous() bool {
	return t.anonymous
}

// KeyType returns the Type of the map keys.
//
// Note that map keys will must always be some type which is representable as a
// string in the IPLD Data Model (e.g. either TypeString or TypeEnum).
func (t TypeMap) KeyType() Type {
	return t.keyType
}

// ValueType returns to the Type of the map values.
func (t TypeMap) ValueType() Type {
	return t.valueType
}

// ValueIsNullable returns a bool describing if the map values are permitted
// to be null.
func (t TypeMap) ValueIsNullable() bool {
	return t.valueNullable
}

// IsAnonymous is returns true if the type was unnamed.  Unnamed types will
// claim to have a Name property like `[Foo]`, and this is not guaranteed
// to be a unique string for all types in the universe.
func (t TypeList) IsAnonymous() bool {
	return t.anonymous
}

// ValueType returns to the Type of the list values.
func (t TypeList) ValueType() Type {
	return t.valueType
}

// ValueIsNullable returns a bool describing if the list values are permitted
// to be null.
func (t TypeList) ValueIsNullable() bool {
	return t.valueNullable
}

// Members returns the list of all types that are possible inhabitants of this union.
func (t TypeUnion) Members() []Type {
	a := make([]Type, len(t.members))
	for i := range t.members {
		a[i] = t.members[i]
	}
	return a
}

func (t TypeUnion) RepresentationStrategy() UnionRepresentation {
	return t.representation
}

func (r UnionRepresentation_Keyed) GetDiscriminant(t Type) string {
	for d, t2 := range r.table {
		if t2 == t {
			return d
		}
	}
	panic("that type isn't a member of this union")
}

// Fields returns a slice of descriptions of the object's fields.
func (t TypeStruct) Fields() []StructField {
	a := make([]StructField, len(t.fields))
	for i := range t.fields {
		a[i] = t.fields[i]
	}
	return a
}

// Field looks up a StructField by name, or returns nil if no such field.
func (t TypeStruct) Field(name string) *StructField {
	if v, ok := t.fieldsMap[name]; ok {
		return &v
	}
	return nil
}

// Parent returns the type information that this field describes a part of.
//
// While in many cases, you may know the parent already from context,
// there may still be situations where want to pass around a field and
// not need to continue passing down the parent type with it; this method
// helps your code be less redundant in such a situation.
// (You'll find this useful for looking up any rename directives, for example,
// when holding onto a field, since that requires looking up information from
// the representation strategy, which is a property of the type as a whole.)
func (f StructField) Parent() *TypeStruct { return f.parent }

// Name returns the string name of this field.  The name is the string that
// will be used as a map key if the structure this field is a member of is
// serialized as a map representation.
func (f StructField) Name() string { return f.name }

// Type returns the Type of this field's value.  Note the field may
// also be unset if it is either Optional or Nullable.
func (f StructField) Type() Type { return f.typ }

// IsOptional returns true if the field is allowed to be absent from the object.
// If IsOptional is false, the field may be absent from the serial representation
// of the object entirely.
//
// Note being optional is different than saying the value is permitted to be null!
// A field may be both nullable and optional simultaneously, or either, or neither.
func (f StructField) IsOptional() bool { return f.optional }

// IsNullable returns true if the field value is allowed to be null.
//
// If is Nullable is false, note that it's still possible that the field value
// will be absent if the field is Optional!  Being nullable is unrelated to
// whether the field's presence is optional as a whole.
//
// Note that a field may be both nullable and optional simultaneously,
// or either, or neither.
func (f StructField) IsNullable() bool { return f.nullable }

// IsMaybe returns true if the field value is allowed to be either null or absent.
//
// This is a simple "or" of the two properties,
// but this method is a shorthand that turns out useful often.
func (f StructField) IsMaybe() bool { return f.nullable || f.optional }

func (t TypeStruct) RepresentationStrategy() StructRepresentation {
	return t.representation
}

func (r StructRepresentation_Map) GetFieldKey(field StructField) string {
	if n, ok := r.renames[field.name]; ok {
		return n
	}
	return field.name
}

func (r StructRepresentation_Stringjoin) GetDelim() string {
	return r.sep
}

// Members returns a slice the strings which are valid inhabitants of this enum.
func (t TypeEnum) Members() []string {
	a := make([]string, len(t.members))
	for i := range t.members {
		a[i] = t.members[i]
	}
	return a
}

// Links can keep a referenced type, which is a hint only about the data on the
// other side of the link, no something that can be explicitly validated without
// loading the link

// HasReferencedType returns true if the link has a hint about the type it references
// false if it's generic
func (t TypeLink) HasReferencedType() bool {
	return t.hasReferencedType
}

// ReferencedType returns the type hint for the node on the other side of the link
func (t TypeLink) ReferencedType() Type {
	return t.referencedType
}
