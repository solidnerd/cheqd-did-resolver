package types

import (
	"fmt"

	cheqd "github.com/cheqd/cheqd-node/x/cheqd/types"
	resource "github.com/cheqd/cheqd-node/x/resource/types"
)

type ResolutionDidDocMetadata struct {
	Created     string            `json:"created,omitempty"`
	Updated     string            `json:"updated,omitempty"`
	Deactivated bool              `json:"deactivated,omitempty"`
	VersionId   string            `json:"versionId,omitempty"`
	Resources   []ResourcePreview `json:"resources,omitempty"`
}

type ResourcePreview struct {
	ResourceURI       string `json:"resourceURI,omitempty"`
	Name              string `json:"name,omitempty"`
	ResourceType      string `json:"resourceType,omitempty"`
	MediaType         string `json:"mediaType,omitempty"`
	Created           string `json:"created,omitempty"`
	Checksum          string `json:"checksum,omitempty"`
	PreviousVersionId string `json:"previousVersionId,omitempty"`
	NextVersionId     string `json:"nextVersionId,omitempty"`
}

func NewResolutionDidDocMetadata(did string, metadata cheqd.Metadata, resources []*resource.ResourceHeader) ResolutionDidDocMetadata {
	newMetadata := ResolutionDidDocMetadata{
		metadata.Created,
		metadata.Updated,
		metadata.Deactivated,
		metadata.VersionId,
		[]ResourcePreview(nil),
	}
	if metadata.Resources == nil {
		return newMetadata
	}
	for _, r := range resources {
		resourcePreview := ResourcePreview{
			did + RESOURCE_PATH + r.Id,
			r.Name,
			r.ResourceType,
			r.MediaType,
			r.Created,
			fmt.Sprintf("%x", r.Checksum),
			r.PreviousVersionId,
			r.NextVersionId,
		}
		newMetadata.Resources = append(newMetadata.Resources, resourcePreview)
	}
	return newMetadata
}

func TransformToFragmentMetadata(metadata ResolutionDidDocMetadata) ResolutionDidDocMetadata {
	metadata.Resources = nil
	return metadata
}
