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
     txt      reads txt file, converts it to JSON and writes it to new file
     arango   reads csv file, converts to user data structure and stores in arangodb
     ddanat   converts data given DDANAT IDs into JSON file with circos coordinates and other data
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

For more detailed descriptions of each command, run `go run main.go [COMMAND] -h`.
