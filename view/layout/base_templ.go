// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "general_spider_controll_panel/types"

func Base(title string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"Secure and reliable file hosting service. Upload, organize, and share your documents, images, videos, and more. Sign up now to keep your files always within reach.\"><meta name=\"keywords\" content=\"file hosting, file sharing, cloud storage, data storage, secure file hosting, filekeeper, drive, mega\"><meta name=\"author\" content=\"Filekeeper\"><link href=\"/public/output.css\" rel=\"stylesheet\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/layout/base.templ`, Line: 15, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><script src=\"https://unpkg.com/htmx.org@2.0.3\"></script><link href=\"https://cdn.jsdelivr.net/npm/remixicon@3.5.0/fonts/remixicon.css\" rel=\"stylesheet\"></head><body><div id=\"toastContainer\" class=\"fixed top-5 right-5 space-y-2\"></div><div id=\"content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><script>\n\t\t\tfunction showToast(message, duration = 5000, type = 'info') {\n\t\t\t\tconst toastContainer = document.getElementById('toastContainer');\n\t\t\t\tconst toast = document.createElement('div');\n\t\t\t\ttoast.className = `text-white p-4 rounded-lg shadow-lg mb-4 transition-all duration-300 ease-in-out transform translate-x-full opacity-0 max-w-[90vw] ${getToastTypeClass(type)}`;\n\t\t\t\ttoast.style.width = 'max-content';\n\t\t\t\ttoast.innerHTML =\n\t\t\t\t\t`<div class=\"text-sm font-medium text-gray-900\">\n\t\t\t\t\t<div class=\"flex items-start\">\n\t\t\t\t\t\t<div class=\"flex-shrink-0 pt-0.5\">\n\t\t\t\t\t\t\t${getToastIcon(type)}\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div class=\"ml-3 flex-1 \">\n\t\t\t\t\t\t\t<p class=\"text-sm font-medium text-gray-900\" title=\"${message}\">\n\t\t\t\t\t\t\t${message}\n\t\t\t\t\t\t\t</p>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t</div>\n\t\t\t\t</div>\n\t\t\t\t`;\n\n\t\t\t\ttoastContainer.appendChild(toast);\n\n\t\t\t\ttoast.offsetHeight;\n\n\t\t\t\ttoast.classList.remove('translate-x-full', 'opacity-0');\n\n\t\t\t\tconst maxWidth = Math.min(400, toast.offsetWidth);\n\t\t\t\ttoast.style.width = `${maxWidth}px`;\n\n\t\t\t\tsetTimeout(() => {\n\t\t\t\t\ttoast.classList.add('translate-x-full', 'opacity-0');\n\t\t\t\t\tsetTimeout(() => {\n\t\t\t\t\t\ttoastContainer.removeChild(toast);\n\t\t\t\t\t}, 300);\n\t\t\t\t}, duration);\n\t\t\t}\n\t\t\tfunction getToastTypeClass(type) {\n\t\t\t\tswitch (type) {\n\t\t\t\t\tcase 'error':\n\t\t\t\t\t\treturn 'bg-red-50';\n\t\t\t\t\tcase 'success':\n\t\t\t\t\t\treturn 'bg-green-50';\n\t\t\t\t\tdefault:\n\t\t\t\t\t\treturn 'bg-blue-50';\n\t\t\t\t}\n\t\t\t}\n\n\t\t\tfunction getToastIcon(type) {\n\t\t\t\tswitch (type) {\n\t\t\t\t\tcase 'error':\n\t\t\t\t\t\treturn '<svg class=\"h-6 w-6 text-red-400\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\" /></svg>';\n\t\t\t\t\tcase 'success':\n\t\t\t\t\t\treturn '<svg class=\"h-6 w-6 text-green-400\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z\" /></svg>';\n\t\t\t\t\tdefault:\n\t\t\t\t\t\treturn '<svg class=\"h-6 w-6 text-blue-400\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\" /></svg>';\n\t\t\t\t}\n\t\t\t}\n\n\t\t\tdocument.body.addEventListener(\"htmx:afterRequest\", function (event) {\n\t\t\t\tif (event.detail.xhr.getResponseHeader('Content-Type')?.includes('application/json')) {\n\t\t\t\t\tconst jsonResponse = JSON.parse(event.detail.xhr.responseText);\n\t\t\t\t\tshowToast(\n\t\t\t\t\t\tjsonResponse.message, \n\t\t\t\t\t\tjsonResponse.timeout !== 0 ? jsonResponse.timeout : 5000, \n\t\t\t\t\t\tjsonResponse.type \n\t\t\t\t\t\t\t? jsonResponse.type \n\t\t\t\t\t\t\t: event.detail.failed \n\t\t\t\t\t\t\t\t? \"error\" \n\t\t\t\t\t\t\t\t: \"other\"\n\t\t\t\t\t);\n\t\t\t\t}\n\t\t\t\tif (!event.detail.xhr.responseText && event.detail.xhr.status === 0) {\n\t\t\t\t\tshowToast(\"No response from the server. Please try again later.\", 5000, 'error');\n\t\t\t\t\treturn;\n\t\t\t\t}\n\t\t});\n\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var navbarItems = []types.NavbarItems{
	types.NavbarItems{
		Key: "Home",
		Url: "/home",
	},
	types.NavbarItems{
		Key: "Spiders",
		Url: "/spiders",
	},
	types.NavbarItems{
		Key: "Config",
		Url: "/config",
	},
	types.NavbarItems{
		Key: "Deploy",
		Url: "/deploy",
	},
	types.NavbarItems{
		Key: "Proxies",
		Url: "/proxies",
	},
	types.NavbarItems{
		Key: "Kafka broker",
		Url: "/kafka/broker",
	},
}

func LeftNavbar(active string) templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-h-screen flex\"><input type=\"checkbox\" id=\"nav-toggle\" class=\"hidden peer\"> <label for=\"nav-toggle\" class=\"fixed top-4 left-4 z-50 p-2 bg-white rounded-md shadow-md cursor-pointer lg:hidden\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6 peer-checked:hidden\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5\"></path></svg> <svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6 hidden peer-checked:block\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M6 18L18 6M6 6l12 12\"></path></svg> <span class=\"sr-only\">Toggle navigation menu</span></label><div class=\"fixed inset-y-0 left-0 transform -translate-x-full peer-checked:translate-x-0 lg:translate-x-0 transition duration-200 ease-in-out z-30 lg:relative lg:inset-0\"><div class=\"w-64 bg-white border-r h-full flex flex-col overflow-y-auto sticky top-0\"><div class=\"flex h-16 items-center border-b px-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6 text-blue-600\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M12 12.75c1.148 0 2.278.08 3.383.237 1.037.146 1.866.966 1.866 2.013 0 3.728-2.35 6.75-5.25 6.75S6.75 18.728 6.75 15c0-1.046.83-1.867 1.866-2.013A24.204 24.204 0 0112 12.75zm0 0c2.883 0 5.647.508 8.207 1.44a23.91 23.91 0 01-1.152 6.06M12 12.75c-2.883 0-5.647.508-8.208 1.44.125 2.104.52 4.136 1.153 6.06M12 12.75a2.25 2.25 0 002.248-2.354M12 12.75a2.25 2.25 0 01-2.248-2.354M12 8.25c.995 0 1.971-.08 2.922-.236.403-.066.74-.358.795-.762a3.778 3.778 0 00-.399-2.25M12 8.25c-.995 0-1.97-.08-2.922-.236-.402-.066-.74-.358-.795-.762a3.734 3.734 0 01.4-2.253M12 8.25a2.25 2.25 0 00-2.248 2.146M12 8.25a2.25 2.25 0 012.248 2.146M8.683 5a6.032 6.032 0 01-1.155-1.002c.07-.63.27-1.222.574-1.747m.581 2.749A3.75 3.75 0 0115.318 5m0 0c.427-.283.815-.62 1.155-.999a4.471 4.471 0 00-.575-1.752M4.921 6a24.048 24.048 0 00-.392 3.314c1.668.546 3.416.914 5.223 1.082M19.08 6c.205 1.08.337 2.187.392 3.314a23.882 23.882 0 01-5.223 1.082\"></path></svg> <span class=\"ml-2 text-lg font-semibold\">Scrapy Dashboard</span></div><nav class=\"p-4 flex-grow\"><h2 class=\"mb-2 px-4 text-lg font-semibold tracking-tight\">Overview</h2><ul class=\"space-y-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range navbarItems {
			if item.Key == active {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 templ.SafeURL = templ.SafeURL(item.Url)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex items-center rounded-lg px-4 py-2 text-gray-700 hover:bg-gray-400 bg-gray-300\"><span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(item.Key)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/layout/base.templ`, Line: 160, Col: 26}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 templ.SafeURL = templ.SafeURL(item.Url)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex items-center rounded-lg px-4 py-2 text-gray-700 hover:bg-gray-100\"><span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(item.Key)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/layout/base.templ`, Line: 166, Col: 26}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></nav></div></div></div><style>\n        #nav-toggle:checked ~ div {\n            transform: translateX(0);\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
