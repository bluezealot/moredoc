package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Category struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Cover     string     `form:"cover" json:"cover,omitempty" gorm:"column:cover;comment:分类封面;type:varchar(255);size:255;"` //
	ParentId  int        `form:"parent_id" json:"parent_id,omitempty" gorm:"column:parent_id;type:int(11);size:11;default:0;index:parent_id_title,unique;index:parent_id;comment:上级ID;"`
	Title     string     `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(64);size:64;index:parent_id_title,unique;comment:分类名称;"`
	DocCount  int        `form:"doc_count" json:"doc_count,omitempty" gorm:"column:doc_count;type:int(11);size:11;default:0;comment:文档统计;"`
	Sort      int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;index:idx_sort;comment:排序，值越大越靠前;"`
	Enable    bool       `form:"enable" json:"enable,omitempty" gorm:"column:enable;type:tinyint(1);size:1;index:idx_enable;default:1;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Category) TableName() string {
	return tablePrefix + "category"
}

// CreateCategory 创建Category
func (m *DBModel) CreateCategory(category *Category) (err error) {
	err = m.db.Create(category).Error
	if err != nil {
		m.logger.Error("CreateCategory", zap.Error(err))
		return
	}
	if category.Cover != "" {
		m.setAttachmentType(AttachmentTypeCategoryCover, category.Id, []string{category.Cover})
	}
	return
}

// UpdateCategory 更新Category，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateCategory(category *Category, updateFields ...string) (err error) {
	db := m.db.Model(category)

	updateFields = m.FilterValidFields(Category{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else {
		db = db.Select(m.GetTableFields(Category{}.TableName()))
	}

	err = db.Where("id = ?", category.Id).Updates(category).Error
	if err != nil {
		m.logger.Error("UpdateCategory", zap.Error(err))
	}

	if category.Cover != "" {
		m.setAttachmentType(AttachmentTypeCategoryCover, category.Id, []string{category.Cover})
	}

	return
}

// GetCategory 根据id获取Category
func (m *DBModel) GetCategory(id interface{}, fields ...string) (category Category, err error) {
	db := m.db

	fields = m.FilterValidFields(Category{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryByParentIdTitle(parentId int, title string, fields ...string) 根据唯一索引获取Category
func (m *DBModel) GetCategoryByParentIdTitle(parentId int, title string, fields ...string) (category Category, err error) {
	db := m.db

	fields = m.FilterValidFields(Category{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	db = db.Where("parent_id = ?", parentId)

	db = db.Where("title = ?", title)

	err = db.First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetCategoryByParentIdTitle", zap.Error(err))
		return
	}
	return
}

type OptionGetCategoryList struct {
	Page         int
	Size         int
	WithCount    bool                      // 是否返回总数
	Ids          []interface{}             // id列表
	SelectFields []string                  // 查询字段
	QueryRange   map[string][2]interface{} // map[field][]{min,max}
	QueryIn      map[string][]interface{}  // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{}  // map[field][]{value1,value2,...}
}

// GetCategoryList 获取Category列表
func (m *DBModel) GetCategoryList(opt *OptionGetCategoryList) (categoryList []Category, total int64, err error) {
	db := m.db.Model(&Category{})
	tableName := Category{}.TableName()

	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetCategoryList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Order("parent_id asc, sort desc, id asc").Find(&categoryList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetCategoryList", zap.Error(err))
	}
	return
}

// DeleteCategory 删除数据
// 分类下存在文档的分类，则不允许删除
// 查找子分类，如果子分类下存在文档，则不允许删除
func (m *DBModel) DeleteCategory(ids []int64) (err error) {
	var (
		children    []Category
		childrenIds []int64
	)

	m.db.Select("id").Where("parent_id in (?) and doc_count = ?", ids, 0).Find(&children)
	if len(children) > 0 {
		for _, v := range children {
			childrenIds = append(childrenIds, v.Id)
		}
		if err = m.DeleteCategory(childrenIds); err != nil {
			return
		}
	}

	err = m.db.Where("id in (?) and doc_count = ?", ids, 0).Delete(&Category{}).Error
	if err != nil {
		m.logger.Error("DeleteCategory", zap.Error(err))
	}
	return
}

func (m *DBModel) CountCategory() (count int64, err error) {
	err = m.db.Model(&Category{}).Count(&count).Error
	if err != nil {
		m.logger.Error("CountCategory", zap.Error(err))
	}
	return
}
