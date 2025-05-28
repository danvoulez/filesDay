import type { LogLine } from './types';
import { writeLogLine } from './writer';

export function createBaseLog(who: string, did: string, thisObj: any, parent_id?: string): LogLine {
  return {
    id: crypto.randomUUID(),
    tenant: "main",
    who,
    did,
    this: thisObj,
    when: new Date().toISOString(),
    confirmed_by: "system",
    if_ok: "execute",
    if_doubt: "review",
    if_not: "cancel",
    status: "pending",
    valid: true
  };
}