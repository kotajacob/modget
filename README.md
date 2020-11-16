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

- `[a]dd mod...`
	- Download and install a mod based on its MODID or Slug.

- `[d]el mod...`
	- Remove and uninstall a mod based on its MODID or Slug.

- `help [command]`
	- Print help about a command and exit.

- `[s]earch string...`
	- Search curseforge for mods based on search terms. Prints mod names, MODID, Version, and a short description.

- `[sh]ow [mod]...`
	- Query and print information about installed mods or a specific mod by MODID or Slug.

- `[u]pdate [mod]...`
	- NOT YET IMPLEMENTED
	- Check each installed mod (in selected directory or current directory) and prompt to install any new mods.

- `[v]ersion`
	- Print modget version and exit.

## global command options

- `-p path`
	- Install mods in a specific path instead of the current directory.

## add command options

- `-m minecraft version`
	- Limit install for a specific minecraft version. Ex: 1.16.3
- `-l loader`
	- Limit install to a specific mod loader. Ex: Fabric

## update command options

- `-m minecraft version`
	- Limit install for a specific minecraft version. Ex: 1.16.3
- `-l loader`
	- Limit install to a specific mod loader. Ex: Fabric

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
