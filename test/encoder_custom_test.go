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

type CustomStruct1 struct {
	Field1 int
}

func (cs *CustomStruct1) MarshalVPack() ([]byte, error) {
	var b velocypack.Builder
	if err := b.AddValue(velocypack.NewStringValue("Hello world")); err != nil {
		return nil, err
	}
	return b.Bytes()
}

func TestEncoderCustomStruct1(t *testing.T) {
	bytes, err := velocypack.Marshal(&CustomStruct1{
		Field1: 999,
	})
	ASSERT_NIL(err, t)
	s := velocypack.Slice(bytes)

	ASSERT_EQ(s.Type(), velocypack.String, t)
	ASSERT_EQ(`"Hello world"`, s.MustJSONString(), t)
}

type CustomStruct2 struct {
	Field CustomStruct1
}

func TestEncoderCustomStruct2(t *testing.T) {
	bytes, err := velocypack.Marshal(CustomStruct2{
		Field: CustomStruct1{
			Field1: 999222,
		},
	})
	ASSERT_NIL(err, t)
	s := velocypack.Slice(bytes)

	ASSERT_EQ(s.Type(), velocypack.Object, t)
	ASSERT_EQ(`{"Field":{"Field1":999222}}`, s.MustJSONString(), t)
}

type CustomStruct3 struct {
	Field *CustomStruct1
}

func TestEncoderCustomStruct3(t *testing.T) {
	bytes, err := velocypack.Marshal(CustomStruct3{
		Field: &CustomStruct1{
			Field1: 999222,
		},
	})
	ASSERT_NIL(err, t)
	s := velocypack.Slice(bytes)

	ASSERT_EQ(s.Type(), velocypack.Object, t)
	ASSERT_EQ(`{"Field":"Hello world"}`, s.MustJSONString(), t)
}