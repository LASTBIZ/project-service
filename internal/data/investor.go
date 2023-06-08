package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"project-service/internal/biz"
	"strconv"
)

type Investor struct {
	ID       uint64 `gorm:"primaryKey"`
	Money    uint64
	FullName string
	UserID   uint64
	Projects []InvestProject
}

type investorRepo struct {
	data *Data
	log  *log.Helper
}

func NewInvestorRepo(data *Data, logger log.Logger) biz.InvestorRepo {
	return &investorRepo{data: data, log: log.NewHelper(logger)}
}

func (i investorRepo) CreateInvestor(ctx context.Context, investor *biz.Investor) (*biz.Investor, error) {
	var iv Investor
	result := i.data.db.Where(&Investor{UserID: investor.UserID}).Find(&iv)
	if result.RowsAffected == 1 {
		return nil, errors.New(500, "INVESTOR_EXISTS", "exists is by user_id "+strconv.Itoa(int(iv.UserID)))
	}

	iv.Money = 0
	iv.FullName = investor.FullName
	iv.UserID = investor.UserID

	if err := i.data.db.Create(&iv); err.Error != nil {
		return nil, errors.InternalServer("ERROR_CREATE_INVESTOR", "error create investor")
	}

	res := ModelToResponseInvestor(iv)
	return res, nil
}

func (i investorRepo) DeleteInvestor(ctx context.Context, id uint64) (bool, error) {
	var investorInfo Investor
	result := i.data.db.Where(&Investor{ID: id}).First(&investorInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("INVESTOR_NOT_FOUND", "investor not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("INVESTOR_NOT_FOUND", "rows null")
	}

	err := i.data.db.Delete(&investorInfo)
	if err.Error != nil {
		return false, errors.InternalServer("INVESTOR_DELETE_ERROR", "investor delete error")
	}

	return true, nil
}

func (i investorRepo) GetInvestorById(ctx context.Context, id uint64) (*biz.Investor, error) {
	var investorInfo Investor
	if err := i.data.db.Where(&Investor{ID: id}).First(&investorInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("INVESTOR_NOT_FOUND", "investor not found")
		}

		return nil, errors.NotFound("INVESTOR_NOT_FOUND", err.Error())
	}

	re := ModelToResponseInvestor(investorInfo)
	return re, nil
}

func (i investorRepo) ListInvestorByProjectId(ctx context.Context, projectId *uint64, pageNum, pageSize int) ([]*biz.ProjectInvestor, int, error) {
	var investorsInfo []InvestProject
	//result := i.data.db.Joins("JOIN project_investors pi ON pi.investor_id = investors.id").Where("pi.project_id = ?", projectId).Find(&investorsInfo)
	result := i.data.db.Where(&InvestProject{ProjectID: *projectId}).Preload("Investor").Omit("Investor.Money").Find(&investorsInfo)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
		}

		return nil, 0, errors.NotFound("PROJECT_NOT_FOUND", err.Error())
	}
	total := int(result.RowsAffected)
	i.data.db.Scopes(paginate(pageNum, pageSize)).Find(&investorsInfo)
	rv := make([]*biz.ProjectInvestor, 0)
	for _, u := range investorsInfo {
		rv = append(rv, ModelToResponseInvestorProject(u))
	}
	return rv, total, nil
}

func (i investorRepo) AddMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error) {
	var investorInfo Investor
	if err := i.data.db.Where(&Investor{ID: id}).First(&investorInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.NotFound("INVESTOR_NOT_FOUND", "investor not found")
		}

		return false, errors.NotFound("INVESTOR_NOT_FOUND", err.Error())
	}

	investorInfo.Money = investorInfo.Money + uint64(money)
	if err := i.data.db.Save(&investorInfo).Error; err != nil {
		return false, errors.InternalServer("INVESTOR_ADD_MONEY_ERROR", "error add money")
	}

	return true, nil
}

func (i investorRepo) RemoveMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error) {
	var investorInfo Investor
	if err := i.data.db.Where(&Investor{ID: id}).First(&investorInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.NotFound("INVESTOR_NOT_FOUND", "investor not found")
		}

		return false, errors.NotFound("INVESTOR_NOT_FOUND", err.Error())
	}

	investorInfo.Money = investorInfo.Money - uint64(money)
	if err := i.data.db.Save(&investorInfo).Error; err != nil {
		return false, errors.InternalServer("INVESTOR_REMOVE_MONEY_ERROR", "error remove money")
	}

	return true, nil
}

func (i investorRepo) SetMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error) {
	var investorInfo Investor
	if err := i.data.db.Where(&Investor{ID: id}).First(&investorInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.NotFound("INVESTOR_NOT_FOUND", "investor not found")
		}

		return false, errors.NotFound("INVESTOR_NOT_FOUND", err.Error())
	}

	investorInfo.Money = uint64(money)
	if err := i.data.db.Save(&investorInfo).Error; err != nil {
		return false, errors.InternalServer("INVESTOR_SET_MONEY_ERROR", "error set money")
	}

	return true, nil
}

func ModelToResponseInvestorProject(iv InvestProject) *biz.ProjectInvestor {
	ivInfoRsp := &biz.ProjectInvestor{
		ID:    iv.ID,
		Money: iv.Money,
		Investor: biz.Investor{
			FullName: iv.Investor.FullName,
		},

		//ProjectID: iv.ProjectID,
	}

	return ivInfoRsp
}

func ModelToResponseInvestor(iv Investor) *biz.Investor {
	ivInfoRsp := &biz.Investor{
		ID:       iv.ID,
		Money:    iv.Money,
		FullName: iv.FullName,
		UserID:   iv.UserID,
		//ProjectID: iv.ProjectID,
	}

	return ivInfoRsp
}
