// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package spiderView

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/view/layout"
)

func Main(title, tableType string) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex-1 overflow-auto\"><div class=\"p-8\"><div class=\"flex flex-col sm:flex-row justify-between items-start sm:items-center mb-6 space-y-4 sm:space-y-0\"><h1 class=\"text-2xl sm:text-3xl font-bold\">Spiders</h1><div class=\"flex flex-col sm:flex-row space-y-2 sm:space-y-0 sm:space-x-4 w-full sm:w-auto\"><a href=\"/deploy\" class=\"bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded transition-colors duration-300 flex items-center justify-center sm:justify-start\"><svg class=\"w-5 h-5 mr-2\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M12 6v6m0 0v6m0-6h6m-6 0H6\"></path></svg> Deploy Spider</a> <button hx-get=\"?action=get-spiders\" hx-target=\"#spiderList\" class=\"bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded transition-colors duration-300 flex items-center justify-center sm:justify-start\"><svg class=\"w-5 h-5 mr-2\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15\"></path></svg> Refresh</button></div></div><div class=\"mx-auto space-y-4\" id=\"spiders-container\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			switch tableType {
			case "schedule":
				templ_7745c5c3_Err = ChangeToScheduleTable().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			default:
				templ_7745c5c3_Err = ChangeToRunningTable().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div>")
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

func ChangeToRunningTable() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid grid-cols-2 gap-2 mb-6\" role=\"group\"><button class=\"w-full py-2 px-4 bg-blue-500 text-white font-semibold rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50\">Running Spiders</button> <button hx-get=\"?action=change-to-scheduled\" hx-target=\"#spiders-container\" hx-push-url=\"?type=schedule\" class=\"w-full py-2 px-4 bg-gray-200 text-gray-700 font-semibold rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50\">Scheduled Spiders</button></div><div class=\"bg-white shadow-md rounded-lg\"><div class=\"px-4 py-3 border-b border-gray-200\"><h2 class=\"text-lg font-semibold text-gray-800\">Active Spiders</h2></div><ul class=\"divide-y divide-gray-200 overflow-auto\" id=\"spiderList\" hx-get=\"?action=get-spiders\" hx-trigger=\"load\" hx-swap=\"innerHTML\"><li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><div class=\"h-4 w-4 rounded-full bg-gray-200 animate-pulse\"></div></div><div class=\"ml-3\"><div class=\"w-32 h-5 bg-gray-200 rounded animate-pulse\"></div><div class=\"mt-1 w-48 h-4 bg-gray-200 rounded animate-pulse\"></div></div></div><div class=\"flex space-x-2\"><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div></div></div></li><li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><div class=\"h-4 w-4 rounded-full bg-gray-200 animate-pulse\"></div></div><div class=\"ml-3\"><div class=\"w-32 h-5 bg-gray-200 rounded animate-pulse\"></div><div class=\"mt-1 w-48 h-4 bg-gray-200 rounded animate-pulse\"></div></div></div><div class=\"flex space-x-2\"><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div></div></div></li><li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><div class=\"h-4 w-4 rounded-full bg-gray-200 animate-pulse\"></div></div><div class=\"ml-3\"><div class=\"w-32 h-5 bg-gray-200 rounded animate-pulse\"></div><div class=\"mt-1 w-48 h-4 bg-gray-200 rounded animate-pulse\"></div></div></div><div class=\"flex space-x-2\"><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div></div></div></li></ul></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ChangeToScheduleTable() templ.Component {
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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid grid-cols-2 gap-2 mb-6\" role=\"group\"><button hx-get=\"?action=change-to-running\" hx-target=\"#spiders-container\" hx-push-url=\"?type=running\" class=\"w-full py-2 px-4 bg-gray-200 text-gray-700 font-semibold rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50\">Running Spiders</button> <button class=\"w-full py-2 px-4 bg-blue-500 text-white font-semibold rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50\">Scheduled Spiders</button></div><div class=\"bg-white shadow-md rounded-lg\"><div class=\"px-4 py-3 border-b border-gray-200\"><h2 class=\"text-lg font-semibold text-gray-800\">Active Spiders</h2></div><ul class=\"divide-y divide-gray-200 overflow-auto\" id=\"spiderList\" hx-get=\"?action=get-scheduled\" hx-trigger=\"load\" hx-swap=\"innerHTML\"><li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><div class=\"h-4 w-4 rounded-full bg-gray-200 animate-pulse\"></div></div><div class=\"ml-3\"><div class=\"w-32 h-5 bg-gray-200 rounded animate-pulse\"></div><div class=\"mt-1 w-48 h-4 bg-gray-200 rounded animate-pulse\"></div></div></div><div class=\"flex space-x-2\"><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div></div></div></li><li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><div class=\"h-4 w-4 rounded-full bg-gray-200 animate-pulse\"></div></div><div class=\"ml-3\"><div class=\"w-32 h-5 bg-gray-200 rounded animate-pulse\"></div><div class=\"mt-1 w-48 h-4 bg-gray-200 rounded animate-pulse\"></div></div></div><div class=\"flex space-x-2\"><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div></div></div></li><li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><div class=\"h-4 w-4 rounded-full bg-gray-200 animate-pulse\"></div></div><div class=\"ml-3\"><div class=\"w-32 h-5 bg-gray-200 rounded animate-pulse\"></div><div class=\"mt-1 w-48 h-4 bg-gray-200 rounded animate-pulse\"></div></div></div><div class=\"flex space-x-2\"><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div><div class=\"w-16 h-8 bg-gray-200 rounded animate-pulse\"></div></div></div></li></ul></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func GetScheduled(crons []*types.Cron) templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if len(crons) == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"px-6 py-12 text-center\"><svg class=\"mx-auto h-12 w-12 text-gray-400\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg><h3 class=\"mt-2 text-sm font-medium text-gray-900\">No spiders scheduled in corn yet</h3><p class=\"mt-1 text-sm text-gray-500\">Schedule a new spider to get started.</p><div class=\"mt-6\"><a href=\"/deploy\" type=\"button\" class=\"inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500\"><svg class=\"-ml-1 mr-2 h-5 w-5\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6v6m0 0v6m0-6h6m-6 0H6\"></path></svg> Schedule New Spider</a></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			for _, cron := range crons {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"p-3 sm:p-4\"><div class=\"flex flex-col sm:flex-row items-start sm:items-center justify-between border border-gray-200 rounded-lg hover:bg-gray-50\"><div class=\"flex items-start space-x-3 p-3 sm:p-4 w-full sm:w-auto\"><div class=\"h-2 w-2 bg-green-400 rounded-full flex-shrink-0 mt-1.5\" aria-hidden=\"true\"></div><div class=\"min-w-0 flex-1\"><div class=\"font-medium flex flex-col sm:flex-row sm:items-center gap-2 min-w-0 truncate\"><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 templ.SafeURL = templ.SafeURL(fmt.Sprintf("%s/schedule/%s", cron.Project, cron.ID))
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-gray-900 font-medium truncate\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(cron.ID)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 210, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div><div class=\"text-sm text-gray-500 mt-1\">Next run: ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(cron.NextRun)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 213, Col: 71}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"text-sm text-gray-500\">Last run: ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(cron.LastRun)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 214, Col: 66}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"text-sm font-medium text-indigo-600 mt-1 countdown\">Countdown: ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(cron.Countdown)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 215, Col: 98}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div><div class=\"p-3 sm:p-4 w-full sm:w-auto flex justify-end\"><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var11 templ.SafeURL = templ.SafeURL(fmt.Sprintf("%s/schedule/%s", cron.Project, cron.ID))
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var11)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"px-3 py-1 text-sm text-indigo-600 bg-indigo-50 rounded-md hover:bg-indigo-100 focus:outline-none focus:ring-2 focus:ring-indigo-500\">Details</a></div></div></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <script>\n\t\t\tfunction parseTime(text) {\n\t\t\t\tlet days = 0, hours = 0, minutes = 0, seconds = 0;\n\t\t\t\t\n\t\t\t\tconst daysMatch = text.match(/(\\d+)\\s*days?/);\n\t\t\t\tconst hoursMatch = text.match(/(\\d+)\\s*hours?/);\n\t\t\t\tconst minutesMatch = text.match(/(\\d+)\\s*minutes?/);\n\t\t\t\tconst secondsMatch = text.match(/(\\d+)\\s*seconds?/);\n\t\t\t\t\n\t\t\t\tif (daysMatch) days = parseInt(daysMatch[1]);\n\t\t\t\tif (hoursMatch) hours = parseInt(hoursMatch[1]);\n\t\t\t\tif (minutesMatch) minutes = parseInt(minutesMatch[1]);\n\t\t\t\tif (secondsMatch) seconds = parseInt(secondsMatch[1]);\n\n\t\t\t\treturn days * 86400 + hours * 3600 + minutes * 60 + seconds;\n\t\t\t}\n\n\t\t\tfunction updateCountdowns() {\n\t\t\t\tconst countdownElements = document.querySelectorAll('.countdown');\n\n\t\t\t\tcountdownElements.forEach((element) => {\n\t\t\t\t\tconst countdownText = element.textContent;\n\n\t\t\t\t\tlet timeRemaining = parseTime(countdownText);\n\n\t\t\t\t\tfunction updateElementCountdown() {\n\t\t\t\t\t\tlet days = Math.floor(timeRemaining / 86400);\n\t\t\t\t\t\tlet hours = Math.floor((timeRemaining % 86400) / 3600);\n\t\t\t\t\t\tlet minutes = Math.floor((timeRemaining % 3600) / 60);\n\t\t\t\t\t\tlet seconds = timeRemaining % 60;\n\n\t\t\t\t\t\tlet formattedTime = 'in ';\n\n\t\t\t\t\t\tif (days > 0) {\n\t\t\t\t\t\t\tformattedTime += `${days} days `;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tif (hours > 0) {\n\t\t\t\t\t\t\tformattedTime += `${hours} hours `;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tif (minutes > 0 || hours > 0 || days > 0) {\n\t\t\t\t\t\t\tformattedTime += `${minutes} minutes `;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tformattedTime += `${seconds} seconds`;\n\n\t\t\t\t\t\telement.textContent = formattedTime.trim();\n\n\t\t\t\t\t\tif (timeRemaining > 0) {\n\t\t\t\t\t\t\ttimeRemaining--;\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tclearInterval(countdownInterval);\n\t\t\t\t\t\t\tlocation.reload(true);\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\n\t\t\t\t\tlet countdownInterval = setInterval(updateElementCountdown, 1000);\n\n\t\t\t\t\tupdateElementCountdown();\n\t\t\t\t});\n\t\t\t}\n\n\t\t\tupdateCountdowns();\n\n\t\t\tdocument.body.addEventListener('htmx:afterSwap', function(event) {\n\t\t\t\tif (event.target.id === 'spiderList') {\n\t\t\t\t\tupdateCountdowns();\n\t\t\t\t}\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func GetSpider(spiders *types.ScrapydResponseGetingSpiders) templ.Component {
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
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if len(spiders.Running) == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"px-6 py-12 text-center\"><svg class=\"mx-auto h-12 w-12 text-gray-400\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg><h3 class=\"mt-2 text-sm font-medium text-gray-900\">No spiders running</h3><p class=\"mt-1 text-sm text-gray-500\">Get started by deploying a new spider.</p><div class=\"mt-6\"><a href=\"/deploy\" type=\"button\" class=\"inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500\"><svg class=\"-ml-1 mr-2 h-5 w-5\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6v6m0 0v6m0-6h6m-6 0H6\"></path></svg> Deploy New Spider</a></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		for _, runningSpider := range spiders.Running {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><span class=\"h-4 w-4 rounded-full bg-green-400 flex items-center justify-center\"><span class=\"h-2 w-2 rounded-full bg-green-600\"></span></span></div><div class=\"ml-3\"><p class=\"text-sm font-medium text-gray-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(runningSpider.Id)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 325, Col: 69}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"text-sm text-gray-500\">Active - Started at ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(runningSpider.StartTime)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 326, Col: 84}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div><div class=\"flex\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 templ.SafeURL = templ.SafeURL(fmt.Sprintf("%s/active/%v", runningSpider.Project, runningSpider.Id))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var15)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-sm bg-blue-100 hover:bg-blue-200 text-blue-800 font-semibold py-1 px-3 rounded-full transition duration-300 ease-in-out\">Details</a></div></div></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		for _, pendingSpider := range spiders.Pending {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><span class=\"h-4 w-4 rounded-full bg-yellow-400 flex items-center justify-center\"><span class=\"h-2 w-2 rounded-full bg-yellow-600\"></span></span></div><div class=\"ml-3\"><p class=\"text-sm font-medium text-gray-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(pendingSpider.Id)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 347, Col: 69}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"text-sm text-gray-500\">Pending deployment</p></div></div><div class=\"flex\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var17 templ.SafeURL = templ.SafeURL(fmt.Sprintf("%s/active/%v", pendingSpider.Project, pendingSpider.Id))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var17)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-sm bg-blue-100 hover:bg-blue-200 text-blue-800 font-semibold py-1 px-3 rounded-full transition duration-300 ease-in-out\">Details</a></div></div></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		for _, finishedSpider := range spiders.Finished {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"px-4 py-3 sm:px-6\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center\"><div class=\"flex-shrink-0\"><span class=\"h-4 w-4 rounded-full bg-gray-400 flex items-center justify-center\"><span class=\"h-2 w-2 rounded-full bg-gray-600\"></span></span></div><div class=\"ml-3\"><p class=\"text-sm font-medium text-gray-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 string
			templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(finishedSpider.Id)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 369, Col: 70}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"text-sm text-gray-500\">Finished - Completed at ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var19 string
			templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(finishedSpider.EndTime)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/spider/spider/spider.templ`, Line: 370, Col: 87}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div><div class=\"flex\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 templ.SafeURL = templ.SafeURL(fmt.Sprintf("%s/active/%v", finishedSpider.Project, finishedSpider.Id))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var20)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-sm bg-blue-100 hover:bg-blue-200 text-blue-800 font-semibold py-1 px-3 rounded-full transition duration-300 ease-in-out\">Details</a></div></div></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
