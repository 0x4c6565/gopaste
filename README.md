# gopaste

A client for pasting text to p.lee.io from stdin/files

## Examples

```
"lolololol" | gopaste
https://p.lee.io/4f27db79-5aed-49f8-a9c9-c94a8cdca7fa
```

```
gopaste -file /tmp/lol.txt
https://p.lee.io/c01cac79-84bd-4fb4-aa99-aa9e0094c150
```

```
'{"hello":true}' | gopaste -syntax "application/ld+json"
https://p.lee.io/f51eeb24-d54a-46b4-85ab-eecef159e91b
```

```
gopaste -get 1658fa7d-3b41-4725-977a-9cf5934dc3b7
lololol
```

## Usage

```
gopaste --help
Usage of /usr/bin/local/gopaste:
  -expires string
        (Optional) Expire type to use for paste
  -file string
        (Optional) File to read from. Stdin is used if not provided
  -get string
        UUID of paste to retrieve
  -getexpires
        Retrieve supported expire types
  -getsyntax
        Retrieve supported syntax
  -syntax string
        (Optional) Syntax to use for paste
```