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

import (
	"testing"
)

// test suite
const (
	suite_TO_STRING_API = "hestiaTESTING ToString API"
	suite_TO_JSON_API   = "hestiaTESTING ToJSON API"
	suite_TO_TOML_API   = "hestiaTESTING ToTOML API"
	suite_TO_YAML_API   = "hestiaTESTING ToYAML API"
)

// all test switches
const (
	// scenario
	cond_SUPPLY_NIL_SCENARIO = "supply scenario parameter as nil"

	// register
	cond_PROPER_REGISTRATION      = "configure registration properly"
	cond_FAULTY_SKIP_REGISTRATION = "configure registration with faulty skip"
	cond_FAULTY_FAIL_REGISTRATION = "configure registration with faulty fail"
	cond_FAULTY_BOTH_REGISTRATION = "configure registration with faulty skip and fail"

	// name
	cond_PROPER_NAME = "configure proper name"
	cond_EMPTY_NAME  = "configure empty name"

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

	// expectations
	expect_PANIC         = "expecting panic"
	expect_OUTPUT_STRING = "expecting string output"
)

// all test values
const (
	/// name
	t_NAME_PROPER = "Test Scenario Object"

	// description
	t_DESCRIPTION_PROPER = "A Test Scenario Object Description"

	// log
	t_LOG_LINE_1 = "Log Line 1"
	t_LOG_LINE_2 = "Log Line 1"

	// switches
	t_SWITCH_1 = "Switch 1"
	t_SWITCH_2 = "Switch 2"
)

func testlib_ConfigureRegistrations(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_REGISTRATION]:
		ts.controller = &testing.T{}
		ts.skip = func() {}
		ts.fail = func() {}
	case s.Switches[cond_FAULTY_SKIP_REGISTRATION]:
		ts.controller = &testing.T{}
		ts.skip = nil
		ts.fail = func() {}
	case s.Switches[cond_FAULTY_FAIL_REGISTRATION]:
		ts.controller = &testing.T{}
		ts.skip = func() {}
		ts.fail = nil
	case s.Switches[cond_FAULTY_BOTH_REGISTRATION]:
		ts.controller = &testing.T{}
		ts.skip = nil
		ts.fail = nil
	}
}

func testlib_ConfigureName(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_NAME]:
		ts.Name = t_NAME_PROPER
	case s.Switches[cond_EMPTY_NAME]:
		ts.Name = ""
	}
}

func testlib_ConfigureDescription(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_DESCRIPTION]:
		ts.Description = t_DESCRIPTION_PROPER
	case s.Switches[cond_EMPTY_DESCRIPTION]:
		ts.Description = ""
	}
}

func testlib_ConfigureLog(s *Scenario, ts *Scenario) {
	switch {
	case s.Switches[cond_PROPER_LOG]:
		ts.Log = []string{t_LOG_LINE_1, t_LOG_LINE_2}
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
			t_SWITCH_1: true,
			t_SWITCH_2: false,
		}
	case s.Switches[cond_EMPTY_SWITCHES]:
		ts.Switches = map[string]bool{}
	case s.Switches[cond_NIL_SWITCHES]:
		ts.Switches = nil
	}
}