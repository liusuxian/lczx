package wdk

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/do"
	"lczx/internal/service"
)

type sWdkAttachment struct{}

func init() {
	service.RegisterWdkAttachment(newWdkAttachment())
}

// 文档库上传附件记录管理服务
func newWdkAttachment() *sWdkAttachment {
	return &sWdkAttachment{}
}

// GetWdkAttachmentRecord 获取文档库上传附件记录
func (s *sWdkAttachment) GetWdkAttachmentRecord(ctx context.Context, projectId uint64) (list []*v1.WdkAttachmentInfo, err error) {
	err = dao.WdkAttachmentRecord.Ctx(ctx).Where(do.WdkAttachmentRecord{ProjectId: projectId}).
		OrderDesc(dao.WdkAttachmentRecord.Columns().CreatedAt).ScanList(&list, "Record")
	if err != nil {
		return
	}
	err = dao.WdkAttachmentFile.Ctx(ctx).Where(dao.WdkAttachmentFile.Columns().Id, gdb.ListItemValuesUnique(list, "Record", "Id")).
		ScanList(&list, "Attachment", "Record", "Id:Id")
	return
}

// AddWdkAttachment 新增文档库上传附件记录
func (s *sWdkAttachment) AddWdkAttachment(ctx context.Context, req *v1.WdkAttachmentAddReq, Attachments []*model.UploadFileInfo) (err error) {
	err = dao.WdkAttachmentRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 保存文档库上传附件记录数据
		var terr error
		terr = s.saveWdkAttachmentRecord(ctx, req, Attachments)
		if terr != nil {
			return terr
		}
		// 设置文档库项目文件上传状态为正常
		terr = service.WdkProject().SetWdkProjectFileUploadStatusNormal(ctx, req.ProjectId)
		return terr
	})
	return
}

// DeleteWdkAttachment 删除文档库上传附件记录
func (s *sWdkAttachment) DeleteWdkAttachment(ctx context.Context, ids []uint64) (err error) {
	err = dao.WdkAttachmentRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除文档库上传附件记录
		var terr error
		_, terr = dao.WdkAttachmentRecord.Ctx(ctx).WhereIn(dao.WdkAttachmentRecord.Columns().Id, ids).Delete()
		if terr != nil {
			return terr
		}
		// 删除文档库上传附件
		_, terr = dao.WdkAttachmentFile.Ctx(ctx).WhereIn(dao.WdkAttachmentFile.Columns().Id, ids).Delete()
		return terr
	})
	return
}

// saveWdkAttachmentRecord 保存文档库上传附件记录数据
func (s *sWdkAttachment) saveWdkAttachmentRecord(ctx context.Context, req *v1.WdkAttachmentAddReq, Attachments []*model.UploadFileInfo) (err error) {
	var recordId int64
	recordId, err = dao.WdkAttachmentRecord.Ctx(ctx).Data(do.WdkAttachmentRecord{
		ProjectId: req.ProjectId,
		Remark:    req.Remark,
	}).FieldsEx(dao.WdkAttachmentRecord.Columns().Id).InsertAndGetId()
	if err != nil {
		return
	}
	// 保存文档库上传附件文件数据
	attachmentFileData := g.List{}
	for _, file := range Attachments {
		attachmentFileData = append(attachmentFileData, g.Map{
			"id":         recordId,
			"name":       file.FileName,
			"origin_url": file.OriginFileUrl,
			"pdf_url":    file.PdfFileUrl,
		})
	}
	_, err = dao.WdkAttachmentFile.Ctx(ctx).Data(attachmentFileData).Batch(len(attachmentFileData)).Insert()
	return
}
