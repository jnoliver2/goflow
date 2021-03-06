package flows

import (
	"encoding/json"
	"time"

	"github.com/nyaruka/goflow/utils"
)

type UUID string

type AssetUUID UUID

type NodeUUID UUID

func (u NodeUUID) String() string { return string(u) }

type ExitUUID UUID

func (u ExitUUID) String() string { return string(u) }

type FlowUUID UUID

func (u FlowUUID) String() string { return string(u) }

type ActionUUID UUID

func (u ActionUUID) String() string { return string(u) }

type ContactUUID UUID

func (u ContactUUID) String() string { return string(u) }

type FieldUUID UUID

func (u FieldUUID) String() string { return string(u) }

type ChannelUUID UUID

func (u ChannelUUID) String() string { return string(u) }

type RunUUID UUID

func (u RunUUID) String() string { return string(u) }

type StepUUID UUID

func (u StepUUID) String() string { return string(u) }

type LabelUUID UUID

func (u LabelUUID) String() string { return string(u) }

type GroupUUID UUID

func (u GroupUUID) String() string { return string(u) }

// RunStatus represents the current status of the engine session
type SessionStatus string

const (
	// SessionStatusActive represents a session that is still active
	SessionStatusActive SessionStatus = "active"

	// SessionStatusCompleted represents a session that has run to completion
	SessionStatusCompleted SessionStatus = "completed"

	// SessionStatusWaiting represents a session which is waiting for something from the caller
	SessionStatusWaiting SessionStatus = "waiting"

	// SessionStatusErrored represents a session that encountered an error
	SessionStatusErrored SessionStatus = "errored"
)

func (r SessionStatus) String() string { return string(r) }

// RunStatus represents the current status of the flow run
type RunStatus string

const (
	// RunStatusActive represents a run that is still active
	RunStatusActive RunStatus = "active"

	// RunStatusCompleted represents a run that has run to completion
	RunStatusCompleted RunStatus = "completed"

	// RunStatusWaiting represents a run which is waiting for something from the caller
	RunStatusWaiting RunStatus = "waiting"

	// RunStatusErrored represents a run that encountered an error
	RunStatusErrored RunStatus = "errored"

	// RunStatusExpired represents a run that expired due to inactivity
	RunStatusExpired RunStatus = "expired"

	// RunStatusInterrupted represents a run that was interrupted by another flow
	RunStatusInterrupted RunStatus = "interrupted"
)

func (r RunStatus) String() string { return string(r) }

type AssetType string

const (
	AssetTypeFlow    AssetType = "flow"
	AssetTypeChannel AssetType = "channel"
)

type Asset interface {
	AssetType() AssetType
	AssetUUID() AssetUUID
	Validate(AssetStore) error
}

type AssetStore interface {
	IncludeAssets(json.RawMessage) error
	GetChannel(ChannelUUID) (Channel, error)
	GetFlow(FlowUUID) (Flow, error)
}

type Flow interface {
	Asset

	UUID() FlowUUID
	Name() string
	Language() utils.Language
	ExpireAfterMinutes() int
	Translations() FlowTranslations

	Nodes() []Node
	GetNode(uuid NodeUUID) Node
}

type Node interface {
	UUID() NodeUUID

	Router() Router
	Actions() []Action
	Exits() []Exit
	Wait() Wait
}

type Action interface {
	UUID() ActionUUID

	Execute(FlowRun, Step) error
	Validate(AssetStore) error
	utils.Typed
}

type Router interface {
	PickRoute(FlowRun, []Exit, Step) (Route, error)
	Validate([]Exit) error
	ResultName() string
	utils.Typed
}

type Route struct {
	exit  ExitUUID
	match string
}

func (r Route) Exit() ExitUUID { return r.exit }
func (r Route) Match() string  { return r.match }

var NoRoute = Route{}

func NewRoute(exit ExitUUID, match string) Route {
	return Route{exit, match}
}

type Exit interface {
	UUID() ExitUUID
	DestinationNodeUUID() NodeUUID
	Name() string
}

type Wait interface {
	utils.Typed

	Begin(FlowRun, Step)
	CanResume(FlowRun, Step) bool
	HasTimedOut() bool
}

// FlowTranslations provide a way to get the Translations for a flow for a specific language
type FlowTranslations interface {
	GetLanguageTranslations(utils.Language) Translations
}

// Translations provide a way to get the translation for a specific language for a uuid/key pair
type Translations interface {
	GetTextArray(uuid UUID, key string) []string
}

type Event interface {
	CreatedOn() time.Time
	SetCreatedOn(time.Time)

	FromCaller() bool
	SetFromCaller(bool)

	Apply(FlowRun) error

	utils.Typed
}

type Input interface {
	utils.VariableResolver
	utils.Typed

	CreatedOn() time.Time
	Channel() Channel
}

type Step interface {
	utils.VariableResolver

	UUID() StepUUID
	NodeUUID() NodeUUID
	ExitUUID() ExitUUID

	ArrivedOn() time.Time
	LeftOn() *time.Time

	Leave(ExitUUID)

	Events() []Event
}

// LogEntry is a container for a new event generated by the engine, i.e. not from the caller
type LogEntry interface {
	Step() Step
	Action() Action
	Event() Event
}

// Session represents the session of a flow run which may contain many runs
type Session interface {
	Assets() AssetStore

	Environment() utils.Environment
	SetEnvironment(utils.Environment)

	Contact() *Contact
	SetContact(*Contact)

	Status() SessionStatus
	SetTrigger(Flow, FlowRun)
	Wait() Wait
	FlowOnStack(FlowUUID) bool

	StartFlow(FlowUUID, []Event) error
	Resume([]Event) error
	Runs() []FlowRun
	GetRun(RunUUID) (FlowRun, error)

	Log() []LogEntry
	LogEvent(Step, Action, Event)
	ClearLog()
}

// FlowRun represents a single run on a flow by a single contact
type FlowRun interface {
	UUID() RunUUID

	Environment() utils.Environment
	Session() Session
	Context() utils.VariableResolver

	Flow() Flow
	Results() *Results

	Contact() *Contact
	SetContact(*Contact)

	SetExtra(utils.JSONFragment)
	Extra() utils.JSONFragment

	Status() RunStatus
	SetStatus(RunStatus)
	Exit(RunStatus)

	Input() Input
	SetInput(Input)

	ApplyEvent(Step, Action, Event) error
	AddError(Step, Action, error)
	AddFatalError(Step, Action, error)

	CreateStep(Node) Step
	Path() []Step
	PathLocation() (Step, Node, error)

	GetText(uuid UUID, key string, native string) string
	GetTextArray(uuid UUID, key string, native []string) []string

	Webhook() *utils.RequestResponse
	SetWebhook(*utils.RequestResponse)

	Child() FlowRunReference
	Parent() FlowRunReference
	Ancestors() []FlowRunReference

	CreatedOn() time.Time
	ExpiresOn() *time.Time
	ResetExpiration(*time.Time)
	ExitedOn() *time.Time
}

// FlowRunReference represents a flow run reference within a flow
type FlowRunReference interface {
	UUID() RunUUID
	Flow() Flow
	Contact() *Contact

	Results() *Results
	Status() RunStatus

	ExpiresOn() *time.Time
	ResetExpiration(*time.Time)
	ExitedOn() *time.Time
}

// ChannelType represents the type of a Channel
type ChannelType string

func (ct ChannelType) String() string { return string(ct) }

// Channel represents a channel for sending and receiving messages
type Channel interface {
	Asset

	UUID() ChannelUUID
	Name() string
	Address() string
	Type() ChannelType
}

// MsgDirection is the direction of a Msg (either in or out)
type MsgDirection string

const (
	// MsgOut represents an outgoing message
	MsgOut MsgDirection = "O"

	// MsgIn represents an incoming message
	MsgIn MsgDirection = "I"
)
