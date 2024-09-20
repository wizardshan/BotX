package gin

import "github.com/gin-gonic/gin/binding"

type Validator interface {
	Validate() error
}

func (c *Context) Validate(obj any) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(obj, b); err != nil {
		return err
	}
	if validator, ok := obj.(Validator); ok {
		return validator.Validate()
	}
	return nil
}
