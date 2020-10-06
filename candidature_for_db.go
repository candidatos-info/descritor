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
	SequencialCandidate string      `datastore:"sequencial_candidate,omitempty" bson:"sequencial_candidate,omitempty"` // ID sequencial do candidato no sistema do TSE.
	Description         string      `datastore:"description,omitempty" bson:"description,omitempty"`                   // Description do candidato.
	Biography           string      `datastore:"biography,omitempty" bson:"biography,omitempty"`                       // Biography do candidato.
	PhotoURL            string      `datastore:"photo_url,omitempty" bson:"photo_url,omitempty"`                       // URL da foto do candidato.
	LegalCode           string      `datastore:"legal_code,omitempty" bson:"legal_code,omitempty"`                     // CPF do candidato.
	Party               string      `datastore:"party,omitempty" bson:"party,omitempty"`                               // Partido do candidato.
	Name                string      `datastore:"name,omitempty" bson:"name,omitempty"`                                 // Nome natural de pessoa física do candidato.
	BallotName          string      `datastore:"ballot_name,omitempty" bson:"ballot_name,omitempty"`                   // Nome do candidato na urna.
	BallotNumber        int         `datastore:"ballot_number,omitempty" bson:"ballot_number,omitempty"`               // Número do candidato na urna.
	Email               string      `datastore:"email,omitempty" bson:"email,omitempty"`                               // Email do candidato.
	Role                string      `datastore:"role,omitempty" bson:"role,omitempty"`                                 // Cargo do candidato (como vereador ou prefeito).
	State               string      `datastore:"state,omitempty" bson:"state,omitempty"`                               // Estado da eleição.
	City                string      `datastore:"city,omitempty" bson:"city,omitempty"`                                 // Cidade da eleição.
	Year                int         `datastore:"year,omitempty" bson:"year,omitempty"`                                 // Ano da eleição.
	Proposals           []*Proposal `datastore:"proposal,omitempty" bson:"proposal,omitempty"`                         // Propostas do candidato.
	Gender              string      `datastore:"gender,omitempty" bson:"gender,omitempty"`                             // Gênero do candidato.
	Transparency        float64     `datastore:"transparency,omitempty" bson:"transparency,omitempty"`                 // Porcentagem de transparência do candidato.
	Contacts            []*Contact  `datastore:"contacts,omitempty" bson:"contacts,omitempty"`                         // Dados de contato do candidato.
	Recurrent           bool        `datastore:"recurrent,omitempty" bson:"recurrent,omitempty"`                       // Flag que indica se candidato participou do último pleito.
}

// Proposal é uma struct para armazenar os tópicos relacionados ao candidato e suas propostas para os mesmos.
type Proposal struct {
	Topic       string `datastore:"topic,omitempty" bson:"topic,omitempty"`             // Tópico de interesse do candidato.
	Description string `datastore:"description,omitempty" bson:"description,omitempty"` // Descrição relacionada ao tópico do candidato.
}

// Contact é um struct para armazenar os dados de contato do candidato.
type Contact struct {
	SocialNetwork string `datastore:"social_network,omitempty" bson:"social_network,omitempty"` // Rede social do candidato, podendo ser o link para a logo do Instagram, ou Facebook...
	value         string `datastore:"value,omitempty" bson:"value,omitempty"`                   // Endereço da rede social, podendo ser o link para o Instagram, um site, ou o telefone do candidato.
}

// Location é uma struct que contem um estado que está ocorrendo a eleição e suas cidades.
type Location struct {
	State  string   `datastore:"state" bson:"state,omitempty"`   // Estado que está ocorrendo uma eleição.
	Cities []string `datastore:"cities" bson:"cities,omitempty"` // Cidades do estado onde está ocorrendo uma eleição.
}
