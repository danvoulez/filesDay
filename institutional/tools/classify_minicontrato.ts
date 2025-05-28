/**
 * Módulo: classify_minicontrato.ts
 * Só LLM pode atribuir o campo "type" ao minicontrato.
 */
import { Minicontrato } from "../types/minicontrato";

export async function classifyMinicontratoLLM(minicontrato: Minicontrato): Promise<{
  type: string;
  confidence: number;
  reasoning?: string;
  origin: string;
}> {
  // Chamada à LLM externa ou serviço institucional
  // Exemplo de payload:
  // {
  //   prompt: `Classifique semanticamente este minicontrato: ${JSON.stringify(minicontrato)}`
  // }
  // Retorno simulado:
  if (minicontrato.did === "pesar_produto") {
    return {
      type: "pesagem",
      confidence: 0.97,
      reasoning: "Campo did é pesar_produto e this contém peso_kg",
      origin: "llm:gpt-4o"
    };
  }
  // ... outros mapeamentos ou chamada real à LLM
  return { type: "", confidence: 0, origin: "llm:unknown" };
}