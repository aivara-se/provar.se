#!/usr/bin/env node

import parse from 'yargs-parser';
import Papa from 'papaparse';
import { getDefaultPeriod, generateFeedback } from '../index.js';

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
const feedbackType = type || t || 'text';
const feedbackCount = parseInt(count || c, 10) || 10;
const feedbackPeriod = parsePeriod(period || p) || getDefaultPeriod();

// Validate mandatory arguments
if (!feedbackType || !feedbackCount || !feedbackPeriod) {
	console.error('Usage: mockgen --type <type> --count <count> --period <period>');
	process.exit(1);
}
// Validate feedback type
const validFeedbackTypes = ['text', 'csat', 'cnps'];
if (!validFeedbackTypes.includes(feedbackType)) {
	console.error(`Invalid feedback type. Please use one of: ${validFeedbackTypes.join(', ')}`);
	process.exit(1);
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
for (let i = 0; i < feedbackCount; i++) {
	const feedback = generateFeedback(feedbackType, feedbackPeriod);
	feedbacks.push({
		type: feedback.type,
		time: feedback.time.toISOString(),
		cnps: 'cnps' in feedback.data ? feedback.data.cnps : '',
		csat: 'csat' in feedback.data ? feedback.data.csat : '',
		text: feedback.data.text || ''
	});
}

// Convert feedbacks to CSV
console.log(Papa.unparse(feedbacks));
