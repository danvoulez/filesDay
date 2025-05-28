/**
 * Validador institucional (trecho relevante)
 * Reforça: type só pode ser preenchido por LLM
 */
export function validateAgainstTemplate(minicontrato: any, source: string = "manual"): string[] {
  const errors: string[] = [];
  // ...validação padrão...
  if ("type" in minicontrato && source !== "llm") {
    errors.push("Campo 'type' só pode ser atribuído por sistema LLM autorizado");
  }
  return errors;
}