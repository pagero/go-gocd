package gocd

// RemoveLinks from the pipeline object for json marshalling.
func (p *RoleCollection) RemoveLinks() {
	p.Links = nil
}

// GetLinks from pipeline
func (p *RoleCollection) GetLinks() *HALLinks {
	return p.Links
}
