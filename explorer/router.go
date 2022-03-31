package explorer

import (
	"github.com/beego/beego/v2/server/web"
)

func GetRouter() web.LinkNamespace {
	//bot := &BotController{}
	//go bot.RunChecks()

	ns := web.NSNamespace("/explorer",
		web.NSRouter("/getcrosstx", &ExplorerController{}, "get:GetCrossTx"),
		web.NSRouter("/getassetstatistic", &ExplorerController{}, "get:GetAssetStatistic"),
		web.NSRouter("/gettransferstatistic", &ExplorerController{}, "get:GetTransferStatistic"),
		web.NSRouter("/getexplorerinfo/", &ExplorerController{}, "get:GetExplorerInfo"),
		web.NSRouter("/getcrosstxlist/", &ExplorerController{}, "post:GetCrossTxList"),
		web.NSRouter("/gettokentxlist/", &ExplorerController{}, "post:GetTokenTxList"),
		web.NSRouter("/getaddresstxlist/", &ExplorerController{}, "post:GetAddressTxList"),
		web.NSRouter("/getlocktokenlist/", &ExplorerController{}, "get:GetLockTokenList"),
		web.NSRouter("/getlocktokeninfo/", &ExplorerController{}, "get:GetLockTokenInfo"),
		web.NSRouter("/getetheffectuser/", &ExplorerController{}, "post:GetEthEffectUser"),
		web.NSRouter("/bot/", &BotController{}, "get:BotPage"),
		web.NSRouter("/bottxs/", &BotController{}, "get:GetTxs"),
		web.NSRouter("/botcheck/", &BotController{}, "get:CheckTxs"),
		web.NSRouter("/botcheckfee/", &BotController{}, "post:CheckFees"),
		web.NSRouter("/botfinishtx/", &BotController{}, "get:FinishTx"),
		web.NSRouter("/botmarkunmarktxaspaid/", &BotController{}, "get:MarkUnMarkTxAsPaid"),
		web.NSRouter("/botlistlargetx/", &BotController{}, "get:ListLargeTxPage"),
		web.NSRouter("/botlistnodestatus/", &BotController{}, "get:ListNodeStatusPage"),
		web.NSRouter("/botignorenodestatusalarm/", &BotController{}, "get:IgnoreNodeStatusAlarm"),
		web.NSRouter("/botlistrelayeraccountstatus/", &BotController{}, "get:ListRelayerAccountStatus"),
	)
	return ns
}
