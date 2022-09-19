//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import model "github.com/ybalexdp/go-generate-sample/domain/model/item"

type ItemGetter interface {
	Get(id int) (item model.ItemModel, err error)
}
