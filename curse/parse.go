/*
Copyright © 2020 Dakota Walsh

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package curse

import (
	"encoding/json"
)

// Date format = RFC3339

// An addon represents a distinct project on curseforge. It contains nearly
// everything you would see when visiting a mod's landing page in a web
// browser. It even contains a list of the latest files uploaded.
type Addon struct {
	Id                     int                     `json:"id"`
	Name                   string                  `json:"name"`
	Authors                []Author                `json:"authors"`
	Attachments            []Attachment            `json:"attachments"`
	WebsiteUrl             string                  `json:"websiteUrl"`
	GameId                 int                     `json:"gameId"`
	Summary                string                  `json:"summary"`
	DefaultFileId          int                     `json:"defaultFileId"`
	DownloadCount          float64                 `json:"downloadCount"`
	LatestFiles            []File                  `json:"latestFiles"`
	Categories             []Category              `json:"categories"`
	Status                 int                     `json:"status"`
	PrimaryCategoryId      int                     `json:"primaryCategoryId"`
	CategorySections       CategorySection         `json:"categorySection"`
	Slug                   string                  `json:"slug"`
	GameVersionLatestFiles []GameVersionLatestFile `json:"gameVersionLatestFiles"`
	IsFeatured             bool                    `json:"isFeatured"`
	PopularityScore        float64                 `json:"popularityScore"`
	GamePopularityRank     int                     `json:"gamePopularityRank"`
	PrimaryLanguage        string                  `json:"primaryLanguage"`
	GameSlug               string                  `json:"gameSlug"`
	GameName               string                  `json:"gameName"`
	PortalName             string                  `json:"portalName"`
	DateModified           string                  `json:"dateModified"`
	DateCreated            string                  `json:"dateCreated"`
	DateReleased           string                  `json:dateReleased"`
	IsAvailable            bool                    `json:isAvailable"`
	IsExperimental         bool                    `json:"isExperimental"`
}

// An author is a user's profile on curseforge. An Addon can have several
// authors, but unforunately the author struct does not contain a list of the
// author's projects.
type Author struct {
	Name              string `json:"name"`
	Url               string `json:"url"`
	ProjectId         int    `json:"projectId"`
	Id                int    `json:"id"`
	ProjectTitleId    string `json:"projectTitleId"`
	ProjectTitleTitle string `json:"projectTitleTitle"`
	UserId            int    `json:"userId"`
	TwitchId          int    `json:"twitchId"`
}

// An attachement is a file uploaded to the Addon page that is NOT the mod
// itself. Normally this will be something like a screenshot or gif.
type Attachment struct {
	Id           int    `json:"id"`
	ProjectId    int    `json:"projectId"`
	Description  string `json:"description"`
	IsDefault    bool   `json:"isDefault"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	Status       int    `json:"status"`
}

// An addon can be in several categories chosed by the authors. This is useful
// for user's discovering new mods. These are the things on the sidebar when
// you're browsing through the mods on curseforge. Not to be confused with
// CategorySection which refers to if the Addon is a mod, modpack,
// resourcepack, world, and so on.
type Category struct {
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	AvatarUrl  string `json:"avatarUrl"`
	ParentId   int    `json:"parentId"`
	RootId     int    `json:"rootId"`
	Project    int    `json:"projectId"`
	AvatarId   int    `json:"avatarId"`
	GameId     int    `json:"gameId"`
}

// An addon only has one CategorySection. The CategorySection refers to if the
// Addon is a mod, modpack, resourcepack, world, and so on. All mods will be in
// the "Mods" CategorySection.
type CategorySection struct {
	Id                      int    `json:"id"`
	GameId                  int    `json:"gameId"`
	Name                    string `json:"name"`
	PackageType             int    `json:"packageType"`
	Path                    string `json:"path"`
	InitialInclusionPattern string `json:"initialInclusionPattern"`
	ExtraIncludePattern     string `json:extraIncludePattern"`
	GameCategoryId          int    `json:"gameCategoryId"`
}

// An addon contains a list of GameVersionLatestFile(s) which essentially just
// tells you which file is the latest for each "GameVersion" this is obviously
// very useful, but it notably doesn't contain any information about the file
// other than the "ProjectFileId" which can then be used to get more info about
// the file in question.
type GameVersionLatestFile struct {
	GameVersion     string `json:"gameVersion"`
	ProjectFileId   int    `json:"projectFileId"`
	ProjectFileName string `json:"projectFileName"`
	FileType        int    `json:"fileType"`
}

// File represents a specific .jar mod file uploaded to curseforge as part of
// an Addon. It has lots of important information about the file and contains a
// DownloadUrl should you want to save it locally.
type File struct {
	Id                      int          `json:"id"`
	DisplayName             string       `json:"displayName"`
	FileName                string       `json:"fileName"`
	FileDate                string       `json:"fileDate"`
	FileLength              int          `json:"fileLength"`
	ReleaseType             int          `json:"releaseType"`
	FileStatus              int          `json:"fileStatus"`
	DownloadUrl             string       `json:"downloadUrl"`
	IsAlternate             bool         `json:"isAlternate"`
	AlternateFileId         int          `json:"alternateFileId"`
	Dependencies            []Dependency `json:"dependencies"`
	IsAvailable             bool         `json:"isAvailable"`
	Modules                 []Module     `json:"modules"`
	PackageFingerprint      int          `json:"packageFingerprint"`
	GameVersion             []string     `json:"gameVersion"`
	InstallMetadata         string       `json:"installMetadata"`
	ServerPackFileId        string       `json:"serverPackFileId"`
	HasInstallScript        bool         `json:"hasInstallScript"`
	GameVersionDateReleased string       `json:"gameVersionDateReleased"`
	GameVersionFlavor       string       `json:"gameVersionFlavor"`
}

// An Addon can mark another Addon as a dependency. In this case the dependency
// should also be automatically fetched.
type Dependency struct {
	AddonId int `json:"addonId"`
	Type    int `json:"type"`
}

// A File contains information about the specific files an folder inside the
// .jar which can be downloaded. Normally a .jar will have a META-INF,
// mcmod.info, pack.mcmeta, and a folder with the class files. This varies with
// different loaders. Additionally a fingerprint is given which could later be
// used for verification.
type Module struct {
	Foldername  string `json:"foldername"`
	Fingerprint int    `json:"fingerprint"`
}

// A Minecraft Version contains information about a particular update of minecraft.
type MinecraftVersion struct {
	Id                    int    `json:"id"`
	GameVersionId         int    `json:"gameVersionId"`
	VersionString         string `json:"versionString"`
	JarDownloadUrl        string `json:"jarDownloadUrl"`
	JsonDownloadUrl       string `json:"jsonDownloadUrl"`
	Approved              bool   `json:"approved"`
	DateModified          string `json:"dateModified"`
	GameVersionTypeId     int    `json:"gameVersionTypeId"`
	GameVersionStatus     int    `json:"gameVersionStatus"`
	GameVersionTypeStatus int    `json:"gameVersionTypeStatus"`
}

// A Modloader defines the properties of one of the modloaders supported by
// curseforge. Currently is seems to only support forge so this isn't very
// useful for fabric mods.
type Modloader struct {
	Name         string `json:"name"`
	GameVersion  string `json:"gameVersion"`
	Latest       bool   `json:"latest"`
	Recommended  bool   `json:"recommended"`
	DateModified string `json:"dateModified"`
}

func parseAddonInfo(b []byte) (Addon, error) {
	var addon Addon
	err := json.Unmarshal(b, &addon)
	return addon, err
}

func parseAddonFileInformation(b []byte) (File, error) {
	var file File
	err := json.Unmarshal(b, &file)
	return file, err
}

func parseAddonFileDownloadURL(b []byte) string {
	return string(b)
}

func parseAddonFiles(b []byte) ([]File, error) {
	var files []File
	err := json.Unmarshal(b, &files)
	return files, err
}

func parseMinecraftVersionList(b []byte) ([]MinecraftVersion, error) {
	var minecraftVersions []MinecraftVersion
	err := json.Unmarshal(b, &minecraftVersions)
	return minecraftVersions, err
}

func parseModloaderList(b []byte) ([]Modloader, error) {
	var modloaders []Modloader
	err := json.Unmarshal(b, &modloaders)
	return modloaders, err
}
