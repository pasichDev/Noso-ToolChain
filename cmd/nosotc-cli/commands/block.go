package commands

import (
	"fmt"
	"github.com/pasichDev/nosotc/internal/app"
	"github.com/pasichDev/nosotc/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	blockCmd = &cobra.Command{
		Use:   "block",
		Short: ".blk file handler (blockchain blocks)",
		Run:   runBlock,
	}
	blkNumber   int
	countBlocks bool
	fullInfo    bool
	onlyOrders  bool
	onlyRewards bool
)

func init() {

	blockCmd.Flags().IntVarP(&blkNumber, "target", "t", 0, "Selecting block (number) to view from the blockchain folder")
	blockCmd.PersistentFlags().BoolVarP(&countBlocks, "count-blocks", "", false, "Print the number of blocks in the blockchain folder and the number of the last block")
	blockCmd.PersistentFlags().BoolVarP(&fullInfo, "full-info", "", false, "Displaying full information about the block (PoS & PoW transactions and payments)")
	blockCmd.PersistentFlags().BoolVarP(&onlyOrders, "only-orders", "", false, "Displays only transactions")
	blockCmd.PersistentFlags().BoolVarP(&onlyRewards, "only-rewards", "", false, "Displaying information only about PoS & PoW transactions and payments ")

	rootCmd.AddCommand(blockCmd)

}

func initCommand() (string, error) {
	blockPath := ""
	bold := "\033[1m"
	reset := "\033[0m"

	if blkNumber != 0 {
		fmt.Printf("📂 Search for block %d in the blockchain folder (NOSODATA/BLOCKS)\n", blkNumber)
		if nosoDataFolder != "" {
			blockPath = filepath.Join(nosoDataFolder, "BLOCKS", strconv.Itoa(blkNumber)+".blk")

		} else {
			fmt.Printf(bold + "\n❌ To search by block number, you need to specify the path to the NOSODATA blockchain folder in the configuration. Otherwise, use the -f (--file) flag to specify the direct path to the block file \n" + reset)
			return "", nil
		}
	} else {
		if nosoFilePath == "" {
			fmt.Printf(bold + "\n❌ You need to specify the path to the NOSODATA blockchain folder in the configuration. Otherwise, use the -f (--file) flag to specify the direct path to the block file \n" + reset)
			return "", nil
		}

		blNum, er := utils.GetBlockNumberForFile(nosoFilePath)
		blkNumber = blNum
		blockPath = nosoFilePath
		if !strings.HasSuffix(filepath.Base(nosoFilePath), ".blk") || blkNumber == 0 || er != nil {
			fmt.Println("❌ File does not exist or is not supported, check if the path is correct")
			return "", nil
		}
	}

	isBlockToData := utils.CheckIfFileExists(blockPath)

	if isBlockToData {
		fmt.Printf("🗂️ Targer block %d found by following this path '%s'\n", blkNumber, blockPath)
	} else {
		fmt.Printf(bold+"\n❌ Block %d is not found in blockchain NOSO \n"+reset, blkNumber)
		return "", nil
	}

	return blockPath, nil
}

func runBlock(cmd *cobra.Command, args []string) {

	blockPath, err := initCommand()

	if err != nil {
		os.Exit(1)
	}

	viewMode := func() app.ViewModeBlock {
		switch {
		case fullInfo:
			return app.FullViewBlock
		case onlyOrders:
			return app.OnlyOrders
		case onlyRewards:
			return app.OnlyRewards
		default:
			return app.DefaultViewBlock
		}
	}()

	switch {

	case countBlocks:

	default:
		app.GetBlockDetail(blockPath, viewMode)

	}
}
