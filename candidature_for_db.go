package descritor

const (
	// CandidaturesCollection é o nome da coleção das candidaturas.
	CandidaturesCollection = "candidatures"
	// LocationsCollection é o nome da coleção de estados e suas cidades
	LocationsCollection = "locations"
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
	SequencialCandidate string   `datastore:"sequencial_candidate,omitempty"` // ID sequencial do candidato no sistema do TSE.
	Description         string   `datastore:"description,omitempty"`          // Description do candidato.
	Biography           string   `datastore:"biography,omitempty"`            // Biography do candidato.
	PhotoURL            string   `datastore:"photo_url,omitempty"`            // URL da foto do candidato.
	LegalCode           string   `datastore:"legal_code,omitempty"`           // CPF do candidato.
	Party               string   `datastore:"party,omitempty"`                // Partido do candidato.
	Name                string   `datastore:"name,omitempty"`                 // Nome natural de pessoa física do candidato.
	BallotName          string   `datastore:"ballot_name,omitempty"`          // Nome do candidato na urna.
	BallotNumber        int      `datastore:"ballot_number,omitempty"`        // Número do candidato na urna.
	Email               string   `datastore:"email,omitempty"`                // Email do candidato.
	Role                string   `datastore:"role,omitempty"`                 // Cargo do candidato (como vereador ou prefeito).
	State               string   `datastore:"state,omitempty"`                // Estado da eleição.
	City                string   `datastore:"city,omitempty"`                 // Cidade da eleição.
	Year                int      `datastore:"year,omitempty"`                 // Ano da eleição.
	Tags                []string `datastore:"tags,omitempty"`                 // Tags do candidato.
	Gender              string   `datastore:"gender,omitempty"`               // Gênero do candidato.
	Transparence        float64  `datastore:"transparence,omitempty"`         // Porcentagem de transparência do candidato.
	Contact             *Contact `datastore:"contact,omitempty"`              // Dados de contato do candidato.
}

// Contact é um struct para armazenar os dados de contato do candidato.
type Contact struct {
	IconURL string `datastore:"icon_url,omitempty"` // Ícone do contato, podendo ser o link para a logo do Instagram, ou Facebook...
	Link    string `datastore:"link,omitempty"`     // Link de contato, podendo ser o link para o Instagram, um site, ou o telefone do candidato.
}

// Location é uma struct que contem um estado que está ocorrendo a eleição e suas cidades.
type Location struct {
	State  string   `datastore:"state"`  // Estado que está ocorrendo uma eleição.
	Cities []string `datastore:"cities"` // Cidades do estado onde está ocorrendo uma eleição.
}
