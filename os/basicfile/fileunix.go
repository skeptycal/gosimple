//go:build (Darwin && ignore) || (FreeBSD && ignore)
// +build Darwin,ignore FreeBSD,ignore

package basicfile

func (f *basicFile) FileUnix() FileUnix {
	return f
}
