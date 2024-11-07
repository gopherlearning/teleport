// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package packagemanager

func teleportAPTRepo(development bool) aptRepoMetadata {
	if development {
		return aptRepoMetadata{
			endpoint:       "https://apt.releases.development.teleport.dev/",
			pubKeyEndpoint: "https://apt.releases.development.teleport.dev/gpg",
			sourceListFile: "/etc/apt/sources.list.d/development.teleport.list",
			pubKeyFile:     "/usr/share/keyrings/development.teleport-archive-keyring.asc",
		}
	}

	return aptRepoMetadata{
		endpoint:       "https://apt.releases.teleport.dev/",
		pubKeyEndpoint: "https://apt.releases.teleport.dev/gpg",
		sourceListFile: "/etc/apt/sources.list.d/teleport.list",
		pubKeyFile:     "/usr/share/keyrings/teleport-archive-keyring.asc",
	}
}

func teleportYUMRepo(development bool) yumRepoMetadata {
	if development {
		return yumRepoMetadata{
			endpoint: "https://yum.releases.development.teleport.dev/",
		}
	}

	return yumRepoMetadata{
		endpoint: "https://yum.releases.teleport.dev/",
	}
}

func teleportZypperRepo(development bool) zypperRepoMetadata {
	if development {
		return zypperRepoMetadata{
			endpoint:       "https://zypper.releases.development.teleport.dev/",
			pubKeyEndpoint: "https://zypper.releases.development.teleport.dev/gpg",
		}
	}

	return zypperRepoMetadata{
		endpoint:       "https://zypper.releases.teleport.dev/",
		pubKeyEndpoint: "https://zypper.releases.teleport.dev/gpg",
	}
}
