package core

import (
	"github.com/anoideaopen/foundation/core/telemetry"
	"github.com/anoideaopen/foundation/proto"
	"github.com/hyperledger/fabric-chaincode-go/v2/shim"
	"go.opentelemetry.io/otel/codes"
)

// InvokeContractMethod calls a Chaincode contract method, processes the arguments, and
// returns the result as bytes.
func (cc *Chaincode) InvokeContractMethod(
	traceCtx telemetry.TraceContext,
	stub shim.ChaincodeStubInterface,
	sender *proto.Address,
	method string,
	args []string,
) ([]byte, error) {
	traceCtx, span := cc.contract.TracingHandler().StartNewSpan(traceCtx, "chaincode.CallMethod")
	defer span.End()

	cc.contract.setEnv(&environment{
		stub:  stub,
		trace: traceCtx,
	})
	defer cc.contract.delEnv()

	span.AddEvent("call")
	result, err := cc.Router().Invoke(stub, method, cc.PrependSender(method, sender, args)...)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	span.SetStatus(codes.Ok, "")
	return result, nil
}
