#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

// Configuration via environment variables
const DATA_DIR = process.env.DATA_DIR || './data';
const OUTPUT_FILE = process.env.OUTPUT_FILE || 'report.md';

function readJsonFiles(directory) {
    const results = [];

    try {
        const files = fs.readdirSync(directory);

        for (const file of files) {
            if (path.extname(file) === '.json') {
                const filePath = path.join(directory, file);
                try {
                    const content = fs.readFileSync(filePath, 'utf8');
                    const data = JSON.parse(content);
                    results.push(data);
                } catch (error) {
                    console.error(`Error reading file ${file}:`, error.message);
                }
            }
        }
    } catch (error) {
        console.error(`Error reading directory ${directory}:`, error.message);
        process.exit(1);
    }

    return results;
}

function formatTime(timeString) {
    // Parse time format like "1m31.157110501s" and round to 2 decimal places
    const match = timeString.match(/(?:(\d+)m)?(?:(\d+(?:\.\d+)?)s)?/);
    if (!match) return timeString;

    const minutes = parseInt(match[1] || 0);
    const seconds = parseFloat(match[2] || 0);

    if (minutes > 0) {
        return `${minutes}m${seconds.toFixed(2)}s`;
    } else {
        return `${seconds.toFixed(2)}s`;
    }
}

function generateMarkdownReport(data) {
    let markdown = `# 🌺🍕 Hawaiian test Results\n\n`;
    markdown += `Generated on: ${new Date().toLocaleString('en-US')}\n\n`;

    if (data.length === 0) {
        markdown += `No data found.\n`;
        return markdown;
    }

    // Table header
    markdown += `| Model | Score | Quantization Format | Temperature | Top P | Total Time |\n`;
    markdown += `|--------|-------|---------------------------|-------------|-------|-------------|\n`;

    // Table rows
    for (const item of data) {
        const modelName = item.model_id;
        const temperature = item.temperature?.toFixed(2) || 'N/A';
        const topP = item.top_p?.toFixed(2) || 'N/A';
        const totalTime = formatTime(item.total_time);
        markdown += `| ${modelName} | ${item.score} | ${item.quantization_format} | ${temperature} | ${topP} | ${totalTime} |\n`;
    }

    // Statistics
    markdown += `\n## Statistics\n\n`;
    markdown += `- **Total number of tested models**: ${data.length}\n`;
    markdown += `- **Average score**: ${(data.reduce((sum, item) => sum + item.score, 0) / data.length).toFixed(2)}\n`;
    markdown += `- **Maximum score**: ${Math.max(...data.map(item => item.score))}\n`;
    markdown += `- **Minimum score**: ${Math.min(...data.map(item => item.score))}\n`;


    return markdown;
}


function main() {
    console.log(`📊 Generating report from: ${DATA_DIR}`);
    console.log(`📄 Output file: ${OUTPUT_FILE}`);

    const data = readJsonFiles(DATA_DIR);
    console.log(`📋 ${data.length} JSON file(s) found`);

    const report = generateMarkdownReport(data);

    try {
        fs.writeFileSync(OUTPUT_FILE, report, 'utf8');
        console.log(`✅ Report generated successfully: ${OUTPUT_FILE}`);
    } catch (error) {
        console.error(`❌ Error writing file:`, error.message);
        process.exit(1);
    }
}

// Display help if requested
if (process.argv.includes('--help') || process.argv.includes('-h')) {
    console.log(`
Usage: node generate-report.js

Environment variables:
  DATA_DIR     Directory containing JSON files (default: ./data)
  OUTPUT_FILE  Output markdown filename (default: report.md)

Examples:
  node generate-report.js
  DATA_DIR=./data OUTPUT_FILE=my-report.md node generate-report.js
    `);
    process.exit(0);
}

main();