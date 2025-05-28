import fs from 'fs';
import { createClient } from '@supabase/supabase-js';
import type { LogLine } from './types';

const SUPABASE_URL = process.env.SUPABASE_URL || '';
const SUPABASE_ANON_KEY = process.env.SUPABASE_ANON_KEY || '';

export function writeLogLine(log: LogLine) {
  // Escreve no arquivo local
  fs.appendFileSync('static/logs/logline.voulezvous.jsonl', JSON.stringify(log) + '\n');
  // Escreve no Supabase (exemplo básico)
  if (SUPABASE_URL && SUPABASE_ANON_KEY) {
    const supabase = createClient(SUPABASE_URL, SUPABASE_ANON_KEY);
    supabase.from('logline').insert([log]);
  }
}