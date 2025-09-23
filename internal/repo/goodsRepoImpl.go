package repo

import (
	"context"
	"mall/internal/dao/db"
	"mall/internal/entity"

	"gorm.io/gorm"
)

type GoodsRepositoryImpl struct {
	goodsDao   db.GoodsDbDao
	goodskuDao db.GoodsSkuDbDao
}

func NewGoodsRepository() GoodsRepository {
	return &GoodsRepositoryImpl{
		goodsDao:   db.NewGoodsDbDao(),
		goodskuDao: db.NewGoodsSkuDbDao(),
	}
}

func (repo *GoodsRepositoryImpl) WithDBInstance(dbInstance *gorm.DB) GoodsRepository {
	repo.goodsDao.WithDBInstance(dbInstance)
	repo.goodskuDao.WithDBInstance(dbInstance)
	return repo
}

func (repo *GoodsRepositoryImpl) CreateGoods(ctx context.Context, goods entity.GoodsInfo) (goodsId int, err error) {
	goodsId, err = repo.goodsDao.CreateGoods(ctx, goods)
	if err != nil {
		return 0, err
	}

	for idx, _ := range goods.SkuInfo {
		goods.SkuInfo[idx].GoodsId = goodsId
	}

	_, err = repo.goodskuDao.CreateGoodsSku(ctx, goods.SkuInfo)
	if err != nil {
		return 0, err
	}
	return
}

func (repo *GoodsRepositoryImpl) UpdateGoods(ctx context.Context, goods entity.GoodsInfo) (int, error) {
	return repo.goodsDao.UpdateGoods(ctx, goods)
}

func (repo *GoodsRepositoryImpl) DeleteGoods(ctx context.Context, goods entity.GoodsInfo) (int, error) {
	return repo.goodsDao.DeleteGoods(ctx, goods)
}

func (repo *GoodsRepositoryImpl) FindGoodsListByCategoryId(ctx context.Context, cateId int) ([]entity.GoodsInfo, error) {
	return repo.goodsDao.FindGoodsListByCategoryId(ctx, cateId)
}

func (repo *GoodsRepositoryImpl) GetGoodsDetailById(ctx context.Context, goodsId int) (entity.GoodsInfo, error) {
	goodsInfo, err := repo.goodsDao.GetGoodsDetailById(ctx, goodsId)
	if err != nil {
		return entity.GoodsInfo{}, err
	}
	skuList, err := repo.goodskuDao.FindGoodsSkuByGoodId(ctx, goodsId)
	if err != nil {
		return entity.GoodsInfo{}, err
	}
	goodsInfo.SkuInfo = skuList
	return goodsInfo, nil
}

func (repo *GoodsRepositoryImpl) SelectGoodsBySkuId(ctx context.Context, skuId int) (goodSku entity.GoodsSkuInfo, err error) {
	return repo.goodskuDao.SelectGoodsBySkuId(ctx, skuId)
}

func (repo *GoodsRepositoryImpl) SelectGoodsBySkuIdForUpdate(ctx context.Context, skuId int) (goodSku entity.GoodsSkuInfo, err error) {
	return repo.goodskuDao.SelectGoodsBySkuIdForUpdate(ctx, skuId)
}

func (repo *GoodsRepositoryImpl) UpdateGoodsSkuStore(ctx context.Context, sku entity.GoodsSkuInfo) (affected int, err error) {
	return repo.goodskuDao.UpdateGoodsSkuStore(ctx, sku)
}

func (repo *GoodsRepositoryImpl) CheckSkuLeftStore(ctx context.Context, skuId int) bool {
	return repo.goodskuDao.CheckSkuLeftStore(ctx, skuId)
}

func (repo *GoodsRepositoryImpl) IncrSkuLeftStore(ctx context.Context, skuId int) {
	repo.goodskuDao.IncrSkuLeftStore(ctx, skuId)
}
