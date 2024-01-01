#!/usr/bin/env node

import parse from 'yargs-parser';
import Papa from 'papaparse';
import { getDefaultPeriod, generateFeedback, FeedbackType } from '../index.js';

function parsePeriod(periodStr: string): [Date, Date] | null {
	if (!periodStr) {
		return null;
	}
	const [start, end] = periodStr.split('-').map((date: string) => new Date(date.trim()));
	return [start, end];
}

// Parse command-line arguments using yargs-parser
const args = process.argv.slice(2);
const parsedArgs = parse(args);

// Extract arguments
const { type, t, count, c, period, p, help, h } = parsedArgs;

// Display help
if (help || h) {
	console.error('Usage: mockgen --type <type> --count <count> --period <period>');
	process.exit(0);
}

// Validate and extract arguments
const feedbackTypes: FeedbackType[] = Array.isArray(type || t) ? type || t : [type || t];
const feedbackCount = parseInt(count || c, 10) || 10;
const feedbackPeriod = parsePeriod(period || p) || getDefaultPeriod();

// Validate mandatory arguments
if (!feedbackTypes.length) {
	feedbackTypes.push('text');
}
if (!feedbackCount || !feedbackPeriod) {
	console.error('Usage: mockgen --type <type> --count <count> --period <period>');
	process.exit(1);
}
// Validate feedback type
const validFeedbackTypes = ['text', 'csat', 'cnps'];
for (const feedbackType of feedbackTypes) {
	if (!validFeedbackTypes.includes(feedbackType)) {
		console.error(`Invalid feedback type. Please use one of: ${validFeedbackTypes.join(', ')}`);
		process.exit(1);
	}
}
// Validate feedback count
if (isNaN(feedbackCount) || feedbackCount < 1) {
	console.error('Invalid feedback count. Please use a number.');
	process.exit(1);
}
// Validate feedback period
if (isNaN(feedbackPeriod[0].getTime())) {
	console.error('Invalid start date format. Please use the date format (YYYY-MM-DD).');
	process.exit(1);
}
if (isNaN(feedbackPeriod[1].getTime())) {
	console.error('Invalid end date format. Please use the date format (YYYY-MM-DD).');
	process.exit(1);
}

// Generate mock feedbacks
const feedbacks = [];
for (const feedbackType of feedbackTypes) {
	for (let i = 0; i < Math.ceil(feedbackCount / feedbackTypes.length); i++) {
		const record = generateFeedback(feedbackType, feedbackPeriod);
		const csvrow: Record<string, unknown> = {
			type: record.type,
			time: record.time.toISOString(),
			cnps: 'cnps' in record.data ? record.data.cnps : '',
			csat: 'csat' in record.data ? record.data.csat : '',
			text: record.data.text || ''
		};
		for (const [key, value] of Object.entries(record.meta)) {
			csvrow[`meta.${key}`] = value;
		}
		for (const [key, value] of Object.entries(record.user)) {
			csvrow[`user.${key}`] = value;
		}
		feedbacks.push(csvrow);
	}
}

// Convert feedbacks to CSV
console.log(Papa.unparse(feedbacks));
