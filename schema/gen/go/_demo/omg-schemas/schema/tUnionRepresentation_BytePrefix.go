package schema

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _UnionRepresentation_BytePrefix struct {
	discriminantTable _Map__TypeName__Int
}
type UnionRepresentation_BytePrefix = *_UnionRepresentation_BytePrefix

func (n _UnionRepresentation_BytePrefix) FieldDiscriminantTable()	Map__TypeName__Int {
	return &n.discriminantTable
}
type _UnionRepresentation_BytePrefix__Maybe struct {
	m schema.Maybe
	v UnionRepresentation_BytePrefix
}
type MaybeUnionRepresentation_BytePrefix = *_UnionRepresentation_BytePrefix__Maybe

func (m MaybeUnionRepresentation_BytePrefix) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeUnionRepresentation_BytePrefix) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeUnionRepresentation_BytePrefix) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeUnionRepresentation_BytePrefix) AsNode() ipld.Node {
	switch m.m {
		case schema.Maybe_Absent:
			return ipld.Absent
		case schema.Maybe_Null:
			return ipld.Null
		case schema.Maybe_Value:
			return m.v
		default:
			panic("unreachable")
	}
}
func (m MaybeUnionRepresentation_BytePrefix) Must() UnionRepresentation_BytePrefix {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__UnionRepresentation_BytePrefix_DiscriminantTable = _String{"discriminantTable"}
)
var _ ipld.Node = (UnionRepresentation_BytePrefix)(&_UnionRepresentation_BytePrefix{})
var _ schema.TypedNode = (UnionRepresentation_BytePrefix)(&_UnionRepresentation_BytePrefix{})
func (UnionRepresentation_BytePrefix) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n UnionRepresentation_BytePrefix) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "discriminantTable":
		return &n.discriminantTable, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n UnionRepresentation_BytePrefix) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (UnionRepresentation_BytePrefix) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix"}.LookupByIndex(0)
}
func (n UnionRepresentation_BytePrefix) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n UnionRepresentation_BytePrefix) MapIterator() ipld.MapIterator {
	return &_UnionRepresentation_BytePrefix__MapItr{n, 0}
}

type _UnionRepresentation_BytePrefix__MapItr struct {
	n UnionRepresentation_BytePrefix
	idx  int
}

func (itr *_UnionRepresentation_BytePrefix__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 1 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__UnionRepresentation_BytePrefix_DiscriminantTable
		v = &itr.n.discriminantTable
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_UnionRepresentation_BytePrefix__MapItr) Done() bool {
	return itr.idx >= 1
}

func (UnionRepresentation_BytePrefix) ListIterator() ipld.ListIterator {
	return nil
}
func (UnionRepresentation_BytePrefix) Length() int {
	return 1
}
func (UnionRepresentation_BytePrefix) IsAbsent() bool {
	return false
}
func (UnionRepresentation_BytePrefix) IsNull() bool {
	return false
}
func (UnionRepresentation_BytePrefix) AsBool() (bool, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix"}.AsBool()
}
func (UnionRepresentation_BytePrefix) AsInt() (int, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix"}.AsInt()
}
func (UnionRepresentation_BytePrefix) AsFloat() (float64, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix"}.AsFloat()
}
func (UnionRepresentation_BytePrefix) AsString() (string, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix"}.AsString()
}
func (UnionRepresentation_BytePrefix) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix"}.AsBytes()
}
func (UnionRepresentation_BytePrefix) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix"}.AsLink()
}
func (UnionRepresentation_BytePrefix) Prototype() ipld.NodePrototype {
	return _UnionRepresentation_BytePrefix__Prototype{}
}
type _UnionRepresentation_BytePrefix__Prototype struct{}

func (_UnionRepresentation_BytePrefix__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _UnionRepresentation_BytePrefix__Builder
	nb.Reset()
	return &nb
}
type _UnionRepresentation_BytePrefix__Builder struct {
	_UnionRepresentation_BytePrefix__Assembler
}
func (nb *_UnionRepresentation_BytePrefix__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_UnionRepresentation_BytePrefix__Builder) Reset() {
	var w _UnionRepresentation_BytePrefix
	var m schema.Maybe
	*nb = _UnionRepresentation_BytePrefix__Builder{_UnionRepresentation_BytePrefix__Assembler{w: &w, m: &m}}
}
type _UnionRepresentation_BytePrefix__Assembler struct {
	w *_UnionRepresentation_BytePrefix
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_discriminantTable _Map__TypeName__Int__Assembler
	}

func (na *_UnionRepresentation_BytePrefix__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_discriminantTable.reset()
}

var (
	fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable = 1 << 0
	fieldBits__UnionRepresentation_BytePrefix_sufficient = 0 + 1 << 0
)
func (na *_UnionRepresentation_BytePrefix__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_UnionRepresentation_BytePrefix{}
	}
	return na, nil
}
func (_UnionRepresentation_BytePrefix__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix"}.BeginList(0)
}
func (na *_UnionRepresentation_BytePrefix__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_UnionRepresentation_BytePrefix__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix"}.AssignBool(false)
}
func (_UnionRepresentation_BytePrefix__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix"}.AssignInt(0)
}
func (_UnionRepresentation_BytePrefix__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix"}.AssignFloat(0)
}
func (_UnionRepresentation_BytePrefix__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix"}.AssignString("")
}
func (_UnionRepresentation_BytePrefix__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix"}.AssignBytes(nil)
}
func (_UnionRepresentation_BytePrefix__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix"}.AssignLink(nil)
}
func (na *_UnionRepresentation_BytePrefix__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_UnionRepresentation_BytePrefix); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "schema.UnionRepresentation_BytePrefix", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_UnionRepresentation_BytePrefix__Assembler) Prototype() ipld.NodePrototype {
	return _UnionRepresentation_BytePrefix__Prototype{}
}
func (ma *_UnionRepresentation_BytePrefix__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_discriminantTable.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	default:
		panic("unreachable")
	}
}
func (ma *_UnionRepresentation_BytePrefix__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}
	switch k {
	case "discriminantTable":
		if ma.s & fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__UnionRepresentation_BytePrefix_DiscriminantTable}
		}
		ma.s += fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable
		ma.state = maState_midValue
		ma.ca_discriminantTable.w = &ma.w.discriminantTable
		ma.ca_discriminantTable.m = &ma.cm
		return &ma.ca_discriminantTable, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"schema.UnionRepresentation_BytePrefix", Key:&_String{k}}
	}
}
func (ma *_UnionRepresentation_BytePrefix__Assembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midKey
	return (*_UnionRepresentation_BytePrefix__KeyAssembler)(ma)
}
func (ma *_UnionRepresentation_BytePrefix__Assembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	switch ma.f {
	case 0:
		ma.ca_discriminantTable.w = &ma.w.discriminantTable
		ma.ca_discriminantTable.m = &ma.cm
		return &ma.ca_discriminantTable
	default:
		panic("unreachable")
	}
}
func (ma *_UnionRepresentation_BytePrefix__Assembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	//FIXME check if all required fields are set
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_UnionRepresentation_BytePrefix__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_UnionRepresentation_BytePrefix__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _UnionRepresentation_BytePrefix__KeyAssembler _UnionRepresentation_BytePrefix__Assembler
func (_UnionRepresentation_BytePrefix__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.KeyAssembler"}.BeginMap(0)
}
func (_UnionRepresentation_BytePrefix__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.KeyAssembler"}.BeginList(0)
}
func (na *_UnionRepresentation_BytePrefix__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.KeyAssembler"}.AssignNull()
}
func (_UnionRepresentation_BytePrefix__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.KeyAssembler"}.AssignBool(false)
}
func (_UnionRepresentation_BytePrefix__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.KeyAssembler"}.AssignInt(0)
}
func (_UnionRepresentation_BytePrefix__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.KeyAssembler"}.AssignFloat(0)
}
func (ka *_UnionRepresentation_BytePrefix__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "discriminantTable":
		if ka.s & fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__UnionRepresentation_BytePrefix_DiscriminantTable}
		}
		ka.s += fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable
		ka.state = maState_expectValue
		ka.f = 0
	default:
		return ipld.ErrInvalidKey{TypeName:"schema.UnionRepresentation_BytePrefix", Key:&_String{k}}
	}
	return nil
}
func (_UnionRepresentation_BytePrefix__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.KeyAssembler"}.AssignBytes(nil)
}
func (_UnionRepresentation_BytePrefix__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.KeyAssembler"}.AssignLink(nil)
}
func (ka *_UnionRepresentation_BytePrefix__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_UnionRepresentation_BytePrefix__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (UnionRepresentation_BytePrefix) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n UnionRepresentation_BytePrefix) Representation() ipld.Node {
	return (*_UnionRepresentation_BytePrefix__Repr)(n)
}
type _UnionRepresentation_BytePrefix__Repr _UnionRepresentation_BytePrefix
var (
	fieldName__UnionRepresentation_BytePrefix_DiscriminantTable_serial = _String{"discriminantTable"}
)
var _ ipld.Node = &_UnionRepresentation_BytePrefix__Repr{}
func (_UnionRepresentation_BytePrefix__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n *_UnionRepresentation_BytePrefix__Repr) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "discriminantTable":
		return n.discriminantTable.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n *_UnionRepresentation_BytePrefix__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (_UnionRepresentation_BytePrefix__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix.Repr"}.LookupByIndex(0)
}
func (n _UnionRepresentation_BytePrefix__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n *_UnionRepresentation_BytePrefix__Repr) MapIterator() ipld.MapIterator {
	return &_UnionRepresentation_BytePrefix__ReprMapItr{n, 0}
}

type _UnionRepresentation_BytePrefix__ReprMapItr struct {
	n   *_UnionRepresentation_BytePrefix__Repr
	idx int
	
}

func (itr *_UnionRepresentation_BytePrefix__ReprMapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
if itr.idx >= 1 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__UnionRepresentation_BytePrefix_DiscriminantTable_serial
		v = itr.n.discriminantTable.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_UnionRepresentation_BytePrefix__ReprMapItr) Done() bool {
	return itr.idx >= 1
}
func (_UnionRepresentation_BytePrefix__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_UnionRepresentation_BytePrefix__Repr) Length() int {
	l := 1
	return l
}
func (_UnionRepresentation_BytePrefix__Repr) IsAbsent() bool {
	return false
}
func (_UnionRepresentation_BytePrefix__Repr) IsNull() bool {
	return false
}
func (_UnionRepresentation_BytePrefix__Repr) AsBool() (bool, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix.Repr"}.AsBool()
}
func (_UnionRepresentation_BytePrefix__Repr) AsInt() (int, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix.Repr"}.AsInt()
}
func (_UnionRepresentation_BytePrefix__Repr) AsFloat() (float64, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix.Repr"}.AsFloat()
}
func (_UnionRepresentation_BytePrefix__Repr) AsString() (string, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix.Repr"}.AsString()
}
func (_UnionRepresentation_BytePrefix__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix.Repr"}.AsBytes()
}
func (_UnionRepresentation_BytePrefix__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.UnionRepresentation_BytePrefix.Repr"}.AsLink()
}
func (_UnionRepresentation_BytePrefix__Repr) Prototype() ipld.NodePrototype {
	return _UnionRepresentation_BytePrefix__ReprPrototype{}
}
type _UnionRepresentation_BytePrefix__ReprPrototype struct{}

func (_UnionRepresentation_BytePrefix__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _UnionRepresentation_BytePrefix__ReprBuilder
	nb.Reset()
	return &nb
}
type _UnionRepresentation_BytePrefix__ReprBuilder struct {
	_UnionRepresentation_BytePrefix__ReprAssembler
}
func (nb *_UnionRepresentation_BytePrefix__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_UnionRepresentation_BytePrefix__ReprBuilder) Reset() {
	var w _UnionRepresentation_BytePrefix
	var m schema.Maybe
	*nb = _UnionRepresentation_BytePrefix__ReprBuilder{_UnionRepresentation_BytePrefix__ReprAssembler{w: &w, m: &m}}
}
type _UnionRepresentation_BytePrefix__ReprAssembler struct {
	w *_UnionRepresentation_BytePrefix
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_discriminantTable _Map__TypeName__Int__ReprAssembler
	}

func (na *_UnionRepresentation_BytePrefix__ReprAssembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_discriminantTable.reset()
}
func (na *_UnionRepresentation_BytePrefix__ReprAssembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_UnionRepresentation_BytePrefix{}
	}
	return na, nil
}
func (_UnionRepresentation_BytePrefix__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix.Repr"}.BeginList(0)
}
func (na *_UnionRepresentation_BytePrefix__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_UnionRepresentation_BytePrefix__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix.Repr"}.AssignBool(false)
}
func (_UnionRepresentation_BytePrefix__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix.Repr"}.AssignInt(0)
}
func (_UnionRepresentation_BytePrefix__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix.Repr"}.AssignFloat(0)
}
func (_UnionRepresentation_BytePrefix__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix.Repr"}.AssignString("")
}
func (_UnionRepresentation_BytePrefix__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix.Repr"}.AssignBytes(nil)
}
func (_UnionRepresentation_BytePrefix__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.UnionRepresentation_BytePrefix.Repr"}.AssignLink(nil)
}
func (na *_UnionRepresentation_BytePrefix__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_UnionRepresentation_BytePrefix); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "schema.UnionRepresentation_BytePrefix.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_UnionRepresentation_BytePrefix__ReprAssembler) Prototype() ipld.NodePrototype {
	return _UnionRepresentation_BytePrefix__ReprPrototype{}
}
func (ma *_UnionRepresentation_BytePrefix__ReprAssembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	default:
		panic("unreachable")
	}
}
func (ma *_UnionRepresentation_BytePrefix__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}
	switch k {
	case "discriminantTable":
		if ma.s & fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__UnionRepresentation_BytePrefix_DiscriminantTable_serial}
		}
		ma.s += fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable
		ma.state = maState_midValue
		ma.ca_discriminantTable.w = &ma.w.discriminantTable
		ma.ca_discriminantTable.m = &ma.cm
		return &ma.ca_discriminantTable, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"schema.UnionRepresentation_BytePrefix.Repr", Key:&_String{k}}
	}
}
func (ma *_UnionRepresentation_BytePrefix__ReprAssembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midKey
	return (*_UnionRepresentation_BytePrefix__ReprKeyAssembler)(ma)
}
func (ma *_UnionRepresentation_BytePrefix__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	switch ma.f {
	case 0:
		ma.ca_discriminantTable.w = &ma.w.discriminantTable
		ma.ca_discriminantTable.m = &ma.cm
		return &ma.ca_discriminantTable
	default:
		panic("unreachable")
	}
}
func (ma *_UnionRepresentation_BytePrefix__ReprAssembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	//FIXME check if all required fields are set
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_UnionRepresentation_BytePrefix__ReprAssembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_UnionRepresentation_BytePrefix__ReprAssembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler repr valueprototype")
}
type _UnionRepresentation_BytePrefix__ReprKeyAssembler _UnionRepresentation_BytePrefix__ReprAssembler
func (_UnionRepresentation_BytePrefix__ReprKeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.Repr.KeyAssembler"}.BeginMap(0)
}
func (_UnionRepresentation_BytePrefix__ReprKeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.Repr.KeyAssembler"}.BeginList(0)
}
func (na *_UnionRepresentation_BytePrefix__ReprKeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.Repr.KeyAssembler"}.AssignNull()
}
func (_UnionRepresentation_BytePrefix__ReprKeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.Repr.KeyAssembler"}.AssignBool(false)
}
func (_UnionRepresentation_BytePrefix__ReprKeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.Repr.KeyAssembler"}.AssignInt(0)
}
func (_UnionRepresentation_BytePrefix__ReprKeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.Repr.KeyAssembler"}.AssignFloat(0)
}
func (ka *_UnionRepresentation_BytePrefix__ReprKeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "discriminantTable":
		if ka.s & fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__UnionRepresentation_BytePrefix_DiscriminantTable_serial}
		}
		ka.s += fieldBit__UnionRepresentation_BytePrefix_DiscriminantTable
		ka.state = maState_expectValue
		ka.f = 0
	default:
		return ipld.ErrInvalidKey{TypeName:"schema.UnionRepresentation_BytePrefix.Repr", Key:&_String{k}}
	}
	return nil
}
func (_UnionRepresentation_BytePrefix__ReprKeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.Repr.KeyAssembler"}.AssignBytes(nil)
}
func (_UnionRepresentation_BytePrefix__ReprKeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"schema.UnionRepresentation_BytePrefix.Repr.KeyAssembler"}.AssignLink(nil)
}
func (ka *_UnionRepresentation_BytePrefix__ReprKeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_UnionRepresentation_BytePrefix__ReprKeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
