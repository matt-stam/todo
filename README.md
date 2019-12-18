# todo-cli

A simple task manager CLI. Following the work of [Steve Francia](https://github.com/spf13) in his OSCON presentation.

Created with [Cobra](https://github.com/spf13/cobra), a popular CLI generating library used in Kubernetes, Docker, Hugo, etc.

Todo-cli will help you get more done in less time. It's designed to make it easy to track, update, and complete your daily tasks.

Compile with 'go build'

Usage:

    todo [command]



Available Commands:

    add         Add a new Todo

    done        Mark Todo as Done

    help        Help about any command

    list        List all Todos



Flags:

      --config string     config file (default is $HOME/.todo.yaml)

      --datafile string   data file to store todos (default "C:/Users/$USER/go/src/github.com/matt-stam/todo/todo-data.json")

    -h, --help              help for todo

    -t, --toggle            Help message for toggle


Use "todo [command] --help" for more information about a command.

subcommand is required
