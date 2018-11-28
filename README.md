# GO File Readers

Golang scripts for reading, writing and converting data.

## Usage

```
NAME:
   go-file-readers - reads and converts data files

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     gff3     reads GFF3 file and extracts data by given type
     csv      reads csv file, converts it to JSON and writes it to new file
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### gff3

_Flags:_

- file (input file)
- type (type of data to extract)
- output (output file)

_Example:_

`go run main.go gff3 -file=canonical_core.gff3 -type=pseudogene -output=pseudogenes.gff3`

### csv

_Flags:_

- file (input file)
- output (output file)

_Example:_

`go run main.go csv -file=users.csv -output=users.json`
