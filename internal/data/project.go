package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"project-service/internal/biz"
	"project-service/internal/utils"
	"time"
)

type Project struct {
	ID uint64 `gorm:"primaryKey"`
	//TODO gormlist string[]
	RoadMapImgURL string
	MainImageUrl  string
	StartDate     time.Time
	EndDate       time.Time
	Location      string
	Name          string
	Description   string `gorm:"type:text"`
	CategoryID    uint32 `gorm:"not null"`
	CurrentBudget uint64
	NeedBudget    uint64
	Category      Category
	Investors     []InvestProject
	Roadmaps      []RoadMap
}

type InvestProject struct {
	ID         uint64 `gorm:"primaryKey"`
	ProjectID  uint64
	InvestorID uint64
	Investor   Investor
	Money      uint64
}

type projectRepo struct {
	data *Data
	log  *log.Helper
}

func NewProjectRepo(data *Data, logger log.Logger) biz.ProjectRepo {
	return &projectRepo{data: data, log: log.NewHelper(logger)}
}

func (p projectRepo) CreateProject(ctx context.Context, project *biz.Project) (*biz.Project, error) {
	pr := Project{}
	pr.RoadMapImgURL = project.RoadMapImgURL
	pr.MainImageUrl = project.MainImageUrl
	pr.StartDate = project.StartDate
	pr.EndDate = project.EndDate
	pr.Location = project.Location
	pr.Name = project.Name
	pr.Description = project.Description
	pr.CategoryID = project.CategoryID
	pr.NeedBudget = project.NeedBudget
	pr.CurrentBudget = 0
	res := p.data.db.Create(&pr)
	if res.Error != nil {
		return nil, errors.InternalServer("CREATE_PROJECT_ERROR", "error create project")
	}

	projectInfoRes := p.modelToResponse(pr)
	return projectInfoRes, nil
}

func (p projectRepo) DeleteProject(ctx context.Context, id uint64) (bool, error) {
	var projectInfo Project
	result := p.data.db.Where(&Project{ID: id}).First(&projectInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("PROJECT_NOT_FOUND", "rows null")
	}

	err := p.data.db.Delete(&projectInfo)
	if err.Error != nil {
		return false, errors.InternalServer("PROJECT_DELETE_ERROR", "project delete error")
	}

	return true, nil
}

func (p projectRepo) UpdateProject(ctx context.Context, project *biz.Project) (bool, error) {
	var projectInfo Project
	result := p.data.db.Where(&Project{ID: project.ID}).First(&projectInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("PROJECT_NOT_FOUND", "rows null")
	}

	projectInfo.RoadMapImgURL = project.RoadMapImgURL
	projectInfo.MainImageUrl = project.MainImageUrl
	projectInfo.StartDate = project.StartDate
	projectInfo.EndDate = project.EndDate
	projectInfo.Location = project.Location
	projectInfo.Name = project.Name
	projectInfo.Description = project.Description
	projectInfo.NeedBudget = project.NeedBudget
	projectInfo.CategoryID = project.CategoryID

	if err := p.data.db.Save(&projectInfo).Error; err != nil {
		return false, errors.InternalServer("PROJECT_UPDATE_ERROR", "error update project")
	}

	return true, nil
}

func (p projectRepo) GetProjectById(ctx context.Context, id uint64) (*biz.Project, error) {
	var projectInfo Project
	if err := p.data.db.Where(&Project{ID: id}).First(&projectInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
		}

		return nil, errors.NotFound("PROJECT_NOT_FOUND", err.Error())
	}

	re := p.modelToResponse(projectInfo)
	return re, nil
}

func (p projectRepo) GetProjectByCategoryID(ctx context.Context, categoryID uint32, keyWords string, pageNum, pageSize int) ([]*biz.Project, int, error) {
	var projectsInfo []Project
	localDB := p.data.db.Model(&Project{})
	if keyWords != "" {
		//queryMap["name"] = "%" + req.KeyWords + "%"
		localDB = localDB.Where("name LIKE ?", "%"+keyWords+"%")
	}

	//var subQuery string
	if categoryID > 0 {
		var category Category
		if result := p.data.db.First(&category, categoryID); result.RowsAffected == 0 {
			return nil, 0, errors.NotFound("CATEGORY_NOT_FOUND", "error not found category")
		}

		//if category.Level == 1 {
		//	subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", categoryID)
		//} else if category.Level == 2 {
		//	subQuery = fmt.Sprintf("select id from category WHERE parent_category_id=%d", categoryID)
		//} else if category.Level == 3 {
		//	subQuery = fmt.Sprintf("select id from category WHERE id=%d", categoryID)
		//}
		localDB = localDB.Where(&Project{CategoryID: categoryID})
	}
	var count int64
	localDB.Count(&count)

	result := localDB.Preload("Category").Scopes(paginate(pageNum, pageSize)).Find(&projectsInfo)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
		}

		return nil, 0, errors.NotFound("PROJECT_NOT_FOUND", err.Error())
	}

	//p.data.db.Scopes(paginate(pageNum, pageSize)).Find(&projectsInfo)
	rv := make([]*biz.Project, 0)
	for _, u := range projectsInfo {
		rv = append(rv, p.modelToResponse(u))
	}
	return rv, int(count), nil
}

func (p projectRepo) InvestProject(ctx context.Context, id uint64, investorID uint64, money int) error {
	var projectInfo Project
	if err := p.data.db.Where(&Project{ID: id}).First(&projectInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFound("PROJECT_NOT_FOUND", "project not found")
		}
		return errors.NotFound("PROJECT_NOT_FOUND", err.Error())
	}

	var investor Investor
	if err := p.data.db.Where(&Investor{UserID: investorID}).Preload("InvestProject").First(&investor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFound("INVESTOR_NOT_FOUND", "investor not found")
		}
		return errors.NotFound("INVESTOR_NOT_FOUND", err.Error())
	}

	if investor.Money < uint64(money) {
		return errors.Conflict("NOT_ENOUGH_MONEY", "not enough money")
	}

	ids := make([]int, 0)
	for _, u := range investor.Projects {
		ids = append(ids, int(u.ID))
	}
	tx := p.data.db.Begin()

	var investProject InvestProject
	if err := tx.Where(&InvestProject{ProjectID: projectInfo.ID, InvestorID: investorID}).First(&investProject).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			investProject.ProjectID = projectInfo.ID
			investProject.InvestorID = investorID
			investProject.Money = uint64(money)
			//todo maybe error catch and rollaback
			tx.Create(&investProject)
		} else {
			tx.Rollback()
			return errors.NotFound("ERROR_INVEST_PROJECT", err.Error())
		}
	}

	investProject.Money += uint64(money)
	if err := tx.Save(&investProject).Error; err != nil {
		tx.Rollback()
		return errors.InternalServer("ERROR_INVEST_PROJECT", "error invest project")
	}

	if !utils.Contains(ids, int(id)) {
		investor.Projects = append(investor.Projects, investProject)
	}
	investor.Money -= uint64(money)
	if err := tx.Save(&investor).Error; err != nil {
		tx.Rollback()
		return errors.InternalServer("ERROR_INVEST_PROJECT", "error invest project")
	}
	projectInfo.CurrentBudget += uint64(money)
	if err := tx.Save(&projectInfo).Error; err != nil {
		tx.Rollback()
		return errors.InternalServer("ERROR_INVEST_PROJECT", "error invest project")
	}

	return tx.Commit().Error
}

func (p projectRepo) modelToResponse(project Project) *biz.Project {
	projectInfoRsp := &biz.Project{
		ID:            project.ID,
		RoadMapImgURL: project.RoadMapImgURL,
		MainImageUrl:  project.MainImageUrl,
		StartDate:     project.StartDate,
		EndDate:       project.EndDate,
		Location:      project.Location,
		Name:          project.Name,
		Description:   project.Description,
		CategoryID:    project.CategoryID,
		Category: biz.Category{
			ID:   project.Category.ID,
			Name: project.Category.Name,
		},
	}

	return projectInfoRsp
}
