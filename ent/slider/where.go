// Code generated by ent, DO NOT EDIT.

package slider

import (
	"kubecit-service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Slider {
	return predicate.Slider(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Slider {
	return predicate.Slider(sql.FieldContainsFold(FieldID, id))
}

// CreateBy applies equality check predicate on the "createBy" field. It's identical to CreateByEQ.
func CreateBy(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldCreateBy, v))
}

// ImageName applies equality check predicate on the "imageName" field. It's identical to ImageNameEQ.
func ImageName(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldImageName, v))
}

// CreateTime applies equality check predicate on the "createTime" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateBy applies equality check predicate on the "updateBy" field. It's identical to UpdateByEQ.
func UpdateBy(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldUpdateBy, v))
}

// ImageRemark applies equality check predicate on the "imageRemark" field. It's identical to ImageRemarkEQ.
func ImageRemark(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldImageRemark, v))
}

// ImageUrl applies equality check predicate on the "imageUrl" field. It's identical to ImageUrlEQ.
func ImageUrl(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldImageUrl, v))
}

// PcHref applies equality check predicate on the "pcHref" field. It's identical to PcHrefEQ.
func PcHref(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldPcHref, v))
}

// UpdateTime applies equality check predicate on the "updateTime" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldUpdateTime, v))
}

// AppHref applies equality check predicate on the "appHref" field. It's identical to AppHrefEQ.
func AppHref(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldAppHref, v))
}

// CreateByEQ applies the EQ predicate on the "createBy" field.
func CreateByEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldCreateBy, v))
}

// CreateByNEQ applies the NEQ predicate on the "createBy" field.
func CreateByNEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldCreateBy, v))
}

// CreateByIn applies the In predicate on the "createBy" field.
func CreateByIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldCreateBy, vs...))
}

// CreateByNotIn applies the NotIn predicate on the "createBy" field.
func CreateByNotIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldCreateBy, vs...))
}

// CreateByGT applies the GT predicate on the "createBy" field.
func CreateByGT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldCreateBy, v))
}

// CreateByGTE applies the GTE predicate on the "createBy" field.
func CreateByGTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldCreateBy, v))
}

// CreateByLT applies the LT predicate on the "createBy" field.
func CreateByLT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldCreateBy, v))
}

// CreateByLTE applies the LTE predicate on the "createBy" field.
func CreateByLTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldCreateBy, v))
}

// CreateByContains applies the Contains predicate on the "createBy" field.
func CreateByContains(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContains(FieldCreateBy, v))
}

// CreateByHasPrefix applies the HasPrefix predicate on the "createBy" field.
func CreateByHasPrefix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasPrefix(FieldCreateBy, v))
}

// CreateByHasSuffix applies the HasSuffix predicate on the "createBy" field.
func CreateByHasSuffix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasSuffix(FieldCreateBy, v))
}

// CreateByEqualFold applies the EqualFold predicate on the "createBy" field.
func CreateByEqualFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEqualFold(FieldCreateBy, v))
}

// CreateByContainsFold applies the ContainsFold predicate on the "createBy" field.
func CreateByContainsFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContainsFold(FieldCreateBy, v))
}

// ImageNameEQ applies the EQ predicate on the "imageName" field.
func ImageNameEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldImageName, v))
}

// ImageNameNEQ applies the NEQ predicate on the "imageName" field.
func ImageNameNEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldImageName, v))
}

// ImageNameIn applies the In predicate on the "imageName" field.
func ImageNameIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldImageName, vs...))
}

// ImageNameNotIn applies the NotIn predicate on the "imageName" field.
func ImageNameNotIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldImageName, vs...))
}

// ImageNameGT applies the GT predicate on the "imageName" field.
func ImageNameGT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldImageName, v))
}

// ImageNameGTE applies the GTE predicate on the "imageName" field.
func ImageNameGTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldImageName, v))
}

// ImageNameLT applies the LT predicate on the "imageName" field.
func ImageNameLT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldImageName, v))
}

// ImageNameLTE applies the LTE predicate on the "imageName" field.
func ImageNameLTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldImageName, v))
}

// ImageNameContains applies the Contains predicate on the "imageName" field.
func ImageNameContains(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContains(FieldImageName, v))
}

// ImageNameHasPrefix applies the HasPrefix predicate on the "imageName" field.
func ImageNameHasPrefix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasPrefix(FieldImageName, v))
}

// ImageNameHasSuffix applies the HasSuffix predicate on the "imageName" field.
func ImageNameHasSuffix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasSuffix(FieldImageName, v))
}

// ImageNameEqualFold applies the EqualFold predicate on the "imageName" field.
func ImageNameEqualFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEqualFold(FieldImageName, v))
}

// ImageNameContainsFold applies the ContainsFold predicate on the "imageName" field.
func ImageNameContainsFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContainsFold(FieldImageName, v))
}

// CreateTimeEQ applies the EQ predicate on the "createTime" field.
func CreateTimeEQ(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "createTime" field.
func CreateTimeNEQ(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "createTime" field.
func CreateTimeIn(vs ...time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "createTime" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "createTime" field.
func CreateTimeGT(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "createTime" field.
func CreateTimeGTE(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "createTime" field.
func CreateTimeLT(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "createTime" field.
func CreateTimeLTE(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateByEQ applies the EQ predicate on the "updateBy" field.
func UpdateByEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldUpdateBy, v))
}

// UpdateByNEQ applies the NEQ predicate on the "updateBy" field.
func UpdateByNEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldUpdateBy, v))
}

// UpdateByIn applies the In predicate on the "updateBy" field.
func UpdateByIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldUpdateBy, vs...))
}

// UpdateByNotIn applies the NotIn predicate on the "updateBy" field.
func UpdateByNotIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldUpdateBy, vs...))
}

// UpdateByGT applies the GT predicate on the "updateBy" field.
func UpdateByGT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldUpdateBy, v))
}

// UpdateByGTE applies the GTE predicate on the "updateBy" field.
func UpdateByGTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldUpdateBy, v))
}

// UpdateByLT applies the LT predicate on the "updateBy" field.
func UpdateByLT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldUpdateBy, v))
}

// UpdateByLTE applies the LTE predicate on the "updateBy" field.
func UpdateByLTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldUpdateBy, v))
}

// UpdateByContains applies the Contains predicate on the "updateBy" field.
func UpdateByContains(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContains(FieldUpdateBy, v))
}

// UpdateByHasPrefix applies the HasPrefix predicate on the "updateBy" field.
func UpdateByHasPrefix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasPrefix(FieldUpdateBy, v))
}

// UpdateByHasSuffix applies the HasSuffix predicate on the "updateBy" field.
func UpdateByHasSuffix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasSuffix(FieldUpdateBy, v))
}

// UpdateByEqualFold applies the EqualFold predicate on the "updateBy" field.
func UpdateByEqualFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEqualFold(FieldUpdateBy, v))
}

// UpdateByContainsFold applies the ContainsFold predicate on the "updateBy" field.
func UpdateByContainsFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContainsFold(FieldUpdateBy, v))
}

// ImageRemarkEQ applies the EQ predicate on the "imageRemark" field.
func ImageRemarkEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldImageRemark, v))
}

// ImageRemarkNEQ applies the NEQ predicate on the "imageRemark" field.
func ImageRemarkNEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldImageRemark, v))
}

// ImageRemarkIn applies the In predicate on the "imageRemark" field.
func ImageRemarkIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldImageRemark, vs...))
}

// ImageRemarkNotIn applies the NotIn predicate on the "imageRemark" field.
func ImageRemarkNotIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldImageRemark, vs...))
}

// ImageRemarkGT applies the GT predicate on the "imageRemark" field.
func ImageRemarkGT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldImageRemark, v))
}

// ImageRemarkGTE applies the GTE predicate on the "imageRemark" field.
func ImageRemarkGTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldImageRemark, v))
}

// ImageRemarkLT applies the LT predicate on the "imageRemark" field.
func ImageRemarkLT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldImageRemark, v))
}

// ImageRemarkLTE applies the LTE predicate on the "imageRemark" field.
func ImageRemarkLTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldImageRemark, v))
}

// ImageRemarkContains applies the Contains predicate on the "imageRemark" field.
func ImageRemarkContains(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContains(FieldImageRemark, v))
}

// ImageRemarkHasPrefix applies the HasPrefix predicate on the "imageRemark" field.
func ImageRemarkHasPrefix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasPrefix(FieldImageRemark, v))
}

// ImageRemarkHasSuffix applies the HasSuffix predicate on the "imageRemark" field.
func ImageRemarkHasSuffix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasSuffix(FieldImageRemark, v))
}

// ImageRemarkEqualFold applies the EqualFold predicate on the "imageRemark" field.
func ImageRemarkEqualFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEqualFold(FieldImageRemark, v))
}

// ImageRemarkContainsFold applies the ContainsFold predicate on the "imageRemark" field.
func ImageRemarkContainsFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContainsFold(FieldImageRemark, v))
}

// ImageUrlEQ applies the EQ predicate on the "imageUrl" field.
func ImageUrlEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldImageUrl, v))
}

// ImageUrlNEQ applies the NEQ predicate on the "imageUrl" field.
func ImageUrlNEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldImageUrl, v))
}

// ImageUrlIn applies the In predicate on the "imageUrl" field.
func ImageUrlIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldImageUrl, vs...))
}

// ImageUrlNotIn applies the NotIn predicate on the "imageUrl" field.
func ImageUrlNotIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldImageUrl, vs...))
}

// ImageUrlGT applies the GT predicate on the "imageUrl" field.
func ImageUrlGT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldImageUrl, v))
}

// ImageUrlGTE applies the GTE predicate on the "imageUrl" field.
func ImageUrlGTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldImageUrl, v))
}

// ImageUrlLT applies the LT predicate on the "imageUrl" field.
func ImageUrlLT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldImageUrl, v))
}

// ImageUrlLTE applies the LTE predicate on the "imageUrl" field.
func ImageUrlLTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldImageUrl, v))
}

// ImageUrlContains applies the Contains predicate on the "imageUrl" field.
func ImageUrlContains(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContains(FieldImageUrl, v))
}

// ImageUrlHasPrefix applies the HasPrefix predicate on the "imageUrl" field.
func ImageUrlHasPrefix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasPrefix(FieldImageUrl, v))
}

// ImageUrlHasSuffix applies the HasSuffix predicate on the "imageUrl" field.
func ImageUrlHasSuffix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasSuffix(FieldImageUrl, v))
}

// ImageUrlEqualFold applies the EqualFold predicate on the "imageUrl" field.
func ImageUrlEqualFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEqualFold(FieldImageUrl, v))
}

// ImageUrlContainsFold applies the ContainsFold predicate on the "imageUrl" field.
func ImageUrlContainsFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContainsFold(FieldImageUrl, v))
}

// PcHrefEQ applies the EQ predicate on the "pcHref" field.
func PcHrefEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldPcHref, v))
}

// PcHrefNEQ applies the NEQ predicate on the "pcHref" field.
func PcHrefNEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldPcHref, v))
}

// PcHrefIn applies the In predicate on the "pcHref" field.
func PcHrefIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldPcHref, vs...))
}

// PcHrefNotIn applies the NotIn predicate on the "pcHref" field.
func PcHrefNotIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldPcHref, vs...))
}

// PcHrefGT applies the GT predicate on the "pcHref" field.
func PcHrefGT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldPcHref, v))
}

// PcHrefGTE applies the GTE predicate on the "pcHref" field.
func PcHrefGTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldPcHref, v))
}

// PcHrefLT applies the LT predicate on the "pcHref" field.
func PcHrefLT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldPcHref, v))
}

// PcHrefLTE applies the LTE predicate on the "pcHref" field.
func PcHrefLTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldPcHref, v))
}

// PcHrefContains applies the Contains predicate on the "pcHref" field.
func PcHrefContains(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContains(FieldPcHref, v))
}

// PcHrefHasPrefix applies the HasPrefix predicate on the "pcHref" field.
func PcHrefHasPrefix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasPrefix(FieldPcHref, v))
}

// PcHrefHasSuffix applies the HasSuffix predicate on the "pcHref" field.
func PcHrefHasSuffix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasSuffix(FieldPcHref, v))
}

// PcHrefEqualFold applies the EqualFold predicate on the "pcHref" field.
func PcHrefEqualFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEqualFold(FieldPcHref, v))
}

// PcHrefContainsFold applies the ContainsFold predicate on the "pcHref" field.
func PcHrefContainsFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContainsFold(FieldPcHref, v))
}

// UpdateTimeEQ applies the EQ predicate on the "updateTime" field.
func UpdateTimeEQ(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "updateTime" field.
func UpdateTimeNEQ(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "updateTime" field.
func UpdateTimeIn(vs ...time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "updateTime" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "updateTime" field.
func UpdateTimeGT(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "updateTime" field.
func UpdateTimeGTE(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "updateTime" field.
func UpdateTimeLT(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "updateTime" field.
func UpdateTimeLTE(v time.Time) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldUpdateTime, v))
}

// AppHrefEQ applies the EQ predicate on the "appHref" field.
func AppHrefEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEQ(FieldAppHref, v))
}

// AppHrefNEQ applies the NEQ predicate on the "appHref" field.
func AppHrefNEQ(v string) predicate.Slider {
	return predicate.Slider(sql.FieldNEQ(FieldAppHref, v))
}

// AppHrefIn applies the In predicate on the "appHref" field.
func AppHrefIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldIn(FieldAppHref, vs...))
}

// AppHrefNotIn applies the NotIn predicate on the "appHref" field.
func AppHrefNotIn(vs ...string) predicate.Slider {
	return predicate.Slider(sql.FieldNotIn(FieldAppHref, vs...))
}

// AppHrefGT applies the GT predicate on the "appHref" field.
func AppHrefGT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGT(FieldAppHref, v))
}

// AppHrefGTE applies the GTE predicate on the "appHref" field.
func AppHrefGTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldGTE(FieldAppHref, v))
}

// AppHrefLT applies the LT predicate on the "appHref" field.
func AppHrefLT(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLT(FieldAppHref, v))
}

// AppHrefLTE applies the LTE predicate on the "appHref" field.
func AppHrefLTE(v string) predicate.Slider {
	return predicate.Slider(sql.FieldLTE(FieldAppHref, v))
}

// AppHrefContains applies the Contains predicate on the "appHref" field.
func AppHrefContains(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContains(FieldAppHref, v))
}

// AppHrefHasPrefix applies the HasPrefix predicate on the "appHref" field.
func AppHrefHasPrefix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasPrefix(FieldAppHref, v))
}

// AppHrefHasSuffix applies the HasSuffix predicate on the "appHref" field.
func AppHrefHasSuffix(v string) predicate.Slider {
	return predicate.Slider(sql.FieldHasSuffix(FieldAppHref, v))
}

// AppHrefEqualFold applies the EqualFold predicate on the "appHref" field.
func AppHrefEqualFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldEqualFold(FieldAppHref, v))
}

// AppHrefContainsFold applies the ContainsFold predicate on the "appHref" field.
func AppHrefContainsFold(v string) predicate.Slider {
	return predicate.Slider(sql.FieldContainsFold(FieldAppHref, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Slider) predicate.Slider {
	return predicate.Slider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Slider) predicate.Slider {
	return predicate.Slider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Slider) predicate.Slider {
	return predicate.Slider(func(s *sql.Selector) {
		p(s.Not())
	})
}
