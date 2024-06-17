package startsession

import (
	"context"
	"testing"

	repositorytest_test "github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid/sessionidtest"
)

func Test_StartSession_Success(t *testing.T) {
	db := repositorytest_test.NewDBStore()

	ports := Ports{
		SessionRepository:  repositorytest_test.ProvideSessionRepository(&db),
		SessionIDGenerator: sessionidtest.ProvideSessionIDGen(),
	}

	request := Request{
		Title: "Test Session",
		Owner: 1337,
	}

	sessionID, err := Start(context.Background(), ports, request)
	if err != nil {
		t.Fatal(err)
	}

	got, err := ports.SessionRepository.FindByID(context.Background(), *sessionID)
	if err != nil {
		t.Fatal(err)
	}

	if got.Owner != 1337 {
		t.Fatalf("was %d expected %d", got.Owner, 1337)
	}

	if got.Title != "Test Session" {
		t.Fatalf("was %s expected %s", got.Title, "Test Session")
	}
}
