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

package hestiaBITS

import (
	internal "hestia/hestiaINTERNAL/hestiaMATH"
)

func S64_TrailingZeros(x uint64) uint64 {
	return internal.S64_TrailingZeros_64Bits(x)
}

func S32_TrailingZeros(x uint32) uint32 {
	return internal.S32_TrailingZeros_32Bits(x)
}

func S16_TrailingZeros(x uint16) uint16 {
	return internal.S16_TrailingZeros_16Bits(x)
}

func S8_TrailingZeros(x uint8) uint8 {
	return internal.S8_TrailingZeros_8Bits(x)
}
