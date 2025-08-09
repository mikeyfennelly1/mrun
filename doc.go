// Package mrun provides a low-level OCI (Open Container Initiative) compliant container runtime.
//
// # Overview
//
// mrun is a low-level OCI (Open Container Initiative) compliant container runtime.
//
// # OCI compliance
//
// By being OCI compliant it has the following responsibilities:
//
//   - Spawns containers on a host based on a container bundle.
//   - Manages the state of the container after startup.
//   - Provides subcommands for administrating entities to manage container state:
//     state, create, start, kill, and delete.
//
// # Prerequisite reading
//
//   - Open Container Initiative Spec: https://github.com/mikeyfennelly1/mrun/tree/main/docs.
//   - mrun architecture: https://github.com/mikeyfennelly1/mrun/tree/main/docs.
package mrun
