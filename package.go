package gomposer

// TODO reanme
type PackageRepository struct {
	client *HttpClient
}

func (r *PackageRepository) Find(packageName string) (*PackageInfo, error) {

	output := &PackageInfo{}

	err := r.client.Request("GET", "/phpunit.json", output)

	return output, err
}

type PackageInfo struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Time        string       `json:"time"`
	Maintainers []Maintainer `json:"maintainers"`
	Versions    []Version    `json:"versions"`
	Type        string       `json:"type"`
	Repository  string       `json:"repository"`
	Downloads   Downloads    `json:"downloads"`
	Favers      int          `json:"favers"`
}

type Maintainer struct {
	Name string `json:"name"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type Source struct {
	Type      string `json:"type"`
	Url       string `json:"url"`
	Reference string `json:"reference"`
	Shasum    string `json:"shasum"`
}

type Downloads struct {
	Total   int `json:"total"`
	Monthly int `json:"monthly"`
	Daily   int `json:daily`
}

type Autoload struct {
	ClassMap []string          `json:"classmap"`
	Files    []string          `json:"files"`
	Psr0     map[string]string `json:"psr-0"`
	Psr4     map[string]string `json:"psr-4"`
}

// TODO rename since this going to be the name composer.json
type Version struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Keywords          []string          `json:"keywords"`
	Homepage          string            `json:"homepage"`
	Version           string            `json:"version"`
	VersionNormalized string            `json:version_normalized`
	License           []string          `json:"license"`
	Authors           Author            `json:"authors"`
	Source            Source            `json:"source"`
	Dist              Source            `json:"dist"`
	Type              string            `json:"string"`
	Time              string            `json:"time"`
	Autoload          Autoload          `json:"autoload"`
	Require           map[string]string `json:"require"`
	RequireDev        map[string]string `json:"require-dev"`
	Suggest           map[string]string `json:"suggest"`
}
