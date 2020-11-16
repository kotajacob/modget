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

// Addon represents a distinct project on curseforge. It contains nearly
// everything you would see when visiting a mod's landing page in a web
// browser. It even contains a list of the latest files uploaded.
type Addon struct {
	ID                     int                     `json:"id"`
	Name                   string                  `json:"name"`
	Authors                []Author                `json:"authors"`
	Attachments            []Attachment            `json:"attachments"`
	WebsiteURL             string                  `json:"websiteUrl"`
	GameID                 int                     `json:"gameId"`
	Summary                string                  `json:"summary"`
	DefaultFileID          int                     `json:"defaultFileId"`
	DownloadCount          float64                 `json:"downloadCount"`
	LatestFiles            []File                  `json:"latestFiles"`
	Categories             []Category              `json:"categories"`
	Status                 int                     `json:"status"`
	PrimaryCategoryID      int                     `json:"primaryCategoryId"`
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
	DateReleased           string                  `json:"dateReleased"`
	IsAvailable            bool                    `json:"isAvailable"`
	IsExperimental         bool                    `json:"isExperimental"`
}

// Author is a user's profile on curseforge. An Addon can have several
// authors, but unforunately the author struct does not contain a list of the
// author's projects.
type Author struct {
	Name              string `json:"name"`
	URL               string `json:"url"`
	ProjectID         int    `json:"projectId"`
	ID                int    `json:"id"`
	ProjectTitleID    string `json:"projectTitleId"`
	ProjectTitleTitle string `json:"projectTitleTitle"`
	UserID            int    `json:"userId"`
	TwitchID          int    `json:"twitchId"`
}

// Attachment is a file uploaded to the Addon page that is NOT the mod
// itself. Normally this will be something like a screenshot or gif.
type Attachment struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"projectId"`
	Description  string `json:"description"`
	IsDefault    bool   `json:"isDefault"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Status       int    `json:"status"`
}

// Category represents one of the sections on Curseforge. Addons can be in
// multiple Categories. They are useful for discovering new mods.
type Category struct {
	CategoryID int    `json:"categoryId"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	AvatarURL  string `json:"avatarUrl"`
	ParentID   int    `json:"parentId"`
	RootID     int    `json:"rootId"`
	Project    int    `json:"projectId"`
	AvatarID   int    `json:"avatarId"`
	GameID     int    `json:"gameId"`
}

// CategorySection refers to if the Addon is a mod, modpack, resourcepack,
// world, and so on. All mods will be in the "Mods" CategorySection.
type CategorySection struct {
	ID                      int    `json:"id"`
	GameID                  int    `json:"gameId"`
	Name                    string `json:"name"`
	PackageType             int    `json:"packageType"`
	Path                    string `json:"path"`
	InitialInclusionPattern string `json:"initialInclusionPattern"`
	ExtraIncludePattern     string `json:"extraIncludePattern"`
	GameCategoryID          int    `json:"gameCategoryId"`
}

// GameVersionLatestFile essentially tells you which file is the latest for
// each "GameVersion" this is obviously very useful, but it notably doesn't
// contain any information about the file other than the "ProjectFileID" which
// can then be used to get more info about the file in question.
type GameVersionLatestFile struct {
	GameVersion     string `json:"gameVersion"`
	ProjectFileID   int    `json:"projectFileId"`
	ProjectFileName string `json:"projectFileName"`
	FileType        int    `json:"fileType"`
}

// File represents a specific .jar mod file uploaded to curseforge as part of
// an Addon. It has lots of important information about the file and contains a
// DownloadURL should you want to save it locally.
type File struct {
	ID                      int          `json:"id"`
	DisplayName             string       `json:"displayName"`
	FileName                string       `json:"fileName"`
	FileDate                string       `json:"fileDate"`
	FileLength              int          `json:"fileLength"`
	ReleaseType             int          `json:"releaseType"`
	FileStatus              int          `json:"fileStatus"`
	DownloadURL             string       `json:"downloadUrl"`
	IsAlternate             bool         `json:"isAlternate"`
	AlternateFileID         int          `json:"alternateFileId"`
	Dependencies            []Dependency `json:"dependencies"`
	IsAvailable             bool         `json:"isAvailable"`
	Modules                 []Module     `json:"modules"`
	PackageFingerprint      int          `json:"packageFingerprint"`
	GameVersion             []string     `json:"gameVersion"`
	InstallMetadata         string       `json:"installMetadata"`
	ServerPackFileID        string       `json:"serverPackFileId"`
	HasInstallScript        bool         `json:"hasInstallScript"`
	GameVersionDateReleased string       `json:"gameVersionDateReleased"`
	GameVersionFlavor       string       `json:"gameVersionFlavor"`
}

// Dependency represents an Addon that is required by a certain other Addon.
type Dependency struct {
	AddonID int `json:"addonId"`
	Type    int `json:"type"`
}

// Module is the content inside a File's .jar such as  META-INF, mcmod.info,
// pack.mcmeta, and a folder with the class files. This varies with different
// loaders. Additionally a fingerprint is given which could later be used for
// verification.
type Module struct {
	Foldername  string `json:"foldername"`
	Fingerprint int    `json:"fingerprint"`
}

// MinecraftVersion contains information about a particular update of minecraft.
type MinecraftVersion struct {
	ID                    int    `json:"id"`
	GameVersionID         int    `json:"gameVersionId"`
	VersionString         string `json:"versionString"`
	JarDownloadURL        string `json:"jarDownloadUrl"`
	JSONDownloadURL       string `json:"jsonDownloadUrl"`
	Approved              bool   `json:"approved"`
	DateModified          string `json:"dateModified"`
	GameVersionTypeID     int    `json:"gameVersionTypeId"`
	GameVersionStatus     int    `json:"gameVersionStatus"`
	GameVersionTypeStatus int    `json:"gameVersionTypeStatus"`
}

// Modloader defines the properties of one of the modloaders supported by
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

func parseAddonSearch(b []byte) ([]Addon, error) {
	var addons []Addon
	err := json.Unmarshal(b, &addons)
	return addons, err
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
