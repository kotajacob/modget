package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Files struct {
	Files []File
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

func parse() {
	b, err := ioutil.ReadFile("pretty.json")
	check(err)

	var files Files
	err = json.Unmarshal(b, &files)
	check(err)

	fmt.Println(files)
}
