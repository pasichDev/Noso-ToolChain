package commands

import (
	"fmt"
	"github.com/Friends-Of-Noso/NosoData-Go/legacy"
	"github.com/Friends-Of-Noso/NosoData-Go/utils"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Display the contents of a wallet file",
	Run:   runWallet,
}

var (
	isShowSecret bool
)

func init() {
	walletCmd.PersistentFlags().BoolVarP(&isShowSecret, "show-secret", "", false, "Display all file secrets")

	rootCmd.AddCommand(walletCmd)
}

func runWallet(cmd *cobra.Command, args []string) {
	bold := "\033[1m"
	reset := "\033[0m"
	walletPath := ""

	// Determine the file to process

	walletPath = nosoFilePath
	if !strings.HasSuffix(filepath.Base(walletPath), ".pkw") {
		fmt.Println("❌ Error: Path to file must end with .pkw.")
		os.Exit(1)
	}

	switch {

	default:
		fmt.Printf("📂 Processing file -> '%s'\n", walletPath)

		wallet := legacy.LegacyWallet{}
		err := wallet.ReadFromFile(nosoFilePath)
		cobra.CheckErr(err)
		fmt.Printf(bold + "💼 Wallet:\n" + reset)
		for i, a := range wallet.Accounts {
			color.Bold.Println("Position:", i+1)
			fmt.Printf("%-12s: %s\n", "Hash", a.Hash.GetString())
			fmt.Printf("%-12s: %s\n", "Custom", a.Custom.GetString())
			if isShowSecret {
				fmt.Printf("%-12s: %s\n", "Pub key", a.PublicKey.GetString())
				fmt.Printf("%-12s: %s\n", "Priv key", a.PrivateKey.GetString())
			}
			fmt.Println(bold + "\n🗄️ The information is not up to date" + reset)
			fmt.Printf("%-12s: %s\n", "Balance", utils.ToNoso(a.Balance))
			fmt.Printf("%-12s: %s\n", "Pending", utils.ToNoso(a.Pending))
			fmt.Printf("%-12s: %s\n", "Score", utils.ToNoso(a.Score))
			fmt.Printf("%-12s: %s\n", "Last Operation", utils.ToNoso(a.LastOperation))
			if wallet.AccountsCount > 1 {
				color.Bold.Println("-----------------------------------\n")
			}
		}
	}
}
