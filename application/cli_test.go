package application_test

import (
	"bytes"
	poker "github.com/igorakimy/go_with_tests/application"
	"io"
	"strings"
	"testing"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("3", "Chris wins")

		poker.NewCLI(in, stdout, game).PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}

		in := userSends("8", "Cleo")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &poker.GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	amountGot := got.Amount
	if amountGot != want.Amount {
		t.Errorf("got amount %d, want %d", amountGot, want.Amount)
	}

	gotScheduledTime := got.At
	if gotScheduledTime != want.At {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, want.At)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t testing.TB, game *poker.GameSpy, numberOfPlayers int) {
	t.Helper()
	if game.StartedWith != numberOfPlayers {
		t.Errorf("game started with %d number of players, but want %d", game.StartedWith, numberOfPlayers)
	}
}

func assertFinishCalledWith(t testing.TB, game *poker.GameSpy, playerName string) {
	t.Helper()
	finishedWith := strings.Replace(game.FinishedWith, " wins", "", -1)
	if finishedWith != playerName {
		t.Errorf("game finished with %s as winner, but want %s", finishedWith, playerName)
	}
}

func assertGameNotStarted(t testing.TB, game *poker.GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Error("the game started when it shouldn't")
	}
}

func userSends(lines ...string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n") + "\n")
}
