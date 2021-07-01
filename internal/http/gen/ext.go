package gen

// Int64　int64の値取得
func (r ID) Int64() int64 {
	return int64(r)
}

// String　stringの値取得
func (r Key) String() string {
	return string(r)
}
