package apicontracts

type OperatorConfig struct {
	Id         string        `json:"id" bson:"_id"`
	ApiVersion string        `json:"apiVersion"`
	Kind       string        `json:"kind"`
	Spec       *OperatorSpec `json:"spec"`
}

type OperatorSpec struct {
	ImagePostfix string         `json:"imagePostfix"`
	Tasks        []OperatorTask `json:"tasks"`
}

type OperatorTask struct {
	Index   uint   `json:"index"`
	Name    string `json:"name"`
	Version string `json:"version"`
	RunOnce bool   `json:"runOnce"`
}

// OperatorJob defines the config for a task assigned to the ror operator
type OperatorJob struct {
	// +kubebuilder:validation:MinLength=1
	ImageName string `json:"imageName" validate:"required"`
	// +kubebuilder:validation:MinLength=1
	ImageTag string `json:"imageTag" validate:"required"`
	// +optional
	RunOnce bool `json:"runOnce"`

	Version string `json:"version"`

	Cmd string `json:"cmd"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:default=3
	BackOffLimit int32 `json:"backoffLimit" validate:"min=1,max=10"`

	// +kubebuilder:validation:Minimum=30
	// +kubebuilder:validation:Maximum=600
	// +kubebuilder:default=180
	TimeOutInSeconds int64 `json:"timeoutInSeconds" validate:"min=30,max=600"`

	// +optional
	Configs []OperatorJobConfig `json:"configs"`
}

type OperatorJobConfigType string

const (
	OperatorJobConfigTypeEnv  OperatorJobConfigType = "env"
	OperatorJobConfigTypeFile OperatorJobConfigType = "file"
)

type OperatorJobConfig struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name" validate:"required"`

	// +kubebuilder:validation:Required
	Type OperatorJobConfigType `json:"type" validate:"required"`

	// +kubebuilder:validation:MinLength=1
	Path string `json:"path,omitempty"`

	// +kubebuilder:validation:Required
	Data map[string]string `json:"data" validate:"required"`
}

// Spesification of properties for a task
type TaskSpec struct {
	// +kubebuilder:validation:Required
	ImageName string `json:"imageName" validate:"required"`

	// +optional
	Cmd string `json:"cmd"`

	// +optional
	EnvVars []KeyValue `json:"envVars"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:default=3
	BackOffLimit int32 `json:"backoffLimit" validate:"min=1,max=10"`

	// +kubebuilder:validation:Minimum=30
	// +kubebuilder:validation:Maximum=600
	// +kubebuilder:default=180
	TimeOutInSeconds int64 `json:"timeoutInSeconds" validate:"min=30,max=600"`

	// +kubebuilder:validation:Required
	Version string `json:"version" validate:"semver"`

	// +optional
	Secret *TaskSecret `json:"secret"`

	// +kubebuilder:validation:Required
	Scripts *TaskScripts `json:"scripts"`
}

type TaskSourceType string

const (
	Unknown TaskSourceType = ""
	Git     TaskSourceType = "git"
	Vault   TaskSourceType = "vault"
)

type TaskScripts struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	FileNameAndData []FileNameAndData `json:"fileNameAndData"`

	// +kubebuilder:validation:Required
	ScriptDirectory string `json:"scriptDirectory"`
}

type TaskSecret struct {
	// +kubebuilder:validation:Required
	FileNameAndData []FileNameAndData `json:"fileNameAndData"`

	// +kubebuilder:validation:MinLength=1
	Path string `json:"path"`

	// +optional
	GitSource *TaskGitSource `json:"gitSource"`

	// +optional
	VaultSource *TaskVaultSource `json:"vaultSource"`
}

type FileNameAndData struct {
	// +kubebuilder:validation:MinLength=1
	FileName string `json:"filename"`

	// +kubebuilder:validation:Required
	Data string `json:"data"`
}

type TaskGitSource struct {
	// +kubebuilder:validation:Required
	Type TaskSourceType `json:"type"`

	// +kubebuilder:validation:Required
	ContentPath string `json:"contentPath"`

	// +kubebuilder:validation:Required
	GitConfig GitConfig `json:"gitConfig"`
}

type TaskVaultSource struct {
	// +kubebuilder:validation:Required
	Type TaskSourceType `json:"type"`

	// +kubebuilder:validation:Required
	VaultPath string `json:"vaultPath"`
}

// Gitconfig description
type GitConfig struct {
	Token string `json:"token"`

	User string `json:"user"`

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=^[https|git].*
	Repository string `json:"repository"`

	// +kubebuilder:validation:Required
	Branch string `json:"branch"`

	// +kubebuilder:validation:Required
	ProjectId int `json:"projectId"`
}

// Key Value struct
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// needed for kubernetes
func (in *OperatorJob) DeepCopyInto(out *OperatorJob) {
	*out = *in
}

// needed for kubernetes
func (in *OperatorJob) DeepCopy() *OperatorJob {
	if in == nil {
		return nil
	}
	out := new(OperatorJob)
	in.DeepCopyInto(out)
	return out
}
