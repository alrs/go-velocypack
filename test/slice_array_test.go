//
// DISCLAIMER
//
// Copyright 2017 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package test

import (
	"testing"

	velocypack "github.com/arangodb/go-velocypack"
)

func TestSliceArrayEmpty(t *testing.T) {
	slice := velocypack.Slice{0x01}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_TRUE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(0), slice.MustLength(), t)
}

func TestSliceArrayCases1(t *testing.T) {
	slice := velocypack.Slice{0x02, 0x05, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases2(t *testing.T) {
	slice := velocypack.Slice{0x02, 0x06, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases3(t *testing.T) {
	slice := velocypack.Slice{0x02, 0x08, 0x00, 0x00, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases4(t *testing.T) {
	slice := velocypack.Slice{0x02, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases5(t *testing.T) {
	slice := velocypack.Slice{0x03, 0x06, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases6(t *testing.T) {
	slice := velocypack.Slice{0x03, 0x08, 0x00, 0x00, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases7(t *testing.T) {
	slice := velocypack.Slice{0x03, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases8(t *testing.T) {
	slice := velocypack.Slice{0x04, 0x08, 0x00, 0x00, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases9(t *testing.T) {
	slice := velocypack.Slice{0x04, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases10(t *testing.T) {
	slice := velocypack.Slice{0x05, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x31, 0x32, 0x33}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases11(t *testing.T) {
	slice := velocypack.Slice{0x06, 0x09, 0x03, 0x31, 0x32, 0x33, 0x03, 0x04, 0x05}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases12(t *testing.T) {
	slice := velocypack.Slice{0x06, 0x0b, 0x03, 0x00, 0x00, 0x31, 0x32, 0x33, 0x05, 0x06, 0x07}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases13(t *testing.T) {
	slice := velocypack.Slice{0x06, 0x0f, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x31, 0x32, 0x33, 0x09, 0x0a, 0x0b}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases14(t *testing.T) {
	slice := velocypack.Slice{0x07, 0x0e, 0x00, 0x03, 0x00, 0x31, 0x32, 0x33, 0x05, 0x00, 0x06, 0x00, 0x07, 0x00}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases15(t *testing.T) {
	slice := velocypack.Slice{0x07, 0x12, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x31, 0x32, 0x33, 0x09, 0x00, 0x0a, 0x00, 0x0b, 0x00}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases16(t *testing.T) {
	slice := velocypack.Slice{0x08, 0x18, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x31, 0x32, 0x33, 0x09, 0x00, 0x00, 0x00,
		0x0a, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCases17(t *testing.T) {
	slice := velocypack.Slice{0x09, 0x2c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x31, 0x32, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(3), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)
}

func TestSliceArrayCasesCompact(t *testing.T) {
	slice := velocypack.Slice{0x13, 0x08, 0x30, 0x31, 0x32, 0x33, 0x34, 0x05}

	ASSERT_EQ(velocypack.Array, slice.Type(), t)
	ASSERT_TRUE(slice.IsArray(), t)
	ASSERT_FALSE(slice.IsEmptyArray(), t)
	ASSERT_EQ(velocypack.ValueLength(len(slice)), slice.MustByteSize(), t)
	ASSERT_EQ(velocypack.ValueLength(5), slice.MustLength(), t)
	ss := slice.MustAt(0)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(0), ss.MustGetInt(), t)

	ss = slice.MustAt(1)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(1), ss.MustGetInt(), t)

	ss = slice.MustAt(4)
	ASSERT_TRUE(ss.IsSmallInt(), t)
	ASSERT_EQ(int64(4), ss.MustGetInt(), t)
}