/* global process */

import { execSync } from 'node:child_process';
import { readFileSync, writeFileSync } from 'node:fs';
import YAML from 'yaml';

const args = process.argv.slice(2);
const openapiFilePath = args[0];
const openapiFileText = readFileSync(openapiFilePath, 'utf8');
const openapiFileYAML = YAML.parse(openapiFileText);

// Generate typescript types from openapi file
execSync(`openapi-typescript ${openapiFilePath} -o ./src/api.type.ts`);

// Parse the openapi file to identify groups and operations
const parsed = { groups: {} };
for (const path of Object.keys(openapiFileYAML.paths)) {
	const pathObject = openapiFileYAML.paths[path];
	for (const method of Object.keys(pathObject)) {
		const methodObject = pathObject[method];
		const [groupName, operationName] = methodObject.operationId.split('_');
		if (!parsed.groups[groupName]) {
			parsed.groups[groupName] = {};
		}
		const group = parsed.groups[groupName];
		const operation = (group[operationName] = { path, method });
		if (methodObject.requestBody) {
			operation.hasRequestBody = true;
		}
		if (methodObject.responses) {
			for (const statusCode of Object.keys(methodObject.responses)) {
				if (statusCode !== 204) {
					operation.hasResponseBody = true;
					operation.responseStatusCode = statusCode;
				}
			}
		}
		if (methodObject.parameters.length) {
			operation.hasPathParameters = true;
			operation.pathParameters = methodObject.parameters
				.filter((p) => p.in === 'path')
				.map((p) => p.name);
		}
	}
}

// Generate api code
const lines = [
	'import type { paths } from "./api.type.ts";',
	'import type { Fetcher } from "./fetcher.ts";'
];

// Generate api type aliases for each group and operation
for (const groupName of Object.keys(parsed.groups)) {
	const group = parsed.groups[groupName];
	lines.push(`\nexport type ${groupName} = {`);
	for (const operationName of Object.keys(group)) {
		lines.push(`${operationName}: {`);
		const operation = group[operationName];
		if (operation.hasRequestBody) {
			lines.push(
				`RequestBody: paths['${operation.path}']['${operation.method}']['requestBody']['content']['application/json']`
			);
		} else {
			lines.push(`RequestBody: void;`);
		}
		if (operation.hasResponseBody) {
			lines.push(
				`ResponseBody: paths['${operation.path}']['${operation.method}']['responses'][${operation.responseStatusCode}]['content']['application/json']`
			);
		} else {
			lines.push(`ResponseBody: void;`);
		}
		lines.push(`};`);
	}
	lines.push(`};`);
}

// Generate code to call these endpoints
for (const groupName of Object.keys(parsed.groups)) {
	const group = parsed.groups[groupName];
	lines.push(`\nexport const create${groupName}Endpoints = (f: Fetcher) => {`);
	lines.push(`return {`);
	for (const operationName of Object.keys(group)) {
		// function head
		lines.push(`${operationName}: async (`);
		const operation = group[operationName];
		if (operation.hasPathParameters) {
			for (const pathParam of operation.pathParameters) {
				lines.push(`${pathParam}: string,`);
			}
		}
		if (operation.hasRequestBody) {
			lines.push(`body: ${groupName}['${operationName}']['RequestBody'],`);
		}
		if (operation.hasResponseBody) {
			lines.push(`): Promise<${groupName}['${operationName}']['ResponseBody']> => {`);
		} else {
			lines.push(`) => {`);
		}
		// function body
		if (operation.hasRequestBody) {
			if (operation.hasResponseBody) {
				lines.push(
					`return f.fetch<${groupName}['${operationName}']['ResponseBody'], ${groupName}['${operationName}']['RequestBody']>(`
				);
			} else {
				lines.push(`return f.fetch<null, ${groupName}['${operationName}']['RequestBody']>(`);
			}
		} else if (operation.hasResponseBody) {
			lines.push(`return f.fetch<${groupName}['${operationName}']['ResponseBody']>(`);
		} else {
			lines.push(`return f.fetch<null>(`);
		}
		lines.push(`'${operation.method.toUpperCase()}',`);
		if (operation.hasPathParameters) {
			let formattedPath = operation.path;
			for (const pathParam of operation.pathParameters) {
				formattedPath = formattedPath.replace(`{${pathParam}}`, `\${${pathParam}}`);
			}
			lines.push(`\`${formattedPath}\`,`);
		} else {
			lines.push(`'${operation.path}',`);
		}
		if (operation.hasRequestBody) {
			lines.push(`body,`);
		}
		lines.push(`);`);
		lines.push(`},`);
	}
	lines.push(`};`);
	lines.push(`};`);
}

writeFileSync('./src/api.code.ts', lines.join('\n'));

// Format code
execSync(`yarn format`);
