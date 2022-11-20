// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
//
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//                  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

// NOTE:
//   1. Test codes are in the public hestiaSTRING package.

package hestiaFMT

import (
	"hestia/hestiaINTERNAL/hestiaMATH"
)

func SN_ParseUINT(input string, base uint64, size uint16) (out uint64, err Error) {
	var errX Error
	var digit uint8
	var data []rune
	var x, multiplier uint64
	var base64 float64
	var omizu *omizuData

	// pick off easy targets
	out = 0
	switch {
	case input == "":
		return out, ERROR_OK
	default:
	}

	data = []rune(input)
	base64 = float64(base)
	omizu, err = __sN_Omizu_Parse(&data, &base64)

	switch {
	case err != ERROR_OK:
		return out, err
	case string(omizu.partial) != "0":
		return out, ERROR_INPUT_INVALID
	case string(omizu.exponent) != "0":
		return out, ERROR_INPUT_INVALID
	case omizu.negativeValue:
		return out, ERROR_INPUT_INVALID
	}

	// process base number
	switch {
	case len(omizu.base) == 0:
		// use provided base
	default:
		multiplier = 1
		for i := len(omizu.base) - 1; i >= 0; i-- {
			digit, errX = _SN_DIGIT_To_NUMBER(omizu.base[i])
			if errX != ERROR_OK {
				err = ERROR_BASE_INVALID
				break
			}

			x += multiplier * uint64(digit)
			multiplier *= 10
		}

		if base != x {
			err = ERROR_BASE_MISMATCHED
			base = x // prioritize parsed base
		}
	}

	// process round number
	multiplier = 1
	x = 0
	for i := len(omizu.round) - 1; i >= 0; i-- {
		digit, errX = _SN_DIGIT_To_NUMBER(omizu.round[i])
		switch {
		case errX != ERROR_OK:
			return 0, ERROR_INPUT_INVALID
		case uint64(digit) >= base:
			return 0, ERROR_INPUT_INVALID
		}

		out += uint64(digit) * multiplier
		multiplier *= base

		if out < x {
			return 0, ERROR_INPUT_OVERFLOW
		}
		x = out
	}

	// resize accordingly
	switch hestiaMATH.S64_Resize(&out, size, false) {
	case hestiaMATH.ERROR_OK:
		return out, ERROR_OK
	case hestiaMATH.ERROR_INPUT_INVALID:
		return 0, ERROR_INPUT_INVALID
	case hestiaMATH.ERROR_INPUT_OUT_OF_RANGE:
		return 0, ERROR_INPUT_OUT_OF_RANGE
	default:
		return 0, ERROR_BAD_EXEC
	}
}

func S8_FormatUINT8(number uint8, base uint8, lettercase Lettercase) (out []rune) {
	var i uint8

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 8 // MAX: 8-bits in base-2
	out = make([]rune, i)

	// process number according to base
	_s8_formatNumber(&out, &base, &number, &i, &lettercase)

	// done conversion return output
	return out[i:]
}

func S8_FormatINT8(input int8, base uint8, lettercase Lettercase) (out []rune) {
	var i, number uint8

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if input is 0 for all bases
	if input == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 8 + 1 // MAX: 8-bits in base-2 + sign when available (-)
	out = make([]rune, i)

	// modulus negative number for division
	if input < 0 {
		number = uint8(input * -1)
	} else {
		number = uint8(input)
	}

	// process number according to base
	_s8_formatNumber(&out, &base, &number, &i, &lettercase)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

func S16_FormatUINT16(number uint16, base uint16, lettercase Lettercase) (out []rune) {
	var i uint16

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 16 // MAX: 16-bits in base-2
	out = make([]rune, i)

	// process number according to base
	_s16_formatNumber(&out, &base, &number, &i, &lettercase)

	// done conversion return output
	return out[i:]
}

func S16_FormatINT16(input int16, base uint16, lettercase Lettercase) (out []rune) {
	var i, number uint16

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if input is 0 for all bases
	if input == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 16 + 1 // MAX: 16-bits in base-2 + sign when available (-)
	out = make([]rune, i)

	// modulus negative number for division
	if input < 0 {
		number = uint16(input * -1)
	} else {
		number = uint16(input)
	}

	// process number according to base
	_s16_formatNumber(&out, &base, &number, &i, &lettercase)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

func S32_FormatUINT32(number uint32, base uint32, lettercase Lettercase) (out []rune) {
	var i uint32

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 32 // MAX: 32-bits in base-2
	out = make([]rune, i)

	// process number according to base
	_s32_formatNumber(&out, &base, &number, &i, &lettercase)

	// done conversion return output
	return out[i:]
}

func S32_FormatINT32(input int32, base uint32, lettercase Lettercase) (out []rune) {
	var i, number uint32

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if input is 0 for all bases
	if input == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 32 + 1 // MAX: 32-bits in base-2 + sign when available (-)
	out = make([]rune, i)

	// modulus negative number for division
	if input < 0 {
		number = uint32(input * -1)
	} else {
		number = uint32(input)
	}

	// process number according to base
	_s32_formatNumber(&out, &base, &number, &i, &lettercase)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

func S64_FormatUINT64(number uint64, base uint64, lettercase Lettercase) (out []rune) {
	var i uint64

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if number is 0 for all bases
	if number == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 64 // MAX: 64-bits in base-2
	out = make([]rune, i)

	// process number according to base
	_s64_formatNumber(&out, &base, &number, &i, &lettercase)

	// done conversion return output
	return out[i:]
}

func S64_FormatINT64(input int64, base uint64, lettercase Lettercase) (out []rune) {
	var i uint64
	var number uint64

	// guard against all processible base number
	if base < 2 || base > 36 { // 36 = 'z'
		panic("base number smaller than 2 or larger than 36!")
	}

	// return early if input is 0 for all bases
	if input == 0 {
		return []rune{'0'}
	}

	// prepare for conversion
	i = 64 + 1 // MAX: 64-bits in base-2 + sign when available (-)
	out = make([]rune, i)

	// modulus negative number for division
	if input < 0 {
		number = uint64(input * -1)
	} else {
		number = uint64(input)
	}

	// process number according to base
	_s64_formatNumber(&out, &base, &number, &i, &lettercase)

	// process negative sign
	if input < 0 {
		i--
		out[i] = '-'
	}

	// done conversion and return output
	return out[i:]
}

//nolint:goconst
func _SN_DIGIT_To_NUMBER(input rune) (uint8, Error) {
	switch input {
	case '0':
		return 0, ERROR_OK
	case '1':
		return 1, ERROR_OK
	case '2':
		return 2, ERROR_OK
	case '3':
		return 3, ERROR_OK
	case '4':
		return 4, ERROR_OK
	case '5':
		return 5, ERROR_OK
	case '6':
		return 6, ERROR_OK
	case '7':
		return 7, ERROR_OK
	case '8':
		return 8, ERROR_OK
	case '9':
		return 9, ERROR_OK
	case 'a', 'A':
		return 10, ERROR_OK
	case 'b', 'B':
		return 11, ERROR_OK
	case 'c', 'C':
		return 12, ERROR_OK
	case 'd', 'D':
		return 13, ERROR_OK
	case 'e', 'E':
		return 14, ERROR_OK
	case 'f', 'F':
		return 15, ERROR_OK
	case 'g', 'G':
		return 16, ERROR_OK
	case 'h', 'H':
		return 17, ERROR_OK
	case 'i', 'I':
		return 18, ERROR_OK
	case 'j', 'J':
		return 19, ERROR_OK
	case 'k', 'K':
		return 20, ERROR_OK
	case 'l', 'L':
		return 21, ERROR_OK
	case 'm', 'M':
		return 22, ERROR_OK
	case 'n', 'N':
		return 23, ERROR_OK
	case 'o', 'O':
		return 24, ERROR_OK
	case 'p', 'P':
		return 25, ERROR_OK
	case 'q', 'Q':
		return 26, ERROR_OK
	case 'r', 'R':
		return 27, ERROR_OK
	case 's', 'S':
		return 28, ERROR_OK
	case 't', 'T':
		return 29, ERROR_OK
	case 'u', 'U':
		return 30, ERROR_OK
	case 'v', 'V':
		return 31, ERROR_OK
	case 'w', 'W':
		return 32, ERROR_OK
	case 'x', 'X':
		return 33, ERROR_OK
	case 'y', 'Y':
		return 34, ERROR_OK
	case 'z', 'Z':
		return 35, ERROR_OK
	default:
		return 0, ERROR_INPUT_INVALID
	}
}

func _SN_NUMBER_To_DIGIT(number uint8, lettercase *Lettercase) rune {
	var charset []rune

	switch *lettercase {
	case LETTERCASE_UPPER:
		charset = []rune("0123456789ABCDEFGHIJKLMONPQRSTUVWXYZ")
	case LETTERCASE_LOWER:
		charset = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	default:
		panic("unknown lettercase!")
	}

	return charset[number]
}

func _s8_formatNumber(buffer *[]rune,
	base *uint8, number *uint8, i *uint8, lettercase *Lettercase) {
	var x uint8

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = _SN_NUMBER_To_DIGIT(*number-x*(*base), lettercase)
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = _SN_NUMBER_To_DIGIT(*number, lettercase)
}

func _s16_formatNumber(buffer *[]rune,
	base *uint16, number *uint16, i *uint16, lettercase *Lettercase) {
	var x uint16

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number-x*(*base)), lettercase)
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number), lettercase)
}

func _s32_formatNumber(buffer *[]rune,
	base *uint32, number *uint32, i *uint32, lettercase *Lettercase) {
	var x uint32

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number-x*(*base)), lettercase)
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number), lettercase)
}

func _s64_formatNumber(buffer *[]rune,
	base *uint64, number *uint64, i *uint64, lettercase *Lettercase) {
	var x uint64

	// process number according to base
	for *number >= *base {
		*i--
		x = *number / *base
		(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number-x*(*base)), lettercase)
		*number = x
	}

	// process last remainder
	*i--
	(*buffer)[*i] = _SN_NUMBER_To_DIGIT(uint8(*number), lettercase)
}
