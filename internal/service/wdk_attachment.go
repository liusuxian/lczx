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
func (s *sWdkAttachment) GetWdkAttachmentRecord(ctx context.Context, projectId uint64) (list []*v1.WdkAttachmentInfo, err error) {
	err = dao.WdkAttachmentRecord.Ctx(ctx).Where(do.WdkAttachmentRecord{ProjectId: projectId}).
		OrderDesc(dao.WdkAttachmentRecord.Columns().CreateAt).ScanList(&list, "Record")
	if err != nil {
		return
	}
	err = dao.WdkAttachmentFile.Ctx(ctx).Where(dao.WdkAttachmentFile.Columns().Id, gdb.ListItemValuesUnique(list, "Record", "Id")).
		ScanList(&list, "Attachment", "Record", "Id:Id")
	return
}

// AddWdkAttachment 新增文档库上传附件记录
func (s *sWdkAttachment) AddWdkAttachment(ctx context.Context, req *v1.WdkAttachmentAddReq, Attachments []*upload.FileInfo) (err error) {
	err = dao.WdkAttachmentRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查新增文档库上传附件记录权限
		var terr error
		_, terr = s.AuthAdd(ctx, req.ProjectId)
		if terr != nil {
			return terr
		}
		// 保存文档库上传附件记录数据
		terr = s.saveWdkAttachmentRecord(ctx, req, Attachments)
		return terr
	})
	return
}

// AuthAdd 检查新增文档库上传附件记录权限
func (s *sWdkAttachment) AuthAdd(ctx context.Context, projectId uint64) (wdkProject *v1.WdkProjectInfo, err error) {
	// 通过文档库项目ID判断文档库项目信息是否存在
	wdkProject, err = WdkProject().GetWdkProjectById(ctx, projectId)
	if err != nil {
		return
	}
	if wdkProject == nil || wdkProject.ProjectInfo == nil {
		err = gerror.Newf(`文档库项目ID[%d]不存在`, projectId)
		return
	}
	// 判断写入权限
	user := Context().Get(ctx).User
	if user.Id != wdkProject.ProjectInfo.PrincipalUid {
		err = gerror.New("抱歉！！！该项目您没有上传附件的权限")
		return
	}
	return
}

// saveWdkAttachmentRecord 保存文档库上传附件记录数据
func (s *sWdkAttachment) saveWdkAttachmentRecord(ctx context.Context, req *v1.WdkAttachmentAddReq, Attachments []*upload.FileInfo) (err error) {
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
