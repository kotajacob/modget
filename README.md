# modget

Modget is a package manager for minecraft curseforge mods. It provides commands
for searching, managing, and querying information about mods. Currently it is
unreleased and in an early development state. The below commands are slowly
being implemented.

## install

modget requires the following build dependencies:

- `golang >= 1.13`
- `scdoc`
- `make`

1. `git clone https://git.sr.ht/~kota/mdtest`
2. `make all`
3. `sudo make install`

## uninstall

1. `sudo make uninstall`

## commands

- `add <MODID/Slug>`
	- Download and install a mod based on its MODID or Slug.

- `del <MODID/Slug>`
	- NOT YET IMPLEMENTED
	- Remove and uninstall a mod based on its MODID or Slug.

- `help <command>`
	- Print help about a command and exit.

- `search <Search String>`
	- Search curseforge for mods based on search terms. Prints mod names, MODID, Version, and a short description.

- `show <MODID/Slug>`
	- NOT YET IMPLEMENTED
	- Query and print more information about a specific mod by MODID Slug.

- `update <MODID/Slug>`
	- NOT YET IMPLEMENTED
	- Check each installed mod (in selected directory or current directory) and prompt to install any new mods.

- `version`
	- Print modget version and exit.

## add command options

- `-m <MINECRAFT VERSION>`
	- Limit install for a specific minecraft version.
- `-l <LOADER>`
	- Limit install to a specific mod loader.

## del command options

- `-auto-remove`
	- NOT YET IMPLEMENTED
	- Search for uneeded dependency mods and remove them.

## update command options

- `-m <MINECRAFT VERSION>`
	- Update to a different minecraft version.

## Slug

Every mod on curseforge has a short name called a slug which is found in the
URL for the mod's page. For example in this url
[https://www.curseforge.com/minecraft/mc-mods/appleskin](https://www.curseforge.com/minecraft/mc-mods/appleskin)
the slug is `appleskin`. The modget commands are smart enough to convert slugs
to MODIDs which makes managing your mods much easier. Internally everything
uses MODID.

## MODID

Every mod on curseforge has a 6 digit MODID which can be found on the mods page
or by using modget's search command. Typically using the Slug for the mod is
easier and more userfriendly, but if that fails you can try the MODID.

## resources

[Send patches](https://git-send-email.io) and questions to
[~kota/modget@lists.sr.ht](https://lists.sr.ht/~kota/modget).

Bugs & todo here: [~kota/modget](https://todo.sr.ht/~kota/modget)
