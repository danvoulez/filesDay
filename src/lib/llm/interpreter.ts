import type { LLMResult } from './types';

export async function callLLM(text: string): Promise<LLMResult> {
  const provider = process.env.LLM_PROVIDER || 'openai';
  const apiKey = process.env.LLM_API_KEY || '';
  if (!apiKey) throw new Error('Missing LLM_API_KEY');
  // Chamada fictícia para LLM
  return {
    logline: { who: "llm", did: "interpret", this: text, status: "valid" }
  };
}