import webapi from "./gocliRequest"
import * as components from "./operationComponents"
export * from "./operationComponents"

/**
 * @description "收藏列表"
 * @param params
 */
export function getCollectionList(params: components.GetCollectionListReqParams) {
	return webapi.get<components.GetCollectionListResp>(`/operation/v1/collection`, params)
}

/**
 * @description "新建收藏"
 * @param req
 */
export function addCollection(req: components.AddCollectionReq) {
	return webapi.post<null>(`/operation/v1/collection`, req)
}

/**
 * @description "取消收藏"
 * @param params
 */
export function deleteCollection(params: components.DeleteCollectionReqParams, id: number) {
	return webapi.delete<null>(`/operation/v1/collection/${id}`, params)
}

/**
 * @description "收藏列表分组"
 * @param params
 */
export function getCollectionGroupList(params: components.GetCollectionGroupListReqParams) {
	return webapi.get<components.GetCollectionGropListResp>(`/operation/v1/collection/group`, params)
}

/**
 * @description "新建收藏分组"
 * @param req
 */
export function addCollectionGroup(req: components.AddCollectionGroup) {
	return webapi.post<null>(`/operation/v1/collection/group`, req)
}

/**
 * @description "删除收藏分组"
 * @param params
 */
export function deleteCollectionGroup(params: components.DeleteCollectionGroupParams, id: number) {
	return webapi.delete<null>(`/operation/v1/collection/group/${id}`, params)
}

/**
 * @description "修改收藏分组"
 * @param req
 */
export function updateCollectionGroup(req: components.UpdateCollectionGroup, id:string) {
	return webapi.put<null>(`/operation/v1/collection/group/${id}`, req)
}

/**
 * @description "获取文章评论"
 * @param params
 */
export function getArticleCommentList(params: components.GetArticleCommentListReqParams) {
	return webapi.get<components.GetArticleCommentListResp>(`/operation/v1/article/comment`, params)
}

/**
 * @description "获取文章评论回复"
 * @param params
 */
export function getArticleCommentReplyList(params: components.GetArticleCommentReplyListReqParams) {
	return webapi.get<components.GetArticleCommentReplyListResp>(`/operation/v1/article/comment/reply`, params)
}

/**
 * @description "获取帖子评论"
 * @param params
 */
export function getPostCommentList(params: components.GetPostCommentListReqParams) {
	return webapi.get<components.GetPostCommentListResp>(`/operation/v1/post/comment`, params)
}

/**
 * @description "获取帖子评论回复"
 * @param params
 */
export function getPostCommentReplyList(params: components.GetPostCommentReplyListReqParams) {
	return webapi.get<components.GetPostCommentReplyListResp>(`/operation/v1/post/comment/reply`, params)
}

/**
 * @description "添加文章评论"
 * @param req
 */
export function addArticleComment(req: components.AddArticleCommentReq) {
	return webapi.post<null>(`/operation/v1/article/comment`, req)
}

/**
 * @description "删除文章评论"
 * @param params
 */
export function deleteArticleComment(params: components.DeleteArticleCommentReqParams, id: number) {
	return webapi.delete<null>(`/operation/v1/article/comment/${id}`, params)
}

/**
 * @description "添加文章评论回复"
 * @param req
 */
export function addArticleCommentReply(req: components.AddArticleCommentReplyReq) {
	return webapi.post<null>(`/operation/v1/article/comment/reply`, req)
}

/**
 * @description "删除文章评论回复"
 * @param params
 */
export function deleteArticleCommentReply(params: components.DeleteArticleCommentReplyReqParams, id: number) {
	return webapi.delete<null>(`/operation/v1/article/comment/reply/${id}`, params)
}

/**
 * @description "添加帖子评论"
 * @param req
 */
export function addPostComment(req: components.AddPostCommentReq) {
	return webapi.post<null>(`/operation/v1/post/comment`, req)
}

/**
 * @description "删除帖子评论"
 * @param params
 */
export function deletePostComment(params: components.DeletePostCommentReqParams, id: number) {
	return webapi.delete<null>(`/operation/v1/post/comment/${id}`, params)
}

/**
 * @description "添加帖子评论回复"
 * @param req
 */
export function addPostCommentReply(req: components.AddPostCommentReplyReq) {
	return webapi.post<null>(`/operation/v1/post/comment/reply`, req)
}

/**
 * @description "删除帖子评论回复"
 * @param params
 */
export function deletePostCommentReply(params: components.DeletePostCommentReplyReqParams, id: number) {
	return webapi.delete<null>(`/operation/v1/post/comment/reply/${id}`, params)
}

/**
 * @description "文章分享"
 * @param req
 */
export function addArticleShare(req: components.AddArticleShareReq) {
	return webapi.post<null>(`/operation/v1/article/share`, req)
}

/**
 * @description "帖子分享"
 * @param req
 */
export function addPostShare(req: components.AddPostShareReq) {
	return webapi.post<null>(`/operation/v1/post/share`, req)
}

/**
 * @description "文章回复评论点踩"
 * @param req
 */
export function handleArticleCommentReplyThumbDown(req: components.HandleArticleCommentReplyThumbDownReq) {
	return webapi.post<null>(`/operation/v1/article/comment/reply/thumb/down`, req)
}

/**
 * @description "文章回复评论点赞"
 * @param req
 */
export function handleArticleCommentReplyThumbUp(req: components.HandleArticleCommentReplyThumbUpReq) {
	return webapi.post<null>(`/operation/v1/article/comment/reply/thumb/up`, req)
}

/**
 * @description "文章评论点踩"
 * @param req
 */
export function handleArticleCommentThumbDown(req: components.HandleArticleCommentThumbDownReq) {
	return webapi.post<null>(`/operation/v1/article/comment/thumb/down`, req)
}

/**
 * @description "文章评论点赞"
 * @param req
 */
export function handleArticleCommentThumbUp(req: components.HandleArticleCommentThumbUpReq) {
	return webapi.post<null>(`/operation/v1/article/comment/thumb/up`, req)
}

/**
 * @description "文章点踩"
 * @param req
 */
export function handleArticleThumbDown(req: components.HandleArticleThumbDownReq) {
	return webapi.post<null>(`/operation/v1/article/thumb/down`, req)
}

/**
 * @description "文章点赞"
 * @param req
 */
export function handleArticleThumbUp(req: components.HandleArticleThumbUpReq) {
	return webapi.post<null>(`/operation/v1/article/thumb/up`, req)
}

/**
 * @description "帖子回复评论点踩"
 * @param req
 */
export function handlePostCommentReplyThumbDown(req: components.HandlePostCommentReplyThumbDownReq) {
	return webapi.post<null>(`/operation/v1/post/comment/reply/thumb/down`, req)
}

/**
 * @description "帖子回复评论点赞"
 * @param req
 */
export function handlePostCommentReplyThumbUp(req: components.HandlePostCommentReplyThumbUpReq) {
	return webapi.post<null>(`/operation/v1/post/comment/reply/thumb/up`, req)
}

/**
 * @description "帖子评论点踩"
 * @param req
 */
export function handlePostCommentThumbDown(req: components.HandlePostCommentThumbDownReq) {
	return webapi.post<null>(`/operation/v1/post/comment/thumb/down`, req)
}

/**
 * @description "帖子评论点赞"
 * @param req
 */
export function handlePostCommentThumbUp(req: components.HandlePostCommentThumbUpReq) {
	return webapi.post<null>(`/operation/v1/post/comment/thumb/up`, req)
}

/**
 * @description "帖子点踩"
 * @param req
 */
export function handlePostThumbDown(req: components.HandlePostThumbDownReq) {
	return webapi.post<null>(`/operation/v1/post/thumb/down`, req)
}

/**
 * @description "帖子点赞"
 * @param req
 */
export function handlePostThumbUp(req: components.HandlePostThumbUpReq) {
	return webapi.post<null>(`/operation/v1/post/thumb/up`, req)
}
