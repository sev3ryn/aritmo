package datatype

import "encoding/json"
import "fmt"

type CurrUpdate struct {
	Id     string  `json:"id"`
	Factor float64 `json:"factor"`
}

type CurrService struct {
	updateCh   <-chan []byte
	Currencies map[string]*SimpleDataType
}

type CurrType struct {
	b             *BaseDataType
	convTempFuncs map[tempType]ConvFunc
}

func (c *CurrService) updateRates() {
	var upd CurrUpdate
	var updJSON []byte
	for {
		updJSON = <-c.updateCh
		err := json.Unmarshal(updJSON, &upd)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cur, ok := c.Currencies[upd.Id]
		if ok {
			cur.Factor = upd.Factor
		}
	}
}

func GetCurrUnits(updateCh <-chan []byte) (currList []DataType) {
	c := &CurrService{updateCh: updateCh, Currencies: map[string]*SimpleDataType{
		"usd": &SimpleDataType{
			b: &BaseDataType{
				Group:       GroupLength,
				Names:       []string{"usd", "$"},
				DisplayName: "usd",
			},
			Factor: 1,
		},
		"eur": &SimpleDataType{
			b: &BaseDataType{
				Group:       GroupLength,
				Names:       []string{"eur"},
				DisplayName: "eur",
			},
			Factor: 0.95111,
		},
		"gbp": &SimpleDataType{
			b: &BaseDataType{
				Group:       GroupLength,
				Names:       []string{"gbp"},
				DisplayName: "gbp",
			},
			Factor: 0.81377,
		},
	}}
	go c.updateRates()

	for _, v := range c.Currencies {
		currList = append(currList, v)
	}
	return currList
}
