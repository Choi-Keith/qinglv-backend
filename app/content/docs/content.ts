import webapi from "./gocliRequest"
import * as components from "./contentComponents"
export * from "./contentComponents"

/**
 * @description "文章列表"
 * @param params
 */
export function getArticleList(params: components.GetArticleListReqParams) {
	return webapi.get<components.GetArticleListResp>(`/content/v1/article`, params)
}

/**
 * @description "根据id获取文章详情"
 * @param params
 */
export function getArticleById(params: components.GetArticleByIdReqParams, id: number) {
	return webapi.get<components.GetArticleByIdResp>(`/content/v1/article/${id}`, params)
}

/**
 * @description "我的文章列表"
 * @param params
 */
export function getUserArticleList(params: components.GetUserArticleListReqParams) {
	return webapi.get<components.GetUserArticleListResp>(`/content/v1/user/article`, params)
}

/**
 * @description "新建文章"
 * @param req
 */
export function addArticle(req: components.AddArticleReq) {
	return webapi.post<null>(`/content/v1/article`, req)
}

/**
 * @description "修改文章"
 * @param req
 */
export function updateArticle(req: components.UpdateArticleReq) {
	return webapi.put<null>(`/content/v1/article`, req)
}

/**
 * @description "删除"
 * @param params
 */
export function deleteArticle(params: components.DeleteArticleReqParams, id: number) {
	return webapi.delete<null>(`/content/v1/article/${id}`, params)
}

/**
 * @description "我关注的文章列表"
 * @param params
 */
export function getFollowingArticleList(params: components.GetArticleListReqParams) {
	return webapi.get<components.GetArticleListResp>(`/content/v1/article/following`, params)
}

/**
 * @description "分类列表"
 * @param params
 */
export function getCategoryList(params: components.GetCategoryListReqParams) {
	return webapi.get<components.GetCategoryListResp>(`/content/v1/category`, params)
}

/**
 * @description "通过Id获取分类详情"
 * @param params
 */
export function getCategoryById(params: components.GetCategoryByIdReqParams, id: number) {
	return webapi.get<components.GetCategoryByIdResp>(`/content/v1/category/${id}`, params)
}

/**
 * @description "新建分类"
 * @param req
 */
export function addCategory(req: components.AddCategoryReq) {
	return webapi.post<null>(`/content/v1/category`, req)
}

/**
 * @description "删除分类"
 * @param params
 */
export function deleteCategory(params: components.DeleteCategoryReqParams, id: number) {
	return webapi.delete<null>(`/content/v1/category/${id}`, params)
}

/**
 * @description "修改分类"
 * @param req
 */
export function updateCategory(req: components.UpdateCategoryReq, id:string) {
	return webapi.put<null>(`/content/v1/category/${id}`, req)
}

/**
 * @description "新建草稿"
 * @param req
 */
export function addDraft(req: components.AddDraftReq) {
	return webapi.post<null>(`/content/v1/draft`, req)
}

/**
 * @description "修改草稿"
 * @param req
 */
export function updateDraft(req: components.UpdateDraftReq) {
	return webapi.put<null>(`/content/v1/draft`, req)
}

/**
 * @description "草稿列表"
 * @param params
 */
export function getDraftList(params: components.GetDraftListReqParams) {
	return webapi.get<components.GetDraftListResp>(`/content/v1/draft`, params)
}

/**
 * @description "删除草稿"
 * @param params
 */
export function deleteDraft(params: components.DeleteDraftReqParams, id: number) {
	return webapi.delete<null>(`/content/v1/draft/${id}`, params)
}

/**
 * @description "草稿详情"
 * @param params
 */
export function getDrafById(params: components.GetDraftByIdReqParams, id: number) {
	return webapi.get<components.GetDraftByIdResp>(`/content/v1/draft/${id}`, params)
}

/**
 * @description "帖子列表"
 * @param params
 */
export function getPostList(params: components.GetPostListReqParams) {
	return webapi.get<components.GetPostListResp>(`/content/v1/post`, params)
}

/**
 * @description "根据id获取帖子详情"
 * @param params
 */
export function getPostById(params: components.GetPostByIdReqParams, id: number) {
	return webapi.get<components.GetPostByIdResp>(`/content/v1/post/${id}`, params)
}

/**
 * @description "我的帖子列表"
 * @param params
 */
export function getUserPostList(params: components.GetUserPostListReqParams) {
	return webapi.get<components.GetUserPostListResp>(`/content/v1/user/post`, params)
}

/**
 * @description "新建帖子"
 * @param req
 */
export function addPost(req: components.AddPostReq) {
	return webapi.post<null>(`/content/v1/post`, req)
}

/**
 * @description "删除"
 * @param params
 */
export function deletePost(params: components.DeletePostReqParams, id: number) {
	return webapi.delete<null>(`/content/v1/post/${id}`, params)
}

/**
 * @description "我关注的帖子列表"
 * @param params
 */
export function getFollowingPostList(params: components.GetPostListReqParams) {
	return webapi.get<components.GetPostListResp>(`/content/v1/post/following`, params)
}

/**
 * @description "标签列表"
 * @param params
 */
export function getTagList(params: components.GetTagListReqParams) {
	return webapi.get<components.GetTagListResp>(`/content/v1/tag`, params)
}

/**
 * @description "通过Id获取标签详情"
 * @param params
 */
export function getTagById(params: components.GetTagByIdReqParams, id: number) {
	return webapi.get<components.GetTagByIdResp>(`/content/v1/tag/${id}`, params)
}

/**
 * @description "通过name获取标签详情"
 * @param params
 */
export function getTagByName(params: components.GetTagByNameReqParams) {
	return webapi.get<components.GetTagByNameResp>(`/content/v1/tag/name`, params)
}

/**
 * @description "新建标签"
 * @param req
 */
export function addTag(req: components.AddTagReq) {
	return webapi.post<null>(`/content/v1/tag`, req)
}

/**
 * @description "删除标签"
 * @param params
 */
export function deleteTag(params: components.DeleteTagReqParams, id: number) {
	return webapi.delete<null>(`/content/v1/tag/${id}`, params)
}

/**
 * @description "修改标签"
 * @param req
 */
export function updateTag(req: components.UpdateTagReq,id:string) {
	return webapi.put<null>(`/content/v1/tag/${id}`, req)
}

/**
 * @description "帖子话题列表"
 * @param params
 */
export function getTopicList(params: components.GetTopicListReqParams) {
	return webapi.get<components.GetTopicListResp>(`/content/v1/topic`, params)
}

/**
 * @description "通过Id获取话题详情"
 * @param params
 */
export function getTopicById(params: components.GetTopicByIdReqParams, id: number) {
	return webapi.get<components.GetTopicByIdResp>(`/content/v1/topic/${id}`, params)
}

/**
 * @description "通过name获取话题详情"
 * @param params
 */
export function getTopicByName(params: components.GetTopicByNameReqParams) {
	return webapi.get<components.GetTopicByNameResp>(`/content/v1/topic/name`, params)
}

/**
 * @description "新建帖子话题"
 * @param req
 */
export function addTopic(req: components.AddTopicReq) {
	return webapi.post<null>(`/content/v1/topic`, req)
}

/**
 * @description "删除帖子话题"
 * @param params
 */
export function deleteTopic(params: components.DeleteTopicReqParams, id: number) {
	return webapi.delete<null>(`/content/v1/topic/${id}`, params)
}

/**
 * @description "修改帖子话题"
 * @param req
 */
export function updateTopic(req: components.UpdateTopicReq,id:string) {
	return webapi.put<null>(`/content/v1/topic/${id}`, req)
}

/**
 * @description "删除图片"
 * @param req
 */
export function removeImages(req: components.RemoveImagesReq) {
	return webapi.delete<null>(`/content/v1/image/remove`, req)
}

/**
 * @description "上传图片"
 * @param params
 */
export function uploadImages(params: components.UploadImagesReqParams, type: number) {
	return webapi.post<components.UploadImagesResp>(`/content/v1/image/upload/${type}`, params)
}
