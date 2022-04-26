package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	v1 "lczx/api/v1"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
)

type sWdkFile struct{}

var (
	insWdkFile = sWdkFile{}
)

// WdkFile 文档库上传文件管理服务
func WdkFile() *sWdkFile {
	return &insWdkFile
}

// GetWdkFileRecord 获取文档库上传文件记录
func (s *sWdkFile) GetWdkFileRecord(ctx context.Context, projectId uint64) (list []*v1.WdkFileRecordInfo, err error) {
	err = dao.WdkFile.Ctx(ctx).Where(do.WdkFile{ProjectId: projectId}).WhereIn(dao.WdkFile.Columns().AuditStatus, []uint{0, 2}).
		OrderDesc(dao.WdkFile.Columns().CreateAt).ScanList(&list, "File")
	if err != nil {
		return
	}
	err = dao.WdkFiletype.Ctx(ctx).Where(dao.WdkFiletype.Columns().FileId, gdb.ListItemValuesUnique(list, "File", "FileId")).
		ScanList(&list, "Filetype", "File", "FileId:FileId")
	return
}
