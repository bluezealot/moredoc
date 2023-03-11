package model

func getPermissions() (permissions []Permission) {
	permissions = []Permission{
		{Title: "查看用户列表", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/ListUser"},
		{Title: "创建用户组", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/CreateGroup"},
		{Title: "查看友链列表", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/ListFriendlink"},
		{Title: "创建友链", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/CreateFriendlink"},
		{Title: "更新友链", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/UpdateFriendlink"},
		{Title: "删除友链", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/DeleteFriendlink"},
		{Title: "查看附件列表", Description: "", Method: "GRPC", Path: "/api.v1.AttachmentAPI/ListAttachment"},
		{Title: "删除附件", Description: "", Method: "GRPC", Path: "/api.v1.AttachmentAPI/DeleteAttachment"},
		{Title: "更新附件信息", Description: "", Method: "GRPC", Path: "/api.v1.AttachmentAPI/UpdateAttachment"},
		{Title: "获取横幅列表", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/ListBanner"},
		{Title: "上传横幅", Description: "", Method: "POST", Path: "/api/v1/upload/banner"},
		{Title: "创建横幅", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/CreateBanner"},
		{Title: "更新横幅", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/UpdateBanner"},
		{Title: "删除横幅", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/DeleteBanner"},
		{Title: "根据ID查询附件", Description: "", Method: "GRPC", Path: "/api.v1.AttachmentAPI/GetAttachment"},
		{Title: "根据ID查询横幅", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/GetBanner"},
		{Title: "根据ID查询友链", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/GetFriendlink"},
		{Title: "删除用户组", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/DeleteGroup"},
		{Title: "更新用户组", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/UpdateGroup"},
		{Title: "获取用户组列表", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/ListGroup"},
		{Title: "获取权限列表", Description: "", Method: "GRPC", Path: "/api.v1.PermissionAPI/ListPermission"},
		{Title: "根据ID查询权限", Description: "", Method: "GRPC", Path: "/api.v1.PermissionAPI/GetPermission"},
		{Title: "更新权限信息", Description: "", Method: "GRPC", Path: "/api.v1.PermissionAPI/UpdatePermission"},
		{Title: "获取用户组授权", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/GetGroupPermission"},
		{Title: "给用户组授权", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/UpdateGroupPermission"},
		{Title: "获取系统配置列表", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/ListConfig"},
		{Title: "更新系统配置", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/UpdateConfig"},
		{Title: "新增用户", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/AddUser"},
		{Title: "设置用户密码与分组", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/SetUser"},
		{Title: "删除用户", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/DeleteUser"},
		{Title: "更新用户资料", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/UpdateUserProfile"},
		{Title: "根据ID查询用户", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/GetUser"},
		{Title: "查询文档分类列表", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/ListCategory"},
		{Title: "根据ID查询文档分类", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/GetCategory"},
		{Title: "新增文档分类", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/CreateCategory"},
		{Title: "更新文档分类", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/UpdateCategory"},
		{Title: "删除文档分类", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/DeleteCategory"},
		{Title: "获取文档列表", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/ListDocument"},
		{Title: "删除文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/DeleteDocument"},
		{Title: "根据ID查询文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/GetDocument"},
		{Title: "更新文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/UpdateDocument"},
		{Title: "获取回收站文档列表", Description: "", Method: "GRPC", Path: "/api.v1.RecycleAPI/ListRecycleDocument"},
		{Title: "恢复回收站文档", Description: "", Method: "GRPC", Path: "/api.v1.RecycleAPI/RecoverRecycleDocument"},
		{Title: "清空回收站文档", Description: "", Method: "GRPC", Path: "/api.v1.RecycleAPI/ClearRecycleDocument"},
		{Title: "删除回收站文档", Description: "", Method: "GRPC", Path: "/api.v1.RecycleAPI/DeleteRecycleDocument"},
		{Title: "文章管理", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/ListArticle"},
		{Title: "更新文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/UpdateArticle"},
		{Title: "删除文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/DeleteArticle"},
		{Title: "创建文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/CreateArticle"},
		{Title: "上传文章图片和音视频", Description: "", Method: "POST", Path: "/api/v1/upload/article"},
		{Title: "上传文档分类封面", Description: "", Method: "POST", Path: "/api/v1/upload/category"},
		{Title: "上传配置图片文件", Description: "", Method: "POST", Path: "/api/v1/upload/config"},
		{Title: "获取评论列表", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/ListComment"},
		{Title: "获取单个评论", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/GetComment"},
		{Title: "批量审核评论", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/CheckComment"},
		{Title: "删除评论", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/DeleteComment"},
		{Title: "推荐文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/SetDocumentRecommend"},
		{Title: "查询举报列表", Description: "", Method: "GRPC", Path: "/api.v1.ReportAPI/ListReport"},
		{Title: "处理举报内容", Description: "", Method: "GRPC", Path: "/api.v1.ReportAPI/UpdateReport"},
		{Title: "删除举报内容", Description: "", Method: "GRPC", Path: "/api.v1.ReportAPI/DeleteReport"},
		{Title: "查看系统信息", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/GetStats"},
		{Title: "更新站点地图sitemap", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/UpdateSitemap"},
		{Title: "获取环境依赖检测", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/GetEnvs"},
		{Title: "更新头像", Description: "", Method: "POST", Path: "/api/v1/upload/avatar"},
		{Title: "失败文档一键重转", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/SetDocumentReconvert"},
	}
	return
}
