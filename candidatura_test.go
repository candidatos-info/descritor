package descritor

import (
	"testing"

	"github.com/matryer/is"
)

func TestCandidatura(t *testing.T) {
	is := is.New(t)
	cand := Candidatura{Legislatura: 2016, Cargo: ExecMunicipal, UF: "AL"}
	j, err := cand.ToJSON()
	is.NoErr(err)
	is.Equal(`{"leg":2016,"cargo":"EM","uf":"AL"}`, string(j))

}

func TestToCSV(t *testing.T) {
	c := []Candidatura{{Legislatura: 2016, Cargo: ExecMunicipal, UF: "AL"}, {Legislatura: 2016, Cargo: LegMunicipal, UF: "PE"}}
	b, err := ToCSV(c)
	is := is.New(t)
	is.NoErr(err)
	is.Equal("leg,cargo,uf\n2016,EM,AL\n2016,LM,PE\n", string(b))
}
