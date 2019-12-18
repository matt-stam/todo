# todo

A simple task manager CLI. Following the work of Steve Francia in his OSCON presentation.

Created with Cobra, a CLI framework used in Kubernetes, Docker, etc.

Todo will help you get more done in less time. It's designed to make it easy to track, update, and complete your daily tasks.

Compile with 'go build'

Usage:

    todo [command]



Available Commands:

    add         Add a new todo

    done        Mark Item as Done

    help        Help about any command

    list        list all todos



Flags:

      --config string     config file (default is $HOME/.todo.yaml)

      --datafile string   data file to store todos (default "C:/Users/$USER/go/src/github.com/matt-stam/todo/todo-data.json")

    -h, --help              help for todo

    -t, --toggle            Help message for toggle


Use "todo [command] --help" for more information about a command.

subcommand is required
