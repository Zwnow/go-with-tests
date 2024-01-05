package pointers

import "testing"

func TestWallet(t *testing.T) {


    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })
    t.Run("withdraw", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        wallet.Withdraw(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })
    t.Run("withdraw insufficient funds", func(t *testing.T) {
        startingBallance := Bitcoin(20)
        wallet := Wallet{startingBallance}
        err := wallet.Withdraw(Bitcoin(100))

        assertBalance(t, wallet, startingBallance)
        assertError(t, err, ErrInsufficientFunds)
    })

}
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertError(t testing.TB, got error, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("wanted an error but didn't get one")
    }

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
