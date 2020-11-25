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

- `[d]elete mod...`
	- Remove and uninstall a mod based on its MODID or Slug.

- `help [command]`
	- Print help about a command and exit.

- `[m]ark [mod]...`
	- Change the update status of a mod.

- `[s]earch string...`
	- Search curseforge for mods based on search terms. Prints mod names, MODID, Version, and a short description.

- `[sh]ow [mod]...`
	- Query and print information about installed mods or a specific mod by MODID or Slug.

- `[u]pdate [mod]...`
	- Check each installed mod (in selected directory or current directory) and prompt to install any new mods.

- `[v]ersion`
	- Print modget version and exit.

## global command options

- `-p path`
	- Install mods in a specific path instead of the current directory.

## add command options

- `-m, --minecraft <minecraft version>`
	- Limit install for a specific minecraft version. Ex: 1.16.3
- `-l --loader <modloader>`
	- Limit install to a specific mod loader. Ex: Fabric

## show command options

- `-l, --oneline`
	- Display mods one per line.

## update command options

- `-m, --minecraft <minecraft version>`
	- Limit install for a specific minecraft version. Ex: 1.16.3
- `-l --loader <modloader>`
	- Limit install to a specific mod loader. Ex: Fabric

## Slug

Every mod on curseforge has a short name called a slug which is found in the
URL for the mod's page. For example in this url
[https://www.curseforge.com/minecraft/mc-mods/appleskin](https://www.curseforge.com/minecraft/mc-mods/appleskin)
the slug is `appleskin`. The modget commands will attempt to convert slugs to
MODIDs which makes managing your mods much easier. Internally everything uses
MODID. Currently this feature is broken for the majority of curseforge mods.

## MODID

Every mod on curseforge has a 6 digit MODID which can be found on the mods page
or by using modget's search command. This is used instead of names when
installing and removing mods as that is what is supported by the curseforge
API.

## resources

[Send patches](https://git-send-email.io) and questions to
[~kota/modget@lists.sr.ht](https://lists.sr.ht/~kota/modget).

Bugs & todo here: [~kota/modget](https://todo.sr.ht/~kota/modget)
