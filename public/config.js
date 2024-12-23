let configs = [{ id: '1', type: 'xpath', value: '', children: [], key: '' }];
let baseUrl = '';
let draggedElement = null;
let originalParent = null;
let originalNextSibling = null;
let version = '1.0';
let hasChanged = false;

function renderConfigs() {
    const container = document.getElementById('config-container');
    container.innerHTML = '';
    renderConfigItems(configs, container);
    updateJsonPreview();
    saveToLocalStorage();
}

function renderConfigItems(items, container) {
    items.forEach(config => {
        const configElement = createConfigElement(config);
        container.appendChild(configElement);
    });
}

function createConfigElement(config) {
    const wrapper = document.createElement('div');
    wrapper.className = 'config-item bg-gray-50 p-2 rounded-md border border-gray-200';
    wrapper.dataset.id = config.id;
    wrapper.dataset.type = config.type;
    wrapper.dataset.key = config.key || '';
    wrapper.dataset.value = config.value || '';
    wrapper.dataset.tag = config.tag || '';
    wrapper.dataset.optional = config.optional.toString();
    wrapper.dataset.dataType = config.dataType || 'str';
    wrapper.draggable = true;

    const content = document.createElement('div');
    content.className = 'flex items-center space-x-2 overflow-x-auto';
    content.innerHTML = `
                <div class="cursor-move text-gray-500 hover:text-gray-700">≡</div>
                <select class="p-1 border border-gray-300 rounded-md">
                    ${['xpath', '_list', '_pagination', '_loop', '@section', 'constraint'].map(type => `<option value="${type}" ${config.type === type ? 'selected' : ''}>${type}</option>`).join('')}
                </select>
                <select class="tag-select p-1 border border-gray-300 rounded-md ${['_loop', '@section'].includes(config.type) ? '' : 'hidden'}">
                    <option value="">Select tag</option>
                    <option value="root" ${config.tag === 'root' ? 'selected' : ''}>root</option>
                    <option value="global" ${config.tag === 'global' ? 'selected' : ''}>global</option>
                    <option value="parent" ${config.tag === 'parent' ? 'selected' : ''}>parent</option>
                </select>
                <input placeholder="Enter key" value="${config.key || ''}" class="p-1 border border-gray-300 rounded-md flex-grow ${['_pagination', '_list'].includes(config.type) ? 'hidden' : ''}">
                <input placeholder="Enter value" value="${config.value || ''}" class="p-1 border border-gray-300 rounded-md flex-grow ${['@section'].includes(config.type) ? 'hidden' : ''}">
                <select class="data-type-select p-1 border border-gray-300 rounded-md ${config.type === 'xpath' ? '' : 'hidden'}">
                    <option value="str" ${config.dataType === 'str' ? 'selected' : ''}>string</option>
                    <option value="int" ${config.dataType === 'int' ? 'selected' : ''}>integer</option>
                    <option value="list" ${config.dataType === 'list' ? 'selected' : ''}>list</option>
                    <option value="timestamp" ${config.dataType === 'timestamp' ? 'selected' : ''}>timestamp</option>
                    <option value="concat-str" ${config.dataType === 'concat-str' ? 'selected' : ''}>concat string</option>
                </select>
                <button type="button" class="add-child px-2 py-1 bg-blue-500 text-white rounded-md hover:bg-blue-600 ${['xpath', 'constraint', '_pagination'].includes(config.type) ? 'hidden' : ''}">+</button>
                <button type="button" class="remove-config px-2 py-1 bg-red-500 text-white rounded-md hover:bg-red-600">-</button>
                <div class="flex items-center optional-continer ${config.type === 'xpath' ? '' : 'hidden'}">
                    <input type="checkbox" class="optional-toggle mr-1" ${config.optional ? 'checked' : ''}>
                    <span class="text-sm">Optional</span>
                </div>
            `;
    wrapper.appendChild(content);

    const typeSelect = content.querySelector('select');
    const tagSelect = content.querySelector('.tag-select');
    const keyInput = content.querySelector('input[placeholder="Enter key"]');
    const valueInput = content.querySelector('input[placeholder="Enter value"]');
    const dataTypeSelect = content.querySelector('.data-type-select');
    const addButton = content.querySelector('button.add-child');
    const removeButton = content.querySelector('button.remove-config');
    const optionalToggle = content.querySelector('.optional-toggle');

    typeSelect.onchange = (e) => updateConfig(wrapper, 'type', e.target.value);
    tagSelect.onchange = (e) => updateConfig(wrapper, 'tag', e.target.value);
    keyInput.onchange = (e) => updateConfig(wrapper, 'key', e.target.value);
    valueInput.onchange = (e) => updateConfig(wrapper, 'value', e.target.value);
    dataTypeSelect.onchange = (e) => updateConfig(wrapper, 'dataType', e.target.value);
    if (addButton) addButton.onclick = () => addChild(config.id);
    removeButton.onclick = () => removeConfig(config.id);
    if (optionalToggle) optionalToggle.onchange = (e) => updateConfig(wrapper, 'optional', e.target.checked);

    wrapper.addEventListener('dragstart', (e) => dragStart(e, config.id));
    wrapper.addEventListener('dragend', dragEnd);
    wrapper.addEventListener('dragover', dragOver);
    wrapper.addEventListener('drop', (e) => drop(e, config.id));

    if (config.children.length > 0 && config.type !== 'xpath') {
        const childrenContainer = document.createElement('div');
        childrenContainer.className = 'ml-4 mt-2 space-y-2';
        renderConfigItems(config.children, childrenContainer);
        wrapper.appendChild(childrenContainer);
    }

    return wrapper;
}

function updateConfig(element, field, value) {
    element.dataset[field] = value;
    if (field === 'type') {
        const keyInput = element.querySelector('input[placeholder="Enter key"]');
        const valueInput = element.querySelector('input[placeholder="Enter value"]');
        const addButton = element.querySelector('button.add-child');
        const tagSelect = element.querySelector('.tag-select');
        const dataTypeSelect = element.querySelector('.data-type-select');
        const optionalToggle = element.querySelector('.optional-toggle');
        const optionalContiner = element.querySelector('.optional-continer');
        const optionalLabel = optionalToggle ? optionalToggle.nextElementSibling : null;

        if (['_pagination', '_list'].includes(value)) {
            keyInput.classList.add('hidden');
            valueInput.classList.remove('hidden');
        } else if (['@section'].includes(value)) {
            keyInput.classList.remove('hidden');
            valueInput.classList.add('hidden');
        } else {
            keyInput.classList.remove('hidden');
            valueInput.classList.remove('hidden');
        }

        if (['_loop', '@section'].includes(value)) {
            tagSelect.classList.remove('hidden');
        } else {
            tagSelect.classList.add('hidden');
        }

        if (value === 'xpath') {
            dataTypeSelect.classList.remove('hidden');
            if (optionalToggle && optionalLabel) {
                optionalContiner.classList.remove('hidden');
                optionalToggle.classList.remove('hidden');
                optionalLabel.classList.remove('hidden');
            }
        } else {
            dataTypeSelect.classList.add('hidden');
            if (optionalToggle && optionalLabel) {
                optionalContiner.classList.add('hidden');
                optionalToggle.classList.add('hidden');
                optionalLabel.classList.add('hidden');
            }
        }

        if (addButton) {
            addButton.classList.toggle('hidden', ['xpath', 'constraint', '_pagination'].includes(value));
        }
    }
    rebuildConfigsFromDOM();
    updateJsonPreview();
    incrementVersion();
    saveToLocalStorage();
}

function rebuildConfigsFromDOM() {
    const container = document.getElementById('config-container');
    configs = buildConfigsFromElements(container.children);
}

function buildConfigsFromElements(elements) {
    const result = [];
    for (const element of elements) {
        if (element.classList.contains('config-item')) {
            const config = {
                id: element.dataset.id,
                type: element.dataset.type,
                key: element.dataset.key,
                value: element.dataset.value,
                tag: element.dataset.tag,
                optional: element.dataset.type === 'xpath' ? element.dataset.optional === 'true' : false,
                dataType: element.dataset.dataType,
                children: []
            };

            const childrenContainer = element.querySelector('div.ml-4');
            if (childrenContainer) {
                config.children = buildConfigsFromElements(childrenContainer.children);
            }
            result.push(config);
        }
    }
    return result;
}

function addChild(parentId) {
    const parentElement = document.querySelector(`.config-item[data-id="${parentId}"]`);
    if (parentElement) {
        const newConfig = { id: Date.now().toString(), type: 'xpath', value: '', children: [], key: '', optional: false, dataType: 'str' };
        const newElement = createConfigElement(newConfig);
        let childrenContainer = parentElement.querySelector('div.ml-4');
        if (!childrenContainer) {
            childrenContainer = document.createElement('div');
            childrenContainer.className = 'ml-4 mt-2 space-y-2';
            parentElement.appendChild(childrenContainer);
        }
        // childrenContainer.appendChild(newElement);
        childrenContainer.insertBefore(newElement, childrenContainer.firstChild);
        rebuildConfigsFromDOM();
        updateJsonPreview();
        incrementVersion();
        saveToLocalStorage();
    }
}

function removeConfig(id) {
    const element = document.querySelector(`.config-item[data-id="${id}"]`);
    if (element) {
        element.remove();
        rebuildConfigsFromDOM();
        updateJsonPreview();
        incrementVersion();
        saveToLocalStorage();
    }
}

function dragStart(e, id) {
    draggedElement = e.target.closest('.config-item');
    originalParent = draggedElement.parentNode;
    originalNextSibling = draggedElement.nextElementSibling;
    e.dataTransfer.setData('text/plain', id);
    requestAnimationFrame(() => {
        draggedElement.classList.add('dragging');
    });
}

function dragEnd(e) {
    if (draggedElement) {
        draggedElement.classList.remove('dragging');
        document.querySelectorAll('.drag-over').forEach(el => el.classList.remove('drag-over'));
        if (!document.body.contains(draggedElement)) {
            if (originalNextSibling) {
                originalParent.insertBefore(draggedElement, originalNextSibling);
            } else {
                originalParent.appendChild(draggedElement);
            }
        }
        rebuildConfigsFromDOM();
        updateJsonPreview();
        saveToLocalStorage();
    }
    draggedElement = null;
    originalParent = null;
    originalNextSibling = null;
}

function dragOver(e) {
    e.preventDefault();
    const closestConfigItem = e.target.closest('.config-item');
    if (closestConfigItem && closestConfigItem !== draggedElement) {
        const rect = closestConfigItem.getBoundingClientRect();
        const midpoint = (rect.top + rect.bottom) / 2;
        document.querySelectorAll('.drag-over').forEach(el => el.classList.remove('drag-over'));
        if (e.clientY < midpoint) {
            closestConfigItem.classList.add('drag-over');
            closestConfigItem.parentNode.insertBefore(draggedElement, closestConfigItem);
        } else {
            closestConfigItem.classList.add('drag-over');
            closestConfigItem.parentNode.insertBefore(draggedElement, closestConfigItem.nextSibling);
        }
    }
}

function drop(e, droppedOnId) {
    e.preventDefault();
    const draggedId = e.dataTransfer.getData('text');
    if (draggedId !== droppedOnId) {
        const droppedOnElement = document.querySelector(`.config-item[data-id="${droppedOnId}"]`);
        if (droppedOnElement && !droppedOnElement.contains(draggedElement)) {
            droppedOnElement.parentNode.insertBefore(draggedElement, droppedOnElement.nextSibling);
        }
    }
    draggedElement.classList.remove('dragging');
    document.querySelectorAll('.drag-over').forEach(el => el.classList.remove('drag-over'));
    draggedElement = null;
    rebuildConfigsFromDOM();
    updateJsonPreview();
    incrementVersion();
    saveToLocalStorage();
}

function cleanupConfigs(configs) {
    const structure = {};
    configs.forEach(config => {
        config.key = config.key.replaceAll(/ /g, `_`);
        config.value = config.value.replaceAll(`"`, `'`)
        if (config.type === 'constraint') {
            structure[config.key] = {
                value: config.value,
                type: config.type
            };
        }
        else if (config.type === 'xpath') {
            structure[config.key] = {
                value: config.value,
                optional: config.optional,
                type: config.dataType
            };
        } else if (config.type === '_list') {
            structure[config.type] = {
                _element: config.value,
                ...cleanupConfigs(config.children)
            };
        } else if (config.type === '_loop') {
            structure[config.type] = {
                _element: config.value,
                _key: config.key,
                _tag: config.tag,
                ...cleanupConfigs(config.children)
            };
        } else if (config.type === '_pagination') {
            structure[config.type] = config.value;
            cleanupConfigs(config.children);
        } else if (config.type === '@section') {
            if (!config.key.startsWith('@')) {
                config.key = '@' + config.key
            }
            structure[config.key] = {
                _tag: config.tag,
                ...cleanupConfigs(config.children)
            };
        }
        document.querySelector(`div[data-id='${config.id}']`).querySelector('input[placeholder="Enter key"]').value = config.key
        document.querySelector(`div[data-id='${config.id}']`).querySelector('input[placeholder="Enter value"]').value = config.value
    });
    return structure;
}

function findSection(sectionKey) {
    function searchSection(items) {
        for (const item of items) {
            if (item.type === '@section' && item.key === sectionKey) {
                return cleanupConfigs([item]);
            }
            if (item.children && item.children.length > 0) {
                const result = searchSection(item.children);
                if (result) return result;
            }
        }
        return null;
    }

    return searchSection(configs);
}

function getSections() {
    const sections = [];
    function findSections(items) {
        items.forEach(item => {
            if (item.type === '@section') {
                sections.push({
                    key: item.key,
                    config: JSON.stringify(cleanupConfigs([item]), null, 2)
                });
            }
            if (item.children.length > 0) {
                findSections(item.children);
            }
        });
    }
    findSections(configs);
    return sections;
}

function incrementVersion() {
    if (!hasChanged) {
        const [major, minor] = version.split('.').map(Number);
        version = `${major}.${minor + 1}`;
        document.getElementById('version-display').textContent = `Version: ${version}`;
        hasChanged = true;
        saveToLocalStorage();
    }
}

function updateJsonPreview() {
    const cleanConfigs = cleanupConfigs(configs);
    const fullConfig = {
        base_url: baseUrl,
        structure: cleanConfigs
    };
    document.getElementById('json-preview').textContent = JSON.stringify(fullConfig, null, 2);
    document.getElementById('json-data').value = JSON.stringify(fullConfig, null, 2);
    document.getElementById('preview-config').value = JSON.stringify(fullConfig, null, 2)
}

function saveToLocalStorage() {
    localStorage.setItem('configDashboardState', JSON.stringify({
        configs: configs,
        baseUrl: baseUrl,
        version: version,
        hasChanged: hasChanged
    }));
}

function loadFromLocalStorage() {
    const savedState = localStorage.getItem('configDashboardState');
    if (savedState) {
        const parsedState = JSON.parse(savedState);
        configs = parsedState.configs;
        baseUrl = parsedState.baseUrl;
        version = parsedState.version || '1.0';
        hasChanged = parsedState.hasChanged || false;
        document.getElementById('base-url').value = baseUrl;
        document.getElementById('version-display').textContent = `Version: ${version}`;
        renderConfigs();
    }
}

function addNewConfig() {
    const newConfig = { id: Date.now().toString(), type: 'xpath', value: '', children: [], key: '', optional: false, dataType: 'str' };
    const newElement = createConfigElement(newConfig);
    document.getElementById('config-container').appendChild(newElement);
    rebuildConfigsFromDOM();
    updateJsonPreview();
    incrementVersion();
    saveToLocalStorage();
}

function downloadConfig() {
    const cleanConfigs = cleanupConfigs(configs);
    const fullConfig = {
        base_url: baseUrl,
        structure: cleanConfigs
    };
    const blob = new Blob([JSON.stringify(fullConfig, null, 2)], { type: 'application/json' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'config.json';
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
}

function loadConfig(e) {
    const file = e.target.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = (event) => {
            try {
                const loadedConfig = JSON.parse(event.target.result);
                baseUrl = loadedConfig.base_url || '';
                document.getElementById('base-url').value = baseUrl;
                configs = convertStructureToConfigs(loadedConfig.structure);
                version = '1.0';
                hasChanged = false;
                document.getElementById('version-display').textContent = `Version: ${version}`;
                renderConfigs();
            } catch (error) {
                console.error('Error parsing JSON:', error);
                alert('Error loading configuration. Please make sure it\'s a valid JSON file.');
            }
        };
        reader.readAsText(file);
    }
}

function convertStructureToConfigs(structure) {
    const result = [];
    for (const [key, value] of Object.entries(structure)) {
        if (typeof value === 'object' && 'value' in value && 'type' in value) {
            if (value.type === 'constraint') {
                result.push({
                    id: Date.now().toString() + Math.random().toString(36).substr(2, 9),
                    type: 'constraint',
                    key: key,
                    value: value.value,
                    optional: false,
                    dataType: 'str',
                    children: []
                });
            } else if ('optional' in value) {
                result.push({
                    id: Date.now().toString() + Math.random().toString(36).substr(2, 9),
                    type: 'xpath',
                    key: key,
                    value: value.value,
                    optional: value.optional,
                    dataType: value.type,
                    children: []
                });
            }
        } else if (key === '_pagination') {
            result.push({
                id: Date.now().toString() + Math.random().toString(36).substr(2, 9),
                type: '_pagination',
                key: value,
                value: value,
                children: [],
                optional: false,
                dataType: 'str'
            });
        } else if (typeof value === 'object') {
            const newConfig = {
                id: Date.now().toString() + Math.random().toString(36).substr(2, 9),
                type: key.startsWith('@') ? '@section' : key,
                key: key.startsWith('@') ? key : '',
                value: value._element || '',
                tag: value._tag || '',
                children: [],
                optional: false,
                dataType: 'str'
            };
            if (key === '_loop') {
                newConfig.key = value._key;
            } else if (key === '_list') {
                newConfig.key = value._element;
            }
            const childConfigs = convertStructureToConfigs(
                Object.fromEntries(
                    Object.entries(value).filter(([k]) => !['_element', '_key', '_tag'].includes(k))
                )
            );
            newConfig.children = childConfigs;
            result.push(newConfig);
        }
    }
    return result;
}

document.getElementById('add-config').addEventListener('click', addNewConfig);
document.getElementById('download-config').addEventListener('click', downloadConfig);
document.getElementById('load-config').addEventListener('change', loadConfig);
document.getElementById('base-url').addEventListener('change', (e) => {
    baseUrl = e.target.value;
    updateJsonPreview();
    saveToLocalStorage();
});
document.getElementById('config-container').addEventListener('dragover', (e) => {
    e.preventDefault();
});

function updatePreviewConfig() {
    const url = document.getElementById('preview-url').value;
    const key = document.getElementById('section-select').value;
    let fullConfig;

    if (!url.startsWith("http://") && !url.startsWith("https://" && typeof url === "string" && url.trim() !== "")) {
        document.getElementById('preview-url').value = "http://" + url;
    }
    
    if (key !== "main") {
        const cleanConfigs = findSection(key);
        fullConfig = {
            base_url: url,
            structure: cleanConfigs
        };
    } else {
        const cleanConfigs = cleanupConfigs(configs);
        fullConfig = {
            base_url: url,
            structure: cleanConfigs
        };
    }

    document.getElementById('preview-config').value = JSON.stringify(fullConfig, null, 2);
}

document.getElementById('preview-url').addEventListener('input', updatePreviewConfig);
document.getElementById('section-select').addEventListener('change', updatePreviewConfig);

loadFromLocalStorage();

const customUrlInput = document.getElementById('custom-url-input');
const dropdownToggle = document.getElementById('dropdown-toggle');
const dropdownMenu = document.getElementById('dropdown-menu');
const baseUrlInput = document.getElementById('base-url');
const urlList = document.getElementById('url-list');
const selectedOption = document.getElementById('selected-option');

function updateSelectedUrl(url, closeDropdown = false) {
    baseUrlInput.value = url;
    customUrlInput.value = url;
    selectedOption.textContent = url;
    if (closeDropdown) {
        dropdownMenu.classList.add('hidden');
    }
}

customUrlInput.addEventListener('input', function(e) {
    updateSelectedUrl(e.target.value);
});

customUrlInput.addEventListener('keydown', function(e) {
    if (e.key === 'Enter') {
        updateSelectedUrl(e.target.value, true);
    }
});

urlList.addEventListener('click', function(e) {
    const listItem = e.target.closest('li');
    if (listItem) {
        const baseUrlInputInList = listItem.querySelector('.base-url-input');
        if (baseUrlInputInList) {
            updateSelectedUrl(baseUrlInputInList.value, true);
        }
    }
});

dropdownToggle.addEventListener('click', () => {
    dropdownMenu.classList.toggle('hidden');
    if (!dropdownMenu.classList.contains('hidden')) {
        customUrlInput.focus();
    }
});

document.addEventListener('click', function(e) {
    if (!dropdownToggle.contains(e.target) && !dropdownMenu.contains(e.target)) {
        dropdownMenu.classList.add('hidden');
    }
});

// Ensure the custom input is updated if the base URL is changed programmatically
baseUrlInput.addEventListener('change', function(e) {
    updateSelectedUrl(e.target.value);
});

if (configs.length === 0) {
    renderConfigs();
}