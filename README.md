# Container
Golang 相关的容器组件实现。

## Set
Golang 实现的集合,可选线程安全。


```
type Set interface {
	// Add 向集合添加元素
	Add(values ...interface{})

	// Remove 从集合移除指定的元素
	Remove(values ...interface{})

	// RemoveAll 从集合移除所有的元素
	RemoveAll()

	// Exists 判断集合是否包含指定元素
	Exists(v interface{}) bool

	// Contains 判断集合是否包含指定的元素,包含所有的元素才会返回 true, 否则返回 false
	Contains(values ...interface{}) bool

	// Len 返回集合元素的长度
	Len() int

	// Values 返回集合所有元素组成的 Slice
	Values() []interface{}

	// Iter 返回集合的所有元素
	Iter() <-chan interface{}

	// Equal 判断和另一个集合是否相等
	Equal(s Set) bool

	// Clone 复制一个集合
	Clone() Set

	// Intersect 交集
	Intersect(s Set) Set

	// Union 并集
	Union(s Set) Set

	// Difference 差集
	Difference(s Set) Set
}
```

通过 set.NewSet() 创建的集合是 *非线程安全* 的。
 
通过 set.NewBlockSet() 创建的集合是 *线程安全* 的。

## Map
Golang 内置的 map 是非线程安全的, 因此重新实现了一个 Map。

```
type Map interface {
	// Set 添加一组键值对
	Set(key, value interface{})

	// Remove 移除指定的 key 及其 value
	Remove(key interface{})

	// RemoveAll 移除所有的 key 及 value
	RemoveAll()

	// Exists 判断指定的 key 是否存在
	Exists(key interface{}) bool

	// Contains 判断指定的 key 列表是否存在,只有当所有的 key 都存在的时候,才会返回 true
	Contains(keys ...interface{}) bool

	// Len 返回元素的个数
	Len() int

	// Value 获取指定 key 的 value
	Value(key interface{}) interface{}

	// Keys 返回由所有 key 组成的 Slice
	Keys() []interface{}

	// Values 返回由所有 value 组成的 Slice
	Values() []interface{}

	// Iter 返回所有 key 及 value
	Iter() <-chan itemValue
}
```

通过 smap.NewMap() 创建的 Map 是 *非线程安全* 的。

通过 smap.NewBlockMap() 创建的 Map 是 *线程安全* 的。

## 其它 - container 包


#### Contains
用于判断 source 对象中是否包含 target 对象。
source 对象可以是 map、slice 或者 array。

```
func Contains(source, target interface{}) bool
```

#### Remove
从 source 对象中移除找到的第一个 target 对象。
source 对象可以是 map 或者 slice。

```
func Remove(source, target interface{}) bool
```

#### RemoveAll
从 source 对象中移除所有的 target 对象。
source 对象可以是 map 或者 slice,由于 map 的 key 是不可重复的,所以本方法一般用于 Slice。

```
func RemoveAll(source, target interface{}) bool
```

#### Index
用于获取 target 对象在 source 对象中的索引位置。
source 对象可以为 slice 或者 array。

```
func Index(source, target interface{}) int
```

#### Indexs
用于获取 target 对象在 source 对象中的所有索引列表。
source 对象可以为 slice 或者 array。

```
func Indexs(source, target interface{}) []int
```