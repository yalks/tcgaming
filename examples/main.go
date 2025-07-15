package main

import (
	"fmt"
	"log"

	// "github.com/goforj/godump"
	// uuid "github.com/google/uuid"

	"github.com/goforj/godump"
	uuid "github.com/google/uuid"
	"github.com/tc-gaming/tcgaming-go"
)

func main() {
	fmt.Println("--------> TC-Gaming Sample Code 天成游戏范例代码 <--------")

	// Initialize configuration (matching Java SDK pattern)
	config := tcgaming.NewConfig(
		"http://www.connect8play.com/doBusiness.do", // API URL
		"jpaycny",          // Merchant Code
		"waxfjAic",         // DES Key (8 bytes)
		"Ssb3wHZJMcGK5EbS", // SHA256 Key
	)

	// Create client with configuration
	client, err := tcgaming.NewClient(config, "CNY")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// For testing, you can uncomment the line below and use the correct URL:
	// client.SetURL("YOUR_ACTUAL_API_URL")

	username := "phoenixGO"
	fmt.Println("Username:", username)
	productType := tcgaming.EG5
	fmt.Println("Product Type:", productType)
	// password := "phoenix"

	// 2.1. CREATE/REGISTER PLAYER API 创建/确认玩家接口
	// fmt.Println("\n執行 --> 2.1. CREATE/REGISTER PLAYER API 创建/确认玩家接口")
	// createResp, err := client.CreateUser(username, password)
	// if err != nil {
	// 	log.Printf("Error creating user: %v", err)
	// } else {
	// 	godump.Dump("Create User:", createResp)
	// }

	// 2.2. UPDATE PASSWORD API 更新密码接口
	// fmt.Println("\n執行 --> 2.2. UPDATE PASSWORD API 更新密码接口")
	// newPassword := "phoenix123"
	// updateResp, err := client.UpdatePassword(username, newPassword)
	// if err != nil {
	// 	log.Printf("Error updating password: %v", err)
	// } else {
	// 	fmt.Printf("Password updated: %+v\n", updateResp)
	// }

	// // 2.3. GET BALANCE API 获取余额接口
	fmt.Println("\n執行 --> 2.3. GET BALANCE API 获取余额接口")
	balanceResp, err := client.GetBalance(username, productType)
	if err != nil {
		log.Printf("Error getting balance: %v", err)
	} else {
		godump.Dump("Balance:", balanceResp)
	}

	// 2.4. FUND TRANSFER API 资金转账接口
	// godump.Dump("執行 --> 2.4. FUND TRANSFER API 资金转账接口")

	referenceNo := fmt.Sprintf("%s%04d", uuid.New().String(), 1)
	fmt.Println("Reference No:", referenceNo)

	// var amount float64 = 10.0

	// transferResp, err := client.UserTransfer(
	// 	username,
	// 	productType,
	// 	tcgaming.FundTypeOut,
	// 	amount,
	// 	referenceNo,
	// )
	// if err != nil {
	// 	log.Printf("Error transferring funds: %v", err)
	// } else {
	// 	godump.Dump("Transfer completed:", transferResp)
	// }

	// 2.5. CHECK TRANSACTION STATUS API 检查交易状态接口
	// godump.Dump("執行 --> 2.5. CHECK TRANSACTION STATUS API 检查交易状态接口")
	// statusResp, err := client.CheckTransfer(productType, referenceNo)
	// if err != nil {
	// 	log.Printf("Error checking transfer status: %v", err)
	// } else {
	// 	godump.Dump("Transfer status:", statusResp)
	// }

	// 2.6. LAUNCH GAME API 启动游戏接口 - 电子游戏
	// fmt.Println("\n執行 --> 2.6. LAUNCH GAME API 启动游戏接口 - 电子游戏")
	// godump.Dump(username, productType, tcgaming.GameModeLive, "EG5353", tcgaming.PlatformAll)
	// fmt.Println("\n")
	// gameResp, err := client.LaunchGameRNG(
	// 	username,
	// 	productType,
	// 	tcgaming.GameModeLive,
	// 	"EG5353",
	// 	tcgaming.PlatformAll,
	// )
	// if err != nil {
	// 	log.Printf("Error launching RNG game: %v", err)
	// } else {
	// 	fmt.Printf("Game launched: %+v\n", gameResp)
	// }

	// // 2.6. LAUNCH GAME API 启动游戏接口 - 彩票游戏
	// fmt.Println("\n執行 --> 2.6. LAUNCH GAME API 启动游戏接口 - 彩票游戏")
	// lotteryResp, err := client.LaunchGameLottery(
	// 	username,
	// 	tcgaming.ProductTypeLottery,
	// 	tcgaming.GameModeLive,
	// 	"Lobby",
	// 	tcgaming.PlatformHTML5,
	// 	"Lobby",
	// )
	// if err != nil {
	// 	log.Printf("Error launching lottery game: %v", err)
	// } else {
	// 	fmt.Printf("Lottery game launched: %+v\n", lotteryResp)
	// }

	// 2.7. GAME LIST API 游戏列表接口
	// fmt.Println("\n執行 --> 2.7. GAME LIST API 游戏列表接口")
	// gameListResp, err := client.GetGameList(
	// 	productType,
	// 	tcgaming.PlatformAll,
	// 	tcgaming.ClientTypeHTML5,
	// 	tcgaming.GameTypePVP,
	// 	1,
	// 	100,
	// )
	// if err != nil {
	// 	log.Printf("Error getting game list: %v", err)
	// } else {
	// 	godump.Dump("Game list raw response:", gameListResp)
	// }

	// 3.1. GET RNG BET DETAILS 获得电子游戏及真人投注详情接口
	// fmt.Println("\n執行 --> 3.1. GET RNG BET DETAILS 获得电子游戏及真人投注详情接口")
	// batchName := "201706262010"
	// betDetailsResp, err := client.GetBetDetails(batchName, 1)
	// if err != nil {
	// 	log.Printf("Error getting bet details: %v", err)
	// } else {
	// 	fmt.Printf("Bet details: %+v\n", betDetailsResp)
	// }

	// 3.2. GET LIVE BET DETAILS BY MEMBER 获得玩家真人投注详情接口
	fmt.Println("\n執行 --> 3.2. GET LIVE BET DETAILS BY MEMBER 获得玩家真人投注详情接口")
	startDate := "2025-07-14 00:00:00"
	endDate := "2025-07-16 00:00:00"
	liveBetResp, err := client.GetLiveBetDetailsByMember(username, startDate, endDate, 1)
	if err != nil {
		log.Printf("Error getting Live bet details: %v", err)
	} else {
		// fmt.Printf("Live bet details: %+v\n", liveBetResp)
		godump.Dump("Live bet details:", liveBetResp)
	}

	// 3.3. GET RNG BET DETAILS BY MEMBER 获得玩家电子游戏投注详情接口
	fmt.Println("\n執行 --> 3.3. GET RNG BET DETAILS BY MEMBER 获得玩家电子游戏投注详情接口")
	rngBetResp, err := client.GetRNGBetDetailsByMember(username, startDate, endDate, 1)
	if err != nil {
		log.Printf("Error getting RNG bet details: %v", err)
	} else {
		godump.Dump("RNG bet details:", rngBetResp)
	}

	// fmt.Println("\n--------> Sample execution completed <--------")
}
