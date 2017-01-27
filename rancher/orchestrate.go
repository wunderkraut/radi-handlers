package rancher

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * Orchestration Handler using Rancher
 */

// Rancher Orchestration Handler
type RancherOrchestrateHandler struct {
	RancherBaseClientHandler
}

// Initialize and activate the Handler
func (orchestrate *RancherOrchestrateHandler) Init() api_operation.Result {
	result := api_operation.BaseResult{}
	result.Set(true, []error{})


	ops := api_operation.Operations{}

	// ops.Add(api_operation.Operation(&RancherOrchestrateUpOperation{BaseRancherServiceOperation: *baseOperation}))
	// ops.Add(api_operation.Operation(&RancherOrchestrateStopOperation{BaseRancherServiceOperation: *baseOperation}))
	// ops.Add(api_operation.Operation(&RancherOrchestrateDownOperation{BaseRancherServiceOperation: *baseOperation}))

	orchestrate.operations = &ops

	return api_operation.Result(&result)
}

// Rturn a string identifier for the Handler (not functionally needed yet)
func (orchestrate *RancherOrchestrateHandler) Id() string {
	return "rancher.orchestrate"
}
