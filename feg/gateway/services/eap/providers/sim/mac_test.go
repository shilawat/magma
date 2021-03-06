/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sim

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GsmFroUmtsTestSet struct {
	ik, ck, xres, kc, sres1, sres2 []byte
}

var (
	testHmac = [20]byte{222, 124, 155, 133, 184, 183, 138, 166, 188, 138, 122, 54, 247, 10, 144, 112, 28, 157, 180, 217}
	rand     = [][]byte{
		{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f},
		{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f},
		{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f},
	}
	sres = [][]byte{
		{0xd1, 0xd2, 0xd3, 0xd4},
		{0xe1, 0xe2, 0xe3, 0xe4},
		{0xf1, 0xf2, 0xf3, 0xf4},
	}
	Kc = [][]byte{
		{0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7},
		{0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7},
		{0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7},
	}
	origIdentity = "1244070100000001@eapsim.foo"
	nonce        = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	mk           = []byte{
		0xe5, 0x76, 0xd5, 0xca, 0x33, 0x2e, 0x99, 0x30, 0x01, 0x8b,
		0xf1, 0xba, 0xee, 0x27, 0x63, 0xc7, 0x95, 0xb3, 0xc7, 0x12}
	k_encr = []byte{0x53, 0x6e, 0x5e, 0xbc, 0x44, 0x65, 0x58, 0x2a, 0xa6, 0xa8, 0xec, 0x99, 0x86, 0xeb, 0xb6, 0x20}
	k_aut  = []byte{0x25, 0xaf, 0x19, 0x42, 0xef, 0xcb, 0xf4, 0xbc, 0x72, 0xb3, 0x94, 0x34, 0x21, 0xf2, 0xa9, 0x74}
	msk    = []byte{
		0x39, 0xd4, 0x5a, 0xea, 0xf4, 0xe3, 0x06, 0x01, 0x98, 0x3e, 0x97, 0x2b, 0x6c, 0xfd, 0x46, 0xd1,
		0xc3, 0x63, 0x77, 0x33, 0x65, 0x69, 0x0d, 0x09, 0xcd, 0x44, 0x97, 0x6b, 0x52, 0x5f, 0x47, 0xd3,
		0xa6, 0x0a, 0x98, 0x5e, 0x95, 0x5c, 0x53, 0xb0, 0x90, 0xb2, 0xe4, 0xb7, 0x37, 0x19, 0x19, 0x6a,
		0x40, 0x25, 0x42, 0x96, 0x8f, 0xd1, 0x4a, 0x88, 0x8f, 0x46, 0xb9, 0xa7, 0x88, 0x6e, 0x44, 0x88,
	}
	emsk = []byte{
		0x59, 0x49, 0xea, 0xb0, 0xff, 0xf6, 0x9d, 0x52, 0x31, 0x5c, 0x6c, 0x63, 0x4f, 0xd1, 0x4a, 0x7f,
		0x0d, 0x52, 0x02, 0x3d, 0x56, 0xf7, 0x96, 0x98, 0xfa, 0x65, 0x96, 0xab, 0xee, 0xd4, 0xf9, 0x3f,
		0xbb, 0x48, 0xeb, 0x53, 0x4d, 0x98, 0x54, 0x14, 0xce, 0xed, 0x0d, 0x9a, 0x8e, 0xd3, 0x3c, 0x38,
		0x7c, 0x9d, 0xfd, 0xab, 0x92, 0xff, 0xbd, 0xf2, 0x40, 0xfc, 0xec, 0xf6, 0x5a, 0x2c, 0x93, 0xb9,
	}
	testData = []byte{0x01, 0x02, 0x01, 0x18, 0x12, 0x0b, 0x00, 0x00, 0x01, 0x0d, 0x00, 0x00,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
		0x81, 0x05, 0x00, 0x00,
		0x9e, 0x18, 0xb0, 0xc2, 0x9a, 0x65, 0x22, 0x63, 0xc0, 0x6e, 0xfb, 0x54, 0xdd, 0x00, 0xa8, 0x95,
		0x82, 0x2d, 0x00, 0x00,
		0x55, 0xf2, 0x93, 0x9b, 0xbd, 0xb1, 0xb1, 0x9e, 0xa1, 0xb4, 0x7f, 0xc0, 0xb3, 0xe0, 0xbe, 0x4c,
		0xab, 0x2c, 0xf7, 0x37, 0x2d, 0x98, 0xe3, 0x02, 0x3c, 0x6b, 0xb9, 0x24, 0x15, 0x72, 0x3d, 0x58,
		0xba, 0xd6, 0x6c, 0xe0, 0x84, 0xe1, 0x01, 0xb6, 0x0f, 0x53, 0x58, 0x35, 0x4b, 0xd4, 0x21, 0x82,
		0x78, 0xae, 0xa7, 0xbf, 0x2c, 0xba, 0xce, 0x33, 0x10, 0x6a, 0xed, 0xdc, 0x62, 0x5b, 0x0c, 0x1d,
		0x5a, 0xa6, 0x7a, 0x41, 0x73, 0x9a, 0xe5, 0xb5, 0x79, 0x50, 0x97, 0x3f, 0xc7, 0xff, 0x83, 0x01,
		0x07, 0x3c, 0x6f, 0x95, 0x31, 0x50, 0xfc, 0x30, 0x3e, 0xa1, 0x52, 0xd1, 0xe1, 0x0a, 0x2d, 0x1f,
		0x4f, 0x52, 0x26, 0xda, 0xa1, 0xee, 0x90, 0x05, 0x47, 0x22, 0x52, 0xbd, 0xb3, 0xb7, 0x1d, 0x6f,
		0x0c, 0x3a, 0x34, 0x90, 0x31, 0x6c, 0x46, 0x92, 0x98, 0x71, 0xbd, 0x45, 0xcd, 0xfd, 0xbc, 0xa6,
		0x11, 0x2f, 0x07, 0xf8, 0xbe, 0x71, 0x79, 0x90, 0xd2, 0x5f, 0x6d, 0xd7, 0xf2, 0xb7, 0xb3, 0x20,
		0xbf, 0x4d, 0x5a, 0x99, 0x2e, 0x88, 0x03, 0x31, 0xd7, 0x29, 0x94, 0x5a, 0xec, 0x75, 0xae, 0x5d,
		0x43, 0xc8, 0xed, 0xa5, 0xfe, 0x62, 0x33, 0xfc, 0xac, 0x49, 0x4e, 0xe6, 0x7a, 0x0d, 0x50, 0x4d,
		0x0b, 0x05, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
	expectedMac = []byte{
		0xfe, 0xf3, 0x24, 0xac,
		0x39, 0x62, 0xb5, 0x9f,
		0x3b, 0xd7, 0x82, 0x53,
		0xae, 0x4d, 0xcb, 0x6a,
	}
	challengeTestData = []byte{
		0x02, 0x02, 0x00, 0x1c, 0x12, 0x0b, 0x00, 0x00, 0x0b, 0x05, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
	expectedChellengeMac = []byte{
		0xf5, 0x6d, 0x64, 0x33,
		0xe6, 0x8e, 0xd2, 0x97,
		0x6a, 0xc1, 0x19, 0x37,
		0xfc, 0x3d, 0x11, 0x54,
	}
	gsmFromUmts = []GsmFroUmtsTestSet{
		{
			ik:    []byte{0xf7, 0x69, 0xbc, 0xd7, 0x51, 0x04, 0x46, 0x04, 0x12, 0x76, 0x72, 0x71, 0x1c, 0x6d, 0x34, 0x41},
			ck:    []byte{0xb4, 0x0b, 0xa9, 0xa3, 0xc5, 0x8b, 0x2a, 0x05, 0xbb, 0xf0, 0xd9, 0x87, 0xb2, 0x1b, 0xf8, 0xcb},
			xres:  []byte{0xa5, 0x42, 0x11, 0xd5, 0xe3, 0xba, 0x50, 0xbf},
			kc:    []byte{0xea, 0xe4, 0xbe, 0x82, 0x3a, 0xf9, 0xa0, 0x8b},
			sres1: []byte{0x46, 0xf8, 0x41, 0x6a},
			sres2: []byte{0xa5, 0x42, 0x11, 0xd5},
		},
		{
			ik:    []byte{0x59, 0xa9, 0x2d, 0x3b, 0x47, 0x6a, 0x04, 0x43, 0x48, 0x70, 0x55, 0xcf, 0x88, 0xb2, 0x30, 0x7b},
			ck:    []byte{0x5d, 0xbd, 0xbb, 0x29, 0x54, 0xe8, 0xf3, 0xcd, 0xe6, 0x65, 0xb0, 0x46, 0x17, 0x9a, 0x50, 0x98},
			xres:  []byte{0x80, 0x11, 0xc4, 0x8c, 0x0c, 0x21, 0x4e, 0xd2},
			kc:    []byte{0xaa, 0x01, 0x73, 0x9b, 0x8c, 0xaa, 0x97, 0x6d},
			sres1: []byte{0x8c, 0x30, 0x8a, 0x5e},
			sres2: []byte{0x80, 0x11, 0xc4, 0x8c},
		},
		{
			ik:    []byte{0x0c, 0x45, 0x24, 0xad, 0xea, 0xc0, 0x41, 0xc4, 0xdd, 0x83, 0x0d, 0x20, 0x85, 0x4f, 0xc4, 0x6b},
			ck:    []byte{0xe2, 0x03, 0xed, 0xb3, 0x97, 0x15, 0x74, 0xf5, 0xa9, 0x4b, 0x0d, 0x61, 0xb8, 0x16, 0x34, 0x5d},
			xres:  []byte{0xf3, 0x65, 0xcd, 0x68, 0x3c, 0xd9, 0x2e, 0x96},
			kc:    []byte{0x9a, 0x8e, 0xc9, 0x5f, 0x40, 0x8c, 0xc5, 0x07},
			sres1: []byte{0xcf, 0xbc, 0xe3, 0xfe},
			sres2: []byte{0xf3, 0x65, 0xcd, 0x68},
		},
	}
)

func TestMacGeneration(t *testing.T) {
	hmac := HmacSha1([]byte("key"), []byte("The quick brown fox jumps over"), []byte(" the lazy dog"))
	t.Logf("Generated HMAC: %v", hmac)
	if !reflect.DeepEqual(hmac, testHmac[:]) {
		t.Fatalf(
			"HMACs don't match.\n\tGenerated HMAC(%d): %v\n\tExpected  HMAC(%d): %v",
			len(hmac), hmac, len(testHmac), testHmac)
	}
	actualMK := MK([]byte(origIdentity), nonce, []byte{0, Version}, []byte{0, Version}, Kc)
	assert.Equal(t, mk, actualMK)

	K_encr, K_aut, MSK, EMSK := MakeKeys([]byte(origIdentity), nonce, []byte{0, Version}, []byte{0, Version}, Kc)
	t.Logf("Generated keys:\n\tK_encr=%v\n\tK_aut=%v\n\tMSK=%v\n\tEMSK=%v", K_encr, K_aut, MSK, EMSK)

	if len(K_encr) != 16 {
		t.Fatalf("Invalid K_encr Len: %d", len(K_encr))
	}
	if len(K_aut) != 16 {
		t.Fatalf("Invalid K_aut Len: %d", len(K_aut))
	}
	if len(MSK) != 64 {
		t.Fatalf("Invalid MSK Len: %d", len(MSK))
	}
	if len(EMSK) != 64 {
		t.Fatalf("Invalid EMSK Len: %d", len(EMSK))
	}
	assert.Equal(t, k_encr, K_encr)
	assert.Equal(t, k_aut, K_aut)
	assert.Equal(t, msk, MSK)
	assert.Equal(t, emsk, EMSK)

	mac := GenMac(testData, nonce, K_aut)
	if len(mac) != 16 {
		t.Fatalf("Invalid MAC Len: %d", len(mac))
	}
	// compare generated MAC with expected
	if !reflect.DeepEqual(mac, []byte(expectedMac)) {
		t.Fatalf(
			"MACs don't match.\n\tGenerated MAC(%d): %v\n\tExpected  MAC(%d): %v",
			len(mac), mac, len(expectedMac), []byte(expectedMac))
	}
	mac = GenChallengeMac(challengeTestData, sres, K_aut)
	if len(mac) != 16 {
		t.Fatalf("Invalid Challenge MAC Len: %d", len(mac))
	}
	// compare generated ChallengeMAC with expected
	if !reflect.DeepEqual(mac, []byte(expectedChellengeMac)) {
		t.Fatalf(
			"Challenge MACs don't match.\n\tGenerated MAC(%d): %v\n\tExpected  MAC(%d): %v",
			len(mac), mac, len(expectedChellengeMac), []byte(expectedChellengeMac))
	}
}

func TestGsmFromUmts(t *testing.T) {
	for _, testCase := range gsmFromUmts {
		kc, sres := GsmFromUmts1(testCase.ck, testCase.ik, testCase.xres)
		assert.Equal(t, testCase.kc, kc)
		assert.Equal(t, testCase.sres1, sres)
		kc, sres = GsmFromUmts2(testCase.ck, testCase.ik, testCase.xres)
		assert.Equal(t, testCase.kc, kc)
		assert.Equal(t, testCase.sres2, sres)
	}
}
