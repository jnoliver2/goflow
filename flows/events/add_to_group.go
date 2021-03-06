package events

import "github.com/nyaruka/goflow/flows"

// TypeAddToGroup is the type of our add to group action
const TypeAddToGroup string = "add_to_group"

// AddToGroupEvent events will be created with the groups a contact should be added to.
//
// ```
//   {
//     "type": "add_to_group",
//     "created_on": "2006-01-02T15:04:05Z",
//     "groups": [{
//	     "uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d",
//       "name": "Survey Audience"
//     }]
//   }
// ```
//
// @event add_to_group
type AddToGroupEvent struct {
	BaseEvent
	Groups []*flows.Group `json:"groups"  validate:"required,min=1,dive,uuid4"`
}

// NewGroupEvent returns a new group event
func NewGroupEvent(groups []*flows.Group) *AddToGroupEvent {
	return &AddToGroupEvent{
		BaseEvent: NewBaseEvent(),
		Groups:    groups,
	}
}

// Type returns the type of this event
func (e *AddToGroupEvent) Type() string { return TypeAddToGroup }

// Apply applies this event to the given run
func (e *AddToGroupEvent) Apply(run flows.FlowRun) error {
	for _, group := range e.Groups {
		run.Contact().AddGroup(group.UUID(), group.Name())
	}
	return nil
}
