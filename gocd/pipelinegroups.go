package gocd

import "context"

// PipelineGroupsService describes the HAL _link resource for the api response object for a pipeline group response.
type PipelineGroupsService service

// PipelineGroups represents a collection of pipeline groups
type PipelineGroups []*PipelineGroup

// PipelineGroup describes a pipeline group API response.
type PipelineGroup struct {
	Name      string      `json:"name"`
	Pipelines []*Pipeline `json:"pipelines"`
}

// List Pipeline groups
func (pgs *PipelineGroupsService) List(ctx context.Context, name string) (*PipelineGroups, *APIResponse, error) {
	type EmbeddedObj struct {
		PipelineGroup []*PipelineGroup `json:"groups"`
	}
	type AllPipelineGroupsResponse struct {
		Embedded EmbeddedObj `json:"_embedded"`
	}
	pg := new(AllPipelineGroupsResponse)
	_, resp, err := pgs.client.getAction(ctx, &APIClientRequest{
		Path:         "admin/pipeline_groups",
		APIVersion:   apiLatest,
		ResponseType: responseTypeJSON,
		ResponseBody: &pg,
	})

	filtered := PipelineGroups{}
	if name != "" && err == nil {
		for _, pipelineGroup := range pg.Embedded.PipelineGroup {
			if pipelineGroup.Name == name {
				filtered = append(filtered, pipelineGroup)
			}
		}
	} else {
		filtered = pg.Embedded.PipelineGroup
	}

	return &filtered, resp, err
}
