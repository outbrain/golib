/*
   Copyright 2014 Shlomi Noach.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package math

func MinInt(i1, i2 int) int {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func MaxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}
