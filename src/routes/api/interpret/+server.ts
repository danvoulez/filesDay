import type { RequestHandler } from '@sveltejs/kit';
import { callLLM } from '$lib/llm/interpreter';

export const POST: RequestHandler = async ({ request }) => {
  const { text } = await request.json();
  const result = await callLLM(text);
  return new Response(JSON.stringify(result));
};