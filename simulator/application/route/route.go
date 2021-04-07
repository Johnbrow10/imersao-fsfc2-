package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Route represents a request of new delivery request
type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

// Position is a type which contains the lat and long
type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// PartialRoutePosition is the actual response which the system will return
//  Posição parcial da onde o di estar na rota
type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

//  quando esse struct for convertido para json tera ja um padrão nessas variaveis

// NewRoute creates a *Route struct
func NewRoute() *Route {
	return &Route{}
}

// LoadPositions loads from a .txt file all positions (lat and long) to the Position attribute of the struct
func (r *Route) LoadPositions() error {
	// Se a rota nao tiver um ID entao traz um erro
	if r.ID == "" {
		return errors.New("route id not informed")
	}
	//  variavel criada para abrir os arquivos txt dependendo do Id da rota
	f, err := os.Open("destinations/" + r.ID + ".txt")
	// twm um if para verificar se a variavel ta vazia ou não
	if err != nil {
		return err
	}
	// ultilizando o defer ele fecha após as ações de cima ser finalizadas
	defer f.Close()
	//  com o scanner para scanear a lista
	scanner := bufio.NewScanner(f)
	//  com esse for ele percorre toda a lista
	for scanner.Scan() {
		// e faz um splint nas lat e long para separar atraves das virgulas
		data := strings.Split(scanner.Text(), ",")
		// convertendo os dados da lat que era em string para float com tamanho em 64bits e no indice 0
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}
		//  e a mesma coisa da longitude
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		// e as positions sao preenchida com um Append na lista cas variaveis de lat e long
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}
	//  e se naotiver nenhum erro continua a aplicação retronando
	return nil
}

// ExportJsonPositions generates a slice of string in Json using PartialRoutePosition struct
// e exportando as posições para criar um Json
func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	//  para sasber quantas posições existe nessa lista
	total := len(r.Positions)
	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		// percorrer todas as posições  da latitude e longitude
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		//  entao se a ultima posiçaõ do array for igual a k que e o controlador das rotas logo as routes e finalizada
		if total-1 == k {
			route.Finished = true
		}
		// pegar uma struct e converter em JSON
		jsonRoute, err := json.Marshal(route)
		//  se tiver erro por pegar algo vazio
		if err != nil {
			return nil, err
		}
		//  se nao pegar erro ele faz um append com uma lista destring que com cada posição tem um Json nele.
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
