import { LogLine } from "../types/logline";

// Função para enviar LogLine ao backend
export async function submitLogLine(logLine: LogLine) {
  const res = await fetch("/api/logline", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(logLine),
  });
  if (!res.ok) throw new Error("Erro ao registrar logline");
  return await res.json();
}