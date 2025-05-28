/**
 * Reconciliador de minicontratos sem type.
 * Só preenche se a origem for LLM.
 */
import { classifyMinicontratoLLM } from "./classify_minicontrato";
import { getUnclassifiedMinicontratos, patchMinicontrato } from "./db";

export async function reconcileTypes() {
  const ghosts = await getUnclassifiedMinicontratos();
  for (const minicontrato of ghosts) {
    const result = await classifyMinicontratoLLM(minicontrato);
    if (result.type && result.confidence > 0.85) {
      await patchMinicontrato(minicontrato.id, {
        type: result.type,
        type_origin: result.origin,
        type_timestamp: new Date().toISOString()
      });
    }
  }
}