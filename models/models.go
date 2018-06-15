package models

// Repository DataModel
type Repository struct {
	ID   uint64
	Name string
}

// Tag Datamodel
type Tag struct {
	ID   uint64
	Name string
}

// Manifest datamodel
type Manifest struct {
	ID     uint64
	digest string
}

// Layer data model
type Layer struct {
	ID     uint64
	digest string
}

// ManifestLayer link  datamodel
type ManifestLayer struct {
	manifestID uint64
	layerID    uint64
}

// LayerManifest link datamodel
type LayerManifest struct {
	layerID    uint64
	manifestID uint64
}
