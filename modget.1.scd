modget(1) [""["modget-VERSION"]]

# NAME

modget - A package manager for minecraft curseforge mods.

# SYNOPSIS

*greetd* [options] <command>

# OPTIONS

*-h, --help*
	Show help message and quit.

*--version*
	Show the version and quit.

*-d, --dir*
	Specify an install directory to work in instead of the current 
	working directory.

*-y, --yes*
	Answer yes to any prompts automatically.

# COMMANDS

*add, install <MODID>*
	Download and install a mod based on its MODID.

*del, remove, uninstall <MODID>*
	Remove and uninstall a mod based on its MODID.

*update <MODID>*
	Check each installed mod (in selected directory or current directory)
	and prompt to install any new mods.

*show <MODID>*
	Query and print more information about a specific mod by MODID.

*search <Search String>*
	Search curseforge for mods based on search terms.
	Prints mod names, MODID, Version, and a short description.

# MODID

Every mod on curseforge has a 6 digit MODID which can be found on the mods page
or by using modget's search command. This is used instead of names when
installing and removing mods as multiple mods could technically share the same
name.

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/modget/
