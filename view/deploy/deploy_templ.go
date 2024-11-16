// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.786
package deployView

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "general_spider_controll_panel/view/layout"

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
			templ_7745c5c3_Err = layout.LeftNavbar("Deploy").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto p-4 max-w-2xl\"><div class=\"mb-6\"><a href=\"/spiders\" class=\"text-blue-600 hover:underline flex items-center\"><i class=\"ri-arrow-left-line mr-2\"></i> Back to Spider List</a></div><h1 class=\"text-3xl font-bold mb-6\">Deploy Spider</h1><form id=\"deployForm\" class=\"space-y-6\"><div><label for=\"base-url-select\" class=\"block text-sm font-medium text-gray-700 mb-1\">Select Base URL</label> <select id=\"base-url-select\" name=\"base_url\" class=\"mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm\" required><option value=\"\">Select a base URL</option></select></div><div><label for=\"config-select\" class=\"block text-sm font-medium text-gray-700 mb-1\">Select Configuration</label> <select id=\"config-select\" name=\"config\" class=\"mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm\" required><option value=\"\">Select a configuration</option></select></div><div id=\"spider-settings\" class=\"space-y-4\"><div><label for=\"download_delay\" class=\"block text-sm font-medium text-gray-700 mb-1\">Download Delay (seconds)</label> <input type=\"number\" id=\"download_delay\" name=\"download_delay\" class=\"mt-1 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md\" placeholder=\"1\" min=\"0\" step=\"0.1\"></div><div><label for=\"concurrent_requests\" class=\"block text-sm font-medium text-gray-700 mb-1\">Concurrent Requests</label> <input type=\"number\" id=\"concurrent_requests\" name=\"concurrent_requests\" class=\"mt-1 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md\" placeholder=\"16\" min=\"1\"></div><div><label for=\"additional_settings\" class=\"block text-sm font-medium text-gray-700 mb-1\">Additional Settings (JSON)</label> <textarea id=\"additional_settings\" name=\"additional_settings\" rows=\"4\" class=\"mt-1 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md\" placeholder=\"{&#34;ROBOTSTXT_OBEY&#34;: true, &#34;USER_AGENT&#34;: &#34;MyBot/1.0&#34;}\"></textarea></div></div><button type=\"submit\" class=\"w-full inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500\">Deploy Spider</button></form><div class=\"mt-6 p-4 bg-yellow-100 rounded-md flex items-start\"><i class=\"ri-alert-line text-yellow-700 mr-2 mt-0.5\"></i><p class=\"text-sm text-yellow-700\">Deploying a spider will consume resources and may incur costs. Make sure you have the necessary permissions and understand the implications before proceeding.</p></div></div></div><script>\n        let spiderConfigs = [];\n\n        fetch('/api/spider-configs')\n            .then(response => response.json())\n            .then(data => {\n                spiderConfigs = data;\n                populateBaseUrls();\n            })\n            .catch(error => console.error('Error fetching spider configs:', error));\n\n        function populateBaseUrls() {\n            const baseUrlSelect = document.getElementById('base-url-select');\n            baseUrlSelect.innerHTML = '<option value=\"\">Select a base URL</option>';\n            spiderConfigs.forEach(config => {\n                const option = document.createElement('option');\n                option.value = config.base_url;\n                option.textContent = config.base_url;\n                baseUrlSelect.appendChild(option);\n            });\n        }\n\n        function populateConfigs(baseUrl) {\n            const configSelect = document.getElementById('config-select');\n            configSelect.innerHTML = '<option value=\"\">Select a configuration</option>';\n            const selectedConfig = spiderConfigs.find(config => config.base_url === baseUrl);\n            if (selectedConfig) {\n                selectedConfig.configs.forEach(configId => {\n                    const option = document.createElement('option');\n                    option.value = configId;\n                    option.textContent = `Config ${configId.slice(0, 8)}...`;\n                    configSelect.appendChild(option);\n                });\n            }\n        }\n\n        document.getElementById('base-url-select').addEventListener('change', (e) => {\n            populateConfigs(e.target.value);\n        });\n\n        document.getElementById('deployForm').addEventListener('submit', (e) => {\n            e.preventDefault();\n            const formData = new FormData(e.target);\n            const data = Object.fromEntries(formData.entries());\n            \n            try {\n                data.additional_settings = JSON.parse(data.additional_settings);\n            } catch (error) {\n                console.error('Invalid JSON in additional settings');\n                alert('Invalid JSON in additional settings. Please check and try again.');\n                return;\n            }\n\n            console.log('Deploying spider with settings:', data);\n            fetch('/api/deploy-spider', {\n                method: 'POST',\n                headers: {\n                    'Content-Type': 'application/json',\n                },\n                body: JSON.stringify(data),\n            })\n            .then(response => response.json())\n            .then(result => {\n                console.log('Deployment result:', result);\n                alert('Spider deployment initiated successfully!');\n            })\n            .catch(error => {\n                console.error('Error deploying spider:', error);\n                alert('Error deploying spider. Please try again.');\n            });\n        });\n    </script>")
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

var _ = templruntime.GeneratedTemplate
