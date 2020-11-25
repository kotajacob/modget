modget(1) [""["modget-VERSION"]]

# NAME

modget - A package manager for minecraft curseforge mods.

# SYNOPSIS

*modget* <command> [options]

# COMMANDS

*add, a <MODID>*
	Download and install a mod based on its MODID.

*delete, d <MODID>*
	Remove and uninstall a mod based on its MODID.

*help <command>*
	Show help message and quit.

*mark, m <MODID>*
	Mark allows you to hold or unhold your mods. When a mod has a hold
	status it will be ignored by the update command and will remain at its
	exact version until the hold is removed or a different version is
	manually installed. By default the hold status of your selected mod(s)
	is toggled and printed.

*search <Search String>*
	Search curseforge for mods based on search terms.
	Prints mod names, MODID, Version, and a short description.

*show, sh <MODID>*
	Query and print more information about a specific mod by MODID.

*update, u <MODID>*
	Check each installed mod (in selected directory or current directory)
	and prompt to install any new mods.

*version*
	Show the version and quit.

# GLOBAL OPTIONS

*-p, --path*
	Specify an install directory to work in instead of the current 
	working directory.

# ADD OPTIONS

*-m, --minecraft <Minecraft Version>*
	Specify a minecraft version when installing mod(s). For example 1.16.3
	or 1.12

*-l, --loader <Modloader>*
	Specify a modloader installing mod(s). For example fabric, forge, rift,
	or liteloader.

# SHOW OPTIONS

*-l, --oneline*
	Display mods one per line.

# UPDATE OPTIONS

*-m, --minecraft <Minecraft Version>*
	Specify a new minecraft version. This will update your database to the
	new minecraft version and check if any/all of your mods have updates.
	Additionally you will be prompted for mods which do not support the
	version and may delete them. For example 1.16.3 or 1.12

*-l, --loader <Modloader>*
	Specify a new modloader. This will update your database to the new
	loader and check if any/all of your mods have updates. Additionally you
	will be prompted for mods which do not support the version and may
	delete them. For example fabric, forge, rift, or liteloader.

## Slug

Every mod on curseforge has a short name called a slug which is found in the
URL for the mod's page. For example in this url
https://www.curseforge.com/minecraft/mc-mods/appleskin the slug is `appleskin`.
The modget commands will attempt to convert slugs to MODIDs which makes
managing your mods much easier. Internally everything uses MODID. Currently
this feature is broken for the majority of curseforge mods.

# MODID

Every mod on curseforge has a 6 digit MODID which can be found on the mods page
or by using modget's search command. This is used instead of names when
installing and removing mods as that is what is supported by the curseforge
API.

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/modget/
