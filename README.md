# Flatpkg

CLI that searches on Flathub directly because the Flatpak CLI is a pain to use.

# Usage

Type the name of the app, and the CLI will search Flathub directly, only getting apps.

```sh
flatpkg install obs
flatpkg search obs
```

There is also a clean command to clear data from uninstalled apps.

```sh
flatpkg clean
```

# Instalation

Run `go build` and place the binary wherever you want.

To get completion working, run the following command:

```sh
flatpkg completion YOUR-SHELL > COMPLETIONS-LOCATION
```

The completion location varies depending on the shell, here are some locations of popular shells:

- fish: `~/.config/fish/completions/`
- bash: `/usr/share/bash-completion/completions/`
