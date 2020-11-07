/* modget Copyright (C) 2020 Dakota Walsh */
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var Version string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func help() {
	fmt.Printf("modget " + Version + "\n")
	fmt.Printf("Usage: modget command\n\n")
	fmt.Printf("modget is a command line package manager for curseforge minecraft mods. It \nprovides commands for searching, managing, and querying information about mods.\n\n")
	fmt.Printf("Commands:\n")
	fmt.Printf("\tadd - Download and install a mod based on its MODID.\n")
	fmt.Printf("\tdel - Remove and uninstall a mod based on its MODID.\n")
	fmt.Printf("\tupdate - Check each installed mod and prompt to install any new mods.\n")
	fmt.Printf("\tshow - Query and print more information about a specific mod by MODID.\n")
	fmt.Printf("\tsearch - Search curseforge for mods based on search terms.\n")

	os.Exit(1)
}

func add(mods []string, mc string, loader string) {
	// DEBUG //
	fmt.Println("  tail:", mods)
	fmt.Println("  mc: ", mc)
	fmt.Println("  loader: ", loader)
	modid, err := strconv.Atoi(mods[0])
	check(err)
	ParseAddonInfo(GetAddonInfo(modid))
	ParseAddonFiles(GetAddonFiles(modid))
}

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addMc := addCmd.String("mc", "", "specify a certain minecraft version")
	addLoader := addCmd.String("loader", "", "specify a certain modloader")

	delCmd := flag.NewFlagSet("del", flag.ExitOnError)

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	if len(os.Args) < 2 {
		help()
	}

	switch os.Args[1] {
	// For each subcommand we parse its own flags.
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'add'")
		mods := addCmd.Args()
		if len(mods) > 0 {
			add(mods, *addMc, *addLoader)
		} else {
			fmt.Println("add subcommand requires at least one mod")
		}
	case "del":
		delCmd.Parse(os.Args[2:])
		// DEBUG //
		fmt.Println("subcommand 'del'")
		fmt.Println("  tail:", delCmd.Args())
	case "update":
		updateCmd.Parse(os.Args[2:])
		// DEBUG //
		fmt.Println("subcommand 'update'")
		fmt.Println("  tail:", updateCmd.Args())
	case "help":
		help()
	case "version":
		fmt.Printf("modget " + Version + "\n")
	default:
		fmt.Printf("unknown subcommand\n\n")
		help()
	}

	/* read the first argument as modid */
	// ParseAddonInfo(GetAddonInfo(modid))
	// ParseAddonFiles(GetAddonFiles(modid))
}
