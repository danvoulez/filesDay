import React from "react";
import { useFlipFlow } from "../hooks/useFlipFlow";
import { LogLine } from "../types/logline";
import { submitLogLine } from "../services/api";
import Suggestions from "./Suggestions";
import FeedbackToast from "./FeedbackToast";

export default function FlipForm() {
  const { step, data, nextStep, reset } = useFlipFlow();
  const [loading, setLoading] = React.useState(false);
  const [feedback, setFeedback] = React.useState("");

  async function handleSubmit() {
    setLoading(true);
    try {
      await submitLogLine(data as LogLine);
      setFeedback("Registro feito com sucesso!");
      reset();
    } catch {
      setFeedback("Erro ao registrar. Tente novamente.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="flip-form">
      {step === 0 && (
        <input
          placeholder="Quem fez?"
          value={data.who || ""}
          onChange={(e) => nextStep({ who: e.target.value })}
        />
      )}
      {step === 1 && (
        <input
          placeholder="O que foi feito?"
          value={data.did || ""}
          onChange={(e) => nextStep({ did: e.target.value })}
        />
      )}
      {step === 2 && (
        <input
          placeholder="Sobre o quê? (opcional)"
          value={data.this || ""}
          onChange={(e) => nextStep({ this: e.target.value })}
        />
      )}
      {step === 3 && (
        <>
          <Suggestions data={data} />
          <button onClick={handleSubmit} disabled={loading}>
            Registrar
          </button>
        </>
      )}
      <FeedbackToast message={feedback} onClose={() => setFeedback("")} />
    </div>
  );
}