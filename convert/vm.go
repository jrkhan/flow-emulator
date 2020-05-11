package convert

import (
	"github.com/dapperlabs/flow-go/engine/execution/computation/virtualmachine"

	sdkConvert "github.com/dapperlabs/flow-emulator/convert/sdk"
	"github.com/dapperlabs/flow-emulator/types"
)

func VMTransactionResultToEmulator(vmTxResult virtualmachine.TransactionResult, txIndex int) types.TransactionResult {

	txID := sdkConvert.FlowIdentifierToSDK(vmTxResult.TransactionID)

	sdkEvents := sdkConvert.RuntimeEventsToSDK(vmTxResult.Events, txID, txIndex)

	return types.TransactionResult{
		TransactionID: txID,
		Error:         vmTxResult.Error,
		Logs:          vmTxResult.Logs,
		Events:        sdkEvents,
	}
}