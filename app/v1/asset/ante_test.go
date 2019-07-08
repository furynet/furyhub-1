package asset

import (
	"testing"

	"github.com/irisnet/irishub/app/v1/auth"
	"github.com/irisnet/irishub/app/v1/bank"
	"github.com/irisnet/irishub/app/v1/params"
	"github.com/irisnet/irishub/codec"
	"github.com/irisnet/irishub/modules/guardian"
	sdk "github.com/irisnet/irishub/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

// TestAssetAnteHandler tests the ante handler of asset
func TestAssetAnteHandler(t *testing.T) {
	ms, accountKey, assetKey, paramskey, paramsTkey := setupMultiStore()

	cdc := codec.New()
	RegisterCodec(cdc)
	auth.RegisterBaseAccount(cdc)

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	guardianKeeper := guardian.Keeper{}
	paramsKeeper := params.NewKeeper(cdc, paramskey, paramsTkey)
	ak := auth.NewAccountKeeper(cdc, accountKey, auth.ProtoBaseAccount)
	bk := bank.NewBaseKeeper(cdc, ak)
	keeper := NewKeeper(cdc, assetKey, bk, guardianKeeper, DefaultCodespace, paramsKeeper.Subspace(DefaultParamSpace))

	// init params
	keeper.Init(ctx)

	// set test accounts
	addr1 := sdk.AccAddress([]byte("addr1"))
	addr2 := sdk.AccAddress([]byte("addr2"))
	acc1 := ak.NewAccountWithAddress(ctx, addr1)
	acc2 := ak.NewAccountWithAddress(ctx, addr2)

	// get asset fees
	gatewayCreateFee := getGatewayCreateFee(ctx, keeper, "mon")
	nativeTokenIssueFee := getTokenIssueFee(ctx, keeper, "sym")
	gatewayTokenIssueFee := getGatewayTokenIssueFee(ctx, keeper, "sym")
	nativeTokenMintFee := getTokenMintFee(ctx, keeper, "sym")

	// construct msgs
	msgCreateGateway := NewMsgCreateGateway(addr1, "mon", "i", "d", "w")
	msgIssueNativeToken := MsgIssueToken{Source: AssetSource(0x00), Symbol: "sym"}
	msgIssueGatewayToken := MsgIssueToken{Source: AssetSource(0x02), Symbol: "sym"}
	msgMintNativeToken := MsgMintToken{TokenId: "i.sym"}
	msgNonAsset1 := sdk.NewTestMsg(addr1)
	msgNonAsset2 := sdk.NewTestMsg(addr2)

	// construct test txs
	tx1 := auth.StdTx{Msgs: []sdk.Msg{msgCreateGateway, msgIssueNativeToken, msgIssueGatewayToken, msgMintNativeToken}}
	tx2 := auth.StdTx{Msgs: []sdk.Msg{msgCreateGateway, msgIssueNativeToken, msgNonAsset1, msgIssueGatewayToken, msgMintNativeToken}}
	tx3 := auth.StdTx{Msgs: []sdk.Msg{msgNonAsset2, msgCreateGateway, msgIssueNativeToken, msgIssueGatewayToken, msgMintNativeToken}}

	// set signers and construct an ante handler
	newCtx := auth.WithSigners(ctx, []auth.Account{acc1, acc2})
	anteHandler := NewAnteHandler(keeper)

	// assert that the ante handler will return with `abort` set to true
	acc1.SetCoins(sdk.Coins{gatewayCreateFee.Plus(nativeTokenIssueFee)})
	_, res, abort := anteHandler(newCtx, tx1, false)
	require.Equal(t, true, abort)
	require.Equal(t, false, res.IsOK())

	// assert that the ante handler will return with `abort` set to true
	acc1.SetCoins(acc1.GetCoins().Plus(sdk.Coins{gatewayTokenIssueFee}))
	_, res, abort = anteHandler(newCtx, tx1, false)
	require.Equal(t, true, abort)
	require.Equal(t, false, res.IsOK())

	// assert that the ante handler will return with `abort` set to false
	acc1.SetCoins(acc1.GetCoins().Plus(sdk.Coins{nativeTokenMintFee}))
	_, res, abort = anteHandler(newCtx, tx1, false)
	require.Equal(t, false, abort)
	require.Equal(t, true, res.IsOK())

	// assert that the ante handler will return with `abort` set to false
	acc1.SetCoins(sdk.Coins{gatewayCreateFee.Plus(nativeTokenIssueFee)})
	_, res, abort = anteHandler(newCtx, tx2, false)
	require.Equal(t, false, abort)
	require.Equal(t, true, res.IsOK())

	// assert that the ante handler will return with `abort` set to false
	acc1.SetCoins(sdk.Coins{})
	_, res, abort = anteHandler(newCtx, tx3, false)
	require.Equal(t, false, abort)
	require.Equal(t, true, res.IsOK())

	// assert that the ante handler will return with `abort` set to true
	newCtx = auth.WithSigners(ctx, []auth.Account{})
	_, res, abort = anteHandler(newCtx, tx3, false)
	require.Equal(t, true, abort)
	require.Equal(t, false, res.IsOK())
}
