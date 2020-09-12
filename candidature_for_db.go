package descritor

import "time"

const (
	// CandidaturesCollection é o nome da coleção das candidaturas.
	CandidaturesCollection = "candidatures"
	// LocationsCollection é o nome da coleção de estados e suas cidades
	LocationsCollection = "locations"
	// AccessTokenCollection é o nome da coleção de códigos de acessos
	AccessTokenCollection = "access_tokens"
)

// VotingCity é a struct que encapsula as candidaturas de uma cidade.
// Para cada eleição salvamos as candidaturas agrupadas todas por cidade, ou seja,
// uma eleição no Brasil vai resultar em 5570 inserçōes uma vez que o Brasil possui
// esta quantidade de cidades (https://pt.wikipedia.org/wiki/Lista_de_estados_brasileiros_por_n%C3%BAmero_de_munic%C3%ADpios).
//
// Essa struct é usada pelo projeto resumidor de banco de dados (https://github.com/candidatos-info/resumidores) para fazer a escrita no banco
// e pelo site (https://github.com/candidatos-info/site) para renderizar dados no frontend.
type VotingCity struct {
	Year       int               `datastore:"year,omitempty"`       // Ano de eleição.
	City       string            `datastore:"city,omitempty"`       // Cidade da eleição.
	State      string            `datastore:"state,omitempty"`      // Estado da eleição.
	Candidates []*CandidateForDB `datastore:"candidates,omitempty"` // Lista com os candidatos da cidade.
}

// CandidateForDB é uma struct que contem alguns dos atributos da struct Candidatura.
// Essa struct é usada somente para persistência em banco e para atender requisitos da UI do site.
type CandidateForDB struct {
	SequencialCandidate string `datastore:"sequencial_candidate,omitempty"` // ID sequencial do candidato no sistema do TSE.
	Site                string `datastore:"site,omitempty"`                 // URL do site do candidato.
	Facebook            string `datastore:"facebook,omitempty"`             // Facebook do candidato.
	Twitter             string `datastore:"twitter,omitempty"`              // Twitter do candidato.
	Instagram           string `datastore:"instagram,omitempty"`            // Instagram do candidato.
	Description         string `datastore:"description,omitempty"`          // Description do candidato.
	Biography           string `datastore:"biography,omitempty"`            // Biography do candidato.
	PhotoURL            string `datastore:"photo_url,omitempty"`            // URL da foto do candidato.
	LegalCode           string `datastore:"legal_code,omitempty"`           // CPF do candidato.
	Party               string `datastore:"party,omitempty"`                // Partido do candidato.
	Name                string `datastore:"name,omitempty"`                 // Nome natural de pessoa física do candidato.
	BallotName          string `datastore:"ballot_name,omitempty"`          // Nome do candidato na urna.
	BallotNumber        int    `datastore:"ballot_number,omitempty"`        // Número do candidato na urna.
	Email               string `datastore:"email,omitempty"`                // Email do candidato.
	Role                string `datastore:"role,omitempty"`                 // Cargo do candidato (como vereador ou prefeito).
}

// Location é uma struct que contem um estado que está ocorrendo a eleição e suas cidades.
type Location struct {
	State  string   `datastore:"state"`  // Estado que está ocorrendo uma eleição.
	Cities []string `datastore:"cities"` // Cidades do estado onde está ocorrendo uma eleição.
}

// AccessToken é uma struct para armazenar os códigos de acessos de perfil dos candidatos
type AccessToken struct {
	Code     string    `datastore:"code,omitempty"` // Código de acesso de perfil.
	IssuedAt time.Time `datastore:"issued_at"`      // Hora que token foi criado.
}
