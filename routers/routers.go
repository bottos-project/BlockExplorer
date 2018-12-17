package routers

import (
	"github.com/BlockExplorer/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	rExplorer := router.Group("/explorer_api")
	// Home
	rExplorer.POST("/comm/queryMajSummaryAuto", controller.HomeGetTotalCount) // 首页统计

	// total
	rExplorer.POST("/cust/queryCustSummaryAuto", controller.AccountTotal) // 账户实时总数

	// List
	rExplorer.POST("/block/queryBlockListAuto", controller.BlockList)               // BlockList
	rExplorer.POST("/cust/queryCustListAuto", controller.AccountList)               // account 列表
	rExplorer.POST("/trade/queryTradeListAuto", controller.TrxTransactionList)      // 交易列表
	rExplorer.POST("/transferl/queryTransferlListAuto", controller.TrxTransferList) // 转账列表

	// Detail
	rExplorer.POST("/block/queryBlockDetailAuto", controller.BlockDetail)          // 区块详情
	rExplorer.POST("/trade/queryTradeDetailAuto", controller.TrxTransactionDetail) // 交易详情
	rExplorer.POST("/cust/queryCustDetailAuto", controller.AccountDetail)          // 账户详情
}
