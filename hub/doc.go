// Package hub contains sources for the hub mechanism in the mrun runtime.
//
// As the name suggests hub/mediator logic centralises dependencies and communications
// between subsystems within the mrun runtime. It decides what the state of the program should
// be during the planning phase, and thereafter, executes that plan.
//
// For example, when the hub is given the command to start, it is given a log level, and a path to configuration.
//
//   - reads the config.json
//   - Plans the steps it needs to do, or more specifically, creates an InitChain from ChainLink.
//     items through the ChainLinkFactory, it glues these together based on the configuration.
//   - Creates a HubContext object in memory, which serves as a static object throughout the lifecycle
//     of the Chain that can be read to discover the context such as log level, and progression through the InitChain.
//   - Executes each ChainLink item in order. For each item, passing a pointer to the HubContext in memory,
//     checking an error on return from the ChainLink, and informing the state subsystem based on what happened in
//     that ChainLink execution. The state subsystem then updates the state.json file for that container.
package hub
