package commands

import "github.com/bmccarson/gator/internal/state"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(state *state.DataStore, arg string) error
}

func Init() map[string]CliCommand {
	commands := make(map[string]CliCommand)

	commands["login"] = CliCommand{
		Name:        "login",
		Description: "set the current user as the user in the config",
		Callback:    Login,
	}

	return commands
}
