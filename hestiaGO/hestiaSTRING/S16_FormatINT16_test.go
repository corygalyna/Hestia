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
	"hestia/hestiaERROR"
	"hestia/hestiaNUMBER"
	"hestia/hestiaTESTING"
	"testing"
)

func test_cases_S16_FormatINT16() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-2 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_2,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-2 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_2,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-2 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_2,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-2 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_2,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-5 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_5,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-5 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_2,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-5 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_2,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-5 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_2,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-8 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_8,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-8 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_8,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-8 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_8,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-8 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_8,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_10,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_10,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_10,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_10,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-12 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_12,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-12 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_12,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-12 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_12,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-12 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_12,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-16 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_16,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-16 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_16,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-16 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_16,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-16 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_16,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-22 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_22,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-22 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_22,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-22 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_22,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-22 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_22,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-36 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_36,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-36 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_36,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-36 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_36,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-36 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_36,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-1 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_1,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-1 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_1,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-1 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_1,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-1 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_1,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-0 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_0,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-0 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_0,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-0 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_0,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-0 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_0,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-37 lowercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_37,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-37 uppercase positive value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_37,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-37 lowercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_37,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-37 uppercase negative value.
`,
			Switches: []string{
				cond_VALUE_INT16,
				cond_BASE_37,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 lowercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_10,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 uppercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_10,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 lowercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_10,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 uppercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_10,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-16 lowercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_16,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-16 uppercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_16,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-16 lowercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_16,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-16 uppercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_16,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-36 lowercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_36,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-36 uppercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_36,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-36 lowercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_36,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-36 uppercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_36,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-37 lowercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_37,
				cond_LOWERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-37 uppercase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_37,
				cond_UPPERCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-37 lowercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_37,
				cond_LOWERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to raise error for base-37 uppercase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_37,
				cond_UPPERCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 unknowncase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_10,
				cond_UNKNOWNCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 unknowncase positive 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_10,
				cond_UNKNOWNCASE,
				cond_POSITIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 unknowncase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_10,
				cond_UNKNOWNCASE,
				cond_NEGATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/S16_FormatINT16 is able to process base-10 unknowncase negative 0 value.
`,
			Switches: []string{
				cond_VALUE_ZERO,
				cond_BASE_10,
				cond_UNKNOWNCASE,
				cond_NEGATIVE,
			},
		},
	}
}

func Test_S16_FormatINT16(t *testing.T) {
	scenarios := test_cases_S16_FormatINT16()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/S16_FormatINT16 API"

		// prepare
		subject := int16(hestiaNUMBER.MAX_INT16)
		switch {
		case hestiaTESTING.HasCondition(s, cond_VALUE_ZERO):
			subject = 0
		case hestiaTESTING.HasCondition(s, cond_NEGATIVE):
			subject = int16(hestiaNUMBER.MIN_INT16)
		}
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: %d", subject))

		base := uint16(create_base(s))
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Base	: %d", base))

		lettercase := create_lettercase(s)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Lettercase	: %d", lettercase))

		// test
		var output string
		err := hestiaERROR.OK
		panick := ""

		_panick := hestiaTESTING.Exec(func() any {
			output, err = S16_FormatINT16(subject, base, lettercase)
			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Output	: '%s'", output))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Error	: %d", err))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Panick	: '%v'", panick))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		if !assert_S16_FormatINT16_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S16_FormatINT16_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S16_FormatINT16_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S16_FormatINT16_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0),
		hestiaTESTING.HasCondition(s, cond_BASE_1),
		hestiaTESTING.HasCondition(s, cond_BASE_37):
		return err == hestiaERROR.DATA_INVALID
	}

	return err == hestiaERROR.OK
}

//nolint:gocognit
func assert_S16_FormatINT16_output(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0),
		hestiaTESTING.HasCondition(s, cond_BASE_1),
		hestiaTESTING.HasCondition(s, cond_BASE_37):
		return output == ""
	case hestiaTESTING.HasCondition(s, cond_VALUE_ZERO):
		return output == "0"
	case hestiaTESTING.HasCondition(s, cond_BASE_2):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output ==
				"-1000000000000000"
		}
		return output == "111111111111111"
	case hestiaTESTING.HasCondition(s, cond_BASE_5):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output == "-2022033"
		}
		return output == "2022032"
	case hestiaTESTING.HasCondition(s, cond_BASE_8):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output == "-100000"
		}
		return output == "77777"
	case hestiaTESTING.HasCondition(s, cond_BASE_10):
		if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
			return output == "-32768"
		}
		return output == "32767"
	case hestiaTESTING.HasCondition(s, cond_BASE_12):
		switch {
		case hestiaTESTING.HasCondition(s, cond_LOWERCASE),
			hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				return output == "-16b68"
			}
			return output == "16b67"
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				return output == "-16B68"
			}
			return output == "16B67"
		}
	case hestiaTESTING.HasCondition(s, cond_BASE_16):
		switch {
		case hestiaTESTING.HasCondition(s, cond_LOWERCASE),
			hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				return output == "-8000"
			}
			return output == "7fff"
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				return output == "-8000"
			}
			return output == "7FFF"
		}
	case hestiaTESTING.HasCondition(s, cond_BASE_22):
		switch {
		case hestiaTESTING.HasCondition(s, cond_LOWERCASE),
			hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				return output == "-31fa"
			}
			return output == "31f9"
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				return output == "-31FA"
			}
			return output == "31F9"
		}
	case hestiaTESTING.HasCondition(s, cond_BASE_36):
		switch {
		case hestiaTESTING.HasCondition(s, cond_LOWERCASE),
			hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				return output == "-pa8"
			}
			return output == "pa7"
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			if hestiaTESTING.HasCondition(s, cond_NEGATIVE) {
				return output == "-PA8"
			}
			return output == "PA7"
		}
	}

	return false
}

func assert_S16_FormatINT16_panick(panick string) bool {
	return panick == ""
}
