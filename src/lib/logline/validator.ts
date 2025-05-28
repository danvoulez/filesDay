import type { LogLine } from './types';

export function validateLogLine(log: LogLine): { valid: boolean; errors: string[] } {
  const errors: string[] = [];
  if (!log.id) errors.push('id missing');
  if (!log.tenant) errors.push('tenant missing');
  if (!log.who) errors.push('who missing');
  if (!log.did) errors.push('did missing');
  if (!log.when) errors.push('when missing');
  if (!log.status) errors.push('status missing');
  return { valid: errors.length === 0, errors };
}