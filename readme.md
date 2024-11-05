# Factorio mod downloader
A simple tool for downloading mods from the factorio mod portal or from zip files hosted on the internet.

examples:
- `.\main.exe d o --name factorissimo-2-notnotmelon --destination . --token <token> --user user --version 3.5.11 --factorioVersion 2.0`
- `.\main.exe d u --name squeakThrough_1.9.0.zip --url https://github.com/Suprcheese/Squeak-Through/releases/download/1.9.0/Squeak.Through_1.9.0.zip --destination .`

# help for the commands
```bash
NAME:
   factorio-mod-downloader - A Cli tool to download factorio mods

USAGE:
   factorio-mod-downloader [global options] command [command options]

AUTHOR:
   James Rhoat <James@Rhoat.com>

COMMANDS:
   download, d  Allows users to download mods
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

COPYRIGHT:
   Rhoat, LLC, 2024
```

```bash
NAME:
   factorio-mod-downloader download - Allows users to download mods

USAGE:
   factorio-mod-downloader download command [command options]

COMMANDS:
   unofficial, u  downloads an unofficial mod from alternate sources uring a web request
   official, o    downloads an official mod from the mod portal
   help, h        Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

```bash
NAME:
   factorio-mod-downloader download official - downloads an official mod from the mod portal

USAGE:
   factorio-mod-downloader download official [command options]

OPTIONS:
   --name value             Name Of the mod you are downloading. Suggested Name_version
   --destination value      Destination where the mods will be downloaded.
   --version value          Version of the mod you want to download
   --factorioVersion value  Version of factorio you are running
   --user value             Name of the user.
   --token value            Token for the factorio user.
   --help, -h               show help
```
```bash
NAME:
   factorio-mod-downloader download unofficial - downloads an unofficial mod from alternate sources uring a web request

USAGE:
   factorio-mod-downloader download unofficial [command options]

OPTIONS:
   --url value          URL that you can download the mod zip package from. E.G. https://github.com/Suprcheese/Squeak-Through/releases/download/1.9.0/Squeak.Through_1.9.0.zip
   --name value         Name Of the mod you are downloading. Suggested Name_version
   --destination value  Destination where the mods will be downloaded.
   --help, -h           show help
```