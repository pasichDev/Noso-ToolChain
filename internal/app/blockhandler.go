package app

import (
	"fmt"
	"github.com/Friends-Of-Noso/NosoData-Go/legacy"
	"github.com/Friends-Of-Noso/NosoData-Go/utils"
	"github.com/gookit/color"
	"time"
)

type ViewModeBlock int64

const (
	DefaultViewBlock ViewModeBlock = 0
	FullViewBlock    ViewModeBlock = 1
	OnlyOrders       ViewModeBlock = 2
	OnlyRewards      ViewModeBlock = 3
)

func GetBlockData(filePath string) (legacy.LegacyBlock, error) {
	block := legacy.LegacyBlock{}
	err := block.ReadFromFile(filePath)
	if err != nil {
		return legacy.LegacyBlock{}, err
	}
	return block, nil

}

func GetBlockDetail(filepath string, viewMode ViewModeBlock) {
	block, err := GetBlockData(filepath)
	if err != nil {

	}
	color.Bold.Println("\n🧱 Block Detail:")
	fmt.Printf("%-20s: %d\n", "Number", block.Number)
	if viewMode != OnlyOrders && viewMode != OnlyRewards {

		fmt.Printf("%-20s: %s\n", "Time Start", time.Unix(block.TimeStart, 0).Format(time.RFC1123))
		fmt.Printf("%-20s: %s\n", "Time End", time.Unix(block.TimeEnd, 0).Format(time.RFC1123))
		fmt.Printf("%-20s: %d seconds\n", "Time Total", block.TimeTotal)
		fmt.Printf("%-20s: %d seconds\n", "Time Last 20", block.TimeLast20)
		fmt.Printf("%-20s: %d\n", "Difficulty", block.Difficulty) // Assuming Difficulty is int32
		fmt.Printf("%-20s: %s\n", "Target Hash", block.TargetHash.GetString())
		fmt.Printf("%-20s: %s\n", "Solution", block.Solution.GetString())
		fmt.Printf("%-20s: %s\n", "Last Block Hash", block.LastBlockHash.GetString())
		fmt.Printf("%-20s: %s\n", "Miner", block.Miner.GetString())
		fmt.Printf("%-20s: %s\n", "Fee", utils.ToNoso(block.Fee))
		fmt.Printf("%-20s: %s\n", "Reward", utils.ToNoso(block.Reward))
		fmt.Printf("%-20s: %d\n", "Transaction Count", block.TransactionsCount)
		if viewMode != FullViewBlock {

			color.Bold.Println("\n💰 Reward:")
			if block.ProofOfStakeRewardCount > 0 {
				fmt.Printf("%-20s: %d\n", "PoW count", block.ProofOfStakeRewardCount)
				fmt.Printf("%-20s: %d\n", "Amount:", utils.ToNoso(block.ProofOfStakeRewardAmount))

			} else {
				fmt.Printf("%-20s: %s\n", "PoW", "none")
			}
			if block.MasterNodeRewardCount > 0 {
				fmt.Printf("%-20s: %d\n", "MN count", block.MasterNodeRewardCount)
				fmt.Printf("%-20s: %s\n", "Amount", utils.ToNoso(block.MasterNodeRewardAmount))

			} else {
				fmt.Printf("%-20s: %s\n", "MN", "none")
			}

		}
	}

	if viewMode == OnlyOrders || viewMode == FullViewBlock {
		if block.TransactionsCount > 0 {
			color.Bold.Println("\n🔄 Transactions:", block.TransactionsCount)
			var n int32
			for n = 0; n < block.TransactionsCount; n++ {
				fmt.Printf("%-20s: %s\n", "OrderID", block.Transactions[n].OrderID.GetString())
				fmt.Printf("%-20s: %s\n", "TransferID", block.Transactions[n].TransferID.GetString())
				fmt.Printf("%-20s: %d\n", "Block", block.Transactions[n].Block)
				fmt.Printf("%-20s: %d\n", "Order lines", block.Transactions[n].OrderLinesCount)
				fmt.Printf("%-20s: %s\n", "Order type", block.Transactions[n].OrderType.GetString())
				fmt.Printf("%-20s: %s\n", "Timestamp", time.Unix(block.Transactions[n].TimeStamp, 0).Format(time.RFC1123))
				fmt.Printf("%-20s: %s\n", "Reference", block.Transactions[n].Reference.GetString())
				fmt.Printf("%-20s: %d\n", "Transfer Index", block.Transactions[n].TransferIndex)
				fmt.Printf("%-20s: %s\n", "Sender", block.Transactions[n].Sender.GetString())
				fmt.Printf("%-20s: %s\n", "Address", block.Transactions[n].Address.GetString())
				fmt.Printf("%-20s: %s\n", "Receiver", block.Transactions[n].Receiver.GetString())
				fmt.Printf("%-20s: %s\n", "Fee", utils.ToNoso(block.Transactions[n].AmountFee))
				fmt.Printf("%-20s: %s\n", "Value", utils.ToNoso(block.Transactions[n].AmountTransfer))
				fmt.Printf("%-20s: %s\n", "Signature", block.Transactions[n].Signature.GetString())
				if block.TransactionsCount > 1 {
					color.Bold.Println("-----------------------------------")
				}
			}
		}
	}
	if viewMode == OnlyRewards || viewMode == FullViewBlock {
		if block.ProofOfStakeRewardCount > 0 {
			color.Bold.Println("\n🌱 PoS Addresses:", block.ProofOfStakeRewardCount)
			fmt.Println("💰 Amount:", utils.ToNoso(block.ProofOfStakeRewardAmount))
			var n int32
			for n = 0; n < block.ProofOfStakeRewardCount; n++ {
				fmt.Printf("%s\n", block.ProofOfStakeRewardAddresses[n].GetString())
			}
		}

		if block.MasterNodeRewardCount > 0 {
			color.Bold.Println("\n🌱 MN Addresses:", block.MasterNodeRewardCount)
			fmt.Println("💰 Amount:", utils.ToNoso(block.MasterNodeRewardAmount))
			var n int32
			for n = 0; n < block.MasterNodeRewardCount; n++ {
				fmt.Printf("%s\n", block.MasterNodeRewardAddresses[n].GetString())
			}
		}
	}

	return

}
