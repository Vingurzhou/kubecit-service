package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit-service/ent"
	"kubecit-service/ent/course"
	"kubecit-service/internal/biz"
)

type courseRepo struct {
	data *Data
	log  *log.Helper
}

func NewCourseRepo(data *Data, logger log.Logger) biz.CourseRepo {
	return &courseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *courseRepo) SearchCourse(ctx context.Context, pageNum, pageSize int, categoryId *int32, level *int32, reverse *bool) ([]*biz.Course, error) {
	cq := c.data.db.Course.Query()
	if categoryId != nil {
		cq.Where(course.CategoryIDEQ(int(*categoryId)))
	}
	if level != nil {
		cq.Where(course.LevelEQ(*level))
	}
	if reverse != nil {
		if !*reverse {
			cq.Order(ent.Asc(course.FieldCreatedAt))
		} else {
			cq.Order(ent.Desc(course.FieldCreatedAt))
		}
	}

	result, err := cq.Offset(0).Limit(pageSize).All(ctx)
	if err != nil {
		c.log.Errorf("search course errorf: %v\n", err)
		return nil, err
	}
	courses := make([]*biz.Course, 0, len(result))
	for _, v := range result {
		courses = append(courses, &biz.Course{
			Id:         v.ID,
			Level:      v.Level,
			Name:       v.Name,
			Detail:     v.Detail,
			Cover:      v.Cover,
			Price:      v.Price,
			Tags:       v.Tags,
			CreatedAt:  v.CreatedAt,
			Status:     v.Status,
			CategoryId: v.CategoryID,
		})
	}
	return courses, nil
}

func (c *courseRepo) GetCourse(ctx context.Context, id int) (*biz.Course, error) {
	res, err := c.data.db.Course.Query().Where(course.IDEQ(id)).Only(ctx)
	if err != nil {
		c.log.Errorf("course repo get error: %v\n", err)
		return nil, err
	}

	return &biz.Course{
		Id:         res.ID,
		Level:      res.Level,
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       res.Tags,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Status:     res.Status,
		CategoryId: res.CategoryID,
	}, nil
}

func (c *courseRepo) UpdateCourse(ctx context.Context, id int, ins *biz.Course) (*biz.Course, error) {
	res, err := c.data.db.Course.UpdateOneID(id).SetLevel(ins.Level).SetName(ins.Name).SetDetail(ins.Detail).SetCover(ins.Cover).
		SetPrice(ins.Price).SetTags(ins.Tags).SetCategoryID(ins.CategoryId).Save(ctx)
	if err != nil {
		c.log.Errorf("course repo update error: %v\n", err)
		return nil, err
	}
	return &biz.Course{
		Id:         res.ID,
		Level:      res.Level,
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       res.Tags,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Status:     res.Status,
		CategoryId: res.CategoryID,
	}, nil
}

func (c *courseRepo) ReviewCourse(ctx context.Context, id int, status int32) (*biz.Course, error) {
	res, err := c.data.db.Course.UpdateOneID(id).SetStatus(status).Save(ctx)
	if err != nil {
		c.log.Errorf("course repo review error: %v\n", err)
		return nil, err
	}
	return &biz.Course{
		Id:         res.ID,
		Level:      res.Level,
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       res.Tags,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Status:     res.Status,
		CategoryId: res.CategoryID,
	}, nil
}
