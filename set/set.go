package set

type Set interface {
	// 向集合添加元素
	Add(values ...interface{})

	// 从集合移除指定的元素
	Remove(values ...interface{})

	// 从集合移除所有的元素
	RemoveAll()

	// 判断集合是否包含指定元素
	Exists(v interface{}) bool

	// 判断集合是否包含指定的元素,包含所有的元素才会返回 true, 否则返回 false
	Contains(values ...interface{}) bool

	// 返回集合元素的长度
	Len() int

	// 返回集合元素组成的 Slice
	Values() []interface{}

	// 交集
	Intersect(s Set) Set

	// 并集
	Union(s Set) Set

	// 差集
	Difference(s Set) Set
}


type set map[interface{}]struct{}

func NewSet(values ...interface{}) Set {
	var s = make(set)
	if len(values) > 0 {
		s.Add(values...)
	}
	return &s
}

func (this *set) Add(values ...interface{}) {
	for _, v := range values {
		(*this)[v] = struct{}{}
	}
}

func (this *set) Remove(values ...interface{}) {
	for _, v := range values {
		delete(*this, v)
	}
}

func (this *set) RemoveAll() {
	for k, _ := range *this {
		delete(*this, k)
	}
}

func (this *set) Exists(v interface{}) bool {
	_, found := (*this)[v];
	return found
}

func (this *set) Contains(values ...interface{}) bool {
	for _, v := range values {
		_, exists := (*this)[v]
		if !exists {
			return false
		}
	}
	return true
}

func (this *set) Len() int {
	return len(*this)
}

func (this *set) Values() []interface{} {
	var s = *this
	var ns = make([]interface{}, 0, this.Len())
	for k, _ := range s {
		ns = append(ns, k)
	}
	return ns
}

func (this *set) Intersect(s Set) Set {
	var ns = NewSet()
	var vs = s.Values()
	for _, v := range vs {
		_, exists := (*this)[v]
		if exists {
			ns.Add(v)
		}
	}
	return ns
}

func (this *set) Union(s Set) Set {
	var ns = NewSet()
	ns.Add(this.Values()...)
	ns.Add(s.Values()...)
	return ns
}

func (this *set) Difference(s Set) Set {
	var ns = NewSet()
	for k, _ := range *this {
		if !s.Contains(k) {
			ns.Add(k)
		}
	}
	return ns
}
