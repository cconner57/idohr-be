package storage

// TODO: Phase 2 - File Storage Manager
// Goal: Abstract where files are saved so we can switch providers easily.

// Service interface defines the methods our application needs.
type Service interface {
	Save(filename string, data []byte) error
	Get(filename string) ([]byte, error)
}

// LocalDisk implements Service using the Raspberry Pi's hard drive (uploads/ folder).
// S3Cloud implements Service using AWS S3 (Future Phase).
