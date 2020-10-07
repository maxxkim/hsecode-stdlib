package set

import (
	"math/rand"
	"fmt"
)

const (
	p               = 0.25
	DefaultMaxLevel = 32
)

type iter interface {
	Next() (ok bool)
	Prev() (ok bool)
	Value() Element
}

type Element interface {
	Less(other Element) bool
	Equal(other Element) bool
}

type Int int

func (v Int) Less(other Element) bool {
	return int(v) < int(other.(Int))
}

func (v Int) Equal(other Element) bool {
	return int(v) == int(other.(Int))
}

type node struct {
	forward    []*node
	backward   *node
	key, value Element
}

type Set struct {
	lessThan func(l, r interface{}) bool
	header   *node
	footer   *node
	length   int
	MaxLevel int
}

type Iterator struct {
	current *node
	list    *Set
	key     interface{}
	value   interface{}
}

func (n *node) next() *node {
	if len(n.forward) == 0 {
		return nil
	}
	return n.forward[0]
}
func (n *node) previous() *node {
	return n.backward
}

func (n *node) hasNext() bool {
	return n.next() != nil
}

func (n *node) hasPrevious() bool {
	return n.previous() != nil
}

func (it *Iterator) Next() bool {
	if it == nil {
		return false
	}
	if !it.current.hasNext() {
		return false
	}

	it.current = it.current.next()
	it.key = it.current.key
	it.value = it.current.value

	return true
}

func (it *Iterator) Prev() bool {
	if it.current == it.list.footer {
		fmt.Println(it.current.value)
	}
	if it == nil {
		return false
	}
	if !it.current.hasPrevious() {
		return false
	}

	it.current = it.current.previous()
	it.key = it.current.key
	it.value = it.current.value

	return true
}
func (it *Iterator) Value() Element {
	return it.value.(Element)
}

func NewCustomMap(lessthan func(l, r interface{}) bool) *Set {
	return &Set{
		lessThan: lessthan,
		header: &node{
			forward: []*node{nil},
		},
		MaxLevel: DefaultMaxLevel,
	}
}

func New() *Set {
	comparator := func(left, right interface{}) bool {
		return left.(Element).Less(right.(Element))
	}
	return NewCustomMap(comparator)
}

func (s *Set) Begin() *Iterator {
	if s.length == 0 {
		return nil
	}
	current := s.header
	return &Iterator{
		current: current,
		list:    s,
		key:     current.key,
		value:   current.value,
	}
}

func (s *Set) getPath(current *node, update []*node, key interface{}) *node {
	depth := len(current.forward) - 1

	for i := depth; i >= 0; i-- {
		for current.forward[i] != nil && s.lessThan(current.forward[i].key, key) {
			current = current.forward[i]
		}
		if update != nil {
			update[i] = current
		}
	}

	return current.next()
}

func (s *Set) Delete(e Element) bool {
	if e == nil {
		panic("nil key is not supported")
	}

	update := make([]*node, s.level()+1, s.effectiveMaxLevel())
	candidate := s.getPath(s.header, update, e)
	if candidate == nil || candidate.key != e {
		return false
	}

	previous := candidate.backward
	if s.footer == candidate {
		s.footer = previous
	}

	next := candidate.next()
	if next != nil {
		next.backward = previous
	}

	for i := 0; i <= s.level() && update[i].forward[i] == candidate; i++ {
		update[i].forward[i] = candidate.forward[i]
	}

	for s.level() > 0 && s.header.forward[s.level()] == nil {
		s.header.forward = s.header.forward[:s.level()]
	}
	s.length--

	return true
}
func (s *Set) End() *Iterator {
	current := s.footer
	if current == nil {
		return nil
	}
	return &Iterator{
		current: current,
		list:    s,
		key:     current.key,
		value:   current.value,
	}
}
func (s *Set) Find(e Element) (Element, bool) {
	candidate := s.getPath(s.header, nil, e)
	if candidate == nil || candidate.key != e {
		return nil, false
	}
	return candidate.value, true
}

func (s *Set) level() int {
	return len(s.header.forward) - 1
}

func (s *Set) effectiveMaxLevel() int {
	return maxInt(s.level(), s.MaxLevel)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (s Set) randomLevel() (n int) {
	for n = 0; n < s.effectiveMaxLevel() && rand.Float64() < p; n++ {
	}
	return n
}

func (s *Set) Insert(e Element) bool {
	if e == nil {
		panic("nil key is not supported")
	}

	update := make([]*node, s.level()+1, s.effectiveMaxLevel()+1)
	candidate := s.getPath(s.header, update, e)

	if candidate != nil && candidate.key == e {
		candidate.value = e
		return false
	}

	newLevel := s.randomLevel()

	if currentLevel := s.level(); newLevel > currentLevel {
		for i := currentLevel + 1; i <= newLevel; i++ {
			update = append(update, s.header)
			s.header.forward = append(s.header.forward, nil)
		}
	}

	newNode := &node{
		forward: make([]*node, newLevel+1, s.effectiveMaxLevel()+1),
		key:     e,
		value:   e,
	}

	if previous := update[0]; previous.key != nil {
		newNode.backward = previous
	}

	for i := 0; i <= newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}

	s.length++

	if newNode.forward[0] != nil {
		if newNode.forward[0].backward != newNode {
			newNode.forward[0].backward = newNode
		}
	}

	if s.footer == nil || s.lessThan(s.footer.key, e) {
		s.footer = newNode
	}
	return true
}

func (s *Set) Len() int {
	return s.length
}

func (s *Set) LowerBound(e Element) *Iterator {
	candidate := s.getPath(s.header, nil, e)

	if candidate == nil {
		return nil
	}
	return &Iterator{current: candidate, list: s, key: e, value: e}
}

func (s *Set) UpperBound(e Element) *Iterator {
	candidate := s.getPath(s.header, nil, e)

	if candidate == nil {
		return nil
	}
	return &Iterator{current: candidate, list: s, key: e, value: e}

}
