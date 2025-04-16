# gopaste

A client for pasting text to p.lee.io from stdin/files. Utilises client-side encryption/decryption

## Examples

* Create paste from file:

```
[lee@lee ~]# gopaste create -f /tmp/lol.txt
https://p.lee.io/c01cac79-84bd-4fb4-aa99-aa9e0094c150#encryptionKey=IsMVUtO06Mtgs25JU4NnOuKi
```

* Create paste from stdin:

```
[lee@lee ~]# echo "lolololol" | gopaste create -f -
https://p.lee.io/3ee20782-e75e-47c5-97fc-ed2fe50a846a#encryptionKey=kJGDMAYagndatCQ4Fi0UgFHa
```

* Create paste with syntax:

```
[lee@lee ~]# echo '{"hello":true}' | gopaste create --syntax "application/ld+json" -f -
https://p.lee.io/f51eeb24-d54a-46b4-85ab-eecef159e91b#encryptionKey=ddGtwmBD40k2hMAej4LmoQHV
```

* Retrieve paste:

```
[lee@lee ~]# gopaste get 3ee20782-e75e-47c5-97fc-ed2fe50a846a -p kJGDMAYagndatCQ4Fi0UgFHa
lolololol
```

## Usage

```
A CLI tool for interacting with the p.lee.io paste service with client-side encryption support

Usage:
  gopaste [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  create       Create a new paste
  get          Get a paste by ID
  help         Help about any command
  list-expires List available expiration options
  list-syntax  List available syntax highlighting options

Flags:
  -h, --help   help for gopaste

Use "gopaste [command] --help" for more information about a command.
```
