package snapshot_agent

import "bytes"

// SnapshotWriter is the main interface for writing snapshots to varius places.
type SnapshotWriter interface {
	// WriteSnapshot writes the snapshot from the reader to the specified location.
	// To add new locations, implement this interface and update the configuration.
	WriteSnapshot(buf *bytes.Buffer, currentTs int64) (succes bool, err error)
}
