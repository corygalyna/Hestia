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

package hestiaSTRING

import (
	"hestia/hestiaINTERNAL/hestiaFMT"
	"hestia/hestiaINTERNAL/hestiaMATH"
	"hestia/hestiaTESTING"
)

// test condition labels
const (
	cond_BROKEN = "value is having broken/faulty data"

	cond_BOOL_TRUE  = "value is boolean 'true'"
	cond_BOOL_FALSE = "value is boolean 'false'"

	cond_BASE_0  = "value in Base-0 number"
	cond_BASE_1  = "value in Base-1 number"
	cond_BASE_2  = "value in Base-2 number"
	cond_BASE_5  = "value in Base-5 number"
	cond_BASE_8  = "value in Base-8 number"
	cond_BASE_10 = "value in Base-10 number"
	cond_BASE_12 = "value in Base-12 number"
	cond_BASE_16 = "value in Base-16 number"
	cond_BASE_22 = "value in Base-22 number"
	cond_BASE_36 = "value in Base-36 number"
	cond_BASE_37 = "value in Base-37 number"

	cond_RESIZE_0  = "resize value to 0-bit"
	cond_RESIZE_2  = "resize value to 2-bits"
	cond_RESIZE_5  = "resize value to 5-bits"
	cond_RESIZE_8  = "resize value to 8-bits"
	cond_RESIZE_10 = "resize value to 10-bits"
	cond_RESIZE_12 = "resize value to 12-bits"
	cond_RESIZE_16 = "resize value to 16-bits"
	cond_RESIZE_22 = "resize value to 22-bits"
	cond_RESIZE_36 = "resize value to 36-bits"
	cond_RESIZE_65 = "resize value to 65-bits"

	cond_PRECISION_0  = "value has precision with automatic sizing"
	cond_PRECISION_1  = "value has precision with 1 unit"
	cond_PRECISION_2  = "value has precision with 2 units"
	cond_PRECISION_4  = "value has precision with 4 units"
	cond_PRECISION_5  = "value has precision with 5 units"
	cond_PRECISION_10 = "value has precision with 10 units"

	cond_NOTATION_SCIENTIFIC_AUTO     = "scientific notation + auto exponent"
	cond_NOTATION_ISO6093NR3_AUTO     = "ISO6093NR3 scientific notation + auto exponent"
	cond_NOTATION_SCIENTIFIC          = "strictly scientific notation (X.XXXX*B±YYYY)"
	cond_NOTATION_ISO6093NR3          = "strictly ISO6093NR3 scientific notation"
	cond_NOTATION_DECIMAL_NO_EXPONENT = "strictly decimal with no exponent (XXXX.YYYY)"
	cond_NOTATION_IEEE754             = "strictly IEEE754 notation"
	cond_NOTATION_UNKNOWN             = "unknown or faulty notation"

	cond_ROUND_NORMAL = "use normal round number"
	cond_ROUND_ZERO   = "use zero round number"
	cond_ROUND_ONE    = "use one (1) round number"

	cond_PARTIAL_NORMAL = "use normal partial number"
	cond_PARTIAL_ZERO   = "use zero partial number"
	cond_PARTIAL_ONE    = "use one (1) partial number"

	cond_EXPONENT_POSITIVE         = "use positive exponent"
	cond_EXPONENT_NEGATIVE         = "use negative exponent"
	cond_EXPONENT_ZERO             = "use zero exponent"
	cond_EXPONENT_LARGER_MANTISSA  = "use exponent larger than mantissa magnitude"
	cond_EXPONENT_SMALLER_MANTISSA = "use exponent smaller than mantissa magnitude"
	cond_EXPONENT_EQUAL_MANTISSA   = "use exponent equal to mantissa magnitude"

	cond_FLOAT_NAN          = "value is float NaN"
	cond_FLOAT_INF_POSITIVE = "value is float positive infinity"
	cond_FLOAT_INF_NEGATIVE = "value is float negative infinity"
	cond_FLOAT_SMALLEST     = "value is the smallest float"

	cond_VALUE_NUMBER = "value is min/max number"
	cond_VALUE_UINT64 = "value is max uint64"
	cond_VALUE_INT64  = "value is min/max int64"
	cond_VALUE_UINT32 = "value is max uint32"
	cond_VALUE_INT32  = "value is min/max int32"
	cond_VALUE_UINT16 = "value is max uint16"
	cond_VALUE_INT16  = "value is min/max int16"
	cond_VALUE_UINT8  = "value is max uint8"
	cond_VALUE_INT8   = "value is min/max int8"
	cond_VALUE_ZERO   = "value is 0"

	cond_POSITIVE = "value is positive"
	cond_NEGATIVE = "value is negative"

	cond_UNKNOWNCASE = "value has unknown-case characters"
	cond_LOWERCASE   = "value has lowercase characters"
	cond_UPPERCASE   = "value has uppercase characters"
)

const (
	sample_STRING = "test string"
	sample_UINT64 = uint64(127)

	string_FLOAT_NAN          = "NaN"
	string_FLOAT_POSITIVE_INF = "+∞"
	string_FLOAT_NEGATIVE_INF = "-∞"

	string_uint64_max        = "18446744073709551616"
	string_uint64_max_broken = "18446744073709551615"

	string_FALSE = "FALSE"
	string_False = "False"
	string_false = "false"
	string_TRUE  = "TRUE"
	string_True  = "True"
	string_true  = "true"
)

func create_lettercase(s *hestiaTESTING.Scenario) hestiaFMT.Lettercase {
	switch {
	case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
		return LETTERCASE_UPPER
	case hestiaTESTING.HasCondition(s, cond_LOWERCASE):
		return LETTERCASE_LOWER
	}

	return 100
}

func create_base(s *hestiaTESTING.Scenario) uint64 {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0):
		return 0
	case hestiaTESTING.HasCondition(s, cond_BASE_1):
		return 1
	case hestiaTESTING.HasCondition(s, cond_BASE_2):
		return 2
	case hestiaTESTING.HasCondition(s, cond_BASE_5):
		return 5
	case hestiaTESTING.HasCondition(s, cond_BASE_8):
		return 8
	case hestiaTESTING.HasCondition(s, cond_BASE_10):
		return 10
	case hestiaTESTING.HasCondition(s, cond_BASE_12):
		return 12
	case hestiaTESTING.HasCondition(s, cond_BASE_16):
		return 16
	case hestiaTESTING.HasCondition(s, cond_BASE_22):
		return 22
	case hestiaTESTING.HasCondition(s, cond_BASE_36):
		return 36
	case hestiaTESTING.HasCondition(s, cond_BASE_37):
		return 37
	default:
		return 100
	}
}

func create_precision(s *hestiaTESTING.Scenario) uint64 {
	switch {
	case hestiaTESTING.HasCondition(s, cond_PRECISION_1):
		return 1
	case hestiaTESTING.HasCondition(s, cond_PRECISION_4):
		return 4
	case hestiaTESTING.HasCondition(s, cond_PRECISION_5):
		return 5
	case hestiaTESTING.HasCondition(s, cond_PRECISION_10):
		return 10
	default:
		return 0
	}
}

func create_notation(s *hestiaTESTING.Scenario) hestiaFMT.Notation {
	switch {
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC_AUTO):
		return NOTATION_SCIENTIFIC_AUTO
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3_AUTO):
		return NOTATION_ISO6093NR3_AUTO
	case hestiaTESTING.HasCondition(s, cond_NOTATION_SCIENTIFIC):
		return NOTATION_SCIENTIFIC
	case hestiaTESTING.HasCondition(s, cond_NOTATION_ISO6093NR3):
		return NOTATION_ISO6093NR3
	case hestiaTESTING.HasCondition(s, cond_NOTATION_DECIMAL_NO_EXPONENT):
		return NOTATION_DECIMAL_NO_EXPONENT
	case hestiaTESTING.HasCondition(s, cond_NOTATION_IEEE754):
		return NOTATION_IEEE754
	case hestiaTESTING.HasCondition(s, cond_NOTATION_UNKNOWN):
		return 100
	default:
		return 0
	}
}

//nolint:gocognit,goconst
func create_float32_subject(s *hestiaTESTING.Scenario) (subject float32) {
	if hestiaTESTING.HasCondition(s, cond_VALUE_ZERO) {
		return 0.0
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		return hestiaMATH.S32_IEEE754_BitsToFloat(
			hestiaMATH.MASK_FLOAT32_EXPONENT | hestiaMATH.MASK_FLOAT32_MANTISSA,
		)
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		return hestiaMATH.S32_IEEE754_BitsToFloat(
			hestiaMATH.MASK_FLOAT32_EXPONENT,
		)
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return hestiaMATH.S32_IEEE754_BitsToFloat(
			hestiaMATH.MASK_FLOAT32_EXPONENT | hestiaMATH.MASK_FLOAT32_SIGN,
		)
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
		subject = 123456.0
	case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
		subject = 0.0
	case hestiaTESTING.HasCondition(s, cond_ROUND_ONE):
		subject = 1.0
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		subject += 0.7891234
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO):
		subject += 0.0
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_ONE):
		subject += 0.1
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_ZERO):
		// do nothing
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
			subject *= 1e-30
		} else {
			subject *= 1e30
		}
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
			subject *= 1e-2
		} else {
			subject *= 1e2
		}
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-13
			} else {
				subject *= 1e13
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-7
			} else {
				subject *= 1e7
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ONE) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-8
			} else {
				subject *= 1e8
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-6
			} else {
				subject *= 1e6
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ONE):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-7
			} else {
				subject *= 1e7
			}
		}
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
		subject *= -1
	}

	return subject
}

//nolint:gocognit
func create_float64_subject(s *hestiaTESTING.Scenario) (subject float64) {
	if hestiaTESTING.HasCondition(s, cond_VALUE_ZERO) {
		return 0.0
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_FLOAT_NAN):
		return hestiaMATH.S64_IEEE754_BitsToFloat(
			hestiaMATH.MASK_FLOAT64_EXPONENT | hestiaMATH.MASK_FLOAT64_MANTISSA,
		)
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_POSITIVE):
		return hestiaMATH.S64_IEEE754_BitsToFloat(
			hestiaMATH.MASK_FLOAT64_EXPONENT,
		)
	case hestiaTESTING.HasCondition(s, cond_FLOAT_INF_NEGATIVE):
		return hestiaMATH.S64_IEEE754_BitsToFloat(
			hestiaMATH.MASK_FLOAT64_EXPONENT | hestiaMATH.MASK_FLOAT64_SIGN,
		)
	case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL):
		subject = 123456.0
	case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO):
		subject = 0.0
	case hestiaTESTING.HasCondition(s, cond_ROUND_ONE):
		subject = 1.0
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
		subject += 0.7891234
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO):
		subject += 0.0
	case hestiaTESTING.HasCondition(s, cond_PARTIAL_ONE):
		subject += 0.1
	}

	switch {
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_ZERO):
		// do nothing
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_LARGER_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
			subject *= 1e-30
		} else {
			subject *= 1e30
		}
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_SMALLER_MANTISSA):
		if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
			subject *= 1e-2
		} else {
			subject *= 1e2
		}
	case hestiaTESTING.HasCondition(s, cond_EXPONENT_EQUAL_MANTISSA):
		switch {
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-13
			} else {
				subject *= 1e13
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ZERO) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-7
			} else {
				subject *= 1e7
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_ONE) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_NORMAL):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-8
			} else {
				subject *= 1e8
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ZERO):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-6
			} else {
				subject *= 1e6
			}
		case hestiaTESTING.HasCondition(s, cond_ROUND_NORMAL) &&
			hestiaTESTING.HasCondition(s, cond_PARTIAL_ONE):
			if hestiaTESTING.HasCondition(s, cond_EXPONENT_NEGATIVE) {
				subject *= 1e-7
			} else {
				subject *= 1e7
			}
		}
	}

	if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
		subject *= -1
	}

	return subject
}

func create_sizes(s *hestiaTESTING.Scenario) uint16 {
	switch {
	case hestiaTESTING.HasCondition(s, cond_RESIZE_0):
		return 0
	case hestiaTESTING.HasCondition(s, cond_RESIZE_2):
		return 2
	case hestiaTESTING.HasCondition(s, cond_RESIZE_5):
		return 5
	case hestiaTESTING.HasCondition(s, cond_RESIZE_8):
		return 8
	case hestiaTESTING.HasCondition(s, cond_RESIZE_10):
		return 10
	case hestiaTESTING.HasCondition(s, cond_RESIZE_12):
		return 12
	case hestiaTESTING.HasCondition(s, cond_RESIZE_16):
		return 16
	case hestiaTESTING.HasCondition(s, cond_RESIZE_22):
		return 22
	case hestiaTESTING.HasCondition(s, cond_RESIZE_36):
		return 36
	case hestiaTESTING.HasCondition(s, cond_RESIZE_65):
		return 65
	default:
		return 64 // don't resize
	}
}
