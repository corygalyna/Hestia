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

func test_cases_S64_FormatUINT64() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-2 lowercase value.
`,
			Switches: []string{
				cond_BASE_2,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-2 uppercase value.
`,
			Switches: []string{
				cond_BASE_2,
				cond_VALUE_UINT64,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-5 lowercase value.
`,
			Switches: []string{
				cond_BASE_5,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-5 uppercase value.
`,
			Switches: []string{
				cond_BASE_5,
				cond_VALUE_UINT64,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-8 lowercase value.
`,
			Switches: []string{
				cond_BASE_8,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-8 uppercase value.
`,
			Switches: []string{
				cond_BASE_8,
				cond_VALUE_UINT64,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-10 lowercase value.
`,
			Switches: []string{
				cond_BASE_10,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-10 uppercase value.
`,
			Switches: []string{
				cond_BASE_10,
				cond_VALUE_UINT64,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-12 lowercase value.
`,
			Switches: []string{
				cond_BASE_12,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-12 uppercase value.
`,
			Switches: []string{
				cond_BASE_12,
				cond_VALUE_UINT64,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-16 lowercase value.
`,
			Switches: []string{
				cond_BASE_16,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-16 uppercase value.
`,
			Switches: []string{
				cond_BASE_16,
				cond_VALUE_UINT64,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-22 lowercase value.
`,
			Switches: []string{
				cond_BASE_22,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-22 uppercase value.
`,
			Switches: []string{
				cond_BASE_22,
				cond_VALUE_UINT64,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-36 lowercase value.
`,
			Switches: []string{
				cond_BASE_36,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-36 uppercase value.
`,
			Switches: []string{
				cond_BASE_36,
				cond_VALUE_UINT64,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to raise error for base-1 lowercase value.
`,
			Switches: []string{
				cond_BASE_1,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to raise error for base-0 lowercase value.
`,
			Switches: []string{
				cond_BASE_0,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to raise error for base-37 lowercase value.
`,
			Switches: []string{
				cond_BASE_37,
				cond_VALUE_UINT64,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-10 uppercase 0 value.
`,
			Switches: []string{
				cond_BASE_10,
				cond_VALUE_ZERO,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-16 uppercase 0 value.
`,
			Switches: []string{
				cond_BASE_16,
				cond_VALUE_ZERO,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-36 uppercase 0 value.
`,
			Switches: []string{
				cond_BASE_36,
				cond_VALUE_ZERO,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to raise error for base-37 lowercase 0 value.
`,
			Switches: []string{
				cond_BASE_37,
				cond_VALUE_ZERO,
				cond_LOWERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to raise error for base-37 uppercase 0 value.
`,
			Switches: []string{
				cond_BASE_37,
				cond_VALUE_ZERO,
				cond_UPPERCASE,
			},
		}, {
			Description: `
test hestiaSTRING/S64_FormatUINT64 is able to process base-10 unknown-case value.
`,
			Switches: []string{
				cond_BASE_10,
				cond_VALUE_UINT64,
				cond_UNKNOWNCASE,
			},
		},
	}
}

func Test_S64_FormatUINT64(t *testing.T) {
	scenarios := test_cases_S64_FormatUINT64()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/S64_FormatUINT64 API"

		// prepare
		subject := uint64(hestiaNUMBER.MAX_UINT64)
		if hestiaTESTING.HasCondition(s, cond_VALUE_ZERO) {
			subject = 0
		}

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: %d", subject))

		base := create_base(s)
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
			output, err = S64_FormatUINT64(subject, base, lettercase)
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

		if !assert_S64_FormatUINT64_panick(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S64_FormatUINT64_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_S64_FormatUINT64_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_S64_FormatUINT64_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0),
		hestiaTESTING.HasCondition(s, cond_BASE_1),
		hestiaTESTING.HasCondition(s, cond_BASE_37):
		return err == hestiaERROR.DATA_INVALID
	}

	return err == hestiaERROR.OK
}

func assert_S64_FormatUINT64_output(s *hestiaTESTING.Scenario, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_BASE_0),
		hestiaTESTING.HasCondition(s, cond_BASE_1),
		hestiaTESTING.HasCondition(s, cond_BASE_37):
		return output == ""
	case hestiaTESTING.HasCondition(s, cond_VALUE_ZERO):
		return output == "0"
	case hestiaTESTING.HasCondition(s, cond_BASE_2):
		return output == "1111111111111111111111111111111111111111111111111111111111111111"
	case hestiaTESTING.HasCondition(s, cond_BASE_5):
		return output == "2214220303114400424121122430"
	case hestiaTESTING.HasCondition(s, cond_BASE_8):
		return output == "1777777777777777777777"
	case hestiaTESTING.HasCondition(s, cond_BASE_10):
		return output == string_uint64_max_broken
	case hestiaTESTING.HasCondition(s, cond_BASE_12):
		switch {
		case hestiaTESTING.HasCondition(s, cond_LOWERCASE),
			hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
			return output == "839365134a2a240713"
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			return output == "839365134A2A240713"
		}
	case hestiaTESTING.HasCondition(s, cond_BASE_16):
		switch {
		case hestiaTESTING.HasCondition(s, cond_LOWERCASE),
			hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
			return output == "ffffffffffffffff"
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			return output == "FFFFFFFFFFFFFFFF"
		}
	case hestiaTESTING.HasCondition(s, cond_BASE_22):
		switch {
		case hestiaTESTING.HasCondition(s, cond_LOWERCASE),
			hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
			return output == "2l4lf104353j8kf"
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			return output == "2L4LF104353J8KF"
		}
	case hestiaTESTING.HasCondition(s, cond_BASE_36):
		switch {
		case hestiaTESTING.HasCondition(s, cond_LOWERCASE),
			hestiaTESTING.HasCondition(s, cond_UNKNOWNCASE):
			return output == "3w5e11264sgsf"
		case hestiaTESTING.HasCondition(s, cond_UPPERCASE):
			return output == "3W5E11264SGSF"
		}
	}

	return false
}

func assert_S64_FormatUINT64_panick(panick string) bool {
	return panick == ""
}
