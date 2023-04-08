package setx

type mode interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type Set[T mode] struct {
	val map[T]struct{}
}

func NewSet[T mode]() *Set[T] {
	return &Set[T]{
		val: make(map[T]struct{}),
	}
}

// Add 添加元素
func (s *Set[T]) Add(key T) *Set[T] {
	s.val[key] = struct{}{}
	return s
}

// Del 删除元素
func (s *Set[T]) Del(key T) *Set[T] {
	delete(s.val, key)
	return s
}

// List 获取全部元素
func (s *Set[T]) List() []T {
	var out []T
	for k := range s.val {
		out = append(out, k)
	}
	return out
}

// IsIn 判断元素是否存在
func (s *Set[T]) IsIn(key T) bool {
	_, ok := s.val[key]
	return ok
}

// IsNil 判断是否为空
func (s *Set[T]) IsNil(key T) bool {
	return len(s.val) == 0
}

// Empty 置空
func (s *Set[T]) Empty() *Set[T] {
	s.val = make(map[T]struct{})
	return s
}

// Intersection 获取交集
func (s *Set[T]) Intersection(in *Set[T]) []T {
	var out []T
	for k := range s.val {
		if _, ok := in.val[k]; ok {
			out = append(out, k)
		}
	}
	return out
}

// Union 获取并集
func (s *Set[T]) Union(in *Set[T]) []T {
	var out []T
	for k := range in.val {
		out = append(out, k)
	}
	out = append(out, s.List()...)
	return out
}

// Difference 获取补集
func (s *Set[T]) Difference(in *Set[T]) []T {
	var out []T
	for k := range in.val {
		if _, ok := in.val[k]; !ok {
			out = append(out, k)
		}
	}
	return out
}
