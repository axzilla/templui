// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/axzilla/goilerplate/internals/config"

func BaseLayout() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full\"><head><title>Goilerplate - Modern UI Components for Go & Templ</title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><style>\n                /* Hide the page initially */\n                .page-loader { display: none; }\n                html.pre-load .page-loader { display: block; }\n                html.pre-load body { display: none; }\n            </style><script>\n                // Immediately set theme and remove pre-load class\n                (function() {\n                    document.documentElement.classList.add('pre-load');\n                    var theme = localStorage.getItem('theme') || 'light';\n                    document.documentElement.classList.toggle('dark', theme === 'dark');\n                    document.addEventListener('DOMContentLoaded', function() {\n                        document.documentElement.classList.remove('pre-load');\n                    });\n                })();\n            </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if config.AppConfig.GoEnv == "production" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Plausible Analytics --> <script defer data-domain=\"goilerplate.com\" src=\"https://plausible.axeladrian.com/js/script.js\"></script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Favicon --><link rel=\"icon\" href=\"/assets/img/gopher.svg\" type=\"image/x-icon\"><!-- Tailwind CSS (local) --><link href=\"/assets/css/output.css\" rel=\"stylesheet\"><!-- Alpine.js --><script defer src=\"https://cdn.jsdelivr.net/npm/@alpinejs/focus@3.x.x/dist/cdn.min.js\"></script><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\"></script><!-- Highlight.js --><link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/base16/woodland.min.css\"><script src=\"https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/languages/go.min.js\"></script><script>hljs.highlightAll();</script></head><body x-data=\"{\n                theme: localStorage.getItem(&#39;theme&#39;) || &#39;light&#39;,\n                sidebarOpen: false,\n                toggleTheme() {\n                    this.theme = this.theme === &#39;dark&#39; ? &#39;light&#39; : &#39;dark&#39;;\n                    localStorage.setItem(&#39;theme&#39;, this.theme);\n                    document.documentElement.classList.toggle(&#39;dark&#39;, this.theme === &#39;dark&#39;);\n                }\n            }\" class=\"h-full flex flex-col transition-colors duration-200\" :class=\"{&#39;bg-white text-black&#39;: theme === &#39;light&#39;, &#39;text-white&#39;: theme === &#39;dark&#39;}\"><div class=\"page-loader\">Loading...</div><div class=\"flex flex-col min-h-screen\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate