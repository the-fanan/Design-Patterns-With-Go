package prototype

import (
	"errors"
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

type ShirtsCache struct {}

func (c *ShirtsCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
		case White:
			newItem := *whitePrototype
			return &newItem, nil
		case Black:
			newItem := *blackPrototype
			return &newItem, nil
		case Blue:
			newItem := *bluePrototype
			return &newItem, nil
		default:
			return nil, errors.New("Not implemented yet")
	}
}


func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}


