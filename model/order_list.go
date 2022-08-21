package model

type OrderList []Order

func (ol OrderList) GetIDs() []string {
	values := make([]string, len(ol))
	for i, o := range ol {
		values[i] = o.ID
	}

	return values
}
