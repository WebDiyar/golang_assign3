package main

import (
	"fmt"
	"sync"
)

type Team struct {
	Name     string
	Location string
	Founded  int
}

type teamSingleton struct {
	teams map[string]*Team
	mu    sync.RWMutex
}

var instance *teamSingleton
var once sync.Once

func GetTeamManager() *teamSingleton {
	once.Do(func() {
		instance = &teamSingleton{
			teams: make(map[string]*Team),
		}
	})
	return instance
}

func (ts *teamSingleton) AddTeam(name, location string, founded int) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if _, exists := ts.teams[name]; exists {
		fmt.Printf("Team '%s' already exists.\n", name)
		return
	}

	team := &Team{
		Name:     name,
		Location: location,
		Founded:  founded,
	}
	ts.teams[name] = team
	fmt.Printf("Team '%s' added.\n", name)
}

func (ts *teamSingleton) GetTeam(name string) *Team {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	team, exists := ts.teams[name]
	if exists {
		return team
	}
	fmt.Printf("Team '%s' not found.\n", name)
	return nil
}

func (ts *teamSingleton) ListTeams() []*Team {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	teams := make([]*Team, 0, len(ts.teams))
	for _, team := range ts.teams {
		teams = append(teams, team)
	}
	return teams
}

func main() {
	manager := GetTeamManager()

	manager.AddTeam("Los Angeles Lakers", "Los Angeles", 1947)
	manager.AddTeam("Golden State Warriors", "San Francisco", 1946)
	manager.AddTeam("Boston Celtics", "Boston", 1946)
	manager.AddTeam("Los Angeles Lakers", "LA", 1947)

	lakers := manager.GetTeam("Los Angeles Lakers")
	if lakers != nil {
		fmt.Printf("Team: %s, Location: %s, Founded: %d\n", lakers.Name, lakers.Location, lakers.Founded)
	}

	teams := manager.ListTeams()
	for _, team := range teams {
		fmt.Printf("Team: %s, Location: %s, Founded: %d\n", team.Name, team.Location, team.Founded)
	}
}