// Code generated by protoc-gen-equal-go. DO NOT EDIT.
// source: internal/testprotos/test/test_public.proto

package test

func (x *PublicImportMessage) Equal(y *PublicImportMessage) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}
	if x == y {
		return true
	}
	return true
}
