// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package domainListView

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	"general_spider_controll_panel/view/layout"
)

func Main(title string) templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex h-screen\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = layout.LeftNavbar("Spiders").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex-1 overflow-auto\"><div class=\"p-8\"><div class=\"flex justify-between items-center mb-6\"><h1 class=\"text-3xl font-bold\">Spiders</h1><button id=\"refreshButton\" class=\"bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded flex items-center\"><i class=\"ri-refresh-line mr-2\"></i> Refresh</button></div><div class=\"bg-white shadow-md rounded-lg\"><div class=\"px-4 py-3 border-b border-gray-200\"><h2 class=\"text-lg font-semibold text-gray-800\">Domain List</h2><div class=\"flex justify-end\"><a href=\"/config\" id=\"addDomainButton\" class=\"bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded flex items-center transition-colors duration-300\"><svg class=\"w-5 h-5 mr-2\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6v6m0 0v6m0-6h6m-6 0H6\"></path></svg> Add Domain</a></div></div><ul hx-get=\"?action=get-domains\" hx-trigger=\"load\" class=\"divide-y divide-gray-200 overflow-auto\"><li class=\"bg-white overflow-hidden mb-4 rounded-lg shadow-sm border border-gray-200 animate-pulse\"><div class=\"p-4\"><div class=\"flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4\"><div class=\"flex items-center gap-3 w-full sm:w-auto\"><div class=\"w-5 h-5 bg-gray-200 rounded-full\"></div><div class=\"flex flex-col sm:flex-row sm:items-center gap-2 min-w-0 flex-grow\"><div class=\"h-5 bg-gray-200 rounded w-40 sm:w-64\"></div><div class=\"h-5 w-16 bg-gray-200 rounded\"></div></div></div><div class=\"h-9 w-full sm:w-28 bg-gray-200 rounded-md\"></div></div></div><div class=\"px-4 py-3 bg-gray-50 border-t border-gray-100\"><div class=\"h-4 bg-gray-200 rounded w-40 mb-2\"></div><div class=\"flex flex-col sm:flex-row sm:items-center justify-between text-sm gap-2\"><div class=\"h-4 bg-gray-200 rounded w-32\"></div><div class=\"h-4 bg-gray-200 rounded w-24\"></div></div></div></li><li class=\"bg-white overflow-hidden mb-4 rounded-lg shadow-sm border border-gray-200 animate-pulse\"><div class=\"p-4\"><div class=\"flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4\"><div class=\"flex items-center gap-3 w-full sm:w-auto\"><div class=\"w-5 h-5 bg-gray-200 rounded-full\"></div><div class=\"flex flex-col sm:flex-row sm:items-center gap-2 min-w-0 flex-grow\"><div class=\"h-5 bg-gray-200 rounded w-40 sm:w-64\"></div><div class=\"h-5 w-16 bg-gray-200 rounded\"></div></div></div><div class=\"h-9 w-full sm:w-28 bg-gray-200 rounded-md\"></div></div></div><div class=\"px-4 py-3 bg-gray-50 border-t border-gray-100\"><div class=\"h-4 bg-gray-200 rounded w-40 mb-2\"></div><div class=\"flex flex-col sm:flex-row sm:items-center justify-between text-sm gap-2\"><div class=\"h-4 bg-gray-200 rounded w-32\"></div><div class=\"h-4 bg-gray-200 rounded w-24\"></div></div></div></li><li class=\"bg-white overflow-hidden mb-4 rounded-lg shadow-sm border border-gray-200 animate-pulse\"><div class=\"p-4\"><div class=\"flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4\"><div class=\"flex items-center gap-3 w-full sm:w-auto\"><div class=\"w-5 h-5 bg-gray-200 rounded-full\"></div><div class=\"flex flex-col sm:flex-row sm:items-center gap-2 min-w-0 flex-grow\"><div class=\"h-5 bg-gray-200 rounded w-40 sm:w-64\"></div><div class=\"h-5 w-16 bg-gray-200 rounded\"></div></div></div><div class=\"h-9 w-full sm:w-28 bg-gray-200 rounded-md\"></div></div></div><div class=\"px-4 py-3 bg-gray-50 border-t border-gray-100\"><div class=\"h-4 bg-gray-200 rounded w-40 mb-2\"></div><div class=\"flex flex-col sm:flex-row sm:items-center justify-between text-sm gap-2\"><div class=\"h-4 bg-gray-200 rounded w-32\"></div><div class=\"h-4 bg-gray-200 rounded w-24\"></div></div></div></li><li class=\"bg-white overflow-hidden mb-4 rounded-lg shadow-sm border border-gray-200 animate-pulse\"><div class=\"p-4\"><div class=\"flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4\"><div class=\"flex items-center gap-3 w-full sm:w-auto\"><div class=\"w-5 h-5 bg-gray-200 rounded-full\"></div><div class=\"flex flex-col sm:flex-row sm:items-center gap-2 min-w-0 flex-grow\"><div class=\"h-5 bg-gray-200 rounded w-40 sm:w-64\"></div><div class=\"h-5 w-16 bg-gray-200 rounded\"></div></div></div><div class=\"h-9 w-full sm:w-28 bg-gray-200 rounded-md\"></div></div></div><div class=\"px-4 py-3 bg-gray-50 border-t border-gray-100\"><div class=\"h-4 bg-gray-200 rounded w-40 mb-2\"></div><div class=\"flex flex-col sm:flex-row sm:items-center justify-between text-sm gap-2\"><div class=\"h-4 bg-gray-200 rounded w-32\"></div><div class=\"h-4 bg-gray-200 rounded w-24\"></div></div></div></li></ul></div></div></div></div><script>\n\t\tfunction parseStartTime(text) {\n\t\t\tlet days = 0, hours = 0, minutes = 0, seconds = 0;\n\n\t\t\tconst daysMatch = text.match(/(\\d+)\\s*days?/);\n\t\t\tconst hoursMatch = text.match(/(\\d+)\\s*hours?/);\n\t\t\tconst minutesMatch = text.match(/(\\d+)\\s*minutes?/);\n\t\t\tconst secondsMatch = text.match(/(\\d+)\\s*seconds?/);\n\n\t\t\tif (daysMatch) days = parseInt(daysMatch[1]);\n\t\t\tif (hoursMatch) hours = parseInt(hoursMatch[1]);\n\t\t\tif (minutesMatch) minutes = parseInt(minutesMatch[1]);\n\t\t\tif (secondsMatch) seconds = parseInt(secondsMatch[1]);\n\n\t\t\treturn days * 86400 + hours * 3600 + minutes * 60 + seconds;\n\t\t}\n\n\t\tfunction updateCountups() {\n\t\t\tconst countupElements = document.querySelectorAll('.countup');\n\t\t\tcountupElements.forEach((element) => {\n\t\t\t\tconst startingText = element.textContent;\n\t\t\t\t\tconsole.log(startingText)\n\t\t\t\tif (startingText === \"Last crawled: Now\") {\n\t\t\t\t\treturn\n\t\t\t\t} else if (startingText === \"Last crawled: Never been run before :)\") {\n\t\t\t\t\treturn\n\t\t\t\t}\n\t\t\t\tlet startTime = parseStartTime(startingText);\n\t\t\t\tfunction updateElementCountup() {\n\t\t\t\tstartTime++;\n\t\t\t\tlet days = Math.floor(startTime / 86400);\n\t\t\t\tlet hours = Math.floor((startTime % 86400) / 3600);\n\t\t\t\tlet minutes = Math.floor((startTime % 3600) / 60);\n\t\t\t\tlet seconds = startTime % 60;\n\t\t\t\tlet formattedTime = '';\n\t\t\t\tif (days > 0) {\n\t\t\t\t\tformattedTime += `${days} days `;\n\t\t\t\t}\n\t\t\t\tif (hours > 0) {\n\t\t\t\t\tformattedTime += `${hours} hours `;\n\t\t\t\t}\n\t\t\t\tif (minutes > 0 || hours > 0 || days > 0) {\n\t\t\t\t\tformattedTime += `${minutes} minutes `;\n\t\t\t\t}\n\t\t\t\tformattedTime += `${seconds} seconds`;\n\t\t\t\telement.textContent = `Last crawled: ${formattedTime.trim()} ago`;\n\t\t\t\t}\n\t\t\t\tsetInterval(updateElementCountup, 1000);\n\t\t\t\tupdateElementCountup();\n\t\t\t});\n\t\t}\n\n\t\tdocument.body.addEventListener('htmx:afterSwap', function(event) {\n\t\t\tupdateCountups();\n\t\t});\n\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base(title).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func GetDomains(domains []*types.DomainStats) templ.Component {
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
		if len(domains) == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"noDomains\" class=\"px-6 py-12 text-center\"><svg class=\"mx-auto h-12 w-12 text-gray-400\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg><h3 class=\"mt-2 text-sm font-medium text-gray-900\">No domains added yet</h3><p class=\"mt-1 text-sm text-gray-500\">Get started by adding your first domain.</p><div class=\"mt-6\"><a href=\"/config\" class=\"inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500\"><svg class=\"-ml-1 mr-2 h-5 w-5\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6v6m0 0v6m0-6h6m-6 0H6\"></path></svg> Add Your First Domain</a></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			for _, domain := range domains {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"bg-white overflow-hidden\"><div class=\"p-4\"><div class=\"flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4\"><div class=\"flex items-center gap-3 w-full sm:w-auto\"><svg class=\"w-5 h-5 text-gray-400 flex-shrink-0\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg><div class=\"flex flex-col sm:flex-row sm:items-center gap-2 min-w-0\"><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 templ.SafeURL = templ.SafeURL(fmt.Sprintf("spiders/%s", domain.Domain))
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-gray-900 font-medium truncate\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(domain.Domain)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/domains.templ`, Line: 211, Col: 134}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div></div><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 templ.SafeURL = templ.SafeURL(fmt.Sprintf("spiders/%s", domain.Domain))
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"inline-flex items-center justify-center px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 w-full sm:w-auto\">View Details</a></div></div><div class=\"px-4 py-3 bg-gray-50 border-t border-gray-100\"><p class=\"text-sm text-gray-500 mb-2 countup\">Last crawled: ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(domain.LastCrawled)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/domains.templ`, Line: 223, Col: 84}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"text-sm text-gray-500 mb-2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(utils.IntToString(domain.ActiveSpider))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/domains.templ`, Line: 224, Col: 82}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" Active Spiders</p><p class=\"text-sm text-gray-500 mb-2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(utils.IntToString(domain.PendingSpider))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/domains.templ`, Line: 225, Col: 83}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" Pending Spiders</p><p class=\"text-sm text-gray-500 mb-2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(utils.IntToString(domain.FinishedSpider))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/domains.templ`, Line: 226, Col: 84}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" Finished Spiders</p><p class=\"text-sm text-gray-500 mb-2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var11 string
				templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(utils.IntToString(domain.ScheduledSpider))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/domains.templ`, Line: 227, Col: 85}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" Scheduled Spiders</p></div></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
