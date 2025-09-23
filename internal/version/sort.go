package version

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/15 上午11:11
* @Package:
 */
type Collection []*Version

func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Less(i, j int) bool {
	return c[i].LessThan(c[j])
}

func (c Collection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
