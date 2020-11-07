package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

type GameVersionLatestFile struct {
	GameVersion     string `json:"gameVersion"`
	ProjectFileId   int    `json:"projectFileId"`
	ProjectFileName string `json:"projectFileName"`
	FileType        int    `json:"fileType"`
}

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

type Dependency struct {
	AddonId int `json:"addonId"`
	Type    int `json:"type"`
}

type Module struct {
	Foldername  string `json:"foldername"`
	Fingerprint int    `json:"fingerprint"`
}

func parseAddonInfo(b []byte) {
	var addon Addon

	err := json.Unmarshal(b, &addon)
	check(err)

	fmt.Println("Addon Name: " + addon.Name)
	fmt.Println("Addon Summary: " + addon.Summary)
	fmt.Printf("Popularity: %f\n", addon.PopularityScore)
	fmt.Printf("Downloads: %f\n", addon.DownloadCount)
}

func parseAddonFiles(b []byte) {
	var files []File

	err := json.Unmarshal(b, &files)
	check(err)

	for i := 0; i < len(files); i++ {
		fmt.Println("File ID: " + strconv.Itoa(files[i].Id))
		fmt.Println("File Date: " + files[i].FileDate)
		fmt.Println("File Name: " + files[i].FileName)
		for ii := 0; ii < len(files[i].GameVersion); ii++ {
			fmt.Println("Game Version: " + files[i].GameVersion[ii])
		}
		fmt.Println("URL: " + files[i].DownloadUrl)
		fmt.Println("")
	}
}
