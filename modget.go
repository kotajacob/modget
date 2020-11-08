/* modget Copyright (C) 2020 Dakota Walsh */
package main

import (
	"flag"
	"fmt"
	"os"

	"git.sr.ht/~kota/modget/commands"
)

var Version string

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
}

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addMc := addCmd.String("mc", "", "specify a certain minecraft version")
	addLoader := addCmd.String("loader", "", "specify a certain modloader")

	delCmd := flag.NewFlagSet("del", flag.ExitOnError)

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	// Print help if no subcommand specified.
	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}

	switch os.Args[1] {
	// For each subcommand we parse its own flags.
	case "add":
		addCmd.Parse(os.Args[2:])
		mods := addCmd.Args()
		err := commands.Add(mods, *addMc, *addLoader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "del":
		delCmd.Parse(os.Args[2:])
		// DEBUG //
		fmt.Println("  tail:", delCmd.Args())
	case "update":
		updateCmd.Parse(os.Args[2:])
		// DEBUG //
		fmt.Println("  tail:", updateCmd.Args())
	case "help":
		help()
		os.Exit(0)
	case "version":
		fmt.Printf("modget " + Version + "\n")
	default:
		fmt.Printf("unknown subcommand\n\n")
		help()
		os.Exit(1)
	}
}
