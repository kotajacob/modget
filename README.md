# modget

Modget is a package manager for minecraft curseforge mods. It provides commands
for searching, managing, and querying information about mods. Currently it is
unreleased and in an early development state. The below commands are slowly
being implemented.

## install

modget requires the following build dependencies:

- `golang`
- `scdoc`
- `make`

1. `git clone https://git.sr.ht/~kota/mdtest`
2. `make all`
3. `sudo make install`

## uninstall

1. `sudo make uninstall`

## commands

- `add <MODID>`
	- Download and install a mod based on its MODID.

- `del <MODID>`
	- Remove and uninstall a mod based on its MODID.

- `update <MODID>`
	- Check each installed mod (in selected directory or current directory) and prompt to install any new mods.

- `show <MODID>`
	- Query and print more information about a specific mod by MODID.

- `search <Search String>`
	- Search curseforge for mods based on search terms. Prints mod names, MODID, Version, and a short description.

## add command options

- `-mc <MINECRAFT VERSION>`
	- Limit install for a specific minecraft version.
- `-v <VERSION>`
	- Limit install to a specific mod version.
- `-l <LOADER>`
	- Limit install to a specific mod loader.

## del command options

- `-auto-remove`
	- Search for uneeded dependency mods and remove them.

## update command options

- `-mc <MINECRAFT VERSION>`
	- Update to a different minecraft version.

## MODID

Every mod on curseforge has a 6 digit MODID which can be found on the mods page
or by using modget's search command. This is used instead of names when
installing and removing mods as multiple mods could technically share the same
name.

## resources

[Send patches](https://git-send-email.io) and questions to
[~kota/modget@lists.sr.ht](https://lists.sr.ht/~kota/modget).

Bugs & todo here: [~kota/modget](https://todo.sr.ht/~kota/modget)
