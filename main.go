package main

import (
	"encoding/json"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/localalbino/playlist/views"
)

type model struct {
	appState views.ViewState
	clientID string
}

// Loads the user's client ID from their config.json file.
// An empty string means loading has failed.
func loadClientID() string {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return ""
	}

	var config struct {
		ClientID string `json:"client_id"`
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return ""
	}

	return config.ClientID
}

func initModel() model {
	return model{
		// Always load auth view first (shows a connecting screen even if no client)
		appState: views.AuthView,
		clientID: loadClientID(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
