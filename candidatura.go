package descritor

import (
	"encoding/json"
	"time"

	"github.com/gocarina/gocsv"
)

// Cargo representa os possíveis cargos pleiteados pelas candidaturas.
type Cargo string

// Cargos pleiteados via candidaturas
const (
	LegMunicipal  = "LM" // Legislativo municipal (a.k.a vereadora)
	ExecMunicipal = "EM" // Executivo municipal (a.k.a prefeita)
)

// Candidato representa os dados de um candidato
type Candidato struct {
	UF              string    `json:"uf_origem" csv:"uf_origem"`               // Identificador (2 caracteres) da unidade federativa de nascimento do candidato.
	Municipio       string    `json:"municipio_origem" csv:"municipio_origem"` // Município de nascimento do candidato.
	Nascimento      time.Time `json:"nascimento" csv:"nascimento"`             // Data de nascimento do candidato.
	TituloEleitoral string    `json:"titulo_eleitoral" csv:"titulo_eleitoral"` // Titulo eleitoral do candidato.
	Genero          string    `json:"genero" csv:"genero"`                     // Gênero do candidato (MASCULINO ou FEMININO).
	GrauInstrucao   string    `json:"grau_instrucao" csv:"grau_instrucao"`     // Grau de instrução do candidato.
	EstadoCivil     string    `json:"estado_civil" csv:"estado_civil"`         // Estado civil do candidato.
	Raca            string    `json:"raca" csv:"raca"`                         // Raca do candidato (como BRANCA ou PARDA).
	Ocupacao        string    `json:"ocupacao" csv:"ocupacao"`                 // Ocupação do candidato (como COMERCIANTE).
	CPF             string    `json:"cpf" csv:"cpf"`                           // CPF do candidato.
	Nome            string    `json:"nome" csv:"nome"`                         // Nome de pessoa física do candidato.
	Email           string    `json:"email" csv:"email"`                       // Email do candidato.
}

// Candidatura representa uma
type Candidatura struct {
	Legislatura       int    `json:"leg" csv:"leg"`                               // Ano eleitoral em que a candidatura foi homologada.
	Cargo             Cargo  `json:"cargo" csv:"cargo"`                           // Cargo sendo pleiteado pela candidatura.
	UF                string `json:"uf" csv:"uf"`                                 // Identificador (2 caracteres) de unidade federativa onde ocorreu a candidatura.
	Turno             int    `json:"turno" csv:"turno"`                           // Turno do pleito (1 ou 2).
	Municipio         string `json:"municipio" csv:"municipio"`                   // Município que ocorreu a eleição.
	NumeroUrna        int    `json:"numero_urna" csv:"numero_urna"`               // Número do candidato na urna.
	NomeUrna          string `json:"nome_urna" csv:"nome_urna"`                   // Nome do candidato na urna.
	Aptidao           string `json:"aptidao" csv:"aptidao"`                       // Aptidao da candidatura (podendo ser APTO ou INAPTO).
	Situacao          string `json:"situacao" csv:"situacao"`                     // Situação do candidato (pondendo ser DEFERIDO ou INDEFERIDO).
	TipoAgremiacao    string `json:"tipo_agremiacao" csv:"tipo_agremiacao"`       // Indica o tipo de agremiação do candidato (podendo ser PARTIDO ISOLADO ou AGREMIAÇÃO).
	NumeroPartido     int    `json:"numero_partio" csv:"numero_partido"`          // Número do partido do candidato.
	LegendaPartido    string `json:"legenda_partido" csv:"legenda_partido"`       // Legenda do partido do candidato.
	NomePartido       string `json:"nome_partido" csv:"nome_partido"`             // Nome do partido do candidato.
	NomeColigacao     string `json:"nome_coligacao" csv:"nome_coligacao"`         // Nome da coligação a qual o candidato pertence.
	PartidosColigacao string `json:"partidos_coligacao" csv:"partidos_coligacao"` // Partidos pertencentes à coligação do candidato.
	SituacaoFinal     string `json:"situacao_final" csv:"situacao_final"`         // Informa a situação ao fim da eleição (como ELEITO, NÃO ELEITO e SUPLENTE).
	DeclarouBens      bool   `json:"declarou_bens" csv:"declarou_bens"`           // Flag que informa se o candidato declarou seus bens na eleição
	Candidato
}

// ToJSON converte a candidatura para JSON.
func (c *Candidatura) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// ToCSV converte a candidatura para um conjunto de strings
func ToCSV(c []Candidatura) ([]byte, error) {
	b, err := gocsv.MarshalBytes(c)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// FromJSON cria uma instância de candidatura preenchida com os dados contidos no JSON passado como parâmetro.
func FromJSON(b []byte) (*Candidatura, error) {
	c := &Candidatura{}
	if err := json.Unmarshal(b, c); err != nil {
		return c, err
	}
	return c, nil
}
