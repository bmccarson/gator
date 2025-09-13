package commands

import "github.com/bmccarson/gator/internal/state"

func Login(state *state.DataStore, userName string) error {
	err := state.CurrentConfig.SetUser(userName)
	if err != nil {
		return err
	}
	return nil
}
