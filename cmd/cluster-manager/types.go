package main

type EnvironmentContext struct {
	BasePath string
	RepoPath string
	RepoUrl  string
}

type ClusterConfigFile struct {
	Cluster               ClusterConfig           `yaml:"cluster"`
	HelmApplications      []*HelmApplication      `yaml:"helmApplications"`
	KustomizeApplications []*KustomizeApplication `yaml:"kustomizeApplications"`
}

type ClusterConfig struct {
	Name          string            `yaml:"name"`
	Server        string            `yaml:"server"`
	AutoSync      *bool             `yaml:"autoSync"`
	ProjectRoles  []*ProjectRole    `yaml:"projectRoles"`
	CascadeDelete *bool             `yaml:"cascadeDelete"`
	RepoUrl       *string           `yaml:"repoURL"`
	Settings      map[string]string `yaml:"settings"`
}

type Application struct {
	Include        *string `yaml:"include"`
	Name           *string `yaml:"name"`
	RepoUrl        *string `yaml:"repoURL"`
	Path           string  `yaml:"path"`
	AutoSync       *bool   `yaml:"autoSync"`
	CascadeDelete  *bool   `yaml:"cascadeDelete"`
	TargetRevision *string `yaml:"targetRevision"`
	Namespace      *string `yaml:"namespace"`
}

type HelmAddon struct {
	Application            `yaml:",inline"`
	ReleaseName            *string                                `yaml:"releaseName"`
	Parameters             map[string]string                      `yaml:"parameters"`
	Settings               map[string]string                      `yaml:"settings"`
	ValueFiles             []string                               `yaml:"valueFiles"`
	Values                 map[interface{}]interface{}            `yaml:"values"`
	Oauth2ProxyIngressHost *string                                `yaml:"oauth2ProxyIngressHost"`
	OverlayDefinitions     map[string]map[interface{}]interface{} `yaml:"overlayDefinitions"`
}

type HelmApplication struct {
	HelmAddon `yaml:",inline"`
	Addon     *string  `yaml:"addon"`
	Overlays  []string `yaml:"overlays"`
}

type KustomizeApplication struct {
	Application `yaml:",inline"`
}

type ProjectRole struct {
	Name        string
	Description string
	Policies    []string
	JwtTokens   []string
}

type ProjectViewModel struct {
	Name         string
	Server       string
	ProjectRoles []ProjectRole
}

type ApplicationViewModel struct {
	Name           string
	Project        string
	CascadeDelete  bool
	RepoUrl        string
	Path           string
	AutoSync       bool
	Server         string
	TargetRevision string

	// helm specific
	Values                 string
	ValueFiles             []string
	ReleaseName            string
	Parameters             map[string]string
	Namespace              string
	OAuth2ProxyIngressHost string
}

type Oauth2ProxyIngress struct {
	Name       string
	Namespace  string
	SecretName string
	Host       string
}

type ObjectsGeneratorViewModel struct {
	Namespaces           []string
	Oauth2ProxyIngresses []Oauth2ProxyIngress
}
