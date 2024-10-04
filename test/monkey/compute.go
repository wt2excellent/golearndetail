// Package monkey ------------------------------------------------------------
// package monkey
// @file      : compute.go
// @author    : WeiTao
// @contact   : 15537588047@163.com
// @time      : 2024/10/4 15:22
// ------------------------------------------------------------
package monkey

func networkCompute(a, b int) (int, error) {
	c := a + b
	return c, nil
}

func compute(a, b int) (int, error) {
	c, err := networkCompute(a, b)
	return c, err
}

type Compute struct{}

func (c *Compute) NetworkCompute(a, b int) (int, error) {
	sum := a + b
	return sum, nil
}

func (c *Compute) Compute(a, b int) (int, error) {
	sum, err := c.NetworkCompute(a, b)
	return sum, err
}
