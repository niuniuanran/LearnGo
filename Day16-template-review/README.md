# Go: the html/template package

## Functions used yesterday

- `func (t *Template) Execute(wr io.Writer, data interface{}) error` applies a **parsed template** to the specified data object, writing the output to wr.
- `func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error` applies the **template associated with t** that has the given name to the specified data object and writes the output to wr.
- `func New(name string) *Template` allocates a new HTML template with the given name.
- `func (t *Template) New(name string) *Template` allocates a new HTML template **associated with** the given one and with the same delimiters.
- `func ParseGlob(pattern string) (*Template, error)` creates a new Template and parses the template definitions from the files identified by the pattern.
- `func (t *Template) ParseGlob(pattern string) (*Template, error)` parses the template definitions in the files identified by the pattern and **associates** the resulting templates with t.

## What happened in Day15's project

1. A `AppConfig` type was created to hold site-wide application configurations, including the cached templates.
2. In `main`, whenever the application runs, an `AppConfig` instance is created.
3. In `main`, create the template cache and add the cache into the `AppConfig` instance
4. In `main`, call `SetAppConfig` to share the app config with the render package. This will ensure the cache persists across renders whenever handling a request.
5. `main` now registers the handle function, and whenever a request comes and a handler function gets invoked, `render.RenderTemplate` gets called, and it will use the template cache to retrieve the already parsed template with the corresponding name.
