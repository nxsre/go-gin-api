package router

import (
	"github.com/nxsre/go-gin-api/internal/api/admin"
	"github.com/nxsre/go-gin-api/internal/api/authorized"
	"github.com/nxsre/go-gin-api/internal/api/config"
	"github.com/nxsre/go-gin-api/internal/api/cron"
	"github.com/nxsre/go-gin-api/internal/api/helper"
	"github.com/nxsre/go-gin-api/internal/api/menu"
	"github.com/nxsre/go-gin-api/internal/api/tool"
	"github.com/nxsre/go-gin-api/internal/pkg/core"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/swag"
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"
	"sync"
)

func setApiRouter(r *resource) {
	// helper
	helperHandler := helper.New(r.logger, r.db, r.cache)

	helpers := r.mux.Group("/helper")
	{
		helpers.GET("/md5/:str", helperHandler.Md5())
		helpers.POST("/sign", helperHandler.Sign())
	}

	// admin
	adminHandler := admin.New(r.logger, r.db, r.cache)

	// swagger
	swaggerHandler := r.mux.Group("/swagger", core.WrapAuthHandler(r.interceptors.CheckLoginCookie))
	{
		swaggerHandler.GET("/*any", swaggerWarp())
	}

	// 需要签名验证，无需登录验证，无需 RBAC 权限验证
	login := r.mux.Group("/api", r.interceptors.CheckSignature())
	{
		login.POST("/login", adminHandler.Login())
	}

	// 需要签名验证、登录验证，无需 RBAC 权限验证
	notRBAC := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature())
	{
		notRBAC.POST("/admin/logout", adminHandler.Logout())
		notRBAC.PATCH("/admin/modify_password", adminHandler.ModifyPassword())
		notRBAC.GET("/admin/info", adminHandler.Detail())
		notRBAC.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())
	}

	// 需要签名验证、登录验证、RBAC 权限验证
	api := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature(), r.interceptors.CheckRBAC())
	{
		// authorized
		authorizedHandler := authorized.New(r.logger, r.db, r.cache)
		api.POST("/authorized", authorizedHandler.Create())
		api.GET("/authorized", authorizedHandler.List())
		api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
		api.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())

		api.POST("/authorized_api", authorizedHandler.CreateAPI())
		api.GET("/authorized_api", authorizedHandler.ListAPI())
		api.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())

		api.POST("/admin", adminHandler.Create())
		api.GET("/admin", adminHandler.List())
		api.PATCH("/admin/used", adminHandler.UpdateUsed())
		api.PATCH("/admin/offline", adminHandler.Offline())
		api.PATCH("/admin/reset_password/:id", core.AliasForRecordMetrics("/api/admin/reset_password"), adminHandler.ResetPassword())
		api.DELETE("/admin/:id", core.AliasForRecordMetrics("/api/admin"), adminHandler.Delete())

		api.POST("/admin/menu", adminHandler.CreateAdminMenu())
		api.GET("/admin/menu/:id", core.AliasForRecordMetrics("/api/admin/menu"), adminHandler.ListAdminMenu())

		// menu
		menuHandler := menu.New(r.logger, r.db, r.cache)
		api.POST("/menu", menuHandler.Create())
		api.GET("/menu", menuHandler.List())
		api.GET("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Detail())
		api.PATCH("/menu/used", menuHandler.UpdateUsed())
		api.PATCH("/menu/sort", menuHandler.UpdateSort())
		api.DELETE("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Delete())
		api.POST("/menu_action", menuHandler.CreateAction())
		api.GET("/menu_action", menuHandler.ListAction())
		api.DELETE("/menu_action/:id", core.AliasForRecordMetrics("/api/menu_action"), menuHandler.DeleteAction())

		// tool
		toolHandler := tool.New(r.logger, r.db, r.cache)
		api.GET("/tool/hashids/encode/:id", core.AliasForRecordMetrics("/api/tool/hashids/encode"), toolHandler.HashIdsEncode())
		api.GET("/tool/hashids/decode/:id", core.AliasForRecordMetrics("/api/tool/hashids/decode"), toolHandler.HashIdsDecode())
		api.POST("/tool/cache/search", toolHandler.SearchCache())
		api.PATCH("/tool/cache/clear", toolHandler.ClearCache())
		api.GET("/tool/data/dbs", toolHandler.Dbs())
		api.POST("/tool/data/tables", toolHandler.Tables())
		api.POST("/tool/data/mysql", toolHandler.SearchMySQL())
		api.POST("/tool/send_message", toolHandler.SendMessage())

		// config
		configHandler := config.New(r.logger, r.db, r.cache)
		api.PATCH("/config/email", configHandler.Email())

		// cron
		cronHandler := cron.New(r.logger, r.db, r.cache, r.cronServer)
		api.POST("/cron", cronHandler.Create())
		api.GET("/cron", cronHandler.List())
		api.GET("/cron/:id", core.AliasForRecordMetrics("/api/cron/detail"), cronHandler.Detail())
		api.POST("/cron/:id", core.AliasForRecordMetrics("/api/cron/modify"), cronHandler.Modify())
		api.PATCH("/cron/used", cronHandler.UpdateUsed())
		api.PATCH("/cron/exec/:id", core.AliasForRecordMetrics("/api/cron/exec"), cronHandler.Execute())

	}
}

func swaggerWarp() func(ctx core.Context) {
	handler := swaggerFiles.NewHandler()
	var config = Config{
		URL:                      "doc.json",
		DocExpansion:             "list",
		InstanceName:             swag.Name,
		Title:                    "Swagger UI",
		DefaultModelsExpandDepth: 1,
		DeepLinking:              true,
		PersistAuthorization:     false,
		Oauth2DefaultClientID:    "",
	}

	var once sync.Once
	// create a template with name
	index, _ := template.New("swagger_index.html").Parse(swaggerIndexTpl)

	var matcher = regexp.MustCompile(`(.*)(index\.html|doc\.json|favicon-16x16\.png|favicon-32x32\.png|/oauth2-redirect\.html|swagger-ui\.css|swagger-ui\.css\.map|swagger-ui\.js|swagger-ui\.js\.map|swagger-ui-bundle\.js|swagger-ui-bundle\.js\.map|swagger-ui-standalone-preset\.js|swagger-ui-standalone-preset\.js\.map)[?|.]*`)
	return func(ctx core.Context) {

		if ctx.Method() != http.MethodGet {
			ctx.AbortWithError(core.Error(http.StatusMethodNotAllowed, 11403, "xxx"))
			return
		}

		matches := matcher.FindStringSubmatch(ctx.URI())

		if len(matches) != 3 {
			ctx.ResponseWriter().WriteHeader(http.StatusNotFound)
			return
		}

		path := matches[2]
		once.Do(func() {
			handler.Prefix = matches[1]
		})

		switch filepath.Ext(path) {
		case ".html":
			ctx.SetHeader("Content-Type", "text/html; charset=utf-8")
		case ".css":
			ctx.SetHeader("Content-Type", "text/css; charset=utf-8")
		case ".js":
			ctx.SetHeader("Content-Type", "application/javascript")
		case ".png":
			ctx.SetHeader("Content-Type", "image/png")
		case ".json":
			ctx.SetHeader("Content-Type", "application/json; charset=utf-8")
		}

		switch path {
		case "index.html":
			_ = index.Execute(ctx.ResponseWriter(), config.toSwaggerConfig())
		case "doc.json":
			doc, err := swag.ReadDoc(config.InstanceName)
			if err != nil {
				ctx.AbortWithError(core.Error(http.StatusInternalServerError, 10500, "xxx"))

				return
			}
			ctx.ResponseWriter().WriteString(doc)
		default:
			handler.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
			//handler.ServeHTTP(ctx.Writer, ctx.Request)
		}
	}
}

const swaggerIndexTpl = `<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>{{.Title}}</title>
  <link rel="stylesheet" type="text/css" href="./swagger-ui.css" >
  <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16" />
  <style>
    html
    {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
    }
    *,
    *:before,
    *:after
    {
        box-sizing: inherit;
    }

    body {
      margin:0;
      background: #fafafa;
    }
  </style>
</head>

<body>

<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="position:absolute;width:0;height:0">
  <defs>
    <symbol viewBox="0 0 20 20" id="unlocked">
          <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V6h2v-.801C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8z"></path>
    </symbol>

    <symbol viewBox="0 0 20 20" id="locked">
      <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8zM12 8H8V5.199C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="close">
      <path d="M14.348 14.849c-.469.469-1.229.469-1.697 0L10 11.819l-2.651 3.029c-.469.469-1.229.469-1.697 0-.469-.469-.469-1.229 0-1.697l2.758-3.15-2.759-3.152c-.469-.469-.469-1.228 0-1.697.469-.469 1.228-.469 1.697 0L10 8.183l2.651-3.031c.469-.469 1.228-.469 1.697 0 .469.469.469 1.229 0 1.697l-2.758 3.152 2.758 3.15c.469.469.469 1.229 0 1.698z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow">
      <path d="M13.25 10L6.109 2.58c-.268-.27-.268-.707 0-.979.268-.27.701-.27.969 0l7.83 7.908c.268.271.268.709 0 .979l-7.83 7.908c-.268.271-.701.27-.969 0-.268-.269-.268-.707 0-.979L13.25 10z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow-down">
      <path d="M17.418 6.109c.272-.268.709-.268.979 0s.271.701 0 .969l-7.908 7.83c-.27.268-.707.268-.979 0l-7.908-7.83c-.27-.268-.27-.701 0-.969.271-.268.709-.268.979 0L10 13.25l7.418-7.141z"/>
    </symbol>


    <symbol viewBox="0 0 24 24" id="jump-to">
      <path d="M19 7v4H5.83l3.58-3.59L8 6l-6 6 6 6 1.41-1.41L5.83 13H21V7z"/>
    </symbol>

    <symbol viewBox="0 0 24 24" id="expand">
      <path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/>
    </symbol>

  </defs>
</svg>

<div id="swagger-ui"></div>

<script src="./swagger-ui-bundle.js"> </script>
<script src="./swagger-ui-standalone-preset.js"> </script>
<script>
window.onload = function() {
  // Build a system
  const ui = SwaggerUIBundle({
    url: "{{.URL}}",
    dom_id: '#swagger-ui',
    validatorUrl: null,
    oauth2RedirectUrl: {{.Oauth2RedirectURL}},
    persistAuthorization: {{.PersistAuthorization}},
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
	layout: "StandaloneLayout",
    docExpansion: "{{.DocExpansion}}",
	deepLinking: {{.DeepLinking}},
	defaultModelsExpandDepth: {{.DefaultModelsExpandDepth}}
  })

  const defaultClientId = "{{.Oauth2DefaultClientID}}";
  if (defaultClientId) {
    ui.initOAuth({
      clientId: defaultClientId
    })
  }

  window.ui = ui
}
</script>
</body>

</html>
`

type swaggerConfig struct {
	URL                      string
	DocExpansion             string
	Title                    string
	Oauth2RedirectURL        template.JS
	DefaultModelsExpandDepth int
	DeepLinking              bool
	PersistAuthorization     bool
	Oauth2DefaultClientID    string
}

// Config stores ginSwagger configuration variables.
type Config struct {
	// The url pointing to API definition (normally swagger.json or swagger.yaml). Default is `doc.json`.
	URL                      string
	DocExpansion             string
	InstanceName             string
	Title                    string
	DefaultModelsExpandDepth int
	DeepLinking              bool
	PersistAuthorization     bool
	Oauth2DefaultClientID    string
}

func (config Config) toSwaggerConfig() swaggerConfig {
	return swaggerConfig{
		URL:                      config.URL,
		DeepLinking:              config.DeepLinking,
		DocExpansion:             config.DocExpansion,
		DefaultModelsExpandDepth: config.DefaultModelsExpandDepth,
		Oauth2RedirectURL: "`${window.location.protocol}//${window.location.host}$" +
			"{window.location.pathname.split('/').slice(0, window.location.pathname.split('/').length - 1).join('/')}" +
			"/oauth2-redirect.html`",
		Title:                 config.Title,
		PersistAuthorization:  config.PersistAuthorization,
		Oauth2DefaultClientID: config.Oauth2DefaultClientID,
	}
}
