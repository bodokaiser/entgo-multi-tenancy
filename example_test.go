package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bodokaiser/entgo-multi-tenancy/ent"
)

func Example_CreateTeams() {
	ctx := context.Background()
	client := open(ctx)
	defer client.Close()

	team1, err := client.Team.
		Create().
		SetName("Facebook").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating team: %v", err)
	}
	fmt.Println(team1)

	team2, err := client.Team.
		Create().
		SetName("Google").
		Save(ctx)
	log.Println(team2)
	if err != nil {
		log.Fatalf("failed creating team: %v", err)
	}
	fmt.Println(team2)

	// Output:
	// Team(id=1, name=Facebook)
	// Team(id=2, name=Google)
}

func open(ctx context.Context) *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
