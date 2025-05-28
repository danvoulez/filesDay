import { createClient } from '@supabase/supabase-js';
import type { LogLine } from './types';
import fs from 'fs';
import { writeLogLine } from './writer';

const SUPABASE_URL = process.env.SUPABASE_URL || '';
const SUPABASE_ANON_KEY = process.env.SUPABASE_ANON_KEY || '';

export function auditLog(log: LogLine) {
  // Auditoria básica: salva log em outro arquivo e insere no supabase
  fs.appendFileSync('static/logs/audit.jsonl', JSON.stringify(log) + '\n');
  if (SUPABASE_URL && SUPABASE_ANON_KEY) {
    const supabase = createClient(SUPABASE_URL, SUPABASE_ANON_KEY);
    supabase.from('audit').insert([log]);
  }
}