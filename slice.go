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

package velocypack

import (
	"encoding/binary"
	"encoding/hex"
	"math"
)

// Slice provides read only access to a VPack value
type Slice []byte

// SliceFromBytes creates a Slice by casting the given byte slice to a Slice.
func SliceFromBytes(v []byte) Slice {
	return Slice(v)
}

// SliceFromHex creates a Slice by decoding the given hex string into a Slice.
// If decoding fails, nil is returned.
func SliceFromHex(v string) Slice {
	if bytes, err := hex.DecodeString(v); err != nil {
		return nil
	} else {
		return Slice(bytes)
	}
}

// String returns a HEX representation of the slice.
func (s Slice) String() string {
	return hex.EncodeToString(s)
}

// ByteSize returns the total byte size for the slice, including the head byte
func (s Slice) ByteSize() (ValueLength, error) {
	h := s[0]
	// check if the type has a fixed length first
	l := fixedTypeLengths[h]
	if l != 0 {
		// return fixed length
		return ValueLength(l), nil
	}

	// types with dynamic lengths need special treatment:
	switch s.Type() {
	case Array, Object:
		{
			if h == 0x13 || h == 0x14 {
				// compact Array or Object
				return readVariableValueLength(s[1:], false), nil
			}

			if h == 0x01 || h == 0x0a {
				// we cannot get here, because the FixedTypeLengths lookup
				// above will have kicked in already. however, the compiler
				// claims we'll be reading across the bounds of the input
				// here...
				return 1, nil
			}

			VELOCYPACK_ASSERT(h > 0x00 && h <= 0x0e)
			return ValueLength(readIntegerNonEmpty(s[1:], widthMap[h])), nil
		}

	case External:
		{
			return 1 + charPtrLength, nil
		}

	case UTCDate:
		{
			return 1 + int64Length, nil
		}

	case Int:
		{
			return ValueLength(1 + (h - 0x1f)), nil
		}

	case String:
		{
			VELOCYPACK_ASSERT(h == 0xbf)
			if h < 0xbf {
				// we cannot get here, because the FixedTypeLengths lookup
				// above will have kicked in already. however, the compiler
				// claims we'll be reading across the bounds of the input
				// here...
				return ValueLength(h) - 0x40, nil
			}
			// long UTF-8 String
			return ValueLength(1 + 8 + readIntegerFixed(s[1:], 8)), nil
		}

	case Binary:
		{
			VELOCYPACK_ASSERT(h >= 0xc0 && h <= 0xc7)
			return ValueLength(1 + ValueLength(h) - 0xbf + ValueLength(readIntegerNonEmpty(s[1:], int(h)-0xbf))), nil
		}

	case BCD:
		{
			if h <= 0xcf {
				// positive BCD
				VELOCYPACK_ASSERT(h >= 0xc8 && h < 0xcf)
				return ValueLength(1 + ValueLength(h) - 0xc7 + ValueLength(readIntegerNonEmpty(s[1:], int(h)-0xc7))), nil
			}

			// negative BCD
			VELOCYPACK_ASSERT(h >= 0xd0 && h < 0xd7)
			return ValueLength(1 + ValueLength(h) - 0xcf + ValueLength(readIntegerNonEmpty(s[1:], int(h)-0xcf))), nil
		}

	case Custom:
		{
			VELOCYPACK_ASSERT(h >= 0xf4)
			switch h {
			case 0xf4:
			case 0xf5:
			case 0xf6:
				{
					return ValueLength(2 + readIntegerFixed(s[1:], 1)), nil
				}

			case 0xf7:
			case 0xf8:
			case 0xf9:
				{
					return ValueLength(3 + readIntegerFixed(s[1:], 2)), nil
				}

			case 0xfa:
			case 0xfb:
			case 0xfc:
				{
					return ValueLength(5 + readIntegerFixed(s[1:], 4)), nil
				}

			case 0xfd:
			case 0xfe:
			case 0xff:
				{
					return ValueLength(9 + readIntegerFixed(s[1:], 8)), nil
				}

			default:
				{
					// fallthrough intentional
				}
			}
		}
	default:
		{
			// fallthrough intentional
		}
	}

	return 0, InternalError{}
}

// MustByteSize returns the total byte size for the slice, including the head byte.
// Panics in case of an error.
func (s Slice) MustByteSize() ValueLength {
	if v, err := s.ByteSize(); err != nil {
		panic(err)
	} else {
		return v
	}
}

// GetBool returns a boolean value from the slice.
// Returns an error if slice is not of type Bool.
func (s Slice) GetBool() (bool, error) {
	if err := s.AssertType(Bool); err != nil {
		return false, WithStack(err)
	}
	return s.IsTrue(), nil
}

// MustGetBool returns a boolean value from the slice.
// Panics if slice is not of type Bool.
func (s Slice) MustGetBool() bool {
	if v, err := s.GetBool(); err != nil {
		panic(err)
	} else {
		return v
	}
}

// GetDouble returns a Double value from the slice.
// Returns an error if slice is not of type Double.
func (s Slice) GetDouble() (float64, error) {
	if err := s.AssertType(Double); err != nil {
		return 0.0, WithStack(err)
	}
	bits := binary.LittleEndian.Uint64(s[1:])
	return math.Float64frombits(bits), nil
}

// MustGetDouble returns a Double value from the slice.
// Panics if slice is not of type Double.
func (s Slice) MustGetDouble() float64 {
	if v, err := s.GetDouble(); err != nil {
		panic(err)
	} else {
		return v
	}
}

// GetInt returns a Int value from the slice.
// Returns an error if slice is not of type Int.
func (s Slice) GetInt() (int64, error) {
	h := s[0]

	if h >= 0x20 && h <= 0x27 {
		// Int  T
		v := readIntegerNonEmpty(s[1:], int(h)-0x1f)
		if h == 0x27 {
			return toInt64(v), nil
		} else {
			vv := int64(v)
			shift := int64(1) << ((h-0x1f)*8 - 1)
			if vv < shift {
				return vv, nil
			} else {
				return vv - (shift << 1), nil
			}
		}
	}

	if h >= 0x28 && h <= 0x2f {
		// UInt
		v, err := s.GetUInt()
		if err != nil {
			return 0, WithStack(err)
		}
		if v > math.MaxInt64 {
			return 0, NumberOutOfRangeError{}
		}
		return int64(v), nil
	}

	if h >= 0x30 && h <= 0x3f {
		// SmallInt
		return s.GetSmallInt()
	}

	return 0, InvalidTypeError{"Expecting type Int"}
}

// MustGetInt returns a Int value from the slice.
// Panics if slice is not of type Int.
func (s Slice) MustGetInt() int64 {
	if v, err := s.GetInt(); err != nil {
		panic(err)
	} else {
		return v
	}
}

// GetUInt returns a UInt value from the slice.
// Returns an error if slice is not of type UInt.
func (s Slice) GetUInt() (uint64, error) {
	h := s[0]

	if h == 0x28 {
		// single byte integer
		return uint64(s[1]), nil
	}

	if h >= 0x29 && h <= 0x2f {
		// UInt
		return readIntegerNonEmpty(s[1:], int(h)-0x27), nil
	}

	if h >= 0x20 && h <= 0x27 {
		// Int
		v, err := s.GetInt()
		if err != nil {
			return 0, WithStack(err)
		}
		if v < 0 {
			return 0, NumberOutOfRangeError{}
		}
		return uint64(v), nil
	}

	if h >= 0x30 && h <= 0x39 {
		// Smallint >= 0
		return uint64(h - 0x30), nil
	}

	if h >= 0x3a && h <= 0x3f {
		// Smallint < 0
		return 0, NumberOutOfRangeError{}
	}

	return 0, InvalidTypeError{"Expecting type UInt"}
}

// MustGetUInt returns a UInt value from the slice.
// Panics if slice is not of type UInt.
func (s Slice) MustGetUInt() uint64 {
	if v, err := s.GetUInt(); err != nil {
		panic(err)
	} else {
		return v
	}
}

// GetSmallInt returns a SmallInt value from the slice.
// Returns an error if slice is not of type SmallInt.
func (s Slice) GetSmallInt() (int64, error) {
	h := s[0]

	if h >= 0x30 && h <= 0x39 {
		// Smallint >= 0
		return int64(h - 0x30), nil
	}

	if h >= 0x3a && h <= 0x3f {
		// Smallint < 0
		return int64(h-0x3a) - 6, nil
	}

	if (h >= 0x20 && h <= 0x27) || (h >= 0x28 && h <= 0x2f) {
		// Int and UInt
		// we'll leave it to the compiler to detect the two ranges above are
		// adjacent
		return s.GetInt()
	}

	return 0, InvalidTypeError{"Expecting type SmallInt"}
}

// MustGetSmallInt returns a SmallInt value from the slice.
// Panics if slice is not of type SmallInt.
func (s Slice) MustGetSmallInt() int64 {
	if v, err := s.GetSmallInt(); err != nil {
		panic(err)
	} else {
		return v
	}
}

// Length return the number of members for an Array or Object object
func (s Slice) Length() (ValueLength, error) {
	if !s.IsArray() && !s.IsObject() {
		return 0, InvalidTypeError{"Expecting type Array or Object"}
	}

	h := s[0]
	if h == 0x01 || h == 0x0a {
		// special case: empty!
		return 0, nil
	}

	if h == 0x13 || h == 0x14 {
		// compact Array or Object
		end := readVariableValueLength(s[1:], false)
		return readVariableValueLength(s[end-1:], true), nil
	}

	offsetSize := uint64(indexEntrySize(h))
	VELOCYPACK_ASSERT(offsetSize > 0)
	end := readIntegerNonEmpty(s[1:], int(offsetSize))

	// find number of items
	if h <= 0x05 { // No offset table or length, need to compute:
		firstSubOffset := s.findDataOffset(h)
		first := s[firstSubOffset:]
		s, err := first.ByteSize()
		if err != nil {
			return 0, WithStack(err)
		}
		if s == 0 {
			return 0, InternalError{}
		}
		return (ValueLength(end) - firstSubOffset) / s, nil
	} else if offsetSize < 8 {
		return ValueLength(readIntegerNonEmpty(s[offsetSize+1:], int(offsetSize))), nil
	}

	return ValueLength(readIntegerNonEmpty(s[end-offsetSize:], int(offsetSize))), nil
}

// MustLength return the number of members for an Array or Object object.
// Panics in case of error.
func (s Slice) MustLength() ValueLength {
	if result, err := s.Length(); err != nil {
		panic(err)
	} else {
		return result
	}
}

func indexEntrySize(head byte) ValueLength {
	VELOCYPACK_ASSERT(head > 0x00 && head <= 0x12)
	return ValueLength(widthMap[head])
}

func (s Slice) findDataOffset(head byte) ValueLength {
	// Must be called for a nonempty array or object at start():
	VELOCYPACK_ASSERT(head <= 0x12)
	fsm := firstSubMap[head]
	if fsm <= 2 && s[2] != 0 {
		return 2
	}
	if fsm <= 3 && s[3] != 0 {
		return 3
	}
	if fsm <= 5 && s[5] != 0 {
		return 5
	}
	return 9
}
