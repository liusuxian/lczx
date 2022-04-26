package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/internal/upload"
)

type sWdkAttachment struct{}

var (
	insWdkAttachment = sWdkAttachment{}
)

// WdkAttachment 文档库上传附件记录管理服务
func WdkAttachment() *sWdkAttachment {
	return &insWdkAttachment
}

// GetWdkAttachmentRecord 获取文档库上传附件记录
func (s *sWdkAttachment) GetWdkAttachmentRecord(ctx context.Context, projectId uint64) (list []*v1.WdkAttachmentRecordInfo, err error) {
	err = dao.WdkAttachmentRecord.Ctx(ctx).Where(do.WdkAttachmentRecord{ProjectId: projectId}).OrderDesc(dao.WdkAttachmentRecord.Columns().CreateAt).
		ScanList(&list, "Record")
	if err != nil {
		return
	}
	err = dao.WdkAttachmentFile.Ctx(ctx).Where(dao.WdkAttachmentFile.Columns().Id, gdb.ListItemValuesUnique(list, "Record", "Id")).
		ScanList(&list, "Attachment", "Record", "Id:Id")
	return
}

// AddWdkAttachmentRecord 新增文档库上传附件记录
func (s *sWdkAttachment) AddWdkAttachmentRecord(ctx context.Context, req *v1.WdkAttachmentRecordAddReq, Attachments []*upload.FileInfo) (err error) {
	err = dao.WdkAttachmentRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 通过文档库项目ID判断文档库项目信息是否存在
		var terr error
		var wdkProjectExists bool
		wdkProjectExists, terr = WdkProject().WdkProjectExistsById(ctx, req.ProjectId)
		if terr != nil {
			return terr
		}
		if !wdkProjectExists {
			terr = gerror.Newf(`文档库项目ID[%d]不存在`, req.ProjectId)
			return terr
		}
		// 写入文档库上传附件记录数据
		var recordId int64
		recordId, terr = dao.WdkAttachmentRecord.Ctx(ctx).Data(do.WdkAttachmentRecord{
			ProjectId: req.ProjectId,
			Remark:    req.Remark,
		}).InsertAndGetId()
		if terr != nil {
			return terr
		}
		// 写入文档库上传附件数据
		attachmentFileData := g.List{}
		for _, file := range Attachments {
			attachmentFileData = append(attachmentFileData, g.Map{
				"id":         recordId,
				"name":       file.FileName,
				"origin_url": file.OriginFileUrl,
				"pdf_url":    file.PdfFileUrl,
			})
		}
		_, terr = dao.WdkAttachmentFile.Ctx(ctx).Data(attachmentFileData).Batch(len(g.List{})).Insert()
		return terr
	})
	return
}

// DeleteWdkAttachmentRecord 删除文档库上传附件记录
func (s *sWdkAttachment) DeleteWdkAttachmentRecord(ctx context.Context, ids []uint64) (err error) {
	err = dao.WdkAttachmentRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var terr error
		_, terr = dao.WdkAttachmentRecord.Ctx(ctx).WhereIn(dao.WdkAttachmentRecord.Columns().Id, ids).Delete()
		if terr != nil {
			return terr
		}
		_, terr = dao.WdkAttachmentFile.Ctx(ctx).WhereIn(dao.WdkAttachmentFile.Columns().Id, ids).Delete()
		return terr
	})
	return
}
