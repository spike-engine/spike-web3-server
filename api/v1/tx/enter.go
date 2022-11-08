package txApi

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/middleware"
	"github.com/spike-engine/spike-web3-server/request"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/service/sign"
	"github.com/spike-engine/spike-web3-server/service/tx"
)

var log = logger.Logger("txApi")

type TxGroup struct {
	hwManager *sign.HotWalletManager
	txSrv     *tx.TxService
}

func NewTxGroup() (TxGroup, error) {
	return TxGroup{
		hwManager: sign.HwManager,
		txSrv:     tx.TxSrv,
	}, nil
}

func (txGroup *TxGroup) InitTxGroup(g *gin.RouterGroup) {
	g.Use(middleware.WhiteListAuth())
	hotWallet := g.Group("hotWallet")
	{
		hotWallet.POST("/mint", txGroup.BatchMint)
		hotWallet.POST("/withdrawNFT", txGroup.BatchWithdrawNFT)
		hotWallet.POST("withdrawToken", txGroup.BatchWithdrawToken)
	}
	client := g.Group("client")
	{
		client.POST("/rechargeToken", txGroup.RechargeToken)
		client.POST("/importNft", txGroup.ImportNft)
	}
}

// @Summary mint nft
// @Produce json
// @Param   order_id  formData string true "game orderId"
// @Param   token_uri formData string true "nft tokenUri"
// @Param   cb        formData string true "game callBack url address"
// @Success 200       {object} response.Response
// @Failure 500       {object} response.Response
// @Router  /tx-api/v1/hotWallet/mint [post]
// @Security api_key
func (txGroup *TxGroup) BatchMint(c *gin.Context) {
	var service request.BatchMintNFTService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
	}
	err = service.BatchMint(txGroup.hwManager)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// @Summary withdraw nft
// @Produce json
// @Param   order_id         formData string true "game orderId"
// @Param   to_address       formData string true "tx toAddress"
// @Param   token_id         formData int    true "nft token id"
// @Param   contract_address formData string true "nft contract address"
// @Param   cb               formData string true "game callBack url address"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /tx-api/v1/hotWallet/withdrawNFT [post]
// @Security api_key
func (txGroup *TxGroup) BatchWithdrawNFT(c *gin.Context) {
	var service request.BatchWithdrawalNFTService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
	}

	err = service.WithdrawNFT(txGroup.hwManager)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// @Summary withdraw token
// @Produce json
// @Param   order_id         formData string true "game orderId"
// @Param   to_address       formData string true "tx toAddress"
// @Param   amount           formData string true "tx token amount"
// @Param   contract_address formData string true "token contract address(native : 0x0000000000000000000000000000000000000000)"
// @Param   cb               formData string true "game callBack url address"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /tx-api/v1/hotWallet/withdrawToken [post]
// @Security api_key
func (txGroup *TxGroup) BatchWithdrawToken(c *gin.Context) {
	var service request.BatchWithdrawalTokenService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
	}

	err = service.WithdrawToken(txGroup.hwManager)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// @Summary recharge token
// @Produce json
// @Param   order_id         formData string true "game orderId"
// @Param   from_address     formData string true "tx fromAddress"
// @Param   amount           formData string true "tx token amount"
// @Param   contract_address formData string true "token contract address(native : 0x0000000000000000000000000000000000000000)"
// @Param   tx_hash          formData string true "tx hash"
// @Param   cb               formData string true "game callBack url address"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /tx-api/v1/client/rechargeToken [post]
// @Security api_key
func (txGroup *TxGroup) RechargeToken(c *gin.Context) {
	var service request.RechargeTokenService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}

	err = service.RechargeToken(txGroup.txSrv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// @Summary import nft
// @Produce json
// @Param   order_id         formData string true "game orderId"
// @Param   from_address     formData string true "tx fromAddress"
// @Param   contract_address formData string true "nft contract address"
// @Param   token_id         formData int    true "nft token id"
// @Param   tx_hash          formData string true "tx hash"
// @Param   cb               formData string true "game callBack url address"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /tx-api/v1/client/importNft [post]
// @Security api_key
func (txGroup *TxGroup) ImportNft(c *gin.Context) {
	var service request.ImportNftService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
	}

	err = service.ImportNft(txGroup.txSrv)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
