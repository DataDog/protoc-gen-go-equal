// Code generated by protoc-gen-equal-go. DO NOT EDIT.
// source: internal/testprotos/test/weak2/test_weak.proto

package weak2

func (x *WeakImportMessage2) Equal(y *WeakImportMessage2) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}
	if x == y {
		return true
	}
	return x.A == y.A
}
