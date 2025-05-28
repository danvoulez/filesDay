import type { RequestHandler } from '@sveltejs/kit';
import fs from 'fs';
import { parseLogLine } from '$lib/logline/parser';
import { writeLogLine } from '$lib/logline/writer';

export const GET: RequestHandler = async ({ url }) => {
  const tenant = url.searchParams.get('tenant');
  const lines = fs.readFileSync(`static/logs/logline.${tenant}.jsonl`, 'utf-8')
    .split('\n').filter(Boolean).map(parseLogLine);
  return new Response(JSON.stringify(lines));
};

export const POST: RequestHandler = async ({ request }) => {
  const log = await request.json();
  writeLogLine(log);
  return new Response(JSON.stringify({ status: "ok" }));
};