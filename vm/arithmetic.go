// Copyright (C) 2011 Göran Weinholt <goran@weinholt.se>
// Copyright (C) 2011 Per Odlund <per.odlund@gmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Arithmetic for conscheme

package conscheme

import (
	"big"
	"fmt"
	"unsafe"
)

const (
	fixnum_max = int(^uint(0) >> (1 + fixnum_shift))
	fixnum_min = -fixnum_max - 1
)

var fixnum_max_Int, fixnum_min_Int *big.Int

func init() {
	fixnum_max_Int = big.NewInt(int64(fixnum_max))
	fixnum_min_Int = big.NewInt(int64(fixnum_min))
}

// Fixnums

func fixnum_p(x Obj) Obj {
	return Make_boolean((uintptr(unsafe.Pointer(x)) & fixnum_mask) == fixnum_tag)
}

func Make_fixnum(x int) Obj {
	// XXX: assumes x fits in a fixnum
	return Obj(unsafe.Pointer(uintptr((x << fixnum_shift) | fixnum_tag)))
}

func fixnum_to_int(x Obj) int {
	return int(uintptr(unsafe.Pointer(x))) >> fixnum_shift
}

func fixnum_add(fx1,fx2 Obj) Obj {
	i1 := uintptr(unsafe.Pointer(fx1))
	i2 := uintptr(unsafe.Pointer(fx2))
	if (i1 & fixnum_mask) != fixnum_tag || (i2 & fixnum_mask) != fixnum_tag {
		panic("bad type")
	}
	r := i1 + i2
	// TODO: how should we do this?
	// if r < fixnum_min || r > fixnum_max {
	// 	panic("result not representable")
	// }

	return Obj(unsafe.Pointer(uintptr(r - fixnum_tag)))
}

//

func make_bignum(x int64) Obj {
	var vv interface{} = big.NewInt(x)
	return Obj(&vv)
}

func bignum_p(x Obj) Obj {
	if (uintptr(unsafe.Pointer(x)) & heap_mask) != heap_tag { return False }

	switch v := (*x).(type) {
	case *big.Int:
		return True
	}
	return False
}



func make_number(x int) Obj {
	v := Make_fixnum(x)
	if fixnum_to_int(v) != x {
		return make_bignum(int64(x))
	}
	return v
}

func number_to_int(x Obj) int {
	if (uintptr(unsafe.Pointer(x)) & fixnum_mask) == fixnum_tag {
		return fixnum_to_int(x)
	}

	if (uintptr(unsafe.Pointer(x)) & heap_mask) != heap_tag {
		panic("bad type")
	}

	switch v := (*x).(type) {
	case *big.Int:
		// XXX:
		return int(v.Int64())
	}
	panic("bad type")
}

func number_p(x Obj) Obj {
	if (uintptr(unsafe.Pointer(x)) & fixnum_mask) == fixnum_tag {
		return True
	}
	if (uintptr(unsafe.Pointer(x)) & heap_mask) != heap_tag {
		return False
	}
	switch v := (*x).(type) {
	case *big.Int:
		return True
	case *big.Rat:
		return True
	}
	return False
}

func number_equal(x,y Obj) Obj {
	xfx := (uintptr(unsafe.Pointer(x)) & fixnum_mask) == fixnum_tag
	yfx := (uintptr(unsafe.Pointer(y)) & fixnum_mask) == fixnum_tag
	if xfx && yfx {	return Make_boolean(x == y) }

	if (!xfx && (uintptr(unsafe.Pointer(x)) & heap_mask) != heap_tag) ||
		(!yfx && (uintptr(unsafe.Pointer(y)) & heap_mask) != heap_tag) {
		panic("bad type")
	}

	if xfx { return number_equal(y,x) }

	switch vx := (*x).(type) {
	case *big.Int:
		if yfx {
			vy := big.NewInt(int64(fixnum_to_int(y)))
			return Make_boolean(vx.Cmp(vy) == 0)
		}
		switch vy := (*y).(type) {
		case *big.Int:
			return Make_boolean(vx.Cmp(vy) == 0)
		case *big.Rat:
			return number_equal(y,x)
		default:
			panic("bad type")
		}
	case *big.Rat:
		// rationals should always have been converted into
		// other types if the denominator is one
		if yfx { return False }
		switch vy := (*y).(type) {
		case *big.Int:
			return False
		case *big.Rat:
			return Make_boolean(vx.Cmp(vy) == 0)
		default:
			panic("bad type")
		}
	}
	panic("bad type")
}

func number_add(x,y Obj) Obj {
	xfx := (uintptr(unsafe.Pointer(x)) & fixnum_mask) == fixnum_tag
	yfx := (uintptr(unsafe.Pointer(y)) & fixnum_mask) == fixnum_tag
	if xfx && yfx {
		i1 := uintptr(unsafe.Pointer(x))
		i2 := uintptr(unsafe.Pointer(y))
		r := (int(i1) >> fixnum_shift) + (int(i2) >> fixnum_shift)
		if r > fixnum_min && r < fixnum_max {
			return Make_fixnum(r)
		} else {
			return wrap(big.NewInt(int64(r)))
		}
	}

	if (!xfx && (uintptr(unsafe.Pointer(x)) & heap_mask) != heap_tag) ||
		(!yfx && (uintptr(unsafe.Pointer(y)) & heap_mask) != heap_tag) {
		panic("bad type")
	}

	if xfx { return number_add(y,x) }

	switch vx := (*x).(type) {
	case *big.Int:
		var z *big.Int = big.NewInt(0)
		if yfx {
			vy := big.NewInt(int64(fixnum_to_int(y)))
			return wrap(z.Add(vx,vy))
		}
		switch vy := (*y).(type) {
		case *big.Int:
			return wrap(z.Add(vx,vy))
		default:
			panic("bad type")
		}
	}
	panic("bad type")
}

func number_subtract(x,y Obj) Obj {
	xfx := (uintptr(unsafe.Pointer(x)) & fixnum_mask) == fixnum_tag
	yfx := (uintptr(unsafe.Pointer(y)) & fixnum_mask) == fixnum_tag
	if xfx && yfx {
		i1 := uintptr(unsafe.Pointer(x))
		i2 := uintptr(unsafe.Pointer(y))
		r := (int(i1) >> fixnum_shift) - (int(i2) >> fixnum_shift)
		if r > fixnum_min && r < fixnum_max {
			return Make_fixnum(r)
		} else {
			return wrap(big.NewInt(int64(r)))
		}
	}

	if (!xfx && (uintptr(unsafe.Pointer(x)) & heap_mask) != heap_tag) ||
		(!yfx && (uintptr(unsafe.Pointer(y)) & heap_mask) != heap_tag) {
		panic("bad type")
	}

	if xfx { return number_subtract(y,x) }

	switch vx := (*x).(type) {
	case *big.Int:
		var z *big.Int = big.NewInt(0)
		if yfx {
			vy := big.NewInt(int64(fixnum_to_int(y)))
			return wrap(z.Sub(vx,vy))
		}
		switch vy := (*y).(type) {
		case *big.Int:
			return wrap(z.Sub(vx,vy))
		default:
			panic("bad type")
		}
	}
	panic("bad type")
}

func number_divide(x,y Obj) Obj {
	xfx := (uintptr(unsafe.Pointer(x)) & fixnum_mask) == fixnum_tag
	yfx := (uintptr(unsafe.Pointer(y)) & fixnum_mask) == fixnum_tag
	if xfx && yfx {
		i1 := int(uintptr(unsafe.Pointer(x))) >> fixnum_shift
		i2 := int(uintptr(unsafe.Pointer(y))) >> fixnum_shift
		// A good optimizer will combine the div and mod into
		// one instruction.
		r, m := i1 / i2, i1 % i2
		if m == 0 && r > fixnum_min && r < fixnum_max {
			return Make_fixnum(r)
		} else {
			return wrap(big.NewRat(int64(i1),int64(i2)))
		}
	}

	if (!xfx && (uintptr(unsafe.Pointer(x)) & heap_mask) != heap_tag) ||
		(!yfx && (uintptr(unsafe.Pointer(y)) & heap_mask) != heap_tag) {
		panic("bad type")
	}

	if xfx { return number_divide(y,x) }

	switch vx := (*x).(type) {
	case *big.Int:
		var z *big.Int = big.NewInt(0)
		if yfx {
			vy := big.NewInt(int64(fixnum_to_int(y)))
			return wrap(z.Add(vx,vy))
		}
		switch vy := (*y).(type) {
		case *big.Int:
			return wrap(z.Add(vx,vy))
		default:
			panic("bad type")
		}
	}
	panic("bad type")
}

func _number_to_string(num Obj, radix Obj) Obj {
	var format string

	switch number_to_int(radix) {
	case 2: format = "%b"
	case 8: format = "%o"
	case 10: format = "%d"
	default: format = "%x"
	}

	if fixnum_p(num) != False {
		return String_string(fmt.Sprintf(format, fixnum_to_int(num)))
	}

	switch v := (*num).(type) {
	case *big.Int:
		return String_string(fmt.Sprintf(format, v))
	case *big.Rat:
		return String_string(fmt.Sprintf(format + "/" + format,
			v.Num(), v.Denom()))
	}

	panic("number->string needs numbers")
}
