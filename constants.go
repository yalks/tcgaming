package tcgaming

const (
	// Product Types
	ProductTypeLottery = 2
	ProductTypeRNG     = 7
	EG5                = 191 // Electronic Games
	PG                 = 98  // Poker Games
	PP                 = 39  // PVP Games

	// Platforms
	PlatformFlash = "flash"
	PlatformHTML5 = "html5"
	PlatformAll   = "all"

	// Client Types
	ClientTypePC    = "pc"
	ClientTypePhone = "phone"
	ClientTypeWeb   = "web"
	ClientTypeHTML5 = "html5"

	// Game Types
	GameTypeRNG  = "RNG"
	GameTypeLive = "LIVE"
	GameTypePVP  = "PVP"

	// Lottery Bet Modes
	LotteryBetModeTraditional       = "Traditional"
	LotteryBetModeTraditionalMobile = "Traditional_Mobile"

	// API Methods (internal use)
	methodCreateUser        = "cm"
	methodUpdatePassword    = "up"
	methodGetBalance        = "gb"
	methodFundTransfer      = "ft"
	methodCheckTransaction  = "cs"
	methodLaunchGame        = "lg"
	methodGetGameList       = "tgl"
	methodPlayerGameRank    = "pgr"
	methodBetDetails        = "bd"
	methodBetDetailsMember  = "bdm"
	methodLottoTransactions = "lmb"
	methodGetLottoCodes     = "glgl"
)
