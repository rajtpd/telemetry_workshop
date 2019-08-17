// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd netbsd openbsd

package route

import (
	"reflect"
	"testing"
)

type parseAddrsTest struct {
	attrs uint
	fn    func(int, []byte) (int, Addr, error)
	b     []byte
	as    []Addr
}

var parseAddrsLittleEndianTests = []parseAddrsTest{
	{
		sysRTA_DST | sysRTA_GATEWAY | sysRTA_NETMASK | sysRTA_BRD,
		parseKernelInetAddr,
		[]byte{
			0x38, 0x12, 0x0, 0x0, 0xff, 0xff, 0xff, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,

			0x38, 0x12, 0x2, 0x0, 0x6, 0x3, 0x6, 0x0,
			0x65, 0x6d, 0x31, 0x0, 0xc, 0x29, 0x66, 0x2c,
			0xdc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,

			0x10, 0x2, 0x0, 0x0, 0xac, 0x10, 0xdc, 0xb4,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,

			0x10, 0x2, 0x0, 0x0, 0xac, 0x10, 0xdc, 0xff,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		},
		[]Addr{
			&LinkAddr{Index: 0},
			&LinkAddr{Index: 2, Name: "em1", Addr: []byte{0x00, 0x0c, 0x29, 0x66, 0x2c, 0xdc}},
			&Inet4Addr{IP: [4]byte{172, 16, 220, 180}},
			nil,
			nil,
			nil,
			nil,
			&Inet4Addr{IP: [4]byte{172, 16, 220, 255}},
		},
	},
	{
		sysRTA_NETMASK | sysRTA_IFP | sysRTA_IFA,
		parseKernelInetAddr,
		[]byte{
			0x7, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0x0,

			0x18, 0x12, 0xa, 0x0, 0x87, 0x8, 0x0, 0x0,
			0x76, 0x6c, 0x61, 0x6e, 0x35, 0x36, 0x38, 0x32,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,

			0x10, 0x2, 0x0, 0x0, 0xa9, 0xfe, 0x0, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		},
		[]Addr{
			nil,
			nil,
			&Inet4Addr{IP: [4]byte{255, 255, 255, 0}},
			nil,
			&LinkAddr{Index: 10, Name: "vlan5682"},
			&Inet4Addr{IP: [4]byte{169, 254, 0, 1}},
			nil,
			nil,
		},
	},
}

func TestParseAddrs(t *testing.T) {
	tests := parseAddrsLittleEndianTests
	if nativeEndian != littleEndian {
		t.Skip("no test for non-little endian machine yet")
	}

	for i, tt := range tests {
		as, err := parseAddrs(tt.attrs, tt.fn, tt.b)
		if err != nil {
			t.Error(i, err)
			continue
		}
		as = as[:8] // the list varies between operating systems
		if !reflect.DeepEqual(as, tt.as) {
			t.Errorf("#%d: got %+v; want %+v", i, as, tt.as)
			continue
		}
	}
}
