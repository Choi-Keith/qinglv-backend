import webapi from "./gocliRequest"
import * as components from "./userComponents"
export * from "./userComponents"

/**
 * @description "黑名单"
 * @param params
 */
export function getBlacklist(params: components.BlacklistReqParams) {
	return webapi.get<components.BlacklistResp>(`/user/v1/blacklist`, params)
}

/**
 * @description "添加黑名单"
 * @param req
 */
export function addBlacklist(req: components.AddBlackItemReq) {
	return webapi.post<null>(`/user/v1/blacklist`, req)
}

/**
 * @description "取消黑名单"
 * @param req
 */
export function delBlacklist(req: components.DelBlackItemReq) {
	return webapi.delete<null>(`/user/v1/blacklist`, req)
}

/**
 * @description "是否黑名单"
 * @param params
 */
export function isBlackItem(params: components.IsBlackItemReqParams) {
	return webapi.get<components.IsBlackItemResp>(`/user/v1/blacklist/is`, params)
}

/**
 * @description "粉丝列表"
 * @param params
 */
export function getFollowerList(params: components.FollowerListReqParams) {
	return webapi.get<components.FollowerListResp>(`/user/v1/follower`, params)
}

/**
 * @description "关注列表"
 * @param params
 */
export function getFollowingList(params: components.FollowingListReqParams) {
	return webapi.get<components.FollowingListResp>(`/user/v1/following`, params)
}

/**
 * @description "关注"
 * @param req
 */
export function addFollowing(req: components.AddFollowingReq) {
	return webapi.post<null>(`/user/v1/following`, req)
}

/**
 * @description "取消关注"
 * @param req
 */
export function delFollowing(req: components.DelFollowingReq) {
	return webapi.delete<null>(`/user/v1/following`, req)
}

/**
 * @description "查看是否关注"
 * @param params
 */
export function isFollowing(params: components.IsFollowingReqParams) {
	return webapi.get<components.IsFollowingResp>(`/user/v1/following/is`, params)
}

/**
 * @description "消息列表"
 * @param params
 * @param req
 */
export function getMessageList(params: components.GetMessageListReqParams, req: components.GetMessageListReq) {
	return webapi.get<components.GetMessageListResp>(`/user/v1/message`, params, req)
}

/**
 * @description "标记已读"
 * @param req
 */
export function readAllMessageReq(req: components.ReadAllMessageReq) {
	return webapi.get<null>(`/user/v1/message/read/all`, req)
}

/**
 * @description "消息条数"
 */
// export function getUnreadMessageCount() {
// 	return webapi.get<components.GetMessageCountResp>(`/user/v1/message/unread/count`)
// }

/**
 * @description "获取角色列表"
 * @param params
 */
export function getRoleList(params: components.RoleListReqParams) {
	return webapi.get<components.RoleListResp>(`/user/v1/role`, params)
}

/**
 * @description "创建角色"
 * @param req
 */
export function addRole(req: components.RoleReq) {
	return webapi.post<null>(`/user/v1/role`, req)
}

/**
 * @description "获取角色详情"
 */
// export function getRoleByID() {
// 	return webapi.get<components.Role>(`/user/v1/role/${id}`)
// }

/**
 * @description "更改角色"
 * @param req
 */
// export function updateRole(req: components.Role) {
// 	return webapi.put<null>(`/user/v1/role/${id}`, req)
// }

/**
 * @description "删除角色"
 */
// export function delRole() {
// 	return webapi.delete<null>(`/user/v1/role/${id}`)
// }

/**
 * @description "邮件验证"
 * @param params
 */
export function verifyEmail(params: components.VerifyEmailReqParams) {
	return webapi.get<components.VerifyEmailResp>(`/user/v1/email/verify`, params)
}

/**
 * @description "登录"
 * @param req
 */
export function login(req: components.LoginReq) {
	return webapi.post<components.LoginResp>(`/user/v1/login`, req)
}

/**
 * @description "获取登录验证码"
 */
// export function loginCaptcha() {
// 	return webapi.get<components.LoginCaptchaResp>(`/user/v1/login/captcha`)
// }

/**
 * @description "注册"
 * @param req
 */
export function register(req: components.RegisterReq) {
	return webapi.post<components.User>(`/user/v1/register`, req)
}

/**
 * @description "用户列表"
 * @param params
 */
export function getUserList(params: components.UserListReqParams) {
	return webapi.get<components.UserListResp>(`/user/v1/user`, params)
}

/**
 * @description "用户详情"
 * @param params
 */
export function getUserById(params: components.GetUserByIdReqParams, id: number) {
	return webapi.get<components.User>(`/user/v1/user/${id}`, params)
}

/**
 * @description "退出"
 */
// export function logout() {
// 	return webapi.post<null>(`/user/v1/logout`)
// }

/**
 * @description "修改密码"
 * @param req
 */
export function updatePassword(req: components.PasswordReq) {
	return webapi.put<null>(`/user/v1/password`, req)
}

/**
 * @description "更改个人信息"
 * @param req
 */
// export function updateUser(req: components.UpdateUserReq) {
// 	return webapi.put<null>(`/user/v1/user/${id}`, req)
// }

/**
 * @description "删除用户"
 * @param params
 */
export function delUser(params: any, id: number) {
	return webapi.delete<null>(`/user/v1/user/${id}`, params)
}

/**
 * @description "修改用户头像"
 * @param params
 */
export function updateAvatar(params: any, id: number) {
	return webapi.put<null>(`/user/v1/user/${id}/avatar`, params)
}

/**
 * @description "修改个人主页背景"
 * @param params
 */
export function updateProfileBg(params: components.UpdateProfileReqParams, id: number) {
	return webapi.put<null>(`/user/v1/user/${id}/background/image`, params)
}

/**
 * @description "禁止用户"
 * @param params
 */
export function banUser(params: any, id: number) {
	return webapi.put<null>(`/user/v1/user/ban/${id}`, params)
}

/**
 * @description "获取当前用户"
 */
// export function getCurrentUser() {
// 	return webapi.get<any>(`/user/v1/user/current`)
// }
