import type { LogLine } from './types';

export function parseLogLine(line: string): LogLine | null {
  try {
    const obj = JSON.parse(line);
    if (typeof obj === 'object' && obj.id && obj.did) {
      return obj as LogLine;
    }
    return null;
  } catch (e) {
    return null;
  }
}