package schema

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Map__TypeName__Int struct {
	m map[_String]*_Int
	t []_Map__TypeName__Int__entry
}
type Map__TypeName__Int = *_Map__TypeName__Int
type _Map__TypeName__Int__entry struct {
	k _String
	v _Int
}
func (n *_Map__TypeName__Int) LookupMaybe(k String) MaybeInt {
	v, ok := n.m[*k]
	if !ok {
		return &_Map__TypeName__Int__valueAbsent
	}
	return &_Int__Maybe{
		m: schema.Maybe_Value,
		v: v,
	}
}

var _Map__TypeName__Int__valueAbsent = _Int__Maybe{m:schema.Maybe_Absent}

// TODO generate also a plain Lookup method that doesn't box and alloc if this type contains non-nullable values!
type _Map__TypeName__Int__Maybe struct {
	m schema.Maybe
	v Map__TypeName__Int
}
type MaybeMap__TypeName__Int = *_Map__TypeName__Int__Maybe

func (m MaybeMap__TypeName__Int) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMap__TypeName__Int) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMap__TypeName__Int) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMap__TypeName__Int) AsNode() ipld.Node {
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
func (m MaybeMap__TypeName__Int) Must() Map__TypeName__Int {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var _ ipld.Node = (Map__TypeName__Int)(&_Map__TypeName__Int{})
var _ schema.TypedNode = (Map__TypeName__Int)(&_Map__TypeName__Int{})
func (Map__TypeName__Int) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n Map__TypeName__Int) LookupByString(k string) (ipld.Node, error) {
	var k2 _String
	if err := (_String__Prototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	v, exists := n.m[k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k)}
	}
	return v, nil
}
func (n Map__TypeName__Int) LookupByNode(k ipld.Node) (ipld.Node, error) {
	k2, ok := k.(String)
	if !ok {
		panic("todo invalid key type error")
		// 'ipld.ErrInvalidKey{TypeName:"schema.Map__TypeName__Int", Key:&_String{k}}' doesn't quite cut it: need room to explain the type, and it's not guaranteed k can be turned into a string at all
	}
	v, exists := n.m[*k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k2.String())}
	}
	return v, nil
}
func (Map__TypeName__Int) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.Map__TypeName__Int"}.LookupByIndex(0)
}
func (n Map__TypeName__Int) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n Map__TypeName__Int) MapIterator() ipld.MapIterator {
	return &_Map__TypeName__Int__MapItr{n, 0}
}

type _Map__TypeName__Int__MapItr struct {
	n Map__TypeName__Int
	idx  int
}

func (itr *_Map__TypeName__Int__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.t) {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	x := &itr.n.t[itr.idx]
	k = &x.k
	v = &x.v
	itr.idx++
	return
}
func (itr *_Map__TypeName__Int__MapItr) Done() bool {
	return itr.idx >= len(itr.n.t)
}

func (Map__TypeName__Int) ListIterator() ipld.ListIterator {
	return nil
}
func (n Map__TypeName__Int) Length() int {
	return len(n.t)
}
func (Map__TypeName__Int) IsAbsent() bool {
	return false
}
func (Map__TypeName__Int) IsNull() bool {
	return false
}
func (Map__TypeName__Int) AsBool() (bool, error) {
	return mixins.Map{"schema.Map__TypeName__Int"}.AsBool()
}
func (Map__TypeName__Int) AsInt() (int, error) {
	return mixins.Map{"schema.Map__TypeName__Int"}.AsInt()
}
func (Map__TypeName__Int) AsFloat() (float64, error) {
	return mixins.Map{"schema.Map__TypeName__Int"}.AsFloat()
}
func (Map__TypeName__Int) AsString() (string, error) {
	return mixins.Map{"schema.Map__TypeName__Int"}.AsString()
}
func (Map__TypeName__Int) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.Map__TypeName__Int"}.AsBytes()
}
func (Map__TypeName__Int) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.Map__TypeName__Int"}.AsLink()
}
func (Map__TypeName__Int) Prototype() ipld.NodePrototype {
	return _Map__TypeName__Int__Prototype{}
}
type _Map__TypeName__Int__Prototype struct{}

func (_Map__TypeName__Int__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Map__TypeName__Int__Builder
	nb.Reset()
	return &nb
}
type _Map__TypeName__Int__Builder struct {
	_Map__TypeName__Int__Assembler
}
func (nb *_Map__TypeName__Int__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Map__TypeName__Int__Builder) Reset() {
	var w _Map__TypeName__Int
	var m schema.Maybe
	*nb = _Map__TypeName__Int__Builder{_Map__TypeName__Int__Assembler{w: &w, m: &m}}
}
type _Map__TypeName__Int__Assembler struct {
	w *_Map__TypeName__Int
	m *schema.Maybe
	state maState

	cm schema.Maybe
	ka _String__Assembler
	va _Int__Assembler
}

func (na *_Map__TypeName__Int__Assembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Map__TypeName__Int__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_Map__TypeName__Int{}
	}
	na.w.m = make(map[_String]*_Int, sizeHint)
	na.w.t = make([]_Map__TypeName__Int__entry, 0, sizeHint)
	return na, nil
}
func (_Map__TypeName__Int__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.Map__TypeName__Int"}.BeginList(0)
}
func (na *_Map__TypeName__Int__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.Map__TypeName__Int"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Map__TypeName__Int__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int"}.AssignBool(false)
}
func (_Map__TypeName__Int__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int"}.AssignInt(0)
}
func (_Map__TypeName__Int__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int"}.AssignFloat(0)
}
func (_Map__TypeName__Int__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int"}.AssignString("")
}
func (_Map__TypeName__Int__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int"}.AssignBytes(nil)
}
func (_Map__TypeName__Int__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int"}.AssignLink(nil)
}
func (na *_Map__TypeName__Int__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Map__TypeName__Int); ok {
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
		return ipld.ErrWrongKind{TypeName: "schema.Map__TypeName__Int", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_Map__TypeName__Int__Assembler) Prototype() ipld.NodePrototype {
	return _Map__TypeName__Int__Prototype{}
}
func (ma *_Map__TypeName__Int__Assembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.w = &tz.v
		ma.va.m = &ma.cm
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__TypeName__Int__Assembler) valueFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.va.w = nil
		ma.cm = schema.Maybe_Absent
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__TypeName__Int__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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

	var k2 _String
	if err := (_String__Prototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Map__TypeName__Int__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.w = &tz.v
	ma.va.m = &ma.cm
	return &ma.va, nil
}
func (ma *_Map__TypeName__Int__Assembler) AssembleKey() ipld.NodeAssembler {
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
	ma.w.t = append(ma.w.t, _Map__TypeName__Int__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Map__TypeName__Int__Assembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		if !ma.keyFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
		} // if tidy success: carry on
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	return &ma.va
}
func (ma *_Map__TypeName__Int__Assembler) Finish() error {
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
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Map__TypeName__Int__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_Map__TypeName__Int__Assembler) ValuePrototype(_ string) ipld.NodePrototype {
	return _Int__Prototype{}
}
func (Map__TypeName__Int) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Map__TypeName__Int) Representation() ipld.Node {
	return (*_Map__TypeName__Int__Repr)(n)
}
type _Map__TypeName__Int__Repr _Map__TypeName__Int
var _ ipld.Node = &_Map__TypeName__Int__Repr{}
func (_Map__TypeName__Int__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (nr *_Map__TypeName__Int__Repr) LookupByString(k string) (ipld.Node, error) {
	v, err := (Map__TypeName__Int)(nr).LookupByString(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Int).Representation(), nil
}
func (nr *_Map__TypeName__Int__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (Map__TypeName__Int)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Int).Representation(), nil
}
func (_Map__TypeName__Int__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"schema.Map__TypeName__Int.Repr"}.LookupByIndex(0)
}
func (n _Map__TypeName__Int__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (nr *_Map__TypeName__Int__Repr) MapIterator() ipld.MapIterator {
	return &_Map__TypeName__Int__ReprMapItr{(Map__TypeName__Int)(nr), 0}
}

type _Map__TypeName__Int__ReprMapItr _Map__TypeName__Int__MapItr

func (itr *_Map__TypeName__Int__ReprMapItr) Next() (k ipld.Node, v ipld.Node, err error) {
	k, v, err = (*_Map__TypeName__Int__MapItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return k, v.(Int).Representation(), nil
}
func (itr *_Map__TypeName__Int__ReprMapItr) Done() bool {
	return (*_Map__TypeName__Int__MapItr)(itr).Done()
}

func (_Map__TypeName__Int__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_Map__TypeName__Int__Repr) Length() int {
	return len(rn.t)
}
func (_Map__TypeName__Int__Repr) IsAbsent() bool {
	return false
}
func (_Map__TypeName__Int__Repr) IsNull() bool {
	return false
}
func (_Map__TypeName__Int__Repr) AsBool() (bool, error) {
	return mixins.Map{"schema.Map__TypeName__Int.Repr"}.AsBool()
}
func (_Map__TypeName__Int__Repr) AsInt() (int, error) {
	return mixins.Map{"schema.Map__TypeName__Int.Repr"}.AsInt()
}
func (_Map__TypeName__Int__Repr) AsFloat() (float64, error) {
	return mixins.Map{"schema.Map__TypeName__Int.Repr"}.AsFloat()
}
func (_Map__TypeName__Int__Repr) AsString() (string, error) {
	return mixins.Map{"schema.Map__TypeName__Int.Repr"}.AsString()
}
func (_Map__TypeName__Int__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"schema.Map__TypeName__Int.Repr"}.AsBytes()
}
func (_Map__TypeName__Int__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"schema.Map__TypeName__Int.Repr"}.AsLink()
}
func (_Map__TypeName__Int__Repr) Prototype() ipld.NodePrototype {
	return _Map__TypeName__Int__ReprPrototype{}
}
type _Map__TypeName__Int__ReprPrototype struct{}

func (_Map__TypeName__Int__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _Map__TypeName__Int__ReprBuilder
	nb.Reset()
	return &nb
}
type _Map__TypeName__Int__ReprBuilder struct {
	_Map__TypeName__Int__ReprAssembler
}
func (nb *_Map__TypeName__Int__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Map__TypeName__Int__ReprBuilder) Reset() {
	var w _Map__TypeName__Int
	var m schema.Maybe
	*nb = _Map__TypeName__Int__ReprBuilder{_Map__TypeName__Int__ReprAssembler{w: &w, m: &m}}
}
type _Map__TypeName__Int__ReprAssembler struct {
	w *_Map__TypeName__Int
	m *schema.Maybe
	state maState

	cm schema.Maybe
	ka _String__ReprAssembler
	va _Int__ReprAssembler
}

func (na *_Map__TypeName__Int__ReprAssembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Map__TypeName__Int__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_Map__TypeName__Int{}
	}
	na.w.m = make(map[_String]*_Int, sizeHint)
	na.w.t = make([]_Map__TypeName__Int__entry, 0, sizeHint)
	return na, nil
}
func (_Map__TypeName__Int__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"schema.Map__TypeName__Int.Repr"}.BeginList(0)
}
func (na *_Map__TypeName__Int__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"schema.Map__TypeName__Int.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Map__TypeName__Int__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int.Repr"}.AssignBool(false)
}
func (_Map__TypeName__Int__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int.Repr"}.AssignInt(0)
}
func (_Map__TypeName__Int__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int.Repr"}.AssignFloat(0)
}
func (_Map__TypeName__Int__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int.Repr"}.AssignString("")
}
func (_Map__TypeName__Int__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int.Repr"}.AssignBytes(nil)
}
func (_Map__TypeName__Int__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"schema.Map__TypeName__Int.Repr"}.AssignLink(nil)
}
func (na *_Map__TypeName__Int__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Map__TypeName__Int); ok {
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
		return ipld.ErrWrongKind{TypeName: "schema.Map__TypeName__Int.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_Map__TypeName__Int__ReprAssembler) Prototype() ipld.NodePrototype {
	return _Map__TypeName__Int__ReprPrototype{}
}
func (ma *_Map__TypeName__Int__ReprAssembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.w = &tz.v
		ma.va.m = &ma.cm
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__TypeName__Int__ReprAssembler) valueFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.va.w = nil
		ma.cm = schema.Maybe_Absent
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__TypeName__Int__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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

	var k2 _String
	if err := (_String__ReprPrototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Map__TypeName__Int__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.w = &tz.v
	ma.va.m = &ma.cm
	return &ma.va, nil
}
func (ma *_Map__TypeName__Int__ReprAssembler) AssembleKey() ipld.NodeAssembler {
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
	ma.w.t = append(ma.w.t, _Map__TypeName__Int__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Map__TypeName__Int__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		if !ma.keyFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
		} // if tidy success: carry on
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	return &ma.va
}
func (ma *_Map__TypeName__Int__ReprAssembler) Finish() error {
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
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Map__TypeName__Int__ReprAssembler) KeyPrototype() ipld.NodePrototype {
	return _String__ReprPrototype{}
}
func (ma *_Map__TypeName__Int__ReprAssembler) ValuePrototype(_ string) ipld.NodePrototype {
	return _Int__ReprPrototype{}
}
