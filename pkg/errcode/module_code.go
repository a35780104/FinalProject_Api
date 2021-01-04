package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")

	ErrorGetArticleFail    = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")

	ErrorGetMemberFail    = NewError(20040001, "獲取一位用戶失敗")
	ErrorGetMembersFail   = NewError(20040002, "獲取多位用戶失敗")
	ErrorCreateMemberFail = NewError(20040003, "創建用戶失敗")
	ErrorUpdateMemberFail = NewError(20040004, "更新用戶失敗")
	ErrorDeleteMemberFail = NewError(20040005, "刪除用戶失敗")

	ErrorGetProductFail    = NewError(20050001, "獲取一個產品失敗")
	ErrorGetProductsFail   = NewError(20050002, "獲取多個產品失敗")
	ErrorCreateProductFail = NewError(20050003, "創建產品失敗")
	ErrorUpdateProductFail = NewError(20050004, "更新產品失敗")
	ErrorDeleteProductFail = NewError(20050005, "刪除產品失敗")

	ErrorGetRecordsFail    = NewError(20060001, "獲取一筆記錄失敗")
	ErrorGetRecordssFail   = NewError(20060002, "獲取多筆記錄失敗")
	ErrorCreateRecordsFail = NewError(20060003, "創建記錄失敗")
	ErrorUpdateRecordsFail = NewError(20060004, "更新記錄失敗")
	ErrorDeleteRecordsFail = NewError(20060005, "刪除記錄失敗")
)
