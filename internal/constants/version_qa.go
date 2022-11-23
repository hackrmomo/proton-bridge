// Copyright (c) 2022 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge. If not, see <https://www.gnu.org/licenses/>.

//go:build build_qa
// +build build_qa

package constants

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

// AppVersion returns the full rendered version of the app (to be used in request headers).
func AppVersion(version string) string {
	ver, _ := semver.MustParse(version).SetPrerelease("dev")

	return fmt.Sprintf("%v-%v@%v", getAPIOS(), AppName, ver.String())
}
