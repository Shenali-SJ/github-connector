package functions

import (
	"encoding/json"
	"github-connector/pkg/model"
)

func ConvertStringToJson(requeststring string) (*model.Requestjsonmodel, error) {

	var jsonModel *model.Requestjsonmodel
	err := json.Unmarshal([]byte(requeststring), jsonModel)

	return jsonModel, err
}
