import type { RequestHandler } from '@sveltejs/kit';
import { validateLogLine } from '$lib/logline/validator';

export const POST: RequestHandler = async ({ request }) => {
  const log = await request.json();
  const result = validateLogLine(log);
  return new Response(JSON.stringify(result));
};