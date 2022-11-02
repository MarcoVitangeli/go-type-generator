# go-type-generator
Create dynamic golang types from a json file.
After creation of the file with the specified types, then the program runs a type checker using golang [AST](https://es.wikipedia.org/wiki/%C3%81rbol_de_sintaxis_abstracta)

## Usage

```sh
git clone https://github.com/MarcoVitangeli/go-type-generator
cd go-type-generator
go build -o parser cmd/cli.go
./parser [-h | --help] <json file path>
```
