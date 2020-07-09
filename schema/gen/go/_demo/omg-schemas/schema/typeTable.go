package schema

// Type is a struct embeding a NodePrototype/Type for every Node implementation in this package.
// One of its major uses is to start the construction of a value.
// You can use it like this:
//
// 		schema.Type.YourTypeName.NewBuilder().BeginMap() //...
//
// and:
//
// 		schema.Type.OtherTypeName.NewBuilder().AssignString("x") // ...
//
var Type typeSlab

type typeSlab struct {
	AnyScalar       _AnyScalar__Prototype
	AnyScalar__Repr _AnyScalar__ReprPrototype
	Bool       _Bool__Prototype
	Bool__Repr _Bool__ReprPrototype
	Bytes       _Bytes__Prototype
	Bytes__Repr _Bytes__ReprPrototype
	EnumRepresentation       _EnumRepresentation__Prototype
	EnumRepresentation__Repr _EnumRepresentation__ReprPrototype
	EnumRepresentation_Int       _EnumRepresentation_Int__Prototype
	EnumRepresentation_Int__Repr _EnumRepresentation_Int__ReprPrototype
	EnumRepresentation_String       _EnumRepresentation_String__Prototype
	EnumRepresentation_String__Repr _EnumRepresentation_String__ReprPrototype
	EnumValue       _EnumValue__Prototype
	EnumValue__Repr _EnumValue__ReprPrototype
	FieldName       _FieldName__Prototype
	FieldName__Repr _FieldName__ReprPrototype
	Float       _Float__Prototype
	Float__Repr _Float__ReprPrototype
	Int       _Int__Prototype
	Int__Repr _Int__ReprPrototype
	ListRepresentation       _ListRepresentation__Prototype
	ListRepresentation__Repr _ListRepresentation__ReprPrototype
	ListRepresentation_List       _ListRepresentation_List__Prototype
	ListRepresentation_List__Repr _ListRepresentation_List__ReprPrototype
	List__FieldName       _List__FieldName__Prototype
	List__FieldName__Repr _List__FieldName__ReprPrototype
	List__TypeName       _List__TypeName__Prototype
	List__TypeName__Repr _List__TypeName__ReprPrototype
	MapRepresentation       _MapRepresentation__Prototype
	MapRepresentation__Repr _MapRepresentation__ReprPrototype
	MapRepresentation_ListPairs       _MapRepresentation_ListPairs__Prototype
	MapRepresentation_ListPairs__Repr _MapRepresentation_ListPairs__ReprPrototype
	MapRepresentation_Map       _MapRepresentation_Map__Prototype
	MapRepresentation_Map__Repr _MapRepresentation_Map__ReprPrototype
	MapRepresentation_StringPairs       _MapRepresentation_StringPairs__Prototype
	MapRepresentation_StringPairs__Repr _MapRepresentation_StringPairs__ReprPrototype
	Map__EnumValue__Unit       _Map__EnumValue__Unit__Prototype
	Map__EnumValue__Unit__Repr _Map__EnumValue__Unit__ReprPrototype
	Map__FieldName__StructField       _Map__FieldName__StructField__Prototype
	Map__FieldName__StructField__Repr _Map__FieldName__StructField__ReprPrototype
	Map__FieldName__StructRepresentation_Map_FieldDetails       _Map__FieldName__StructRepresentation_Map_FieldDetails__Prototype
	Map__FieldName__StructRepresentation_Map_FieldDetails__Repr _Map__FieldName__StructRepresentation_Map_FieldDetails__ReprPrototype
	Map__String__TypeName       _Map__String__TypeName__Prototype
	Map__String__TypeName__Repr _Map__String__TypeName__ReprPrototype
	Map__TypeName__Int       _Map__TypeName__Int__Prototype
	Map__TypeName__Int__Repr _Map__TypeName__Int__ReprPrototype
	RepresentationKind       _RepresentationKind__Prototype
	RepresentationKind__Repr _RepresentationKind__ReprPrototype
	SchemaMap       _SchemaMap__Prototype
	SchemaMap__Repr _SchemaMap__ReprPrototype
	String       _String__Prototype
	String__Repr _String__ReprPrototype
	StructField       _StructField__Prototype
	StructField__Repr _StructField__ReprPrototype
	StructRepresentation       _StructRepresentation__Prototype
	StructRepresentation__Repr _StructRepresentation__ReprPrototype
	StructRepresentation_ListPairs       _StructRepresentation_ListPairs__Prototype
	StructRepresentation_ListPairs__Repr _StructRepresentation_ListPairs__ReprPrototype
	StructRepresentation_Map       _StructRepresentation_Map__Prototype
	StructRepresentation_Map__Repr _StructRepresentation_Map__ReprPrototype
	StructRepresentation_Map_FieldDetails       _StructRepresentation_Map_FieldDetails__Prototype
	StructRepresentation_Map_FieldDetails__Repr _StructRepresentation_Map_FieldDetails__ReprPrototype
	StructRepresentation_StringJoin       _StructRepresentation_StringJoin__Prototype
	StructRepresentation_StringJoin__Repr _StructRepresentation_StringJoin__ReprPrototype
	StructRepresentation_StringPairs       _StructRepresentation_StringPairs__Prototype
	StructRepresentation_StringPairs__Repr _StructRepresentation_StringPairs__ReprPrototype
	StructRepresentation_Tuple       _StructRepresentation_Tuple__Prototype
	StructRepresentation_Tuple__Repr _StructRepresentation_Tuple__ReprPrototype
	TypeBool       _TypeBool__Prototype
	TypeBool__Repr _TypeBool__ReprPrototype
	TypeBytes       _TypeBytes__Prototype
	TypeBytes__Repr _TypeBytes__ReprPrototype
	TypeCopy       _TypeCopy__Prototype
	TypeCopy__Repr _TypeCopy__ReprPrototype
	TypeDefn       _TypeDefn__Prototype
	TypeDefn__Repr _TypeDefn__ReprPrototype
	TypeDefnInline       _TypeDefnInline__Prototype
	TypeDefnInline__Repr _TypeDefnInline__ReprPrototype
	TypeEnum       _TypeEnum__Prototype
	TypeEnum__Repr _TypeEnum__ReprPrototype
	TypeFloat       _TypeFloat__Prototype
	TypeFloat__Repr _TypeFloat__ReprPrototype
	TypeInt       _TypeInt__Prototype
	TypeInt__Repr _TypeInt__ReprPrototype
	TypeLink       _TypeLink__Prototype
	TypeLink__Repr _TypeLink__ReprPrototype
	TypeList       _TypeList__Prototype
	TypeList__Repr _TypeList__ReprPrototype
	TypeMap       _TypeMap__Prototype
	TypeMap__Repr _TypeMap__ReprPrototype
	TypeName       _TypeName__Prototype
	TypeName__Repr _TypeName__ReprPrototype
	TypeNameOrInlineDefn       _TypeNameOrInlineDefn__Prototype
	TypeNameOrInlineDefn__Repr _TypeNameOrInlineDefn__ReprPrototype
	TypeString       _TypeString__Prototype
	TypeString__Repr _TypeString__ReprPrototype
	TypeStruct       _TypeStruct__Prototype
	TypeStruct__Repr _TypeStruct__ReprPrototype
	TypeUnion       _TypeUnion__Prototype
	TypeUnion__Repr _TypeUnion__ReprPrototype
	UnionRepresentation       _UnionRepresentation__Prototype
	UnionRepresentation__Repr _UnionRepresentation__ReprPrototype
	UnionRepresentation_BytePrefix       _UnionRepresentation_BytePrefix__Prototype
	UnionRepresentation_BytePrefix__Repr _UnionRepresentation_BytePrefix__ReprPrototype
	UnionRepresentation_Envelope       _UnionRepresentation_Envelope__Prototype
	UnionRepresentation_Envelope__Repr _UnionRepresentation_Envelope__ReprPrototype
	UnionRepresentation_Inline       _UnionRepresentation_Inline__Prototype
	UnionRepresentation_Inline__Repr _UnionRepresentation_Inline__ReprPrototype
	UnionRepresentation_Keyed       _UnionRepresentation_Keyed__Prototype
	UnionRepresentation_Keyed__Repr _UnionRepresentation_Keyed__ReprPrototype
	UnionRepresentation_Kinded       _UnionRepresentation_Kinded__Prototype
	UnionRepresentation_Kinded__Repr _UnionRepresentation_Kinded__ReprPrototype
	Unit       _Unit__Prototype
	Unit__Repr _Unit__ReprPrototype
}
