# gopaste

A client for pasting text to p.lee.io from stdin/files. Utilising client-side encryption/decryption

## Examples



```
[root@lee ~]# echo "lolololol" | gopaste
https://p.lee.io/3ee20782-e75e-47c5-97fc-ed2fe50a846a#encryptionKey=kJGDMAYagndatCQ4Fi0UgFHa
```

```
[root@lee ~]# gopaste -get 3ee20782-e75e-47c5-97fc-ed2fe50a846a -decrypt kJGDMAYagndatCQ4Fi0UgFHa
lolololol
```

```
[root@lee ~]# gopaste -file /tmp/lol.txt
https://p.lee.io/c01cac79-84bd-4fb4-aa99-aa9e0094c150#encryptionKey=IsMVUtO06Mtgs25JU4NnOuKi
```

```
[root@lee ~]# echo '{"hello":true}' | gopaste -syntax "application/ld+json"
https://p.lee.io/f51eeb24-d54a-46b4-85ab-eecef159e91b#encryptionKey=ddGtwmBD40k2hMAej4LmoQHV
```

## Usage

```
gopaste --help
Usage of /usr/local/bin/gopaste:
  -decrypt string
        Decryption key for retrieving encrypted pastes (client-side)
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
