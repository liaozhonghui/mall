package repo

import (
	"context"
	"mall/internal/entity"

	"gorm.io/gorm"
)

type GoodsRepository interface {
	WithDBInstance(*gorm.DB) GoodsRepository
	CreateGoods(ctx context.Context, goods entity.GoodsInfo) (int, error)
	UpdateGoods(ctx context.Context, goods entity.GoodsInfo) (int, error)
	DeleteGoods(ctx context.Context, goods entity.GoodsInfo) (int, error)
	FindGoodsListByCategoryId(ctx context.Context, cateId int) ([]entity.GoodsInfo, error)
	GetGoodsDetailById(ctx context.Context, goodsId int) (entity.GoodsInfo, error)
	SelectGoodsBySkuId(ctx context.Context, skuId int) (goodSku entity.GoodsSkuInfo, err error)
	SelectGoodsBySkuIdForUpdate(ctx context.Context, skuId int) (goodSku entity.GoodsSkuInfo, err error)
	UpdateGoodsSkuStore(ctx context.Context, sku entity.GoodsSkuInfo) (affected int, err error)
	CheckSkuLeftStore(ctx context.Context, skuId int) bool
	IncrSkuLeftStore(ctx context.Context, skuId int)
}
