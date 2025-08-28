// Package state contains an API for the mrun state subsystem.
//
// # Scopes
//
// There are specific scopes that are applied in the state
// subsystem. Namely the states:
//
//  1. Global
//  2. Container
//
// # Global Scope
//
// The global scope is used for operations that are NOT container specific
// but rather reside within the aperture of the entire system, and require
// or invoke consideration for the consideration for other container states.
//
// # Container Scope
//
// The container scope is used for operations that ARE container specific.
//
// # mrun ps
//
// The `mrun ps` command gets special consideration in the context of the state subsystem.
package state
