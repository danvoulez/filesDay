import React from "react";
import { LogLine } from "../types/logline";

// Sugestões dummy - trocar por IA real depois
const suggestions = [
  { did: "aceitar_termos", label: "Aceitar Termos" },
  { did: "pesar_produto", label: "Pesar Produto" },
  { did: "fallback_ativado", label: "Fallback Ativado" },
];

export default function Suggestions({ data }: { data: Partial<LogLine> }) {
  return (
    <div className="suggestions">
      <span>Sugestões:</span>
      {suggestions.map((s) => (
        <button key={s.did}>{s.label}</button>
      ))}
    </div>
  );
}