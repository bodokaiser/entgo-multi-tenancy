package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	_ "github.com/bodokaiser/entgo-multi-tenancy/ent/runtime"

	"github.com/bodokaiser/entgo-multi-tenancy/auth"
	"github.com/bodokaiser/entgo-multi-tenancy/ent"
)

func Example() {
	ctx := context.Background()
	client := open(ctx)
	defer client.Close()

	user, err := client.User.
		Create().
		SetEmail("bodo.kaiser@example.org").
		SetUsername("bodo.kaiser").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	fmt.Println(user)

	ctx = auth.WithUser(ctx, user)

	err = client.Team.
		Create().
		SetName("Facebook").
		Exec(ctx)
	if err != nil {
		log.Fatalf("failed creating team: %v", err)
	}

	err = client.Team.
		Create().
		SetName("Google").
		Exec(ctx)
	if err != nil {
		log.Fatalf("failed creating team: %v", err)
	}

	err = client.Member.
		Create().
		SetOwner(true).
		SetUser(user).
		SetTeamID(2).
		Exec(ctx)
	if err != nil {
		log.Fatalf("failed creating member: %v", err)
	}

	teams, err := client.Team.
		Query().
		All(ctx)
	fmt.Println(teams)

	// Output:
	// User(id=1, username=bodo.kaiser, email=bodo.kaiser@example.org)
	// [Team(id=2, name=Google)]
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
