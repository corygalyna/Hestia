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

package hestiaTESTING

// test suite
const (
	suite_LOGF_API         = "hestiaTESTING Logf API"
	suite_HAS_EXECUTED_API = "hestiaTESTING HasExecuted API"
	suite_CONCLUDE_API     = "hestiaTESTING Conclude API"
	suite_CONCLUSION_API   = "hestiaTESTING Conclusion API"
	suite_TO_STRING_API    = "hestiaTESTING ToString API"
	suite_TO_TOML_API      = "hestiaTESTING ToTOML API"
)

// all test switches
const (
	// scenario
	cond_SUPPLY_NIL_SCENARIO = "supply scenario parameter as nil"

	// name
	cond_PROPER_NAME = "configure proper name"
	cond_EMPTY_NAME  = "configure empty name"

	// verdict
	cond_PROPER_VERDICT  = "configure VERDICT_PASS verdict"
	cond_FAIL_VERDICT    = "configure VERDICT_FAIL verdict"
	cond_SKIP_VERDICT    = "configure VERDICT_SKIP verdict"
	cond_UNKNOWN_VERDICT = "configure priv_VERDICT_UNKNOWN verdict"

	// description
	cond_PROPER_DESCRIPTION = "configure proper description"
	cond_EMPTY_DESCRIPTION  = "configure empty description"

	// log
	cond_PROPER_LOG = "configure proper log list"
	cond_EMPTY_LOG  = "configure empty log list"
	cond_NIL_LOG    = "configure nil log list"

	// switches
	cond_PROPER_SWITCHES = "configure proper switches list"
	cond_EMPTY_SWITCHES  = "configure empty switches list"
	cond_NIL_SWITCHES    = "configure nil switches list"

	// format
	cond_PROPER_STRING_FORMAT = "use proper string format"
	cond_EMPTY_STRING_FORMAT  = "use empty string format"

	// args
	cond_PROPER_STRING_ARGUMENTS = "use proper string arguments"
	cond_EMPTY_STRING_ARGUMENTS  = "use empty string arguments"
	cond_NIL_STRING_ARGUMENTS    = "use nil string arguments"

	// expectations
	expect_PANIC         = "expecting panic"
	expect_OUTPUT_STRING = "expecting string output"
)

// all test values
const (
	/// name
	value_NAME_PROPER = "Test Scenario Object"

	// description
	value_DESCRIPTION_PROPER = "A Test Scenario Object Description"

	// log
	value_LOG_LINE_1 = "Log Line 1"
	value_LOG_LINE_2 = "Log Line 1"

	// switches
	value_SWITCH_1 = "Switch 1"
	value_SWITCH_2 = "Switch 2"

	// string format
	value_FORMAT_1                 = "formatted %v %v %v"
	value_FORMAT_1_LOGF_SUCCESSFUL = "formatted proper1 5 true"
	value_FORMAT_2_LOGF_SUCCESSFUL = "formatted %!v(MISSING) %!v(MISSING) %!v(MISSING)"
	value_FORMAT_3_LOGF_SUCCESSFUL = "formatted <nil> %!v(MISSING) %!v(MISSING)"

	// string arguments
	value_ARG_1                  = "proper1"
	value_ARG_2                  = 5
	value_ARG_3                  = true
	value_ARG_1_LOGLN_SUCCESSFUL = "proper1 5 true\n"
	value_ARG_2_LOGLN_SUCCESSFUL = "\n"
	value_ARG_3_LOGLN_SUCCESSFUL = "<nil>\n"

	// verdict
	value_DEFAULT_VERDICT Verdict = 125
)

func testlib_AssertOutputString(s *Scenario, output string) bool {
	if output != "" {
		return s.Switches[expect_OUTPUT_STRING]
	}

	return !s.Switches[expect_OUTPUT_STRING]
}

func testlib_AssertPanic(s *Scenario, panick string) bool {
	if panick != "" {
		return s.Switches[expect_PANIC]
	}

	return !s.Switches[expect_PANIC]
}

func testlib_AssertConclude(s, ts *Scenario) bool {
	switch {
	case ts.verdict == value_DEFAULT_VERDICT:
		return s.Switches[cond_SUPPLY_NIL_SCENARIO] || s.Switches[expect_PANIC]
	case ts.verdict == VERDICT_PASS:
		return s.Switches[cond_PROPER_VERDICT]
	case ts.verdict == VERDICT_FAIL:
		return s.Switches[cond_FAIL_VERDICT]
	case ts.verdict == VERDICT_SKIP:
		return s.Switches[cond_SKIP_VERDICT]
	}

	return false
}

func testlib_GenerateVerdict(s *Scenario) Verdict {
	switch {
	case s.Switches[cond_PROPER_VERDICT]:
		return VERDICT_PASS
	case s.Switches[cond_FAIL_VERDICT]:
		return VERDICT_FAIL
	case s.Switches[cond_SKIP_VERDICT]:
		return VERDICT_SKIP
	default:
		return priv_VERDICT_UNKNOWN
	}
}

func testlib_AssertConclusion(s *Scenario, output Verdict) bool {
	switch {
	case output == priv_VERDICT_UNKNOWN:
		return s.Switches[cond_UNKNOWN_VERDICT]
	case output == VERDICT_PASS:
		return s.Switches[cond_PROPER_VERDICT]
	case output == VERDICT_SKIP:
		if !s.Switches[cond_PROPER_VERDICT] || s.Switches[cond_SUPPLY_NIL_SCENARIO] {
			return true
		}
	}

	return false
}

func testlib_AssertHasExecuted(s *Scenario, output bool) bool {
	if !output {
		switch {
		case s.Switches[cond_SUPPLY_NIL_SCENARIO],
			s.Switches[cond_UNKNOWN_VERDICT]:
			return true
		}
	}
	// is true

	return s.Switches[cond_PROPER_VERDICT]
}

func testlib_ConfigureVerdict(s, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_VERDICT]:
		ts.verdict = VERDICT_PASS
	case s.Switches[cond_UNKNOWN_VERDICT]:
		ts.verdict = priv_VERDICT_UNKNOWN
	default:
		ts.verdict = VERDICT_PASS
	}
}

func testlib_AssertLog(s, ts *Scenario) bool {
	if ts == nil {
		return s.Switches[cond_SUPPLY_NIL_SCENARIO]
	}
	// ts is now available

	if s.Switches[expect_PANIC] {
		switch {
		case s.Switches[cond_EMPTY_STRING_FORMAT],
			s.Switches[cond_SUPPLY_NIL_SCENARIO]:
			return true
		}
	}
	// no longer in panic expectation.

	if ts.Log == nil {
		return false
	}
	// ts.Log is now available

	for _, v := range ts.Log {
		switch v {
		case value_FORMAT_1_LOGF_SUCCESSFUL:
			if s.Switches[cond_PROPER_STRING_FORMAT] &&
				s.Switches[cond_PROPER_STRING_ARGUMENTS] {
				return true
			}
		case value_FORMAT_2_LOGF_SUCCESSFUL:
			if s.Switches[cond_PROPER_STRING_FORMAT] &&
				s.Switches[cond_EMPTY_STRING_ARGUMENTS] {
				return true
			}
		case value_FORMAT_3_LOGF_SUCCESSFUL:
			if s.Switches[cond_PROPER_STRING_FORMAT] &&
				s.Switches[cond_NIL_STRING_ARGUMENTS] {
				return true
			}
		case value_ARG_1_LOGLN_SUCCESSFUL:
			if s.Switches[cond_PROPER_STRING_ARGUMENTS] {
				return true
			}
		case value_ARG_2_LOGLN_SUCCESSFUL:
			if s.Switches[cond_EMPTY_STRING_ARGUMENTS] {
				return true
			}
		case value_ARG_3_LOGLN_SUCCESSFUL:
			if s.Switches[cond_NIL_STRING_ARGUMENTS] {
				return true
			}
		}
	}

	return false
}

func testlib_CreateStringArguments(s *Scenario) []any {
	switch {
	case s.Switches[cond_PROPER_STRING_ARGUMENTS]:
		return []any{value_ARG_1, value_ARG_2, value_ARG_3}
	case s.Switches[cond_EMPTY_STRING_ARGUMENTS]:
		return []any{}
	case s.Switches[cond_NIL_STRING_ARGUMENTS]:
		return nil
	}

	return nil
}

func testlib_CreateStringFormat(s *Scenario) string {
	switch {
	case s.Switches[cond_PROPER_STRING_FORMAT]:
		return value_FORMAT_1
	case s.Switches[cond_EMPTY_STRING_FORMAT]:
		return ""
	}

	return ""
}

func testlib_ConfigureName(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_NAME]:
		ts.Name = value_NAME_PROPER
	case s.Switches[cond_EMPTY_NAME]:
		ts.Name = ""
	}
}

func testlib_ConfigureDescription(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_DESCRIPTION]:
		ts.Description = value_DESCRIPTION_PROPER
	case s.Switches[cond_EMPTY_DESCRIPTION]:
		ts.Description = ""
	}
}

func testlib_ConfigureLog(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_LOG]:
		ts.Log = []string{value_LOG_LINE_1, value_LOG_LINE_2}
	case s.Switches[cond_EMPTY_LOG]:
		ts.Log = []string{}
	case s.Switches[cond_NIL_LOG]:
		ts.Log = nil
	}
}

func testlib_ConfigureSwitches(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_SWITCHES]:
		ts.Switches = map[string]bool{
			value_SWITCH_1: true,
			value_SWITCH_2: false,
		}
	case s.Switches[cond_EMPTY_SWITCHES]:
		ts.Switches = map[string]bool{}
	case s.Switches[cond_NIL_SWITCHES]:
		ts.Switches = nil
	}
}