package createplayer

import (
	"context"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
)

func Test_CreatePlayer_Success(t *testing.T) {
	ctx := context.Background()
	dbStore := repositorytest.NewDBStore()

	ports := Ports{
		PlayerRepository: repositorytest.ProvidePlayerRepository(&dbStore),
	}

	if err := Create(ctx, ports, Request{Name: "Maggus"}); err != nil {
		t.Fatal(err)
	}

	p, err := ports.PlayerRepository.FindByID(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}

	if p.ID != 1 {
		t.Fatalf("id was %d expected %d", p.ID, 1)
	}

	if p.Name != "Maggus" {
		t.Fatalf("id was %s expected %s", p.Name, "Maggus")
	}

	if p.CreatedAt.Day() != time.Now().Day() {
		t.Fatalf("created day was %d expected %d", p.CreatedAt.Day(), time.Now().Day())
	}

}
