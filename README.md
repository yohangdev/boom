# Boom CLI
Learning proof of concept/case study CLI application using Go.
```bash
$ go get
$ go install
```
```bash
$ boom

NAME:
   boom - make an explosive entrance

USAGE:
   boom [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
   doctor   System check
   create   Create new project
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```
## Commands
### Create New Project
```bash
$ boom create

NAME:
   boom create - Create new project

USAGE:
   boom create command [command options] [arguments...]

COMMANDS:
   laravel    Create new backend Laravel project
   microsite  Create new frontend microsite project
   help, h    Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```