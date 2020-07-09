package schema

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _MapRepresentation_StringPairs struct {
	innerDelim _String
	entryDelim _String
}
type MapRepresentation_StringPairs = *_MapRepresentation_StringPairs

func (n _MapRepresentation_StringPairs) FieldInnerDelim()	String {
	return &n.innerDelim
}
func (n _MapRepresentation_StringPairs) FieldEntryDelim()	String {
	return &n.entryDelim
}
type _MapRepresentation_StringPairs__Maybe struct {
	m schema.Maybe
	v MapRepresentation_StringPairs
}
type MaybeMapRepresentation_StringPairs = *_MapRepresentation_StringPairs__Maybe

func (m MaybeMapRepresentation_StringPairs) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMapRepresentation_StringPairs) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMapRepresentation_StringPairs) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMapRepresentation_StringPairs) AsNode() ipld.Node {
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
func (m MaybeMapRepresentation_StringPairs) Must() MapRepresentation_StringPairs {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__MapRepresentation_StringPairs_InnerDelim = _String{"innerDelim"}
	fieldName__MapRepresentation_StringPairs_EntryDelim = _String{"entryDelim"}
)
var _ ipld.Node = (MapRepresentation_StringPairs)(&_MapRepresentation_StringPairs{})
var _ schema.TypedNode = (MapRepresentation_StringPairs)(&_MapRepresentation_StringPairs{})
func (MapRepresentation_StringPairs) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n MapRepresentation_StringPairs) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "innerDelim":
		return &n.innerDelim, nil
	case "entryDelim":
		return &n.entryDelim, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n MapRepresentation_StringPairs) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (MapRepresentation_StringPairs) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs"}.LookupByIndex(0)
}
func (n MapRepresentation_StringPairs) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n MapRepresentation_StringPairs) MapIterator() ipld.MapIterator {
	return &_MapRepresentation_StringPairs__MapItr{n, 0}
}

type _MapRepresentation_StringPairs__MapItr struct {
	n MapRepresentation_StringPairs
	idx  int
}

func (itr *_MapRepresentation_StringPairs__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__MapRepresentation_StringPairs_InnerDelim
		v = &itr.n.innerDelim
	case 1:
		k = &fieldName__MapRepresentation_StringPairs_EntryDelim
		v = &itr.n.entryDelim
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MapRepresentation_StringPairs__MapItr) Done() bool {
	return itr.idx >= 2
}

func (MapRepresentation_StringPairs) ListIterator() ipld.ListIterator {
	return nil
}
func (MapRepresentation_StringPairs) Length() int {
	return 2
}
func (MapRepresentation_StringPairs) IsAbsent() bool {
	return false
}
func (MapRepresentation_StringPairs) IsNull() bool {
	return false
}
func (MapRepresentation_StringPairs) AsBool() (bool, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs"}.AsBool()
}
func (MapRepresentation_StringPairs) AsInt() (int, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs"}.AsInt()
}
func (MapRepresentation_StringPairs) AsFloat() (float64, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs"}.AsFloat()
}
func (MapRepresentation_StringPairs) AsString() (string, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs"}.AsString()
}
func (MapRepresentation_StringPairs) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs"}.AsBytes()
}
func (MapRepresentation_StringPairs) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs"}.AsLink()
}
func (MapRepresentation_StringPairs) Prototype() ipld.NodePrototype {
	return _MapRepresentation_StringPairs__Prototype{}
}
type _MapRepresentation_StringPairs__Prototype struct{}

func (_MapRepresentation_StringPairs__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _MapRepresentation_StringPairs__Builder
	nb.Reset()
	return &nb
}
type _MapRepresentation_StringPairs__Builder struct {
	_MapRepresentation_StringPairs__Assembler
}
func (nb *_MapRepresentation_StringPairs__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MapRepresentation_StringPairs__Builder) Reset() {
	var w _MapRepresentation_StringPairs
	var m schema.Maybe
	*nb = _MapRepresentation_StringPairs__Builder{_MapRepresentation_StringPairs__Assembler{w: &w, m: &m}}
}
type _MapRepresentation_StringPairs__Assembler struct {
	w *_MapRepresentation_StringPairs
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_innerDelim _String__Assembler
	ca_entryDelim _String__Assembler
	}

func (na *_MapRepresentation_StringPairs__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_innerDelim.reset()
	na.ca_entryDelim.reset()
}

var (
	fieldBit__MapRepresentation_StringPairs_InnerDelim = 1 << 0
	fieldBit__MapRepresentation_StringPairs_EntryDelim = 1 << 1
	fieldBits__MapRepresentation_StringPairs_sufficient = 0 + 1 << 0 + 1 << 1
)
func (na *_MapRepresentation_StringPairs__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MapRepresentation_StringPairs{}
	}
	return na, nil
}
func (_MapRepresentation_StringPairs__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs"}.BeginList(0)
}
func (na *_MapRepresentation_StringPairs__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.MapRepresentation_StringPairs"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MapRepresentation_StringPairs__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs"}.AssignBool(false)
}
func (_MapRepresentation_StringPairs__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs"}.AssignInt(0)
}
func (_MapRepresentation_StringPairs__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs"}.AssignFloat(0)
}
func (_MapRepresentation_StringPairs__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs"}.AssignString("")
}
func (_MapRepresentation_StringPairs__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs"}.AssignBytes(nil)
}
func (_MapRepresentation_StringPairs__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs"}.AssignLink(nil)
}
func (na *_MapRepresentation_StringPairs__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MapRepresentation_StringPairs); ok {
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
		return ipld.ErrWrongKind{TypeName: "schema.MapRepresentation_StringPairs", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_MapRepresentation_StringPairs__Assembler) Prototype() ipld.NodePrototype {
	return _MapRepresentation_StringPairs__Prototype{}
}
func (ma *_MapRepresentation_StringPairs__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_innerDelim.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_entryDelim.w = nil
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
func (ma *_MapRepresentation_StringPairs__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "innerDelim":
		if ma.s & fieldBit__MapRepresentation_StringPairs_InnerDelim != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MapRepresentation_StringPairs_InnerDelim}
		}
		ma.s += fieldBit__MapRepresentation_StringPairs_InnerDelim
		ma.state = maState_midValue
		ma.ca_innerDelim.w = &ma.w.innerDelim
		ma.ca_innerDelim.m = &ma.cm
		return &ma.ca_innerDelim, nil
	case "entryDelim":
		if ma.s & fieldBit__MapRepresentation_StringPairs_EntryDelim != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MapRepresentation_StringPairs_EntryDelim}
		}
		ma.s += fieldBit__MapRepresentation_StringPairs_EntryDelim
		ma.state = maState_midValue
		ma.ca_entryDelim.w = &ma.w.entryDelim
		ma.ca_entryDelim.m = &ma.cm
		return &ma.ca_entryDelim, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"schema.MapRepresentation_StringPairs", Key:&_String{k}}
	}
}
func (ma *_MapRepresentation_StringPairs__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_MapRepresentation_StringPairs__KeyAssembler)(ma)
}
func (ma *_MapRepresentation_StringPairs__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_innerDelim.w = &ma.w.innerDelim
		ma.ca_innerDelim.m = &ma.cm
		return &ma.ca_innerDelim
	case 1:
		ma.ca_entryDelim.w = &ma.w.entryDelim
		ma.ca_entryDelim.m = &ma.cm
		return &ma.ca_entryDelim
	default:
		panic("unreachable")
	}
}
func (ma *_MapRepresentation_StringPairs__Assembler) Finish() error {
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
func (ma *_MapRepresentation_StringPairs__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_MapRepresentation_StringPairs__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _MapRepresentation_StringPairs__KeyAssembler _MapRepresentation_StringPairs__Assembler
func (_MapRepresentation_StringPairs__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.KeyAssembler"}.BeginMap(0)
}
func (_MapRepresentation_StringPairs__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.KeyAssembler"}.BeginList(0)
}
func (na *_MapRepresentation_StringPairs__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.KeyAssembler"}.AssignNull()
}
func (_MapRepresentation_StringPairs__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.KeyAssembler"}.AssignBool(false)
}
func (_MapRepresentation_StringPairs__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.KeyAssembler"}.AssignInt(0)
}
func (_MapRepresentation_StringPairs__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.KeyAssembler"}.AssignFloat(0)
}
func (ka *_MapRepresentation_StringPairs__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "innerDelim":
		if ka.s & fieldBit__MapRepresentation_StringPairs_InnerDelim != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MapRepresentation_StringPairs_InnerDelim}
		}
		ka.s += fieldBit__MapRepresentation_StringPairs_InnerDelim
		ka.state = maState_expectValue
		ka.f = 0
	case "entryDelim":
		if ka.s & fieldBit__MapRepresentation_StringPairs_EntryDelim != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MapRepresentation_StringPairs_EntryDelim}
		}
		ka.s += fieldBit__MapRepresentation_StringPairs_EntryDelim
		ka.state = maState_expectValue
		ka.f = 1
	default:
		return ipld.ErrInvalidKey{TypeName:"schema.MapRepresentation_StringPairs", Key:&_String{k}}
	}
	return nil
}
func (_MapRepresentation_StringPairs__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.KeyAssembler"}.AssignBytes(nil)
}
func (_MapRepresentation_StringPairs__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.KeyAssembler"}.AssignLink(nil)
}
func (ka *_MapRepresentation_StringPairs__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_MapRepresentation_StringPairs__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (MapRepresentation_StringPairs) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n MapRepresentation_StringPairs) Representation() ipld.Node {
	return (*_MapRepresentation_StringPairs__Repr)(n)
}
type _MapRepresentation_StringPairs__Repr _MapRepresentation_StringPairs
var (
	fieldName__MapRepresentation_StringPairs_InnerDelim_serial = _String{"innerDelim"}
	fieldName__MapRepresentation_StringPairs_EntryDelim_serial = _String{"entryDelim"}
)
var _ ipld.Node = &_MapRepresentation_StringPairs__Repr{}
func (_MapRepresentation_StringPairs__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n *_MapRepresentation_StringPairs__Repr) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "innerDelim":
		return n.innerDelim.Representation(), nil
	case "entryDelim":
		return n.entryDelim.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n *_MapRepresentation_StringPairs__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (_MapRepresentation_StringPairs__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs.Repr"}.LookupByIndex(0)
}
func (n _MapRepresentation_StringPairs__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n *_MapRepresentation_StringPairs__Repr) MapIterator() ipld.MapIterator {
	return &_MapRepresentation_StringPairs__ReprMapItr{n, 0}
}

type _MapRepresentation_StringPairs__ReprMapItr struct {
	n   *_MapRepresentation_StringPairs__Repr
	idx int
	
}

func (itr *_MapRepresentation_StringPairs__ReprMapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__MapRepresentation_StringPairs_InnerDelim_serial
		v = itr.n.innerDelim.Representation()
	case 1:
		k = &fieldName__MapRepresentation_StringPairs_EntryDelim_serial
		v = itr.n.entryDelim.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MapRepresentation_StringPairs__ReprMapItr) Done() bool {
	return itr.idx >= 2
}
func (_MapRepresentation_StringPairs__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_MapRepresentation_StringPairs__Repr) Length() int {
	l := 2
	return l
}
func (_MapRepresentation_StringPairs__Repr) IsAbsent() bool {
	return false
}
func (_MapRepresentation_StringPairs__Repr) IsNull() bool {
	return false
}
func (_MapRepresentation_StringPairs__Repr) AsBool() (bool, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs.Repr"}.AsBool()
}
func (_MapRepresentation_StringPairs__Repr) AsInt() (int, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs.Repr"}.AsInt()
}
func (_MapRepresentation_StringPairs__Repr) AsFloat() (float64, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs.Repr"}.AsFloat()
}
func (_MapRepresentation_StringPairs__Repr) AsString() (string, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs.Repr"}.AsString()
}
func (_MapRepresentation_StringPairs__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs.Repr"}.AsBytes()
}
func (_MapRepresentation_StringPairs__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.MapRepresentation_StringPairs.Repr"}.AsLink()
}
func (_MapRepresentation_StringPairs__Repr) Prototype() ipld.NodePrototype {
	return _MapRepresentation_StringPairs__ReprPrototype{}
}
type _MapRepresentation_StringPairs__ReprPrototype struct{}

func (_MapRepresentation_StringPairs__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _MapRepresentation_StringPairs__ReprBuilder
	nb.Reset()
	return &nb
}
type _MapRepresentation_StringPairs__ReprBuilder struct {
	_MapRepresentation_StringPairs__ReprAssembler
}
func (nb *_MapRepresentation_StringPairs__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MapRepresentation_StringPairs__ReprBuilder) Reset() {
	var w _MapRepresentation_StringPairs
	var m schema.Maybe
	*nb = _MapRepresentation_StringPairs__ReprBuilder{_MapRepresentation_StringPairs__ReprAssembler{w: &w, m: &m}}
}
type _MapRepresentation_StringPairs__ReprAssembler struct {
	w *_MapRepresentation_StringPairs
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_innerDelim _String__ReprAssembler
	ca_entryDelim _String__ReprAssembler
	}

func (na *_MapRepresentation_StringPairs__ReprAssembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_innerDelim.reset()
	na.ca_entryDelim.reset()
}
func (na *_MapRepresentation_StringPairs__ReprAssembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MapRepresentation_StringPairs{}
	}
	return na, nil
}
func (_MapRepresentation_StringPairs__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs.Repr"}.BeginList(0)
}
func (na *_MapRepresentation_StringPairs__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.MapRepresentation_StringPairs.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MapRepresentation_StringPairs__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs.Repr"}.AssignBool(false)
}
func (_MapRepresentation_StringPairs__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs.Repr"}.AssignInt(0)
}
func (_MapRepresentation_StringPairs__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs.Repr"}.AssignFloat(0)
}
func (_MapRepresentation_StringPairs__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs.Repr"}.AssignString("")
}
func (_MapRepresentation_StringPairs__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs.Repr"}.AssignBytes(nil)
}
func (_MapRepresentation_StringPairs__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.MapRepresentation_StringPairs.Repr"}.AssignLink(nil)
}
func (na *_MapRepresentation_StringPairs__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MapRepresentation_StringPairs); ok {
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
		return ipld.ErrWrongKind{TypeName: "schema.MapRepresentation_StringPairs.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_MapRepresentation_StringPairs__ReprAssembler) Prototype() ipld.NodePrototype {
	return _MapRepresentation_StringPairs__ReprPrototype{}
}
func (ma *_MapRepresentation_StringPairs__ReprAssembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
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
func (ma *_MapRepresentation_StringPairs__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "innerDelim":
		if ma.s & fieldBit__MapRepresentation_StringPairs_InnerDelim != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MapRepresentation_StringPairs_InnerDelim_serial}
		}
		ma.s += fieldBit__MapRepresentation_StringPairs_InnerDelim
		ma.state = maState_midValue
		ma.ca_innerDelim.w = &ma.w.innerDelim
		ma.ca_innerDelim.m = &ma.cm
		return &ma.ca_innerDelim, nil
	case "entryDelim":
		if ma.s & fieldBit__MapRepresentation_StringPairs_EntryDelim != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MapRepresentation_StringPairs_EntryDelim_serial}
		}
		ma.s += fieldBit__MapRepresentation_StringPairs_EntryDelim
		ma.state = maState_midValue
		ma.ca_entryDelim.w = &ma.w.entryDelim
		ma.ca_entryDelim.m = &ma.cm
		return &ma.ca_entryDelim, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"schema.MapRepresentation_StringPairs.Repr", Key:&_String{k}}
	}
}
func (ma *_MapRepresentation_StringPairs__ReprAssembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_MapRepresentation_StringPairs__ReprKeyAssembler)(ma)
}
func (ma *_MapRepresentation_StringPairs__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_innerDelim.w = &ma.w.innerDelim
		ma.ca_innerDelim.m = &ma.cm
		return &ma.ca_innerDelim
	case 1:
		ma.ca_entryDelim.w = &ma.w.entryDelim
		ma.ca_entryDelim.m = &ma.cm
		return &ma.ca_entryDelim
	default:
		panic("unreachable")
	}
}
func (ma *_MapRepresentation_StringPairs__ReprAssembler) Finish() error {
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
func (ma *_MapRepresentation_StringPairs__ReprAssembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_MapRepresentation_StringPairs__ReprAssembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler repr valueprototype")
}
type _MapRepresentation_StringPairs__ReprKeyAssembler _MapRepresentation_StringPairs__ReprAssembler
func (_MapRepresentation_StringPairs__ReprKeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.Repr.KeyAssembler"}.BeginMap(0)
}
func (_MapRepresentation_StringPairs__ReprKeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.Repr.KeyAssembler"}.BeginList(0)
}
func (na *_MapRepresentation_StringPairs__ReprKeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.Repr.KeyAssembler"}.AssignNull()
}
func (_MapRepresentation_StringPairs__ReprKeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.Repr.KeyAssembler"}.AssignBool(false)
}
func (_MapRepresentation_StringPairs__ReprKeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.Repr.KeyAssembler"}.AssignInt(0)
}
func (_MapRepresentation_StringPairs__ReprKeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.Repr.KeyAssembler"}.AssignFloat(0)
}
func (ka *_MapRepresentation_StringPairs__ReprKeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "innerDelim":
		if ka.s & fieldBit__MapRepresentation_StringPairs_InnerDelim != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MapRepresentation_StringPairs_InnerDelim_serial}
		}
		ka.s += fieldBit__MapRepresentation_StringPairs_InnerDelim
		ka.state = maState_expectValue
		ka.f = 0
	case "entryDelim":
		if ka.s & fieldBit__MapRepresentation_StringPairs_EntryDelim != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MapRepresentation_StringPairs_EntryDelim_serial}
		}
		ka.s += fieldBit__MapRepresentation_StringPairs_EntryDelim
		ka.state = maState_expectValue
		ka.f = 1
	default:
		return ipld.ErrInvalidKey{TypeName:"schema.MapRepresentation_StringPairs.Repr", Key:&_String{k}}
	}
	return nil
}
func (_MapRepresentation_StringPairs__ReprKeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.Repr.KeyAssembler"}.AssignBytes(nil)
}
func (_MapRepresentation_StringPairs__ReprKeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"schema.MapRepresentation_StringPairs.Repr.KeyAssembler"}.AssignLink(nil)
}
func (ka *_MapRepresentation_StringPairs__ReprKeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_MapRepresentation_StringPairs__ReprKeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
