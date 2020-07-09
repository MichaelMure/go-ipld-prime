package schema

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _TypeEnum struct {
	typ _Map__EnumValue__Unit
	representation _EnumRepresentation
}
type TypeEnum = *_TypeEnum

func (n _TypeEnum) FieldTyp()	Map__EnumValue__Unit {
	return &n.typ
}
func (n _TypeEnum) FieldRepresentation()	EnumRepresentation {
	return &n.representation
}
type _TypeEnum__Maybe struct {
	m schema.Maybe
	v TypeEnum
}
type MaybeTypeEnum = *_TypeEnum__Maybe

func (m MaybeTypeEnum) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeTypeEnum) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeTypeEnum) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeTypeEnum) AsNode() ipld.Node {
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
func (m MaybeTypeEnum) Must() TypeEnum {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__TypeEnum_Typ = _String{"typ"}
	fieldName__TypeEnum_Representation = _String{"representation"}
)
var _ ipld.Node = (TypeEnum)(&_TypeEnum{})
var _ schema.TypedNode = (TypeEnum)(&_TypeEnum{})
func (TypeEnum) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n TypeEnum) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "typ":
		return &n.typ, nil
	case "representation":
		return &n.representation, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n TypeEnum) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (TypeEnum) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.TypeEnum"}.LookupByIndex(0)
}
func (n TypeEnum) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n TypeEnum) MapIterator() ipld.MapIterator {
	return &_TypeEnum__MapItr{n, 0}
}

type _TypeEnum__MapItr struct {
	n TypeEnum
	idx  int
}

func (itr *_TypeEnum__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__TypeEnum_Typ
		v = &itr.n.typ
	case 1:
		k = &fieldName__TypeEnum_Representation
		v = &itr.n.representation
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_TypeEnum__MapItr) Done() bool {
	return itr.idx >= 2
}

func (TypeEnum) ListIterator() ipld.ListIterator {
	return nil
}
func (TypeEnum) Length() int {
	return 2
}
func (TypeEnum) IsAbsent() bool {
	return false
}
func (TypeEnum) IsNull() bool {
	return false
}
func (TypeEnum) AsBool() (bool, error) {
	return mixins.Map{"schema.TypeEnum"}.AsBool()
}
func (TypeEnum) AsInt() (int, error) {
	return mixins.Map{"schema.TypeEnum"}.AsInt()
}
func (TypeEnum) AsFloat() (float64, error) {
	return mixins.Map{"schema.TypeEnum"}.AsFloat()
}
func (TypeEnum) AsString() (string, error) {
	return mixins.Map{"schema.TypeEnum"}.AsString()
}
func (TypeEnum) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.TypeEnum"}.AsBytes()
}
func (TypeEnum) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.TypeEnum"}.AsLink()
}
func (TypeEnum) Prototype() ipld.NodePrototype {
	return _TypeEnum__Prototype{}
}
type _TypeEnum__Prototype struct{}

func (_TypeEnum__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _TypeEnum__Builder
	nb.Reset()
	return &nb
}
type _TypeEnum__Builder struct {
	_TypeEnum__Assembler
}
func (nb *_TypeEnum__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_TypeEnum__Builder) Reset() {
	var w _TypeEnum
	var m schema.Maybe
	*nb = _TypeEnum__Builder{_TypeEnum__Assembler{w: &w, m: &m}}
}
type _TypeEnum__Assembler struct {
	w *_TypeEnum
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_typ _Map__EnumValue__Unit__Assembler
	ca_representation _EnumRepresentation__Assembler
	}

func (na *_TypeEnum__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_typ.reset()
	na.ca_representation.reset()
}

var (
	fieldBit__TypeEnum_Typ = 1 << 0
	fieldBit__TypeEnum_Representation = 1 << 1
	fieldBits__TypeEnum_sufficient = 0 + 1 << 0 + 1 << 1
)
func (na *_TypeEnum__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_TypeEnum{}
	}
	return na, nil
}
func (_TypeEnum__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.TypeEnum"}.BeginList(0)
}
func (na *_TypeEnum__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.TypeEnum"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_TypeEnum__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.TypeEnum"}.AssignBool(false)
}
func (_TypeEnum__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.TypeEnum"}.AssignInt(0)
}
func (_TypeEnum__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.TypeEnum"}.AssignFloat(0)
}
func (_TypeEnum__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.TypeEnum"}.AssignString("")
}
func (_TypeEnum__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.TypeEnum"}.AssignBytes(nil)
}
func (_TypeEnum__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.TypeEnum"}.AssignLink(nil)
}
func (na *_TypeEnum__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_TypeEnum); ok {
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
		return ipld.ErrWrongKind{TypeName: "schema.TypeEnum", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_TypeEnum__Assembler) Prototype() ipld.NodePrototype {
	return _TypeEnum__Prototype{}
}
func (ma *_TypeEnum__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_typ.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_representation.w = nil
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
func (ma *_TypeEnum__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "typ":
		if ma.s & fieldBit__TypeEnum_Typ != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__TypeEnum_Typ}
		}
		ma.s += fieldBit__TypeEnum_Typ
		ma.state = maState_midValue
		ma.ca_typ.w = &ma.w.typ
		ma.ca_typ.m = &ma.cm
		return &ma.ca_typ, nil
	case "representation":
		if ma.s & fieldBit__TypeEnum_Representation != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__TypeEnum_Representation}
		}
		ma.s += fieldBit__TypeEnum_Representation
		ma.state = maState_midValue
		ma.ca_representation.w = &ma.w.representation
		ma.ca_representation.m = &ma.cm
		return &ma.ca_representation, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"schema.TypeEnum", Key:&_String{k}}
	}
}
func (ma *_TypeEnum__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_TypeEnum__KeyAssembler)(ma)
}
func (ma *_TypeEnum__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_typ.w = &ma.w.typ
		ma.ca_typ.m = &ma.cm
		return &ma.ca_typ
	case 1:
		ma.ca_representation.w = &ma.w.representation
		ma.ca_representation.m = &ma.cm
		return &ma.ca_representation
	default:
		panic("unreachable")
	}
}
func (ma *_TypeEnum__Assembler) Finish() error {
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
func (ma *_TypeEnum__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_TypeEnum__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _TypeEnum__KeyAssembler _TypeEnum__Assembler
func (_TypeEnum__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"schema.TypeEnum.KeyAssembler"}.BeginMap(0)
}
func (_TypeEnum__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"schema.TypeEnum.KeyAssembler"}.BeginList(0)
}
func (na *_TypeEnum__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"schema.TypeEnum.KeyAssembler"}.AssignNull()
}
func (_TypeEnum__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"schema.TypeEnum.KeyAssembler"}.AssignBool(false)
}
func (_TypeEnum__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"schema.TypeEnum.KeyAssembler"}.AssignInt(0)
}
func (_TypeEnum__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"schema.TypeEnum.KeyAssembler"}.AssignFloat(0)
}
func (ka *_TypeEnum__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "typ":
		if ka.s & fieldBit__TypeEnum_Typ != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__TypeEnum_Typ}
		}
		ka.s += fieldBit__TypeEnum_Typ
		ka.state = maState_expectValue
		ka.f = 0
	case "representation":
		if ka.s & fieldBit__TypeEnum_Representation != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__TypeEnum_Representation}
		}
		ka.s += fieldBit__TypeEnum_Representation
		ka.state = maState_expectValue
		ka.f = 1
	default:
		return ipld.ErrInvalidKey{TypeName:"schema.TypeEnum", Key:&_String{k}}
	}
	return nil
}
func (_TypeEnum__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"schema.TypeEnum.KeyAssembler"}.AssignBytes(nil)
}
func (_TypeEnum__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"schema.TypeEnum.KeyAssembler"}.AssignLink(nil)
}
func (ka *_TypeEnum__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_TypeEnum__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (TypeEnum) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n TypeEnum) Representation() ipld.Node {
	return (*_TypeEnum__Repr)(n)
}
type _TypeEnum__Repr _TypeEnum
var (
	fieldName__TypeEnum_Typ_serial = _String{"typ"}
	fieldName__TypeEnum_Representation_serial = _String{"representation"}
)
var _ ipld.Node = &_TypeEnum__Repr{}
func (_TypeEnum__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n *_TypeEnum__Repr) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "typ":
		return n.typ.Representation(), nil
	case "representation":
		return n.representation.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n *_TypeEnum__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (_TypeEnum__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.TypeEnum.Repr"}.LookupByIndex(0)
}
func (n _TypeEnum__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n *_TypeEnum__Repr) MapIterator() ipld.MapIterator {
	return &_TypeEnum__ReprMapItr{n, 0}
}

type _TypeEnum__ReprMapItr struct {
	n   *_TypeEnum__Repr
	idx int
	
}

func (itr *_TypeEnum__ReprMapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__TypeEnum_Typ_serial
		v = itr.n.typ.Representation()
	case 1:
		k = &fieldName__TypeEnum_Representation_serial
		v = itr.n.representation.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_TypeEnum__ReprMapItr) Done() bool {
	return itr.idx >= 2
}
func (_TypeEnum__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_TypeEnum__Repr) Length() int {
	l := 2
	return l
}
func (_TypeEnum__Repr) IsAbsent() bool {
	return false
}
func (_TypeEnum__Repr) IsNull() bool {
	return false
}
func (_TypeEnum__Repr) AsBool() (bool, error) {
	return mixins.Map{"schema.TypeEnum.Repr"}.AsBool()
}
func (_TypeEnum__Repr) AsInt() (int, error) {
	return mixins.Map{"schema.TypeEnum.Repr"}.AsInt()
}
func (_TypeEnum__Repr) AsFloat() (float64, error) {
	return mixins.Map{"schema.TypeEnum.Repr"}.AsFloat()
}
func (_TypeEnum__Repr) AsString() (string, error) {
	return mixins.Map{"schema.TypeEnum.Repr"}.AsString()
}
func (_TypeEnum__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.TypeEnum.Repr"}.AsBytes()
}
func (_TypeEnum__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.TypeEnum.Repr"}.AsLink()
}
func (_TypeEnum__Repr) Prototype() ipld.NodePrototype {
	return _TypeEnum__ReprPrototype{}
}
type _TypeEnum__ReprPrototype struct{}

func (_TypeEnum__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _TypeEnum__ReprBuilder
	nb.Reset()
	return &nb
}
type _TypeEnum__ReprBuilder struct {
	_TypeEnum__ReprAssembler
}
func (nb *_TypeEnum__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_TypeEnum__ReprBuilder) Reset() {
	var w _TypeEnum
	var m schema.Maybe
	*nb = _TypeEnum__ReprBuilder{_TypeEnum__ReprAssembler{w: &w, m: &m}}
}
type _TypeEnum__ReprAssembler struct {
	w *_TypeEnum
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_typ _Map__EnumValue__Unit__ReprAssembler
	ca_representation _EnumRepresentation__ReprAssembler
	}

func (na *_TypeEnum__ReprAssembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_typ.reset()
	na.ca_representation.reset()
}
func (na *_TypeEnum__ReprAssembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_TypeEnum{}
	}
	return na, nil
}
func (_TypeEnum__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.TypeEnum.Repr"}.BeginList(0)
}
func (na *_TypeEnum__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.TypeEnum.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_TypeEnum__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.TypeEnum.Repr"}.AssignBool(false)
}
func (_TypeEnum__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.TypeEnum.Repr"}.AssignInt(0)
}
func (_TypeEnum__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.TypeEnum.Repr"}.AssignFloat(0)
}
func (_TypeEnum__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.TypeEnum.Repr"}.AssignString("")
}
func (_TypeEnum__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.TypeEnum.Repr"}.AssignBytes(nil)
}
func (_TypeEnum__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.TypeEnum.Repr"}.AssignLink(nil)
}
func (na *_TypeEnum__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_TypeEnum); ok {
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
		return ipld.ErrWrongKind{TypeName: "schema.TypeEnum.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_TypeEnum__ReprAssembler) Prototype() ipld.NodePrototype {
	return _TypeEnum__ReprPrototype{}
}
func (ma *_TypeEnum__ReprAssembler) valueFinishTidy() bool {
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
func (ma *_TypeEnum__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "typ":
		if ma.s & fieldBit__TypeEnum_Typ != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__TypeEnum_Typ_serial}
		}
		ma.s += fieldBit__TypeEnum_Typ
		ma.state = maState_midValue
		ma.ca_typ.w = &ma.w.typ
		ma.ca_typ.m = &ma.cm
		return &ma.ca_typ, nil
	case "representation":
		if ma.s & fieldBit__TypeEnum_Representation != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__TypeEnum_Representation_serial}
		}
		ma.s += fieldBit__TypeEnum_Representation
		ma.state = maState_midValue
		ma.ca_representation.w = &ma.w.representation
		ma.ca_representation.m = &ma.cm
		return &ma.ca_representation, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"schema.TypeEnum.Repr", Key:&_String{k}}
	}
}
func (ma *_TypeEnum__ReprAssembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_TypeEnum__ReprKeyAssembler)(ma)
}
func (ma *_TypeEnum__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_typ.w = &ma.w.typ
		ma.ca_typ.m = &ma.cm
		return &ma.ca_typ
	case 1:
		ma.ca_representation.w = &ma.w.representation
		ma.ca_representation.m = &ma.cm
		return &ma.ca_representation
	default:
		panic("unreachable")
	}
}
func (ma *_TypeEnum__ReprAssembler) Finish() error {
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
func (ma *_TypeEnum__ReprAssembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_TypeEnum__ReprAssembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler repr valueprototype")
}
type _TypeEnum__ReprKeyAssembler _TypeEnum__ReprAssembler
func (_TypeEnum__ReprKeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"schema.TypeEnum.Repr.KeyAssembler"}.BeginMap(0)
}
func (_TypeEnum__ReprKeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"schema.TypeEnum.Repr.KeyAssembler"}.BeginList(0)
}
func (na *_TypeEnum__ReprKeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"schema.TypeEnum.Repr.KeyAssembler"}.AssignNull()
}
func (_TypeEnum__ReprKeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"schema.TypeEnum.Repr.KeyAssembler"}.AssignBool(false)
}
func (_TypeEnum__ReprKeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"schema.TypeEnum.Repr.KeyAssembler"}.AssignInt(0)
}
func (_TypeEnum__ReprKeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"schema.TypeEnum.Repr.KeyAssembler"}.AssignFloat(0)
}
func (ka *_TypeEnum__ReprKeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "typ":
		if ka.s & fieldBit__TypeEnum_Typ != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__TypeEnum_Typ_serial}
		}
		ka.s += fieldBit__TypeEnum_Typ
		ka.state = maState_expectValue
		ka.f = 0
	case "representation":
		if ka.s & fieldBit__TypeEnum_Representation != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__TypeEnum_Representation_serial}
		}
		ka.s += fieldBit__TypeEnum_Representation
		ka.state = maState_expectValue
		ka.f = 1
	default:
		return ipld.ErrInvalidKey{TypeName:"schema.TypeEnum.Repr", Key:&_String{k}}
	}
	return nil
}
func (_TypeEnum__ReprKeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"schema.TypeEnum.Repr.KeyAssembler"}.AssignBytes(nil)
}
func (_TypeEnum__ReprKeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"schema.TypeEnum.Repr.KeyAssembler"}.AssignLink(nil)
}
func (ka *_TypeEnum__ReprKeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_TypeEnum__ReprKeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
