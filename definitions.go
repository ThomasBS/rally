package rally

type reference struct {
	ReferenceUrl string `json:"_ref"`
}

type queryReference struct {
	ReferenceUrl string `json:"_ref"`
	Count        int
}

type queryResult struct {
	TotalResultCount int
	StartIndex       int
	PageSize         int
}

type AllowedAttributeValue struct {
	PersistableObject

	AttributeDefinitionReference reference `json:"AttributeDefinition"`
	StringValue                  string
	ValueIndex                   int
}

type AllowedAttributeValueQuery struct {
	queryResult

	Result []AllowedAttributeValue
}

type AllowedQueryOperator struct {
	PersistableObject

	OperatorName string
}

type AllowedQueryOperatorQuery struct {
	queryResult

	Result []AllowedQueryOperator
}

type Artifact struct {
	WorkspaceDomainObject

	ChangesetsQueryReference queryReference `json:"Changesets"`
	Description              string

	// ConversationPost
	DiscussionQueryReference queryReference `json:"Discussion"`

	DisplayColor                 string
	Expedite                     bool
	FormattedId                  string `json:"FormattedID"`
	LastUpdateDate               string
	LatestDiscussionAgeInMinutes int
	MilestonesQueryReference     queryReference `json:"Milestones"`
	Name                         string
	Notes                        string
	OwnerReference               reference `json:"Owner"`
	ProjectReference             reference `json:"Project"`
	Ready                        bool
	RevisionHistoryReference     reference      `json:"RevisionHistory"`
	TagsQueryReference           queryReference `json:"Tags"`
}

type ArtifactQuery struct {
	queryResult

	Result []Artifact
}

type Attachment struct {
	WorkspaceDomainObject

	ArtifactReference reference `json:"Artifact"`

	// AttachmentContent
	ContentReference reference `json:"Content"`

	ContentType             string
	Description             string
	Name                    string
	Size                    int
	TestCaseResultReference reference `json:"TestCaseResult"`
	UserReference           reference `json:"User"`
}

type AttachmentQuery struct {
	queryResult

	Result []Attachment
}

type AttachmentContent struct {
	WorkspaceDomainObject

	Content string
}

type AttributeDefinition struct {
	WorkspaceDomainObject

	// AllowedQueryOperatorQuery
	AllowedQueryOperatorsQueryReference queryReference `json:"AllowedQueryOperators"`

	// AllowedAttributeValueQuery
	AllowedValuesQueryReference queryReference `json:"AllowedValues"`

	// TypeDefinition
	AllowedValueTypeReference reference `json:"AllowedValueType"`

	AttributeType       string
	Constrained         bool
	Custom              bool
	ElementName         string
	Filterable          bool
	Hidden              bool
	MaxFractionalDigits int
	MaxLength           int
	Name                string
	Note                string
	Owned               bool
	ReadOnly            bool
	RealAttributeType   string
	Required            bool
	Sortable            bool
	SystemRequired      bool
	Type                string

	// TypeDefinition
	TypeDefinitionReference reference `json:"TypeDefinition"`

	VisibleOnlyToAdmins bool
}

type AttributeDefinitionQuery struct {
	queryResult

	Result []AttributeDefinition
}

type Blocker struct {
	DomainObject

	// User
	BlockedByReference reference `json:"BlockedBy"`

	// Artifact
	WorkProductReference reference `json:"WorkProduct"`
}

type Build struct {
	WorkspaceDomainObject

	BuildDefinitionReference reference      `json:"BuildDefinition"`
	ChangesetsQueryReference queryReference `json:"Changesets"`
	Duration                 float64
	Message                  string
	Number                   string
	Start                    string
	Status                   string
	Uri                      string
}

type BuildQuery struct {
	queryResult

	Result []Build
}

type BuildDefinition struct {
	WorkspaceDomainObject

	BuildsQueryReference queryReference `json:"Builds"`
	Description          string

	// Build
	LastBuildReference reference `json:"LastBuild"`

	LastStatus             string
	Name                   string
	ProjectReference       reference `json:"Project"`
	ProjectsQueryReference reference `json:"Projects"`
	Uri                    string
}

type Change struct {
	WorkspaceDomainObject

	Action             string
	Base               string
	ChangesetReference reference `json:"Changeset"`
	Extension          string
	PathAndFilename    string
	Uri                string
}

type ChangeQuery struct {
	queryResult

	Result []Change
}

type Changeset struct {
	WorkspaceDomainObject

	ArtifactsQueryReference queryReference `json:"Artifacts"`

	// User
	AuthorReference reference `json:"Author"`

	BuildsQueryReference   queryReference `json:"Builds"`
	ChangesQueryReference  queryReference `json:"Changes"`
	CommitTimestamp        string
	Message                string
	Name                   string
	Revision               string
	SCMRepositoryReference reference `json:"SCMRepository"`
	Uri                    string
}

type ChangesetQuery struct {
	queryResult

	Result []Changeset
}

type ConversationPost struct {
	WorkspaceDomainObject

	ArtifactReference reference `json:"Artifact"`
	PostNumber        int
	Text              string
	UserReference     reference `json:"User"`
}

type ConversationPostQuery struct {
	queryResult

	Result []ConversationPost
}

type Defect struct {
	SchedulableArtifact

	AcceptedDate              string
	AffectsDoc                bool
	AttachmentsQueryReference queryReference `json:"Attachments"`
	Blocked                   bool
	BlockedReason             string
	BlockerReference          reference `json:"Blocker"`
	ClosedDate                string
	DefectSuiteQueryReference queryReference `json:"DefectSuites"`
	DragAndDropRank           string

	// Defect
	DuplicatesQueryReference queryReference `json:"Duplicates"`

	Environment          string
	FixedInBuild         string
	FoundInBuild         string
	InProgressDate       string
	IterationReference   reference `json:"Iteration"`
	OpenedDate           string
	Package              string
	PlanEstimate         float64
	Priority             string
	Recycled             bool
	ReleaseReference     reference `json:"Release"`
	ReleaseNote          bool
	RequirementReference reference `json:"Requirement"`
	Resolution           string
	SalesforceCaseId     string `json:"SalesforceCaseID"`
	SalesforceCaseNumber string
	Severity             string
	State                string

	// User
	SubmittedByReference reference `json:"SubmittedBy"`

	TargetBuild             string
	TargetDate              string
	TaskActualTotal         float64
	TaskEstimateTotal       float64
	TaskRemainingTotal      float64
	TasksQueryReference     queryReference `json:"Tasks"`
	TaskStatus              string
	TestCaseReference       reference      `json:"TestCase"`
	TestCaseResultReference reference      `json:"TestCaseResult"`
	TestCasesQueryReference queryReference `json:"TestCases"`
	TestCaseStatus          string
	VerifiedInBuild         string
}

type DefectQuery struct {
	queryResult

	Result []Defect
}

type DefectSuite struct {
	SchedulableArtifact

	AcceptedDate              string
	AttachmentsQueryReference queryReference `json:"Attachments"`
	Blocked                   bool
	BlockedReason             string
	BlockerReference          reference `json:"Blocker"`
	ClosedDefectCount         int
	DefectsQueryReference     reference `json:"Defects"`
	DefectStatus              string
	DragAndDropRank           string
	InProgressDate            string
	IterationReference        reference `json:"Iteration"`
	Package                   string
	PlanEstimate              float64
	ReleaseReference          reference `json:"Release"`
	TaskActualTotal           float64
	TaskEstimateTotal         float64
	TaskRemainingTotal        float64
	TasksQueryReference       queryReference `json:"Tasks"`
	TaskStatus                string
	TotalDefectCount          int
}

type DefectSuiteQuery struct {
	queryResult

	Result []DefectSuite
}

type DomainObject struct {
	PersistableObject

	SubscriptionReference reference `json:"Subscription"`
}

// https://rally1.rallydev.com/slm/doc/webservice/objectModel.sp#HierarchicalRequirement
type HierarchicalRequirement struct {
	Requirement

	AcceptedDate     string
	Blocked          bool
	BlockedReason    string
	BlockerReference reference `json:"Blocker"`

	// HierarchicalRequirement
	ChildrenQueryReference queryReference `json:"Children"`

	DefectsQueryReference queryReference `json:"Defects"`
	DefectStatus          string
	DirectChildrenCount   int
	DragAndDropRank       string
	HasParent             bool
	InProgressDate        string
	IterationReference    reference `json:"Iteration"`

	// HierarchicalRequirement
	ParentReference reference `json:"Parent"`

	PlanEstimate float64

	// HierarchicalRequirement
	PredecessorsQueryReference queryReference `json:"Predecessors"`

	Recycled         bool
	ReleaseReference reference `json:"Release"`

	// HierarchicalRequirement
	SuccessorsQueryReference queryReference `json:"Successors"`

	TaskActualTotal         float64
	TaskEstimateTotal       float64
	TaskRemainingTotal      float64
	TasksQueryReference     queryReference `json:"Tasks"`
	TaskStatus              string
	TestCasesQueryReference queryReference `json:"TestCases"`
	TestCaseStatus          string
}

type HierarchicalRequirementQuery struct {
	queryResult

	Result []HierarchicalRequirement
}

type Iteration struct {
	WorkspaceDomainObject

	EndDate                               string
	Name                                  string
	Notes                                 string
	PlanEstimate                          float64
	PlannedVelocity                       float64
	ProjectReference                      reference `json:"Project"`
	RevisionHistoryReference              reference `json:"RevisionHistory"`
	StartDate                             string
	State                                 string
	TaskActualTotal                       float64
	TaskEstimateTotal                     float64
	TaskRemainingTotal                    float64
	Theme                                 string
	UserIterationCapacitiesQueryReference queryReference `json:"UserIterationCapacities"`
}

type Milestone struct {
	WorkspaceDomainObject

	// Artifact
	ArtifactsQueryReference queryReference `json:"Artifacts"`

	DisplayColor string
	FormattedId  string `json:"FormattedID"`
	Name         string
	Notes        string

	// ProjectQuery
	ProjectsQueryReference queryReference `json:"Projects"`

	RevisionHistoryReference reference `json:"RevisionHistory"`
	TargetDate               string

	// Project
	TargetProjectReference reference `json:"TargetProject"`

	TotalArtifactCount int
}

type MilestoneQuery struct {
	queryResult

	Result []Milestone
}

type PersistableObject struct {
	CreationDate string
	ObjectId     string `json:"ObjectID"`
	VersionId    string
}

type Project struct {
	WorkspaceDomainObject

	BuildDefinitionsQueryReference queryReference `json:"BuildDefinitions"`

	// ProjectQuery
	ChildrenQueryReference queryReference `json:"Children"`

	Description string

	// UserQuery
	EditorsQueryReference queryReference `json:"Editors"`

	// IterationQuery
	IterationsQueryReference queryReference `json:"Iterations"`

	Name  string
	Notes string

	// User
	OwnerReference reference `json:"Owner"`

	// Project
	ParentReference reference `json:"Parent"`

	// ReleaseQuery
	ReleasesQueryReference queryReference `json:"Releases"`

	SchemaVersion string
	State         string

	// UserQuery
	TeamMembersQueryReference queryReference `json:"TeamMembers"`
}

type ProjectQuery struct {
	queryResult

	Result []Project
}

type Release struct {
	WorkspaceDomainObject

	Accepted                     float64
	GrossEstimateConversionRatio float64
	Name                         string
	Notes                        string
	PlanEstimate                 float64
	PlannedVelocity              float64
	ProjectReference             reference `json:"Project"`
	ReleaseDate                  string
	ReleaseStartDate             string
	RevisionHistoryReference     reference `json:"RevisionHistory"`
	State                        string
	TaskActualTotal              float64
	TaskEstimateTotal            float64
	TaskRemainingTotal           float64
	Theme                        string
	Version                      string
}

type Requirement struct {
	SchedulableArtifact

	// AttachmentQuery
	AttachmentsQueryReference queryReference `json:"Attachments"`

	Package string
}

type Revision struct {
	WorkspaceDomainObject

	Description              string
	RevisionHistoryReference reference `json:"RevisionHistory"`
	RevisionNumber           int
	UserReference            reference `json:"User"`
}

type RevisionQuery struct {
	queryResult

	Result []Revision
}

type RevisionHistory struct {
	WorkspaceDomainObject

	RevisionsQueryReference queryReference `json:"Revisions"`
}

type SCMRepository struct {
	WorkspaceDomainObject

	Description            string
	Name                   string
	ProjectsQueryReference queryReference `json:"Projects"`
	SCMType                string
	Uri                    string
}

type SchedulableArtifact struct {
	Artifact

	AcceptedDate        string
	Blocked             bool
	BlockedReason       string
	BlockerReference    reference `json:"Blocker"`
	DragAndDropRank     string
	IterationReference  reference `json:"Iteration"`
	PlanEstimate        float64
	ReleaseReference    reference `json:"Release"`
	ScheduleState       string
	TaskActualTotal     float64
	TaskEstimateTotal   float64
	TaskRemainingTotal  float64
	TasksQueryReference queryReference `json:"Tasks"`
}

type Subscription struct {
	PersistableObject

	ExpirationDate           string
	MaximumCustomUserFields  int
	MaximumProjects          int
	Modules                  string
	Name                     string
	PasswordExpirationDays   int
	PreviousPasswordCount    int
	ProjectHierarchyEnabled  bool
	RevisionHistoryReference reference `json:"RevisionHistory"`
	SessionTimeoutSeconds    int
	StoryHierarchyEnabled    bool
	StoryHierarchyType       string
	SubscriptionId           int `json:"SubscriptionID"`
	SubscriptionType         string

	// Workspace
	WorkspacesQueryReference queryReference `json:"Workspaces"`
}

type Tag struct {
	WorkspaceDomainObject

	Archived bool
	Name     string
}

type TagQuery struct {
	queryResult

	Result []Tag
}

type Task struct {
	Artifact

	Actuals                   float64
	AttachmentsQueryReference queryReference `json:"Attachments"`
	Blocked                   bool
	BlockedReason             string
	DragAndDropRank           string
	Estimate                  float64
	IterationReference        reference `json:"Iteration"`
	ProjectReference          reference `json:"Project"`
	Recycled                  bool
	ReleaseReference          reference `json:"Release"`
	State                     string
	TaskIndex                 int
	TimeSpent                 float64
	ToDo                      float64

	// SchedulableArtifact
	WorkProductReference reference `json:"WorkProduct"`
}

type TaskQuery struct {
	queryResult

	Results []Task
}

type TestCase struct {
	Artifact

	AttachmentsQueryReference queryReference `json:"Attachments"`
	DefectsQueryReference     queryReference `json:"Defects"`
	DefectStatus              string
	DragAndDropRank           string
	LastBuild                 string
	LastRun                   string
	LastVerdict               string
	Method                    string
	Objective                 string
	Package                   string
	PostConditions            string
	PreConditions             string
	Priority                  string
	Recycled                  bool

	// TestCaseResultQuery
	ResultsQueryReference queryReference `json:"Results"`

	Risk string

	// TestCaseStepQuery
	StepsQueryReference queryReference `json:"Steps"`

	Type                     string
	ValidationExpectedResult string
	ValidationInput          string

	// Artifact
	WorkProductReference reference `json:"WorkProduct"`
}

type TestCaseQuery struct {
	queryResult

	Result []TestCase
}

type TestCaseResult struct {
	WorkspaceDomainObject

	AttachmentsQueryReference queryReference `json:"Attachments"`
	Build                     string
	Date                      string
	Duration                  float64
	Notes                     string
	TestCaseReference         reference `json:"TestCase"`

	// User
	TesterReference reference `json:"Tester"`

	Verdict string
}

type TestCaseResultQuery struct {
	queryResult

	Result []TestCaseResult
}

type TestCaseStep struct {
	WorkspaceDomainObject

	ExpectedResult    string
	Input             string
	StepIndex         int
	TestCaseReference reference `json:"TestCase"`
}

type TestCaseStepQuery struct {
	queryResult

	Result []TestCaseStep
}

type TypeDefinition struct {
	WorkspaceDomainObject

	Abstract bool

	// AttributeDefinitionQuery
	AttributesQueryReference queryReference `json:"Attributes"`

	Creatable   bool
	Deletable   bool
	DisplayName string
	ElementName string
	IdPrefix    string `json:"IDPrefix"`
	Name        string
	Note        string

	// TypeDefinition
	ParentReference reference `json:"Parent"`

	Queryable  bool
	ReadOnly   bool
	Restorable bool

	// RevisionHistory
	RevisionHistoryReference reference `json:"RevisionHistory"`

	TypePath     string
	UserListable bool
}

type TypeDefinitionQuery struct {
	queryResult

	Result []TypeDefinition
}

type User struct {
	DomainObject

	CostCenter               string
	Department               string
	Disabled                 bool
	DisplayName              string
	EmailAddress             string
	FirstName                string
	LandingPage              string
	LastLoginDate            string
	LastName                 string
	LastPasswordUpdateDate   string
	MiddleName               string
	NetworkId                string `json:"NetworkID"`
	OfficeLocation           string
	OnpremLdapUsername       string
	Phone                    string
	Planner                  bool
	RevisionHistoryReference reference `json:"RevisionHistory"`
	Role                     string
	ShortDisplayName         string
	SubscriptionAdmin        bool
	SubscriptionPermission   string

	// ProjectQuery
	TeamMembershipsQueryReference queryReference `json:"TeamMemberships"`

	UserName string

	// UserPermissionQuery
	UserPermissionsQueryReference queryReference `json:"UserPermissions"`

	UserProfileReference reference `json:"UserProfile"`
	WorkspacePermission  string
}

type UserIterationCapacity struct {
	WorkspaceDomainObject

	Capacity           float64
	IterationReference reference `json:"Iteration"`
	Load               float64
	ProjectReference   reference `json:"Project"`
	TaskEstimates      float64
	UserReference      reference `json:"User"`
}

type UserIterationCapacityQuery struct {
	queryResult

	Results []UserIterationCapacity
}

type UserPermission struct {
	DomainObject

	CustomObjectId string `json:"CustomObjectID"`
	Name           string
	Role           string
	UserReference  reference `json:"User"`
}

type UserPermissionQuery struct {
	queryResult

	Result []UserPermission
}

type UserProfile struct {
	DomainObject

	DateFormat                     string
	DateTimeFormat                 string
	DefaultDetailPageToViewingMode bool

	// Project
	DefaultProjectReference reference `json:"DefaultProject"`

	// Workspace
	DefaultWorkspaceReference reference `json:"DefaultWorkspace"`

	EmailNotificationEnabled bool
	SessionTimeoutSeconds    int
	SessionTimeoutWarning    bool
	TimeZone                 string
	WelcomePageHidden        bool
}

type Workspace struct {
	DomainObject

	Description string
	Name        string
	Notes       string

	// User
	OwnerReference reference `json:"Owner"`

	// ProjectQuery
	ProjectsQueryReference queryReference `json:"Projects"`

	// RevisionHistory
	RevisionHistoryReference reference `json:"RevisionHistory"`

	SchemaVersion string
	State         string
	Style         string

	// TypeDefinitionQuery
	TypeDefinitionsQueryReference queryReference `json:"TypeDefinitions"`

	// WorkspaceConfiguration
	WorkspaceConfigurationReference reference `json:"WorkspaceConfiguration"`
}

type WorkspaceConfiguration struct {
	WorkspaceDomainObject

	BuildandChangesetEnabled  bool
	DateFormat                string
	DateTimeFormat            string
	DragDropRankingEnabled    bool
	IterationEstimateUnitName string
	ReleaseEstimateUnitName   string
	TaskUnitName              string
	TimeTrackerEnabled        bool
	TimeZone                  string
	WorkDays                  string
}

type WorkspaceDomainObject struct {
	DomainObject

	WorkspaceReference reference `json:"Workspace"`
}
