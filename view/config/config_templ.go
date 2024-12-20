// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package configView

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "general_spider_controll_panel/view/layout"

func Main(title string, urls []string) templ.Component {
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
			templ_7745c5c3_Err = layout.LeftNavbar("Config").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto p-4 bg-white rounded-lg shadow-md mt-8\"><h1 class=\"text-2xl font-bold mb-4\">Configuration Dashboard</h1><div id=\"modal\" class=\"bg-gray-600 bg-opacity-75 overflow-y-auto h-full w-full fixed inset-0 flex items-center justify-center opacity-0 transition-opacity duration-300 z-50 hidden\"><div class=\"bg-white rounded-lg shadow-xl w-full max-w-md mx-4\"><div class=\"border-b px-4 py-2\"><h3 class=\"text-lg font-semibold text-gray-900\">Scrapy Crawl Preview</h3></div><div class=\"p-4\"><div class=\"mt-2\"><div class=\"mb-4\"><label for=\"preview-url\" class=\"block text-sm font-medium text-gray-700\">URL to Test</label> <input type=\"text\" name=\"base-url\" id=\"preview-url\" class=\"w-full p-2 border border-gray-300 rounded-md\" placeholder=\"Enter URL to test\"></div><div class=\"mb-4\"><label for=\"section-select\" class=\"block text-sm font-medium text-gray-700\">Select @section</label> <select id=\"section-select\" name=\"section-select\" class=\"mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm\"></select></div><button id=\"run-preview\" class=\"w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-blue-600 text-base font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:w-auto sm:text-sm\" hx-post=\"?action=test-config\" hx-include=\"#preview-config, #preview-url\" hx-target=\"#config-runner-container\">Run Preview</button></div><input type=\"hidden\" id=\"preview-config\" name=\"jsonData\"><div class=\"mt-2\"><p class=\"text-sm text-gray-500 mb-4\">This is a preview of a single item from your Scrapy crawl.</p><pre id=\"config-runner-container\" class=\"bg-gray-100 p-4 rounded-md overflow-auto max-h-[400px] text-sm\"><div class=\"animate-pulse space-y-2\"><div class=\"h-4 bg-gray-200 rounded w-3/4\"></div><div class=\"h-4 bg-gray-200 rounded\"></div><div class=\"h-4 bg-gray-200 rounded\"></div><div class=\"h-4 bg-gray-200 rounded w-5/6\"></div><div class=\"h-4 bg-gray-200 rounded w-1/2\"></div></div></pre></div></div><div class=\"border-t px-4 py-2 flex justify-end\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, hidePreviewModal())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 templ.ComponentScript = hidePreviewModal()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline\">Close</button></div></div></div><form hx-post=\"?action=save-config\" hx-include=\"#json-data, #name, #description, #base-url\"><input name=\"base-url\" id=\"base-url\" class=\"hidden\" value=\"\"><div class=\"mb-4\"><div class=\"relative\" id=\"dropdown-container\"><button id=\"dropdown-toggle\" type=\"button\" class=\"w-full px-4 py-2 text-left bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500\"><span id=\"selected-option\">Select a website...</span> <svg class=\"w-5 h-5 ml-2 -mr-1 absolute right-2 top-3\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z\" clip-rule=\"evenodd\"></path></svg></button><div id=\"dropdown-menu\" class=\"absolute z-10 w-full mt-1 bg-white shadow-lg rounded-md hidden\"><div class=\"p-2\"><input type=\"text\" id=\"custom-url-input\" placeholder=\"Enter custom URL...\" value=\"asdasd\" class=\"w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500\"></div><ul id=\"url-list\" class=\"max-h-60 overflow-auto\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, url := range urls {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"px-4 py-2 hover:bg-gray-100 cursor-pointer flex items-center\"><input value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(url)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/config/config.templ`, Line: 74, Col: 28}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"hidden base-url-input\"> <svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5 mr-2 text-gray-500\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg><div><div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(url)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/config/config.templ`, Line: 79, Col: 21}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"text-sm text-gray-500\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(url)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/config/config.templ`, Line: 80, Col: 51}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></div></div><div class=\"mb-4\"><label for=\"name\" class=\"block font-semibold mb-1\">Name</label> <input id=\"name\" name=\"name\" type=\"text\" required placeholder=\"Enter the config name : config-v1\" class=\"w-full p-2 border border-gray-300 rounded-md\"></div><div class=\"mb-4\"><label for=\"configType\" class=\"block font-semibold mb-1\">Config type</label> <select id=\"configType\" name=\"configType\" required class=\"w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500\"><option value=\"\">Choose an config type</option> <option value=\"marketplace\">Marketplace</option> <option value=\"news\">News</option> <option value=\"forum\">Forum</option></select></div><div class=\"mb-4\"><label for=\"description\" class=\"block font-semibold mb-1\">Description (Opsional)</label> <textarea id=\"description\" name=\"description\" rows=\"4\" class=\"block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500\" placeholder=\"What is this config about ?\"></textarea></div><div class=\"flex gap-4 mb-4\"><button type=\"button\" id=\"add-config\" class=\"px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600\">Add Field</button> <label class=\"px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 cursor-pointer\">Load from Disk <input id=\"load-config\" type=\"file\" accept=\".json\" class=\"hidden\"></label> <button type=\"button\" id=\"download-config\" class=\"px-4 py-2 bg-purple-500 text-white rounded-md hover:bg-purple-600\">Download Configuration</button> <button type=\"submit\" id=\"deploy-config\" class=\"px-4 py-2 bg-purple-500 text-white rounded-md hover:bg-purple-600\">Deploy Configuration</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, showPreviewModal())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"button\" id=\"test-config\" onClick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 templ.ComponentScript = showPreviewModal()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"px-4 py-2 bg-purple-500 text-white rounded-md hover:bg-purple-600 disabled:bg-purple-300 disabled:cursor-not-allowed disabled:opacity-50\">Test Configuration</button></div><span id=\"version-display\" class=\"ml-4 text-sm text-gray-500 mb-2\">Version : 1.1</span><div id=\"config-container\" class=\"space-y-2\"></div><div class=\"mt-4\"><div class=\"mt-8\"><h2 class=\"text-xl font-bold mb-2\">JSON Preview</h2><pre id=\"json-preview\" class=\"bg-gray-800 text-white p-4 rounded-md overflow-x-auto\"></pre></div><input type=\"hidden\" id=\"json-data\" name=\"jsonData\"></div></form><br><br><br><br><br></div></div><script src=\"/public/config.js\"></script> <script>\n            // document.body.addEventListener('htmx:beforeRequest', function(event) {\n            //     if (event.target && event.target.id == \"test-config\") {\n            //         document.getElementById('test-config').disabled = true;\n            //         document.getElementById('test-config').textContent = 'Submitting...';\n            //         const modal = document.getElementById('modal');\n            //         modal.classList.remove('hidden');\n            //         setTimeout(() => {\n            //             modal.classList.remove('opacity-0');\n            //             modalContent.classList.remove('-translate-y-full', 'scale-95', 'opacity-0');\n            //         }, 50);\n            //     }\n            // });\n            document.querySelector('button').addEventListener('click', () => {\n                const preContent = document.getElementById('json-preview').innerText;\n                document.getElementById('json-data').value = preContent;\n            });\n    </script>")
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

func showPreviewModal() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_showPreviewModal_ebf4`,
		Function: `function __templ_showPreviewModal_ebf4(){const modal = document.getElementById('modal');
	const sectionSelect = document.getElementById('section-select');

    modal.classList.remove('hidden');

	sectionSelect.innerHTML = '';
    const sections = getSections();
	const option = document.createElement('option');
    option.value = "main";
    option.textContent = "main";
    sectionSelect.appendChild(option);
    sections.forEach(section => {
        const option = document.createElement('option');
        option.value = section.key;
        option.textContent = section.key;
        sectionSelect.appendChild(option);
    });

    setTimeout(() => {
        modal.classList.remove('opacity-0');
    }, 50);
}`,
		Call:       templ.SafeScript(`__templ_showPreviewModal_ebf4`),
		CallInline: templ.SafeScriptInline(`__templ_showPreviewModal_ebf4`),
	}
}

func hidePreviewModal() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_hidePreviewModal_5d19`,
		Function: `function __templ_hidePreviewModal_5d19(){const modal = document.getElementById('modal');

    modal.classList.add('opacity-0');
    setTimeout(() => {
        modal.classList.add('hidden');
    }, 300);
}`,
		Call:       templ.SafeScript(`__templ_hidePreviewModal_5d19`),
		CallInline: templ.SafeScriptInline(`__templ_hidePreviewModal_5d19`),
	}
}

var _ = templruntime.GeneratedTemplate
