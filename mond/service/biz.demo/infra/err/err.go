package xerr

import merr "github.com/tangbo/twatt/mond/wind/err"

var (
	LoginErr = merr.NewError(19001001, "登录失败，token验证不合法", merr.Abnormal)
)
