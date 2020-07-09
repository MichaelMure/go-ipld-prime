package schema

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _TypeString struct {
}
type TypeString = *_TypeString

type _TypeString__Maybe struct {
	m schema.Maybe
	v TypeString
}
type MaybeTypeString = *_TypeString__Maybe

func (m MaybeTypeString) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeTypeString) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeTypeString) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeTypeString) AsNode() ipld.Node {
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
func (m MaybeTypeString) Must() TypeString {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
)
var _ ipld.Node = (TypeString)(&_TypeString{})
var _ schema.TypedNode = (TypeString)(&_TypeString{})
func (TypeString) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n TypeString) LookupByString(key string) (ipld.Node, error) {
	switch key {
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n TypeString) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (TypeString) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.TypeString"}.LookupByIndex(0)
}
func (n TypeString) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n TypeString) MapIterator() ipld.MapIterator {
	return &_TypeString__MapItr{n, 0}
}

type _TypeString__MapItr struct {
	n TypeString
	idx  int
}

func (itr *_TypeString__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 0 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_TypeString__MapItr) Done() bool {
	return itr.idx >= 0
}

func (TypeString) ListIterator() ipld.ListIterator {
	return nil
}
func (TypeString) Length() int {
	return 0
}
func (TypeString) IsAbsent() bool {
	return false
}
func (TypeString) IsNull() bool {
	return false
}
func (TypeString) AsBool() (bool, error) {
	return mixins.Map{"schema.TypeString"}.AsBool()
}
func (TypeString) AsInt() (int, error) {
	return mixins.Map{"schema.TypeString"}.AsInt()
}
func (TypeString) AsFloat() (float64, error) {
	return mixins.Map{"schema.TypeString"}.AsFloat()
}
func (TypeString) AsString() (string, error) {
	return mixins.Map{"schema.TypeString"}.AsString()
}
func (TypeString) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.TypeString"}.AsBytes()
}
func (TypeString) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.TypeString"}.AsLink()
}
func (TypeString) Prototype() ipld.NodePrototype {
	return _TypeString__Prototype{}
}
type _TypeString__Prototype struct{}

func (_TypeString__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _TypeString__Builder
	nb.Reset()
	return &nb
}
type _TypeString__Builder struct {
	_TypeString__Assembler
}
func (nb *_TypeString__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_TypeString__Builder) Reset() {
	var w _TypeString
	var m schema.Maybe
	*nb = _TypeString__Builder{_TypeString__Assembler{w: &w, m: &m}}
}
type _TypeString__Assembler struct {
	w *_TypeString
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	}

func (na *_TypeString__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
}

var (
	fieldBits__TypeString_sufficient = 0
)
func (na *_TypeString__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_TypeString{}
	}
	return na, nil
}
func (_TypeString__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.TypeString"}.BeginList(0)
}
func (na *_TypeString__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.TypeString"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_TypeString__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.TypeString"}.AssignBool(false)
}
func (_TypeString__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.TypeString"}.AssignInt(0)
}
func (_TypeString__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.TypeString"}.AssignFloat(0)
}
func (_TypeString__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.TypeString"}.AssignString("")
}
func (_TypeString__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.TypeString"}.AssignBytes(nil)
}
func (_TypeString__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.TypeString"}.AssignLink(nil)
}
func (na *_TypeString__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_TypeString); ok {
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
		return ipld.ErrWrongKind{TypeName: "schema.TypeString", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_TypeString__Assembler) Prototype() ipld.NodePrototype {
	return _TypeString__Prototype{}
}
func (ma *_TypeString__Assembler) valueFinishTidy() bool {
	switch ma.f {
	default:
		panic("unreachable")
	}
}
func (ma *_TypeString__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"schema.TypeString", Key:&_String{k}}
	}
}
func (ma *_TypeString__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_TypeString__KeyAssembler)(ma)
}
func (ma *_TypeString__Assembler) AssembleValue() ipld.NodeAssembler {
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
	default:
		panic("unreachable")
	}
}
func (ma *_TypeString__Assembler) Finish() error {
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
func (ma *_TypeString__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_TypeString__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _TypeString__KeyAssembler _TypeString__Assembler
func (_TypeString__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"schema.TypeString.KeyAssembler"}.BeginMap(0)
}
func (_TypeString__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"schema.TypeString.KeyAssembler"}.BeginList(0)
}
func (na *_TypeString__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"schema.TypeString.KeyAssembler"}.AssignNull()
}
func (_TypeString__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"schema.TypeString.KeyAssembler"}.AssignBool(false)
}
func (_TypeString__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"schema.TypeString.KeyAssembler"}.AssignInt(0)
}
func (_TypeString__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"schema.TypeString.KeyAssembler"}.AssignFloat(0)
}
func (ka *_TypeString__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	default:
		return ipld.ErrInvalidKey{TypeName:"schema.TypeString", Key:&_String{k}}
	}
	return nil
}
func (_TypeString__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"schema.TypeString.KeyAssembler"}.AssignBytes(nil)
}
func (_TypeString__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"schema.TypeString.KeyAssembler"}.AssignLink(nil)
}
func (ka *_TypeString__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_TypeString__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (TypeString) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n TypeString) Representation() ipld.Node {
	return (*_TypeString__Repr)(n)
}
type _TypeString__Repr _TypeString
var (
)
var _ ipld.Node = &_TypeString__Repr{}
func (_TypeString__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n *_TypeString__Repr) LookupByString(key string) (ipld.Node, error) {
	switch key {
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (n *_TypeString__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (_TypeString__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.TypeString.Repr"}.LookupByIndex(0)
}
func (n _TypeString__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n *_TypeString__Repr) MapIterator() ipld.MapIterator {
	return &_TypeString__ReprMapItr{n, 0}
}

type _TypeString__ReprMapItr struct {
	n   *_TypeString__Repr
	idx int
	
}

func (itr *_TypeString__ReprMapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
if itr.idx >= 0 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_TypeString__ReprMapItr) Done() bool {
	return itr.idx >= 0
}
func (_TypeString__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_TypeString__Repr) Length() int {
	l := 0
	return l
}
func (_TypeString__Repr) IsAbsent() bool {
	return false
}
func (_TypeString__Repr) IsNull() bool {
	return false
}
func (_TypeString__Repr) AsBool() (bool, error) {
	return mixins.Map{"schema.TypeString.Repr"}.AsBool()
}
func (_TypeString__Repr) AsInt() (int, error) {
	return mixins.Map{"schema.TypeString.Repr"}.AsInt()
}
func (_TypeString__Repr) AsFloat() (float64, error) {
	return mixins.Map{"schema.TypeString.Repr"}.AsFloat()
}
func (_TypeString__Repr) AsString() (string, error) {
	return mixins.Map{"schema.TypeString.Repr"}.AsString()
}
func (_TypeString__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.TypeString.Repr"}.AsBytes()
}
func (_TypeString__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.TypeString.Repr"}.AsLink()
}
func (_TypeString__Repr) Prototype() ipld.NodePrototype {
	return _TypeString__ReprPrototype{}
}
type _TypeString__ReprPrototype struct{}

func (_TypeString__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _TypeString__ReprBuilder
	nb.Reset()
	return &nb
}
type _TypeString__ReprBuilder struct {
	_TypeString__ReprAssembler
}
func (nb *_TypeString__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_TypeString__ReprBuilder) Reset() {
	var w _TypeString
	var m schema.Maybe
	*nb = _TypeString__ReprBuilder{_TypeString__ReprAssembler{w: &w, m: &m}}
}
type _TypeString__ReprAssembler struct {
	w *_TypeString
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	}

func (na *_TypeString__ReprAssembler) reset() {
	na.state = maState_initial
	na.s = 0
}
func (na *_TypeString__ReprAssembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_TypeString{}
	}
	return na, nil
}
func (_TypeString__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.TypeString.Repr"}.BeginList(0)
}
func (na *_TypeString__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.TypeString.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_TypeString__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.TypeString.Repr"}.AssignBool(false)
}
func (_TypeString__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.TypeString.Repr"}.AssignInt(0)
}
func (_TypeString__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.TypeString.Repr"}.AssignFloat(0)
}
func (_TypeString__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.TypeString.Repr"}.AssignString("")
}
func (_TypeString__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.TypeString.Repr"}.AssignBytes(nil)
}
func (_TypeString__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.TypeString.Repr"}.AssignLink(nil)
}
func (na *_TypeString__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_TypeString); ok {
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
		return ipld.ErrWrongKind{TypeName: "schema.TypeString.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_TypeString__ReprAssembler) Prototype() ipld.NodePrototype {
	return _TypeString__ReprPrototype{}
}
func (ma *_TypeString__ReprAssembler) valueFinishTidy() bool {
	switch ma.f {
	default:
		panic("unreachable")
	}
}
func (ma *_TypeString__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"schema.TypeString.Repr", Key:&_String{k}}
	}
}
func (ma *_TypeString__ReprAssembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_TypeString__ReprKeyAssembler)(ma)
}
func (ma *_TypeString__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
	default:
		panic("unreachable")
	}
}
func (ma *_TypeString__ReprAssembler) Finish() error {
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
func (ma *_TypeString__ReprAssembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_TypeString__ReprAssembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler repr valueprototype")
}
type _TypeString__ReprKeyAssembler _TypeString__ReprAssembler
func (_TypeString__ReprKeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"schema.TypeString.Repr.KeyAssembler"}.BeginMap(0)
}
func (_TypeString__ReprKeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"schema.TypeString.Repr.KeyAssembler"}.BeginList(0)
}
func (na *_TypeString__ReprKeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"schema.TypeString.Repr.KeyAssembler"}.AssignNull()
}
func (_TypeString__ReprKeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"schema.TypeString.Repr.KeyAssembler"}.AssignBool(false)
}
func (_TypeString__ReprKeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"schema.TypeString.Repr.KeyAssembler"}.AssignInt(0)
}
func (_TypeString__ReprKeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"schema.TypeString.Repr.KeyAssembler"}.AssignFloat(0)
}
func (ka *_TypeString__ReprKeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	default:
		return ipld.ErrInvalidKey{TypeName:"schema.TypeString.Repr", Key:&_String{k}}
	}
	return nil
}
func (_TypeString__ReprKeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"schema.TypeString.Repr.KeyAssembler"}.AssignBytes(nil)
}
func (_TypeString__ReprKeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"schema.TypeString.Repr.KeyAssembler"}.AssignLink(nil)
}
func (ka *_TypeString__ReprKeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_TypeString__ReprKeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
