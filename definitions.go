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

	AttributeDefinitionReference reference `json:"AttributeDefinition"` // AttributeDefinition

	StringValue string
	ValueIndex  int
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

	ChangesetsQueryReference queryReference `json:"Changesets"`      // ChangesetQuery
	DiscussionQueryReference queryReference `json:"Discussion"`      // ConversationPostQuery
	MilestonesQueryReference queryReference `json:"Milestones"`      // MilestoneQuery
	OwnerReference           reference      `json:"Owner"`           // Owner
	ProjectReference         reference      `json:"Project"`         // Project
	RevisionHistoryReference reference      `json:"RevisionHistory"` // RevisionHistory
	TagsQueryReference       queryReference `json:"Tags"`            // TagQuery

	Description                  string
	DisplayColor                 string
	Expedite                     bool
	FormattedId                  string `json:"FormattedID"`
	LastUpdateDate               string
	LatestDiscussionAgeInMinutes int
	Name                         string
	Notes                        string
	Ready                        bool
}

type ArtifactQuery struct {
	queryResult

	Result []Artifact
}

type Attachment struct {
	WorkspaceDomainObject

	ArtifactReference       reference `json:"Artifact"`       // Artifact
	ContentReference        reference `json:"Content"`        // AttachmentContent
	TestCaseResultReference reference `json:"TestCaseResult"` // TestCaseResult
	UserReference           reference `json:"User"`           // User

	ContentType string
	Description string
	Name        string
	Size        int
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

	AllowedQueryOperatorsQueryReference queryReference `json:"AllowedQueryOperators"` // AllowedQueryOperatorQuery
	AllowedValuesQueryReference         queryReference `json:"AllowedValues"`         // AllowedAttributeValueQuery
	AllowedValueTypeReference           reference      `json:"AllowedValueType"`      // TypeDefinition
	TypeDefinitionReference             reference      `json:"TypeDefinition"`        // TypeDefinition

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
	VisibleOnlyToAdmins bool
}

type AttributeDefinitionQuery struct {
	queryResult

	Result []AttributeDefinition
}

type Blocker struct {
	DomainObject

	BlockedByReference   reference `json:"BlockedBy"`   // User
	WorkProductReference reference `json:"WorkProduct"` // Artifact
}

type Build struct {
	WorkspaceDomainObject

	BuildDefinitionReference reference      `json:"BuildDefinition"` // BuildDefinition
	ChangesetsQueryReference queryReference `json:"Changesets"`      // Changeset

	Duration float64
	Message  string
	Number   string
	Start    string
	Status   string
	Uri      string
}

type BuildQuery struct {
	queryResult

	Result []Build
}

type BuildDefinition struct {
	WorkspaceDomainObject

	BuildsQueryReference   queryReference `json:"Builds"`    // BuildQuery
	LastBuildReference     reference      `json:"LastBuild"` // Build
	ProjectReference       reference      `json:"Project"`   // Project
	ProjectsQueryReference reference      `json:"Projects"`  // ProjectQuery

	Description string
	LastStatus  string
	Name        string
	Uri         string
}

type Change struct {
	WorkspaceDomainObject

	ChangesetReference reference `json:"Changeset"` // Changeset

	Action          string
	Base            string
	Extension       string
	PathAndFilename string
	Uri             string
}

type ChangeQuery struct {
	queryResult

	Result []Change
}

type Changeset struct {
	WorkspaceDomainObject

	ArtifactsQueryReference queryReference `json:"Artifacts"`     // ArtifactQuery
	AuthorReference         reference      `json:"Author"`        // User
	BuildsQueryReference    queryReference `json:"Builds"`        // BuildQuery
	ChangesQueryReference   queryReference `json:"Changes"`       // ChangeQuery
	SCMRepositoryReference  reference      `json:"SCMRepository"` // SCMRepository

	CommitTimestamp string
	Message         string
	Name            string
	Revision        string
	Uri             string
}

type ChangesetQuery struct {
	queryResult

	Result []Changeset
}

type ConversationPost struct {
	WorkspaceDomainObject

	ArtifactReference reference `json:"Artifact"` // Artifact
	UserReference     reference `json:"User"`     // User

	PostNumber int
	Text       string
}

type ConversationPostQuery struct {
	queryResult

	Result []ConversationPost
}

type Defect struct {
	SchedulableArtifact

	AttachmentsQueryReference queryReference `json:"Attachments"`    // AttachmentQuery
	BlockerReference          reference      `json:"Blocker"`        // Blocker
	DefectSuiteQueryReference queryReference `json:"DefectSuites"`   // DefectSuiteQuery
	DuplicatesQueryReference  queryReference `json:"Duplicates"`     // DefectQuery
	IterationReference        reference      `json:"Iteration"`      // Iteration
	ReleaseReference          reference      `json:"Release"`        // Release
	RequirementReference      reference      `json:"Requirement"`    // Requirement
	SubmittedByReference      reference      `json:"SubmittedBy"`    // User
	TasksQueryReference       queryReference `json:"Tasks"`          // TaskQuery
	TestCaseReference         reference      `json:"TestCase"`       // TestCase
	TestCaseResultReference   reference      `json:"TestCaseResult"` // TestCaseResult
	TestCasesQueryReference   queryReference `json:"TestCases"`      // TestCaseQuery

	AcceptedDate         string
	AffectsDoc           bool
	Blocked              bool
	BlockedReason        string
	ClosedDate           string
	DragAndDropRank      string
	Environment          string
	FixedInBuild         string
	FoundInBuild         string
	InProgressDate       string
	OpenedDate           string
	Package              string
	PlanEstimate         float64
	Priority             string
	Recycled             bool
	ReleaseNote          bool
	Resolution           string
	SalesforceCaseId     string `json:"SalesforceCaseID"`
	SalesforceCaseNumber string
	Severity             string
	State                string
	TargetBuild          string
	TargetDate           string
	TaskActualTotal      float64
	TaskEstimateTotal    float64
	TaskRemainingTotal   float64
	TaskStatus           string
	TestCaseStatus       string
	VerifiedInBuild      string
}

type DefectQuery struct {
	queryResult

	Result []Defect
}

type DefectSuite struct {
	SchedulableArtifact

	AttachmentsQueryReference queryReference `json:"Attachments"` // AttachmentQuery
	BlockerReference          reference      `json:"Blocker"`     // Blocker
	DefectsQueryReference     reference      `json:"Defects"`     // DefectQuery
	IterationReference        reference      `json:"Iteration"`   // Iteration
	ReleaseReference          reference      `json:"Release"`     // Release
	TasksQueryReference       queryReference `json:"Tasks"`       // TaskQuery

	AcceptedDate       string
	Blocked            bool
	BlockedReason      string
	ClosedDefectCount  int
	DefectStatus       string
	DragAndDropRank    string
	InProgressDate     string
	Package            string
	PlanEstimate       float64
	TaskActualTotal    float64
	TaskEstimateTotal  float64
	TaskRemainingTotal float64
	TaskStatus         string
	TotalDefectCount   int
}

type DefectSuiteQuery struct {
	queryResult

	Result []DefectSuite
}

type DomainObject struct {
	PersistableObject

	SubscriptionReference reference `json:"Subscription"` // Subscription
}

type HierarchicalRequirement struct {
	Requirement

	BlockerReference           reference      `json:"Blocker"`      // Blocker
	ChildrenQueryReference     queryReference `json:"Children"`     // HierarchicalRequirementQuery
	DefectsQueryReference      queryReference `json:"Defects"`      // DefectQuery
	IterationReference         reference      `json:"Iteration"`    // Iteration
	ParentReference            reference      `json:"Parent"`       // HierarchicalRequirement
	PredecessorsQueryReference queryReference `json:"Predecessors"` // HierarchicalRequirementQuery
	ReleaseReference           reference      `json:"Release"`      // Release
	SuccessorsQueryReference   queryReference `json:"Successors"`   // HierarchicalRequirementQuery
	TasksQueryReference        queryReference `json:"Tasks"`        // TaskQuery
	TestCasesQueryReference    queryReference `json:"TestCases"`    // TestCaseQuery

	AcceptedDate        string
	Blocked             bool
	BlockedReason       string
	DefectStatus        string
	DirectChildrenCount int
	DragAndDropRank     string
	HasParent           bool
	InProgressDate      string
	PlanEstimate        float64
	Recycled            bool
	TaskActualTotal     float64
	TaskEstimateTotal   float64
	TaskRemainingTotal  float64
	TaskStatus          string
	TestCaseStatus      string
}

type HierarchicalRequirementQuery struct {
	queryResult

	Result []HierarchicalRequirement
}

type Iteration struct {
	WorkspaceDomainObject

	ProjectReference                      reference      `json:"Project"`                 // Project
	RevisionHistoryReference              reference      `json:"RevisionHistory"`         // RevisionHistory
	UserIterationCapacitiesQueryReference queryReference `json:"UserIterationCapacities"` // UserIterationCapacityQuery

	EndDate            string
	Name               string
	Notes              string
	PlanEstimate       float64
	PlannedVelocity    float64
	StartDate          string
	State              string
	TaskActualTotal    float64
	TaskEstimateTotal  float64
	TaskRemainingTotal float64
	Theme              string
}

type Milestone struct {
	WorkspaceDomainObject

	ArtifactsQueryReference  queryReference `json:"Artifacts"`       // ArtifactQuery
	ProjectsQueryReference   queryReference `json:"Projects"`        // ProjectQuery
	RevisionHistoryReference reference      `json:"RevisionHistory"` // RevisionHistory
	TargetProjectReference   reference      `json:"TargetProject"`   // Project

	DisplayColor       string
	FormattedId        string `json:"FormattedID"`
	Name               string
	Notes              string
	TargetDate         string
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

	BuildDefinitionsQueryReference queryReference `json:"BuildDefinitions"` // BuildDefinitionQuery
	ChildrenQueryReference         queryReference `json:"Children"`         // ProjectQuery
	EditorsQueryReference          queryReference `json:"Editors"`          // UserQuery
	IterationsQueryReference       queryReference `json:"Iterations"`       // IterationQuery
	OwnerReference                 reference      `json:"Owner"`            // User
	ParentReference                reference      `json:"Parent"`           // Project
	ReleasesQueryReference         queryReference `json:"Releases"`         // ReleaseQuery
	TeamMembersQueryReference      queryReference `json:"TeamMembers"`      // UserQuery

	Description   string
	Name          string
	Notes         string
	SchemaVersion string
	State         string
}

type ProjectQuery struct {
	queryResult

	Result []Project
}

type Release struct {
	WorkspaceDomainObject

	ProjectReference         reference `json:"Project"`         // Project
	RevisionHistoryReference reference `json:"RevisionHistory"` // RevisionHistory

	Accepted                     float64
	GrossEstimateConversionRatio float64
	Name                         string
	Notes                        string
	PlanEstimate                 float64
	PlannedVelocity              float64
	ReleaseDate                  string
	ReleaseStartDate             string
	State                        string
	TaskActualTotal              float64
	TaskEstimateTotal            float64
	TaskRemainingTotal           float64
	Theme                        string
	Version                      string
}

type Requirement struct {
	SchedulableArtifact

	AttachmentsQueryReference queryReference `json:"Attachments"` // AttachmentQuery

	Package string
}

type Revision struct {
	WorkspaceDomainObject

	RevisionHistoryReference reference `json:"RevisionHistory"` // RevisionHistory
	UserReference            reference `json:"User"`            // User

	Description    string
	RevisionNumber int
}

type RevisionQuery struct {
	queryResult

	Result []Revision
}

type RevisionHistory struct {
	WorkspaceDomainObject

	RevisionsQueryReference queryReference `json:"Revisions"` // RevisionQuery
}

type SCMRepository struct {
	WorkspaceDomainObject

	ProjectsQueryReference queryReference `json:"Projects"` // ProjectQuery

	Description string
	Name        string
	SCMType     string
	Uri         string
}

type SchedulableArtifact struct {
	Artifact

	BlockerReference    reference      `json:"Blocker"`   // Blocker
	IterationReference  reference      `json:"Iteration"` // Iteration
	ReleaseReference    reference      `json:"Release"`   // Release
	TasksQueryReference queryReference `json:"Tasks"`     // TaskQuery

	AcceptedDate       string
	Blocked            bool
	BlockedReason      string
	DragAndDropRank    string
	PlanEstimate       float64
	ScheduleState      string
	TaskActualTotal    float64
	TaskEstimateTotal  float64
	TaskRemainingTotal float64
}

type Subscription struct {
	PersistableObject

	RevisionHistoryReference reference      `json:"RevisionHistory"` // RevisionHistory
	WorkspacesQueryReference queryReference `json:"Workspaces"`      // WorkspaceQuery

	ExpirationDate          string
	MaximumCustomUserFields int
	MaximumProjects         int
	Modules                 string
	Name                    string
	PasswordExpirationDays  int
	PreviousPasswordCount   int
	ProjectHierarchyEnabled bool
	SessionTimeoutSeconds   int
	StoryHierarchyEnabled   bool
	StoryHierarchyType      string
	SubscriptionId          int `json:"SubscriptionID"`
	SubscriptionType        string
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

	AttachmentsQueryReference queryReference `json:"Attachments"` // AttachmentQuery
	IterationReference        reference      `json:"Iteration"`   // Iteration
	ProjectReference          reference      `json:"Project"`     // Project
	ReleaseReference          reference      `json:"Release"`     // Release
	WorkProductReference      reference      `json:"WorkProduct"` // SchedulableArtifact

	Actuals         float64
	Blocked         bool
	BlockedReason   string
	DragAndDropRank string
	Estimate        float64
	Recycled        bool
	State           string
	TaskIndex       int
	TimeSpent       float64
	ToDo            float64
}

type TaskQuery struct {
	queryResult

	Results []Task
}

type TestCase struct {
	Artifact

	AttachmentsQueryReference queryReference `json:"Attachments"` // AttachmentQuery
	DefectsQueryReference     queryReference `json:"Defects"`     // DefectQuery
	ResultsQueryReference     queryReference `json:"Results"`     // TestCaseResultQuery
	StepsQueryReference       queryReference `json:"Steps"`       // TestCaseStepQuery
	WorkProductReference      reference      `json:"WorkProduct"` // Artifact

	DefectStatus             string
	DragAndDropRank          string
	LastBuild                string
	LastRun                  string
	LastVerdict              string
	Method                   string
	Objective                string
	Package                  string
	PostConditions           string
	PreConditions            string
	Priority                 string
	Recycled                 bool
	Risk                     string
	Type                     string
	ValidationExpectedResult string
	ValidationInput          string
}

type TestCaseQuery struct {
	queryResult

	Result []TestCase
}

type TestCaseResult struct {
	WorkspaceDomainObject

	AttachmentsQueryReference queryReference `json:"Attachments"` // AttachmentQuery
	TestCaseReference         reference      `json:"TestCase"`    // TestCase
	TesterReference           reference      `json:"Tester"`      // User

	Build    string
	Date     string
	Duration float64
	Notes    string
	Verdict  string
}

type TestCaseResultQuery struct {
	queryResult

	Result []TestCaseResult
}

type TestCaseStep struct {
	WorkspaceDomainObject

	TestCaseReference reference `json:"TestCase"` // TestCase

	ExpectedResult string
	Input          string
	StepIndex      int
}

type TestCaseStepQuery struct {
	queryResult

	Result []TestCaseStep
}

type TypeDefinition struct {
	WorkspaceDomainObject

	AttributesQueryReference queryReference `json:"Attributes"`      // AttributeDefinitionQuery
	ParentReference          reference      `json:"Parent"`          // TypeDefinition
	RevisionHistoryReference reference      `json:"RevisionHistory"` // RevisionHistory

	Abstract     bool
	Creatable    bool
	Deletable    bool
	DisplayName  string
	ElementName  string
	IdPrefix     string `json:"IDPrefix"`
	Name         string
	Note         string
	Queryable    bool
	ReadOnly     bool
	Restorable   bool
	TypePath     string
	UserListable bool
}

type TypeDefinitionQuery struct {
	queryResult

	Result []TypeDefinition
}

type User struct {
	DomainObject

	RevisionHistoryReference      reference      `json:"RevisionHistory"` // RevisionHistory
	TeamMembershipsQueryReference queryReference `json:"TeamMemberships"` // ProjectQuery
	UserProfileReference          reference      `json:"UserProfile"`     // UserProfile
	UserPermissionsQueryReference queryReference `json:"UserPermissions"` // UserPermissionQuery

	CostCenter             string
	Department             string
	Disabled               bool
	DisplayName            string
	EmailAddress           string
	FirstName              string
	LandingPage            string
	LastLoginDate          string
	LastName               string
	LastPasswordUpdateDate string
	MiddleName             string
	NetworkId              string `json:"NetworkID"`
	OfficeLocation         string
	OnpremLdapUsername     string
	Phone                  string
	Planner                bool
	Role                   string
	ShortDisplayName       string
	SubscriptionAdmin      bool
	SubscriptionPermission string
	UserName               string
	WorkspacePermission    string
}

type UserIterationCapacity struct {
	WorkspaceDomainObject

	IterationReference reference `json:"Iteration"` // Iteration
	ProjectReference   reference `json:"Project"`   // Project
	UserReference      reference `json:"User"`      // User

	Capacity      float64
	Load          float64
	TaskEstimates float64
}

type UserIterationCapacityQuery struct {
	queryResult

	Results []UserIterationCapacity
}

type UserPermission struct {
	DomainObject

	UserReference reference `json:"User"` // User

	CustomObjectId string `json:"CustomObjectID"`
	Name           string
	Role           string
}

type UserPermissionQuery struct {
	queryResult

	Result []UserPermission
}

type UserProfile struct {
	DomainObject

	DefaultProjectReference   reference `json:"DefaultProject"`   // Project
	DefaultWorkspaceReference reference `json:"DefaultWorkspace"` // Workspace

	DateFormat                     string
	DateTimeFormat                 string
	DefaultDetailPageToViewingMode bool
	EmailNotificationEnabled       bool
	SessionTimeoutSeconds          int
	SessionTimeoutWarning          bool
	TimeZone                       string
	WelcomePageHidden              bool
}

type Workspace struct {
	DomainObject

	OwnerReference                  reference      `json:"Owner"`                  // User
	RevisionHistoryReference        reference      `json:"RevisionHistory"`        // RevisionHistory
	ProjectsQueryReference          queryReference `json:"Projects"`               // ProjectQuery
	TypeDefinitionsQueryReference   queryReference `json:"TypeDefinitions"`        // TypeDefinitionQuery
	WorkspaceConfigurationReference reference      `json:"WorkspaceConfiguration"` // WorkspaceConfiguration

	Description   string
	Name          string
	Notes         string
	SchemaVersion string
	State         string
	Style         string
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

	WorkspaceReference reference `json:"Workspace"` // Workspace
}
