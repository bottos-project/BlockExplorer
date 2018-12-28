package routers

import (
	"github.com/BlockExplorer/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// rExplorer := router.Group("/explorer_api")
	rExplorer := router.Group("/")
	// Home
	rExplorer.POST("/comm/queryMajSummaryAuto", controller.HomeGetTotalCount) // 首页统计
	rExplorer.POST("/common/queryFuzzyInfoAuto", controller.HomeSearch)       // 首页搜索

	// total
	rExplorer.POST("/cust/queryCustSummaryAuto", controller.AccountTotal) // 账户实时总数

	// List
	rExplorer.POST("/block/queryBlockListAuto", controller.BlockList)                                        // BlockList
	rExplorer.POST("/cust/queryCustListAuto", controller.AccountList)                                        // account 列表
	rExplorer.POST("/trade/queryTradeListAuto", controller.TrxTransactionList)                               // 交易列表
	rExplorer.POST("/transferl/queryTransferlListAuto", controller.TrxTransferList)                          // 转账列表
	rExplorer.POST("/transferl/queryPersonalTransferlListAuto", controller.TrxPersonalTransferList)          // 个人转账列表
	rExplorer.POST("/transaction/queryPersonalTransactionByMethod", controller.TexPersonTransactionByMethod) // 根据方法名查询列表

	// Detail
	rExplorer.POST("/block/queryBlockDetailAuto", controller.BlockDetail)          // 区块详情
	rExplorer.POST("/trade/queryTradeDetailAuto", controller.TrxTransactionDetail) // 交易详情
	rExplorer.POST("/cust/queryCustDetailAuto", controller.AccountDetail)          // 账户详情

	// Statistics
	rExplorer.POST("/trade/queryLastFTTradeSummaryAuto", controller.StatisticTransaction) // transaction statistic
	rExplorer.POST("/cust/queryRigiterSummaryAuto", controller.StatisticAccount)          // account increase

	// statistic
	rExplorer.POST("/test/accountstatistic", controller.TestStatisticAccount)

	// nodes
	rExplorer.POST("/superNode/queryNodeListAuto", controller.NodeProductList)   // 生产节点
	rExplorer.POST("/node/queryNodeRankAuto", controller.NodeServiceList)        // 服务节点
	rExplorer.POST("/superNode/queryNodeDetailAuto", controller.NodeSuperDetail) // 服务节点
}
