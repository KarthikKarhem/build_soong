// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package java

import (
	"android/soong/android"
)

func (library *Library) AndroidMk() (ret android.AndroidMkData, err error) {
	ret.Class = "JAVA_LIBRARIES"
	ret.OutputFile = android.OptionalPathForPath(library.outputFile)
	return
}

func (prebuilt *Prebuilt) AndroidMk() (ret android.AndroidMkData, err error) {
	ret.Class = "JAVA_LIBRARIES"
	ret.OutputFile = android.OptionalPathForPath(prebuilt.classpathFile)
	return
}
