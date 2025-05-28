import fs from 'fs';
import lz4 from 'lz4';

export function compressOldLogs(logsDir = 'static/logs') {
  const files = fs.readdirSync(logsDir).filter(f => f.endsWith('.jsonl'));
  for (const file of files) {
    const input = fs.createReadStream(`${logsDir}/${file}`);
    const output = fs.createWriteStream(`${logsDir}/${file}.lz4`);
    input.pipe(lz4.createEncoderStream()).pipe(output);
  }
}