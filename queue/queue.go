package queue

import (
	"sort"
)

var elems map[string]Elem
func init()  {
	elems = make(map[string]Elem, 0)
}

type Elem interface {
	GetValue() int
	GetKey() string
}

func NewWithPrio(nodes ...Elem) (Queue, interface{}) {
	q := Queue{}
	for _, n := range nodes {
		q.Insert(n)
	}
	return q, elems
}

type Queue []Elem

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].GetValue() < q[j].GetValue()
}

func (q Queue) Swap(i, j int){
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Sort() {
	sort.Sort(q)
}

func (q *Queue) Pop() (n Elem) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Insert(n Elem) {
	*q = append(*q, n)
	q.Sort()
	elems[n.GetKey()] = n
}
