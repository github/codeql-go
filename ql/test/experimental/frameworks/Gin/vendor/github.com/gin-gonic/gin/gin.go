package gin

import "github.com/gin-gonic/gin/binding"

type Context struct {
	Params Params

	// Accepted defines a list of manually accepted formats for content negotiation.
	Accepted []string
}

func (c *Context) GetHeader(key string) string {
	return ""
}
func (c *Context) QueryArray(key string) []string {
	return nil
}
func (c *Context) Query(key string) string {
	return ""
}
func (c *Context) PostFormArray(key string) []string {
	return nil
}
func (c *Context) PostForm(key string) string {
	return ""
}
func (c *Context) Param(key string) string {
	return ""
}
func (c *Context) GetStringSlice(key string) (ss []string) {
	return
}
func (c *Context) GetString(key string) (s string) {
	return
}
func (c *Context) GetRawData() ([]byte, error) {
	return nil, nil
}
func (c *Context) ClientIP() string {
	return ""
}
func (c *Context) ContentType() string {
	return ""
}
func (c *Context) Cookie(name string) (string, error) {
	return "", nil
}
func (c *Context) GetQueryArray(key string) ([]string, bool) {
	return []string{}, false
}
func (c *Context) GetQuery(key string) (string, bool) {
	return "", false
}
func (c *Context) GetPostFormArray(key string) ([]string, bool) {
	return []string{}, false
}
func (c *Context) GetPostForm(key string) (string, bool) {
	return "", false
}
func (c *Context) DefaultPostForm(key, defaultValue string) string {
	return defaultValue
}
func (c *Context) DefaultQuery(key, defaultValue string) string {
	return defaultValue
}
func (c *Context) GetPostFormMap(key string) (map[string]string, bool) {
	return nil, false
}
func (c *Context) GetQueryMap(key string) (map[string]string, bool) {
	return nil, false
}
func (c *Context) GetStringMap(key string) (sm map[string]interface{}) {
	return
}
func (c *Context) GetStringMapString(key string) (sms map[string]string) {
	return
}
func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	return
}
func (c *Context) PostFormMap(key string) map[string]string {
	return nil
}
func (c *Context) QueryMap(key string) map[string]string {
	return nil
}

type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// Get returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) (va string) {
	va, _ = ps.Get(name)
	return
}

func (c *Context) BindYAML(obj interface{}) error {
	return nil
}
func (c *Context) BindXML(obj interface{}) error {
	return nil
}
func (c *Context) BindWith(obj interface{}, b binding.Binding) error {
	return nil
}
func (c *Context) BindUri(obj interface{}) error {
	return nil
}
func (c *Context) BindQuery(obj interface{}) error {
	return nil
}
func (c *Context) MustBindWith(obj interface{}, b binding.Binding) error {
	return nil
}
func (c *Context) BindJSON(obj interface{}) error {
	return nil
}
func (c *Context) Bind(obj interface{}) error {
	return nil
}
func (c *Context) ShouldBind(obj interface{}) error {
	return nil
}
func (c *Context) ShouldBindBodyWith(obj interface{}, bb binding.BindingBody) (err error) {
	return nil
}
func (c *Context) ShouldBindJSON(obj interface{}) error {
	return nil
}
func (c *Context) ShouldBindQuery(obj interface{}) error {
	return nil
}
func (c *Context) ShouldBindUri(obj interface{}) error {
	return nil
}
func (c *Context) ShouldBindWith(obj interface{}, b binding.Binding) error {
	return nil
}
func (c *Context) ShouldBindXML(obj interface{}) error {
	return nil
}
func (c *Context) ShouldBindYAML(obj interface{}) error {
	return nil
}
